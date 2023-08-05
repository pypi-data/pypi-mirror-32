import yaml
import logging
from dli.client.s3 import Client
from dli.client.s3_token_refresher import make_s3_token_refresher
from dli.siren import siren_to_entity, siren_to_dict

from dli.client.exceptions import PackageNotFoundException, DatasetNotFoundException, DownloadFailed

logger = logging.getLogger(__name__)


def to_dict(o, remove_fields=None):
    import inspect
    if remove_fields is None:
        remove_fields = []
    return {
        f: getattr(o, f)
        for f in o.__class__.__dict__.keys()
        if f not in remove_fields and
            not f.startswith('_') and
            hasattr(o, f) and
            not inspect.ismethod(getattr(o, f))
    }


class DatasetFunctions(object):

    def register_dataset_metadata(
        self,
        package_id,
        description,
        files,
        version=1,
        format='csv',
        keywords=None,
        tags=None,
        sources=None
    ):
        """
        Convenience method to register metadata and create a new dataset.
        This function WILL NOT upload files

        :param package_id: The id of the package this dataset belongs to. Can
                 be provided in dataset definition
        :param description: Description of the dataset
        :param files: Path of the files or folders to register
        :param version: Version for this dataset, ``1`` by default
        :param format: Format for this dataset, ``csv`` by default.
        :param keywords: Keywords to be associated with the dataset.
        :param tags: Tags to be associated with the dataset
        :param sources: Sources for lineage definition
        :return: The newly registered dataset
        """
        if not files:
            raise Exception("No files to register have been provided.")

        info = {
            'metadata': {
                'description': description,
                'format': format,
                'keywords': keywords or [],
                'tags': tags or {},
                'version': version,
                'lineage': {
                    'generatedBy': 'SDK upload',
                    'sources': sources or []
                },
                'files': files
            }
        }

        return self._register_dataset(info, package_id=package_id)

    def register_dataset(
        self,
        package_id,
        description,
        s3_prefix,
        files,
        version=1,
        format='csv',
        keywords=None,
        tags=None,
        sources=None
    ):
        """
        Convenience method to register metadata and create a new dataset.
        This function will perform an upload of the files to one
        of the supported data stores

        Supported data stores:
        - s3

        :param package_id: The id of the package this dataset belongs to. Can
                 be provided in dataset definition
        :param description: Description of the dataset
        :param s3_prefix: location for the files in the destination
        :param files: Path of the files or folders to register
        :param version: Version for this dataset, ``1`` by default
        :param format: Format for this dataset, ``csv`` by default.
        :param keywords: Keywords to be associated with the dataset.
        :param tags: Tags to be associated with the dataset
        :param sources: Sources for lineage definition
        :return: The newly registered dataset
        """
        if not files:
            raise Exception("No files to register have been provided.")

        info = {
            'metadata': {
                'description': description,
                'format': format,
                'keywords': keywords or [],
                'tags': tags or {},
                'version': version,
                'lineage': {
                    'generatedBy': 'SDK upload',
                    'sources': sources or []
                },
            },
            'uploads': [
                {
                    'files': files,
                    'target': {'s3': {'prefix': s3_prefix}}
                }
            ]
        }

        return self._register_dataset(info, package_id=package_id)

    def register_dataset_with_config_file(self, path, package_id=None):
        """
        Register a dataset using a YAML file which defines a dataset

        :param path: The path to a YAML file defining a dataset
        :param package_id: The id of the package this dataset belongs to. Can
                 be provided in dataset definition
        :return: A dictionary containing the location of the newly registered
                 dataset
        """
        with open(path, 'r') as stream:
            try:
                info = yaml.load(stream)
                return self._register_dataset(info, package_id)
            except yaml.YAMLError as exc:
                logger.exception("Error: %s.", exc)

    def update_dataset_metadata(
        self,
        dataset_id,
        description=None,
        files=None,
        version=1,
        format='csv',
        keywords=None,
        tags=None,
        sources=None
    ):
        """
        Update a dataset using the given metadata.
        This function will NOT upload files, it just as a means to ammend
        mistakes in a previously submitted dataset instance.

        All arguments are optional, meaning arguments passed as None will not be modified

        :param dataset_id: The id of the dataset to update
        :param description: Description of the dataset
        :param files: Path of the files to upload
        :param s3_prefix: The key under which the files will be stored in S3
        :param keywords: Keywords to be associated with the dataset.
        :param tags: Tags to be associated with the dataset

        :return: The updated dataset
        """
        info = {
            'description': description,
            'files': files,
            'format': format,
            'keywords': keywords,
            'tags': tags,
            'version': version
        }

        if sources:
            info['lineage'] = {
                'generatedBy': 'SDK upload',
                'sources': sources
            }

        info = {k: v for k, v in info.items() if v is not None}

        if not info:
            raise Exception("At least one value needs to be updated has to be provided ")

        return self._update_dataset(info, dataset_id=dataset_id)

    def _update_dataset(self, metadata, dataset_id=None):
        """
        Update a dataset using the given metadata

        :param metadata: A dict representing the updated dataset
        :param dataset_id: The id of the dataset to update
        :return: A dictionary containing the location of the newly registered
                 dataset
        :return: The updated dataset
        """
        if not dataset_id:
            if "datasetId" in metadata:
                dataset_id = metadata["datasetId"]

        if not dataset_id:
            raise Exception("Dataset Id must be provided as a parameter, or as part of metadata")

        dataset = self._get_dataset(dataset_id)
        if not dataset:
            raise Exception("No dataset found with id %s" % dataset_id)

        dataset_as_dict = siren_to_dict(dataset)
        dataset_as_dict['datasetId'] = dataset_id

        # tags are converted to an array when saving the document
        # so they need to be converted back to a dict before updating
        dataset_as_dict["tags"] = {tag["key"]: tag["value"] for tag in dataset_as_dict["tags"]}

        if 'id' in dataset_as_dict:
            del dataset_as_dict['id']

        dataset_as_dict.update(metadata)
        dataset.update_dataset(__json=dataset_as_dict)
        return self.get_dataset(dataset_id)

    def _register_dataset(self, info, package_id=None):
        """
        Manually register a dataset

        :param info: A dictionary containing a `metadata` entry, which
                 specifies the dataset metadata, as well as an 'uploads' entry,
                 to specify the files which should be uploaded and recorded with
                 the metadata.
        :param package_id: The id of the package this dataset belongs to. Can
                 be provided in dataset definition
        :return:
        """
        uploaded_files = []
        metadata = info['metadata']

        if not package_id:
            if "packageId" in metadata:
                package_id = metadata["packageId"]

        if not package_id:
            raise Exception("Package Id must be provided as a parameter, or as part of metadata")

        package = self._get_package(package_id)

        if hasattr(package, 'dataStorage'):
            if package.dataStorage == 'S3':
                if not hasattr(package, 's3Bucket'):
                    raise Exception("There is no bucket associated with the package {}".format(package.id))

                s3_bucket = package.s3Bucket

                if 'uploads' in info:
                    for upload in info['uploads']:
                        uploaded_files.append(self._process_upload(upload, package_id, s3_bucket))

                    flattened = [item for sublist in uploaded_files for item in sublist]
                    metadata['files'] = ["s3://{}".format(f) for f in flattened]

        if not package:
            raise PackageNotFoundException("Package with id {} not found".format(package_id))
        dataset = package.add_dataset(__json=metadata)
        return siren_to_entity(dataset)

    def delete_dataset(self, dataset_id):
        """
        Marks a dataset as deleted.

        :param info: the unique id for the dataset we want to delete.
        :return:
        """
        dataset = self._get_dataset(dataset_id)
        if not dataset:
            raise Exception("No dataset found with id: %s" % dataset_id)

        dataset.delete_dataset(dataset_id=dataset_id)

    def download_dataset(self, dataset_id, destination):
        """
        Helper function that downloads all files
        registered in a dataset into a given destination.

        This function is only supported for data-lake managed s3 buckets,
        otherwise an error will be displayed.

        Currently supports:
          - s3

        :param dataset_id: The id of the dataset we want to download files from
        :param destination: Target location where to store the files (expected to be a directory)
        """

        # get the s3 keys
        # this requires access to be granted
        dataset = self._get_dataset(dataset_id)
        if not dataset:
            raise DatasetNotFoundException("No dataset found with id %s" % dataset_id)

        keys = self.get_s3_access_keys_for_package(dataset.packageId)
        s3_access = Client(
            keys['accessKeyId'],
            keys['secretAccessKey'],
            keys['sessionToken']
        )

        # for each file/folder in the dataset, attempt to download the file
        # rather than failing at the same error, keep to download as much as possible
        # and fail at the end.
        failed = []
        for file in dataset.files:
            try:
                s3_access.download_file(file, destination)
            except Exception:
                logger.exception("Failed to download file `%s` from dataset `%s`", file, dataset_id)
                failed.append(file)

        if failed:
            raise DownloadFailed(
                "Some files in this dataset could not be downloaded, "
                "see logs for detailed information. Failed:\n%s"
                % "\n".join(failed)
            )

    def _process_upload(self, upload, package_id, s3_bucket):
        """
        Given an `upload` dict spec, process the files by uploading them to
        the specified target

        Currently supports:
          - s3

        :param upload: A dictionary specifying an upload `target` dict and a
                 list of files to upload
                 Example:
                 {
                     'files': ['DrowningsPerYearVsNicolasCageMovies.csv']
                     'target': {
                         's3': {
                             'prefix': '/spurious-correlations/'
                         }
                     }
                 }
        :param package_id: The id of the package this dataset belongs to. Can
                 be provided in dataset definition
        :param s3_bucket: The name of the S3 bucket you wish to upload files to
        :return: A list of paths of the uploaded files
        """
        files = upload['files']
        target = upload['target']

        if 's3' in target:
            prefix = target['s3']['prefix']
            s3_location = f'{s3_bucket}/{prefix}'
            return self._process_s3_upload(files, s3_location, package_id)
        else:
            raise Exception("Only S3 uploads are currently supported")

    def _process_s3_upload(self, files, s3_location, package_id):
        s3_access_keys = self.get_s3_access_keys_for_package(package_id)
        token_refresher = make_s3_token_refresher(self, package_id)
        s3_client = Client(s3_access_keys['accessKeyId'], s3_access_keys['secretAccessKey'], s3_access_keys['sessionToken'])
        return s3_client.upload_files_to_s3(files, s3_location, token_refresher)

    def get_dataset(self, dataset_id):
        dataset = self._get_dataset(dataset_id)
        if not dataset:
            return

        return siren_to_entity(dataset)

    def _get_dataset(self, dataset_id):
        return self.get_root_siren().get_dataset(dataset_id=dataset_id)

    def add_files_to_dataset(self, dataset_id, s3_prefix, files):
        """
        Upload files to existing dataset

        :param dataset_id: The id of the dataset to be updated
        :param s3_prefix: Location for the files in the destination s3 bucket
        :param files: List of files to be added to the dataset

        :return: The updated dataset
        """
        dataset = self.get_dataset(dataset_id)
        if dataset:
            pkg = self.get_package(dataset.packageId)
            if pkg:
                s3_location = f'{pkg.s3Bucket}/{s3_prefix}'
                uploaded_files = [f's3://{f}' for f in self._process_s3_upload(files, s3_location, pkg.id)]

                if dataset.files:
                    uploaded_files.extend(dataset.files)

                metadata_update = {'files': uploaded_files}

                return self._update_dataset(metadata_update, dataset_id)
            else:
                raise PackageNotFoundException(f'Package {dataset.packageId} associated with the dataset {dataset_id} does not exist!')
        else:
            raise DatasetNotFoundException(f'Dataset with id {dataset_id} not found')

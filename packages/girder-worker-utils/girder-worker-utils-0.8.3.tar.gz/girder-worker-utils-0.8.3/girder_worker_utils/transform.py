import abc

import six


@six.add_metaclass(abc.ABCMeta)
class Transform(object):
    def _repr_model_(self):
        """Return as representation of the object suitable for storing in mongo.

        This function retuns a string representation of the object
        that is stored in a girder Job Model's 'args' or 'kwargs'
        fields.  It is for display purposes only.

        """
        return str(self)

    def cleanup(self):
        pass

    @abc.abstractmethod
    def transform(self):
        pass


@six.add_metaclass(abc.ABCMeta)
class ResultTransform(Transform):
    @abc.abstractmethod
    def transform(self, data):
        pass

# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class User(RESTObject):
    """ Represents a User in the 

        Notes:
            A User represents the owner of some certificates.
    """

    def __init__(self, **kwargs):
        """ Initializes a User instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> user = User(id=u'xxxx-xxx-xxx-xxx', name=u'User')
              >>> user = User(data=my_dict)
        """

        super(User, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._certificate = None
        self._certificateexpirationdate = None
        self._certificatestatus = None
        self._createdat = None
        self._deleted = None
        self._description = None
        self._email = None
        self._key = None
        self._name = None
        self._namespace = None
        self._normalizedtags = None
        self._parentauthenticator = None
        self._parentid = None
        self._parenttype = None
        self._status = None
        self._suborganizations = None
        self._updatedat = None
        self._username = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="certificate", remote_name="certificate")
        self.expose_attribute(local_name="certificateExpirationDate", remote_name="certificateExpirationDate")
        self.expose_attribute(local_name="certificateStatus", remote_name="certificateStatus")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="email", remote_name="email")
        self.expose_attribute(local_name="key", remote_name="key")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="parentAuthenticator", remote_name="parentAuthenticator")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="subOrganizations", remote_name="subOrganizations")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")
        self.expose_attribute(local_name="userName", remote_name="userName")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        
        return self.ID
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        
        self.ID = ID
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return userIdentity

    # Properties
    @property
    def ID(self):
        """ Get ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        return self._id

    @ID.setter
    def ID(self, value):
        """ Set ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        self._id = value
    
    @property
    def annotation(self):
        """ Get annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        return self._annotation

    @annotation.setter
    def annotation(self, value):
        """ Set annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        self._annotation = value
    
    @property
    def associatedTags(self):
        """ Get associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        return self._associatedtags

    @associatedTags.setter
    def associatedTags(self, value):
        """ Set associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        self._associatedtags = value
    
    @property
    def certificate(self):
        """ Get certificate value.

          Notes:
              Certificate provides a certificate for the user

              
        """
        return self._certificate

    @certificate.setter
    def certificate(self, value):
        """ Set certificate value.

          Notes:
              Certificate provides a certificate for the user

              
        """
        self._certificate = value
    
    @property
    def certificateExpirationDate(self):
        """ Get certificateExpirationDate value.

          Notes:
              CertificateExpirationDate indicates the expiration day for the certificate.

              
        """
        return self._certificateexpirationdate

    @certificateExpirationDate.setter
    def certificateExpirationDate(self, value):
        """ Set certificateExpirationDate value.

          Notes:
              CertificateExpirationDate indicates the expiration day for the certificate.

              
        """
        self._certificateexpirationdate = value
    
    @property
    def certificateStatus(self):
        """ Get certificateStatus value.

          Notes:
              CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.

              
        """
        return self._certificatestatus

    @certificateStatus.setter
    def certificateStatus(self, value):
        """ Set certificateStatus value.

          Notes:
              CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.

              
        """
        self._certificatestatus = value
    
    @property
    def createdAt(self):
        """ Get createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        return self._createdat

    @createdAt.setter
    def createdAt(self, value):
        """ Set createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        self._createdat = value
    
    @property
    def deleted(self):
        """ Get deleted value.

          Notes:
              Deleted marks if the entity has been deleted.

              
        """
        return self._deleted

    @deleted.setter
    def deleted(self, value):
        """ Set deleted value.

          Notes:
              Deleted marks if the entity has been deleted.

              
        """
        self._deleted = value
    
    @property
    def description(self):
        """ Get description value.

          Notes:
              Description is the description of the object.

              
        """
        return self._description

    @description.setter
    def description(self, value):
        """ Set description value.

          Notes:
              Description is the description of the object.

              
        """
        self._description = value
    
    @property
    def email(self):
        """ Get email value.

          Notes:
              e-mail address of the user

              
        """
        return self._email

    @email.setter
    def email(self, value):
        """ Set email value.

          Notes:
              e-mail address of the user

              
        """
        self._email = value
    
    @property
    def key(self):
        """ Get key value.

          Notes:
              Key provides the key for the user. Only available at create or update time.

              
        """
        return self._key

    @key.setter
    def key(self, value):
        """ Set key value.

          Notes:
              Key provides the key for the user. Only available at create or update time.

              
        """
        self._key = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the entity

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the entity

              
        """
        self._name = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        self._namespace = value
    
    @property
    def normalizedTags(self):
        """ Get normalizedTags value.

          Notes:
              NormalizedTags contains the list of normalized tags of the entities

              
        """
        return self._normalizedtags

    @normalizedTags.setter
    def normalizedTags(self, value):
        """ Set normalizedTags value.

          Notes:
              NormalizedTags contains the list of normalized tags of the entities

              
        """
        self._normalizedtags = value
    
    @property
    def parentAuthenticator(self):
        """ Get parentAuthenticator value.

          Notes:
              ParentAuthenticator is an Internal attribute that points to the parent authenticator.

              
        """
        return self._parentauthenticator

    @parentAuthenticator.setter
    def parentAuthenticator(self, value):
        """ Set parentAuthenticator value.

          Notes:
              ParentAuthenticator is an Internal attribute that points to the parent authenticator.

              
        """
        self._parentauthenticator = value
    
    @property
    def parentID(self):
        """ Get parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        return self._parentid

    @parentID.setter
    def parentID(self, value):
        """ Set parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        self._parentid = value
    
    @property
    def parentType(self):
        """ Get parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        return self._parenttype

    @parentType.setter
    def parentType(self, value):
        """ Set parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        self._parenttype = value
    
    @property
    def status(self):
        """ Get status value.

          Notes:
              Status of an entity

              
        """
        return self._status

    @status.setter
    def status(self, value):
        """ Set status value.

          Notes:
              Status of an entity

              
        """
        self._status = value
    
    @property
    def subOrganizations(self):
        """ Get subOrganizations value.

          Notes:
              OU attribute for the generated certificates

              
        """
        return self._suborganizations

    @subOrganizations.setter
    def subOrganizations(self, value):
        """ Set subOrganizations value.

          Notes:
              OU attribute for the generated certificates

              
        """
        self._suborganizations = value
    
    @property
    def updatedAt(self):
        """ Get updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        return self._updatedat

    @updatedAt.setter
    def updatedAt(self, value):
        """ Set updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        self._updatedat = value
    
    @property
    def userName(self):
        """ Get userName value.

          Notes:
              CommonName (CN) for the user certificate

              
        """
        return self._username

    @userName.setter
    def userName(self, value):
        """ Set userName value.

          Notes:
              CommonName (CN) for the user certificate

              
        """
        self._username = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_string_in_list("certificateStatus", self.certificateStatus, ["RENEW", "REVOKED", "VALID"], false)

        if err:
            errors.append(err)

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        err = validate_maximum_length("userName", self.userName, 64, false)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # userIdentity represents the Identity of the object
userIdentity = {"name": "user", "category": "users", "constructor": User}
/**
 * @file DistribGeometricDistribution.cpp
 * @brief Implementation of the DistribGeometricDistribution class.
 * @author SBMLTeam
 *
 * <!--------------------------------------------------------------------------
 * This file is part of libSBML. Please visit http://sbml.org for more
 * information about SBML, and the latest version of libSBML.
 *
 * Copyright (C) 2013-2018 jointly by the following organizations:
 * 1. California Institute of Technology, Pasadena, CA, USA
 * 2. EMBL European Bioinformatics Institute (EMBL-EBI), Hinxton, UK
 * 3. University of Heidelberg, Heidelberg, Germany
 *
 * Copyright (C) 2009-2013 jointly by the following organizations:
 * 1. California Institute of Technology, Pasadena, CA, USA
 * 2. EMBL European Bioinformatics Institute (EMBL-EBI), Hinxton, UK
 *
 * Copyright (C) 2006-2008 by the California Institute of Technology,
 * Pasadena, CA, USA
 *
 * Copyright (C) 2002-2005 jointly by the following organizations:
 * 1. California Institute of Technology, Pasadena, CA, USA
 * 2. Japan Science and Technology Agency, Japan
 *
 * This library is free software; you can redistribute it and/or modify it
 * under the terms of the GNU Lesser General Public License as published by the
 * Free Software Foundation. A copy of the license agreement is provided in the
 * file named "LICENSE.txt" included with this software distribution and also
 * available online as http://sbml.org/software/libsbml/license.html
 * ------------------------------------------------------------------------ -->
 */
#include <sbml/packages/distrib/sbml/DistribGeometricDistribution.h>
#include <sbml/packages/distrib/validator/DistribSBMLError.h>
#include <sbml/util/ElementFilter.h>


using namespace std;



LIBSBML_CPP_NAMESPACE_BEGIN




#ifdef __cplusplus


/*
 * Creates a new DistribGeometricDistribution using the given SBML Level,
 * Version and &ldquo;distrib&rdquo; package version.
 */
DistribGeometricDistribution::DistribGeometricDistribution(unsigned int level,
                                                           unsigned int
                                                             version,
                                                           unsigned int
                                                             pkgVersion)
  : DistribDiscreteUnivariateDistribution(level, version)
  , mProbability (NULL)
{
  setSBMLNamespacesAndOwn(new DistribPkgNamespaces(level, version,
    pkgVersion));
  connectToChild();
}


/*
 * Creates a new DistribGeometricDistribution using the given
 * DistribPkgNamespaces object.
 */
DistribGeometricDistribution::DistribGeometricDistribution(DistribPkgNamespaces
  *distribns)
  : DistribDiscreteUnivariateDistribution(distribns)
  , mProbability (NULL)
{
  setElementNamespace(distribns->getURI());
  connectToChild();
  loadPlugins(distribns);
}


/*
 * Copy constructor for DistribGeometricDistribution.
 */
DistribGeometricDistribution::DistribGeometricDistribution(const
  DistribGeometricDistribution& orig)
  : DistribDiscreteUnivariateDistribution( orig )
  , mProbability ( NULL )
{
  if (orig.mProbability != NULL)
  {
    mProbability = orig.mProbability->clone();
  }

  connectToChild();
}


/*
 * Assignment operator for DistribGeometricDistribution.
 */
DistribGeometricDistribution&
DistribGeometricDistribution::operator=(const DistribGeometricDistribution&
  rhs)
{
  if (&rhs != this)
  {
    DistribDiscreteUnivariateDistribution::operator=(rhs);
    delete mProbability;
    if (rhs.mProbability != NULL)
    {
      mProbability = rhs.mProbability->clone();
    }
    else
    {
      mProbability = NULL;
    }

    connectToChild();
  }

  return *this;
}


/*
 * Creates and returns a deep copy of this DistribGeometricDistribution object.
 */
DistribGeometricDistribution*
DistribGeometricDistribution::clone() const
{
  return new DistribGeometricDistribution(*this);
}


/*
 * Destructor for DistribGeometricDistribution.
 */
DistribGeometricDistribution::~DistribGeometricDistribution()
{
  delete mProbability;
  mProbability = NULL;
}


/*
 * Returns the value of the "id" attribute of this
 * DistribGeometricDistribution.
 */
const std::string&
DistribGeometricDistribution::getId() const
{
  return mId;
}


/*
 * Returns the value of the "name" attribute of this
 * DistribGeometricDistribution.
 */
const std::string&
DistribGeometricDistribution::getName() const
{
  return mName;
}


/*
 * Predicate returning @c true if this DistribGeometricDistribution's "id"
 * attribute is set.
 */
bool
DistribGeometricDistribution::isSetId() const
{
  return (mId.empty() == false);
}


/*
 * Predicate returning @c true if this DistribGeometricDistribution's "name"
 * attribute is set.
 */
bool
DistribGeometricDistribution::isSetName() const
{
  return (mName.empty() == false);
}


/*
 * Sets the value of the "id" attribute of this DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setId(const std::string& id)
{
  return SyntaxChecker::checkAndSetSId(id, mId);
}


/*
 * Sets the value of the "name" attribute of this DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setName(const std::string& name)
{
  mName = name;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Unsets the value of the "id" attribute of this DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::unsetId()
{
  mId.erase();

  if (mId.empty() == true)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else
  {
    return LIBSBML_OPERATION_FAILED;
  }
}


/*
 * Unsets the value of the "name" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::unsetName()
{
  mName.erase();

  if (mName.empty() == true)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else
  {
    return LIBSBML_OPERATION_FAILED;
  }
}


/*
 * Returns the value of the "probability" element of this
 * DistribGeometricDistribution.
 */
const DistribUncertValue*
DistribGeometricDistribution::getProbability() const
{
  return mProbability;
}


/*
 * Returns the value of the "probability" element of this
 * DistribGeometricDistribution.
 */
DistribUncertValue*
DistribGeometricDistribution::getProbability()
{
  return mProbability;
}


/*
 * Predicate returning @c true if this DistribGeometricDistribution's
 * "probability" element is set.
 */
bool
DistribGeometricDistribution::isSetProbability() const
{
  return (mProbability != NULL);
}


/*
 * Sets the value of the "probability" element of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setProbability(const DistribUncertValue*
  probability)
{
  if (probability == NULL)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else if (probability->hasRequiredElements() == false)
  {
    return LIBSBML_INVALID_OBJECT;
  }
  else if (getLevel() != probability->getLevel())
  {
    return LIBSBML_LEVEL_MISMATCH;
  }
  else if (getVersion() != probability->getVersion())
  {
    return LIBSBML_VERSION_MISMATCH;
  }
  else if (getPackageVersion() != probability->getPackageVersion())
  {
    return LIBSBML_PKG_VERSION_MISMATCH;
  }
  else
  {
    delete mProbability;
    mProbability = (probability != NULL) ?
      static_cast<DistribUncertValue*>(probability->clone()) : NULL;
    if (mProbability != NULL) mProbability->connectToParent(this);
    return LIBSBML_OPERATION_SUCCESS;
  }
}


/*
 * Creates a new DistribUncertValue object, adds it to this
 * DistribGeometricDistribution object and returns the DistribUncertValue
 * object created.
 */
DistribUncertValue*
DistribGeometricDistribution::createProbability()
{
  if (mProbability != NULL)
  {
    delete mProbability;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mProbability = new DistribUncertValue(distribns);

  mProbability->setElementName("probability");

  delete distribns;

  connectToChild();

  return mProbability;
}


/*
 * Unsets the value of the "probability" element of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::unsetProbability()
{
  delete mProbability;
  mProbability = NULL;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Returns the XML element name of this DistribGeometricDistribution object.
 */
const std::string&
DistribGeometricDistribution::getElementName() const
{
  static const string name = "geometricDistribution";
  return name;
}


/*
 * Returns the libSBML type code for this DistribGeometricDistribution object.
 */
int
DistribGeometricDistribution::getTypeCode() const
{
  return SBML_DISTRIB_GEOMETRICLDISTRIBUTION;
}


/*
 * Predicate returning @c true if all the required attributes for this
 * DistribGeometricDistribution object have been set.
 */
bool
DistribGeometricDistribution::hasRequiredAttributes() const
{
  bool allPresent =
    DistribDiscreteUnivariateDistribution::hasRequiredAttributes();

  return allPresent;
}


/*
 * Predicate returning @c true if all the required elements for this
 * DistribGeometricDistribution object have been set.
 */
bool
DistribGeometricDistribution::hasRequiredElements() const
{
  bool allPresent =
    DistribDiscreteUnivariateDistribution::hasRequiredElements();

  if (isSetProbability() == false)
  {
    allPresent = false;
  }

  return allPresent;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Write any contained elements
 */
void
DistribGeometricDistribution::writeElements(XMLOutputStream& stream) const
{
  DistribDiscreteUnivariateDistribution::writeElements(stream);

  if (isSetProbability() == true)
  {
    mProbability->write(stream);
  }

  SBase::writeExtensionElements(stream);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Accepts the given SBMLVisitor
 */
bool
DistribGeometricDistribution::accept(SBMLVisitor& v) const
{
  v.visit(*this);

  if (mProbability != NULL)
  {
    mProbability->accept(v);
  }

  v.leave(*this);
  return true;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the parent SBMLDocument
 */
void
DistribGeometricDistribution::setSBMLDocument(SBMLDocument* d)
{
  DistribDiscreteUnivariateDistribution::setSBMLDocument(d);

  if (mProbability != NULL)
  {
    mProbability->setSBMLDocument(d);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Connects to child elements
 */
void
DistribGeometricDistribution::connectToChild()
{
  DistribDiscreteUnivariateDistribution::connectToChild();

  if (mProbability != NULL)
  {
    mProbability->connectToParent(this);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Enables/disables the given package with this element
 */
void
DistribGeometricDistribution::enablePackageInternal(const std::string& pkgURI,
                                                    const std::string&
                                                      pkgPrefix,
                                                    bool flag)
{
  DistribDiscreteUnivariateDistribution::enablePackageInternal(pkgURI,
    pkgPrefix, flag);

  if (isSetProbability())
  {
    mProbability->enablePackageInternal(pkgURI, pkgPrefix, flag);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Updates the namespaces when setLevelVersion is used
 */
void
DistribGeometricDistribution::updateSBMLNamespace(const std::string& package,
                                                  unsigned int level,
                                                  unsigned int version)
{
  DistribDiscreteUnivariateDistribution::updateSBMLNamespace(package, level,
    version);

  if (mProbability != NULL)
  {
    mProbability->updateSBMLNamespace(package, level, version);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::getAttribute(const std::string& attributeName,
                                           bool& value) const
{
  int return_value =
    DistribDiscreteUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::getAttribute(const std::string& attributeName,
                                           int& value) const
{
  int return_value =
    DistribDiscreteUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::getAttribute(const std::string& attributeName,
                                           double& value) const
{
  int return_value =
    DistribDiscreteUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::getAttribute(const std::string& attributeName,
                                           unsigned int& value) const
{
  int return_value =
    DistribDiscreteUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::getAttribute(const std::string& attributeName,
                                           std::string& value) const
{
  int return_value =
    DistribDiscreteUnivariateDistribution::getAttribute(attributeName, value);

  if (return_value == LIBSBML_OPERATION_SUCCESS)
  {
    return return_value;
  }

  if (attributeName == "id")
  {
    value = getId();
    return_value = LIBSBML_OPERATION_SUCCESS;
  }
  else if (attributeName == "name")
  {
    value = getName();
    return_value = LIBSBML_OPERATION_SUCCESS;
  }

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Predicate returning @c true if this DistribGeometricDistribution's attribute
 * "attributeName" is set.
 */
bool
DistribGeometricDistribution::isSetAttribute(const std::string& attributeName)
  const
{
  bool value =
    DistribDiscreteUnivariateDistribution::isSetAttribute(attributeName);

  if (attributeName == "id")
  {
    value = isSetId();
  }
  else if (attributeName == "name")
  {
    value = isSetName();
  }

  return value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setAttribute(const std::string& attributeName,
                                           bool value)
{
  int return_value =
    DistribDiscreteUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setAttribute(const std::string& attributeName,
                                           int value)
{
  int return_value =
    DistribDiscreteUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setAttribute(const std::string& attributeName,
                                           double value)
{
  int return_value =
    DistribDiscreteUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setAttribute(const std::string& attributeName,
                                           unsigned int value)
{
  int return_value =
    DistribDiscreteUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::setAttribute(const std::string& attributeName,
                                           const std::string& value)
{
  int return_value =
    DistribDiscreteUnivariateDistribution::setAttribute(attributeName, value);

  if (attributeName == "id")
  {
    return_value = setId(value);
  }
  else if (attributeName == "name")
  {
    return_value = setName(value);
  }

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Unsets the value of the "attributeName" attribute of this
 * DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::unsetAttribute(const std::string& attributeName)
{
  int value =
    DistribDiscreteUnivariateDistribution::unsetAttribute(attributeName);

  if (attributeName == "id")
  {
    value = unsetId();
  }
  else if (attributeName == "name")
  {
    value = unsetName();
  }

  return value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Creates and returns an new "elementName" object in this
 * DistribGeometricDistribution.
 */
SBase*
DistribGeometricDistribution::createChildObject(const std::string& elementName)
{
  DistribDiscreteUnivariateDistribution* obj = NULL;

  if (elementName == "probability")
  {
    return createProbability();
  }

  return obj;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Adds a new "elementName" object to this DistribGeometricDistribution.
 */
int
DistribGeometricDistribution::addChildObject(const std::string& elementName,
                                             const SBase* element)
{
  if (elementName == "probability" && element->getTypeCode() ==
    SBML_DISTRIB_UNCERTVALUE)
  {
    return setProbability((const DistribUncertValue*)(element));
  }

  return LIBSBML_OPERATION_FAILED;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Removes and returns the new "elementName" object with the given id in this
 * DistribGeometricDistribution.
 */
SBase*
DistribGeometricDistribution::removeChildObject(const std::string& elementName,
                                                const std::string& id)
{
  if (elementName == "probability")
  {
    DistribUncertValue * obj = getProbability();
    if (unsetProbability() == LIBSBML_OPERATION_SUCCESS) return obj;
  }

  return NULL;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the number of "elementName" in this DistribGeometricDistribution.
 */
unsigned int
DistribGeometricDistribution::getNumObjects(const std::string& elementName)
{
  unsigned int n = 0;

  if (elementName == "probability")
  {
    if (isSetProbability())
    {
      return 1;
    }
  }

  return n;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the nth object of "objectName" in this DistribGeometricDistribution.
 */
SBase*
DistribGeometricDistribution::getObject(const std::string& elementName,
                                        unsigned int index)
{
  SBase* obj = NULL;

  if (elementName == "probability")
  {
    return getProbability();
  }

  return obj;
}

/** @endcond */


/*
 * Returns the first child element that has the given @p id in the model-wide
 * SId namespace, or @c NULL if no such object is found.
 */
SBase*
DistribGeometricDistribution::getElementBySId(const std::string& id)
{
  if (id.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mProbability != NULL)
  {
    if (mProbability->getId() == id)
    {
      return mProbability;
    }

    obj = mProbability->getElementBySId(id);
    if (obj != NULL)
    {
      return obj;
    }
  }

  return obj;
}


/*
 * Returns the first child element that has the given @p metaid, or @c NULL if
 * no such object is found.
 */
SBase*
DistribGeometricDistribution::getElementByMetaId(const std::string& metaid)
{
  if (metaid.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mProbability != NULL)
  {
    if (mProbability->getMetaId() == metaid)
    {
      return mProbability;
    }

    obj = mProbability->getElementByMetaId(metaid);
    if (obj != NULL)
    {
      return obj;
    }
  }

  return obj;
}


/*
 * Returns a List of all child SBase objects, including those nested to an
 * arbitrary depth.
 */
List*
DistribGeometricDistribution::getAllElements(ElementFilter* filter)
{
  List* ret = new List();
  List* sublist = NULL;

  ADD_FILTERED_POINTER(ret, sublist, mProbability, filter);


  ADD_FILTERED_FROM_PLUGIN(ret, sublist, filter);

  return ret;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Creates a new object from the next XMLToken on the XMLInputStream
 */
SBase*
DistribGeometricDistribution::createObject(XMLInputStream& stream)
{
  SBase* obj = DistribDiscreteUnivariateDistribution::createObject(stream);

  const std::string& name = stream.peek().getName();

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());

  if (name == "probability")
  {
    if (isSetProbability())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribGeometricDistributionAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mProbability;
    mProbability = new DistribUncertValue(distribns);
    mProbability->setElementName(name);
    obj = mProbability;
  }

  delete distribns;

  connectToChild();

  return obj;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Adds the expected attributes for this element
 */
void
DistribGeometricDistribution::addExpectedAttributes(ExpectedAttributes&
  attributes)
{
  DistribDiscreteUnivariateDistribution::addExpectedAttributes(attributes);

  unsigned int level = getLevel();
  unsigned int coreVersion = getVersion();
  unsigned int pkgVersion = getPackageVersion();

  if (level == 3 && coreVersion == 1 && pkgVersion == 1)
  {
    attributes.add("id");
    attributes.add("name");
  }

  if (level == 3 && coreVersion == 2 && pkgVersion == 1)
  {
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Reads the expected attributes into the member data variables
 */
void
DistribGeometricDistribution::readAttributes(const XMLAttributes& attributes,
                                             const ExpectedAttributes&
                                               expectedAttributes)
{
  unsigned int level = getLevel();
  unsigned int version = getVersion();
  unsigned int pkgVersion = getPackageVersion();
  unsigned int numErrs;
  bool assigned = false;
  SBMLErrorLog* log = getErrorLog();

  DistribDiscreteUnivariateDistribution::readAttributes(attributes,
    expectedAttributes);

  if (log)
  {
    numErrs = log->getNumErrors();

    for (int n = numErrs-1; n >= 0; n--)
    {
      if (log->getError(n)->getErrorId() == UnknownPackageAttribute)
      {
        const std::string details = log->getError(n)->getMessage();
        log->remove(UnknownPackageAttribute);
        log->logPackageError("distrib",
          DistribDistribGeometricDistributionAllowedAttributes, pkgVersion,
            level, version, details);
      }
      else if (log->getError(n)->getErrorId() == UnknownCoreAttribute)
      {
        const std::string details = log->getError(n)->getMessage();
        log->remove(UnknownCoreAttribute);
        log->logPackageError("distrib",
          DistribDistribGeometricDistributionAllowedCoreAttributes, pkgVersion,
            level, version, details);
      }
    }
  }

  if (level == 3 && version == 1 && pkgVersion == 1)
  {
    readL3V1V1Attributes(attributes);
  }

  if (level == 3 && version == 2 && pkgVersion == 1)
  {
    readL3V2V1Attributes(attributes);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Reads the expected attributes into the member data variables
 */
void
DistribGeometricDistribution::readL3V1V1Attributes(const XMLAttributes&
  attributes)
{
  unsigned int level = getLevel();
  unsigned int version = getVersion();
  bool assigned = false;
  unsigned int pkgVersion = getPackageVersion();
  SBMLErrorLog* log = getErrorLog();

  // 
  // id SId (use = "optional" )
  // 

  XMLTriple tripleID("id", mURI, getPrefix());
  assigned = attributes.readInto(tripleID, mId);

  if (assigned == true)
  {
    if (mId.empty() == true)
    {
      logEmptyString(mId, level, version, "<DistribGeometricDistribution>");
    }
    else if (SyntaxChecker::isValidSBMLSId(mId) == false)
    {
      log->logPackageError("distrib", DistribIdSyntaxRule, pkgVersion, level,
        version, "The id on the <" + getElementName() + "> is '" + mId + "', "
          "which does not conform to the syntax.", getLine(), getColumn());
    }
  }

  // 
  // name string (use = "optional" )
  // 

  XMLTriple tripleNAME("name", mURI, getPrefix());
  assigned = attributes.readInto(tripleNAME, mName);

  if (assigned == true)
  {
    if (mName.empty() == true)
    {
      logEmptyString(mName, level, version, "<DistribGeometricDistribution>");
    }
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Reads the expected attributes into the member data variables
 */
void
DistribGeometricDistribution::readL3V2V1Attributes(const XMLAttributes&
  attributes)
{
  unsigned int level = getLevel();
  unsigned int version = getVersion();
  bool assigned = false;
  unsigned int pkgVersion = getPackageVersion();
  SBMLErrorLog* log = getErrorLog();

  // 
  // id SId (use = "optional" )
  // 

  assigned = attributes.readInto("id", mId);

  if (assigned == true)
  {
    if (mId.empty() == true)
    {
      logEmptyString(mId, level, version, "<DistribGeometricDistribution>");
    }
    else if (SyntaxChecker::isValidSBMLSId(mId) == false)
    {
      log->logPackageError("distrib", DistribIdSyntaxRule, pkgVersion, level,
        version, "The id on the <" + getElementName() + "> is '" + mId + "', "
          "which does not conform to the syntax.", getLine(), getColumn());
    }
  }

  // 
  // name string (use = "optional" )
  // 

  // read by SBase;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Writes the attributes to the stream
 */
void
DistribGeometricDistribution::writeAttributes(XMLOutputStream& stream) const
{
  DistribDiscreteUnivariateDistribution::writeAttributes(stream);

  unsigned int level = getLevel();
  unsigned int version = getVersion();
  unsigned int pkgVersion = getPackageVersion();

  if (level == 3 && version == 1 && pkgVersion == 1)
  {
    writeL3V1V1Attributes(stream);
  }

  if (level == 3 && version == 2 && pkgVersion == 1)
  {
    writeL3V2V1Attributes(stream);
  }

  SBase::writeExtensionAttributes(stream);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Writes the attributes to the stream
 */
void
DistribGeometricDistribution::writeL3V1V1Attributes(XMLOutputStream& stream)
  const
{
  if (isSetId() == true)
  {
    stream.writeAttribute("id", getPrefix(), mId);
  }

  if (isSetName() == true)
  {
    stream.writeAttribute("name", getPrefix(), mName);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Writes the attributes to the stream
 */
void
DistribGeometricDistribution::writeL3V2V1Attributes(XMLOutputStream& stream)
  const
{
}

/** @endcond */




#endif /* __cplusplus */


/*
 * Creates a new DistribGeometricDistribution_t using the given SBML Level,
 * Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribGeometricDistribution_t *
DistribGeometricDistribution_create(unsigned int level,
                                    unsigned int version,
                                    unsigned int pkgVersion)
{
  return new DistribGeometricDistribution(level, version, pkgVersion);
}


/*
 * Creates and returns a deep copy of this DistribGeometricDistribution_t
 * object.
 */
LIBSBML_EXTERN
DistribGeometricDistribution_t*
DistribGeometricDistribution_clone(const DistribGeometricDistribution_t* dgd)
{
  if (dgd != NULL)
  {
    return static_cast<DistribGeometricDistribution_t*>(dgd->clone());
  }
  else
  {
    return NULL;
  }
}


/*
 * Frees this DistribGeometricDistribution_t object.
 */
LIBSBML_EXTERN
void
DistribGeometricDistribution_free(DistribGeometricDistribution_t* dgd)
{
  if (dgd != NULL)
  {
    delete dgd;
  }
}


/*
 * Returns the value of the "id" attribute of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
char *
DistribGeometricDistribution_getId(const DistribGeometricDistribution_t * dgd)
{
  if (dgd == NULL)
  {
    return NULL;
  }

  return dgd->getId().empty() ? NULL : safe_strdup(dgd->getId().c_str());
}


/*
 * Returns the value of the "name" attribute of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
char *
DistribGeometricDistribution_getName(const DistribGeometricDistribution_t *
  dgd)
{
  if (dgd == NULL)
  {
    return NULL;
  }

  return dgd->getName().empty() ? NULL : safe_strdup(dgd->getName().c_str());
}


/*
 * Predicate returning @c 1 (true) if this DistribGeometricDistribution_t's
 * "id" attribute is set.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_isSetId(const DistribGeometricDistribution_t *
  dgd)
{
  return (dgd != NULL) ? static_cast<int>(dgd->isSetId()) : 0;
}


/*
 * Predicate returning @c 1 (true) if this DistribGeometricDistribution_t's
 * "name" attribute is set.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_isSetName(const DistribGeometricDistribution_t *
  dgd)
{
  return (dgd != NULL) ? static_cast<int>(dgd->isSetName()) : 0;
}


/*
 * Sets the value of the "id" attribute of this DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_setId(DistribGeometricDistribution_t * dgd,
                                   const char * id)
{
  return (dgd != NULL) ? dgd->setId(id) : LIBSBML_INVALID_OBJECT;
}


/*
 * Sets the value of the "name" attribute of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_setName(DistribGeometricDistribution_t * dgd,
                                     const char * name)
{
  return (dgd != NULL) ? dgd->setName(name) : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "id" attribute of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_unsetId(DistribGeometricDistribution_t * dgd)
{
  return (dgd != NULL) ? dgd->unsetId() : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "name" attribute of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_unsetName(DistribGeometricDistribution_t * dgd)
{
  return (dgd != NULL) ? dgd->unsetName() : LIBSBML_INVALID_OBJECT;
}


/*
 * Returns the value of the "probability" element of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
const DistribUncertValue_t*
DistribGeometricDistribution_getProbability(const
  DistribGeometricDistribution_t * dgd)
{
  if (dgd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dgd->getProbability());
}


/*
 * Predicate returning @c 1 (true) if this DistribGeometricDistribution_t's
 * "probability" element is set.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_isSetProbability(const
  DistribGeometricDistribution_t * dgd)
{
  return (dgd != NULL) ? static_cast<int>(dgd->isSetProbability()) : 0;
}


/*
 * Sets the value of the "probability" element of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_setProbability(
                                            DistribGeometricDistribution_t *
                                              dgd,
                                            const DistribUncertValue_t*
                                              probability)
{
  return (dgd != NULL) ? dgd->setProbability(probability) :
    LIBSBML_INVALID_OBJECT;
}


/*
 * Creates a new DistribUncertValue_t object, adds it to this
 * DistribGeometricDistribution_t object and returns the DistribUncertValue_t
 * object created.
 */
LIBSBML_EXTERN
DistribUncertValue_t*
DistribGeometricDistribution_createProbability(DistribGeometricDistribution_t*
  dgd)
{
  if (dgd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dgd->createProbability());
}


/*
 * Unsets the value of the "probability" element of this
 * DistribGeometricDistribution_t.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_unsetProbability(DistribGeometricDistribution_t *
  dgd)
{
  return (dgd != NULL) ? dgd->unsetProbability() : LIBSBML_INVALID_OBJECT;
}


/*
 * Predicate returning @c 1 (true) if all the required attributes for this
 * DistribGeometricDistribution_t object have been set.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_hasRequiredAttributes(const
  DistribGeometricDistribution_t * dgd)
{
  return (dgd != NULL) ? static_cast<int>(dgd->hasRequiredAttributes()) : 0;
}


/*
 * Predicate returning @c 1 (true) if all the required elements for this
 * DistribGeometricDistribution_t object have been set.
 */
LIBSBML_EXTERN
int
DistribGeometricDistribution_hasRequiredElements(const
  DistribGeometricDistribution_t * dgd)
{
  return (dgd != NULL) ? static_cast<int>(dgd->hasRequiredElements()) : 0;
}




LIBSBML_CPP_NAMESPACE_END



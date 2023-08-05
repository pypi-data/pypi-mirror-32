/**
 * @file DistribBernoulliDistribution.cpp
 * @brief Implementation of the DistribBernoulliDistribution class.
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
#include <sbml/packages/distrib/sbml/DistribBernoulliDistribution.h>
#include <sbml/packages/distrib/validator/DistribSBMLError.h>
#include <sbml/util/ElementFilter.h>


using namespace std;



LIBSBML_CPP_NAMESPACE_BEGIN




#ifdef __cplusplus


/*
 * Creates a new DistribBernoulliDistribution using the given SBML Level,
 * Version and &ldquo;distrib&rdquo; package version.
 */
DistribBernoulliDistribution::DistribBernoulliDistribution(unsigned int level,
                                                           unsigned int
                                                             version,
                                                           unsigned int
                                                             pkgVersion)
  : DistribCategoricalUnivariateDistribution(level, version)
  , mProb (NULL)
{
  setSBMLNamespacesAndOwn(new DistribPkgNamespaces(level, version,
    pkgVersion));
  connectToChild();
}


/*
 * Creates a new DistribBernoulliDistribution using the given
 * DistribPkgNamespaces object.
 */
DistribBernoulliDistribution::DistribBernoulliDistribution(DistribPkgNamespaces
  *distribns)
  : DistribCategoricalUnivariateDistribution(distribns)
  , mProb (NULL)
{
  setElementNamespace(distribns->getURI());
  connectToChild();
  loadPlugins(distribns);
}


/*
 * Copy constructor for DistribBernoulliDistribution.
 */
DistribBernoulliDistribution::DistribBernoulliDistribution(const
  DistribBernoulliDistribution& orig)
  : DistribCategoricalUnivariateDistribution( orig )
  , mProb ( NULL )
{
  if (orig.mProb != NULL)
  {
    mProb = orig.mProb->clone();
  }

  connectToChild();
}


/*
 * Assignment operator for DistribBernoulliDistribution.
 */
DistribBernoulliDistribution&
DistribBernoulliDistribution::operator=(const DistribBernoulliDistribution&
  rhs)
{
  if (&rhs != this)
  {
    DistribCategoricalUnivariateDistribution::operator=(rhs);
    delete mProb;
    if (rhs.mProb != NULL)
    {
      mProb = rhs.mProb->clone();
    }
    else
    {
      mProb = NULL;
    }

    connectToChild();
  }

  return *this;
}


/*
 * Creates and returns a deep copy of this DistribBernoulliDistribution object.
 */
DistribBernoulliDistribution*
DistribBernoulliDistribution::clone() const
{
  return new DistribBernoulliDistribution(*this);
}


/*
 * Destructor for DistribBernoulliDistribution.
 */
DistribBernoulliDistribution::~DistribBernoulliDistribution()
{
  delete mProb;
  mProb = NULL;
}


/*
 * Returns the value of the "id" attribute of this
 * DistribBernoulliDistribution.
 */
const std::string&
DistribBernoulliDistribution::getId() const
{
  return mId;
}


/*
 * Returns the value of the "name" attribute of this
 * DistribBernoulliDistribution.
 */
const std::string&
DistribBernoulliDistribution::getName() const
{
  return mName;
}


/*
 * Predicate returning @c true if this DistribBernoulliDistribution's "id"
 * attribute is set.
 */
bool
DistribBernoulliDistribution::isSetId() const
{
  return (mId.empty() == false);
}


/*
 * Predicate returning @c true if this DistribBernoulliDistribution's "name"
 * attribute is set.
 */
bool
DistribBernoulliDistribution::isSetName() const
{
  return (mName.empty() == false);
}


/*
 * Sets the value of the "id" attribute of this DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setId(const std::string& id)
{
  return SyntaxChecker::checkAndSetSId(id, mId);
}


/*
 * Sets the value of the "name" attribute of this DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setName(const std::string& name)
{
  mName = name;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Unsets the value of the "id" attribute of this DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::unsetId()
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
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::unsetName()
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
 * Returns the value of the "prob" element of this
 * DistribBernoulliDistribution.
 */
const DistribUncertValue*
DistribBernoulliDistribution::getProb() const
{
  return mProb;
}


/*
 * Returns the value of the "prob" element of this
 * DistribBernoulliDistribution.
 */
DistribUncertValue*
DistribBernoulliDistribution::getProb()
{
  return mProb;
}


/*
 * Predicate returning @c true if this DistribBernoulliDistribution's "prob"
 * element is set.
 */
bool
DistribBernoulliDistribution::isSetProb() const
{
  return (mProb != NULL);
}


/*
 * Sets the value of the "prob" element of this DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setProb(const DistribUncertValue* prob)
{
  if (prob == NULL)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else if (prob->hasRequiredElements() == false)
  {
    return LIBSBML_INVALID_OBJECT;
  }
  else if (getLevel() != prob->getLevel())
  {
    return LIBSBML_LEVEL_MISMATCH;
  }
  else if (getVersion() != prob->getVersion())
  {
    return LIBSBML_VERSION_MISMATCH;
  }
  else if (getPackageVersion() != prob->getPackageVersion())
  {
    return LIBSBML_PKG_VERSION_MISMATCH;
  }
  else
  {
    delete mProb;
    mProb = (prob != NULL) ? static_cast<DistribUncertValue*>(prob->clone()) :
      NULL;
    if (mProb != NULL) mProb->connectToParent(this);
    return LIBSBML_OPERATION_SUCCESS;
  }
}


/*
 * Creates a new DistribUncertValue object, adds it to this
 * DistribBernoulliDistribution object and returns the DistribUncertValue
 * object created.
 */
DistribUncertValue*
DistribBernoulliDistribution::createProb()
{
  if (mProb != NULL)
  {
    delete mProb;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mProb = new DistribUncertValue(distribns);

  mProb->setElementName("prob");

  delete distribns;

  connectToChild();

  return mProb;
}


/*
 * Unsets the value of the "prob" element of this DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::unsetProb()
{
  delete mProb;
  mProb = NULL;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Returns the XML element name of this DistribBernoulliDistribution object.
 */
const std::string&
DistribBernoulliDistribution::getElementName() const
{
  static const string name = "bernoulliDistribution";
  return name;
}


/*
 * Returns the libSBML type code for this DistribBernoulliDistribution object.
 */
int
DistribBernoulliDistribution::getTypeCode() const
{
  return SBML_DISTRIB_BERNOULLIDISTRIBUTION;
}


/*
 * Predicate returning @c true if all the required attributes for this
 * DistribBernoulliDistribution object have been set.
 */
bool
DistribBernoulliDistribution::hasRequiredAttributes() const
{
  bool allPresent =
    DistribCategoricalUnivariateDistribution::hasRequiredAttributes();

  return allPresent;
}


/*
 * Predicate returning @c true if all the required elements for this
 * DistribBernoulliDistribution object have been set.
 */
bool
DistribBernoulliDistribution::hasRequiredElements() const
{
  bool allPresent =
    DistribCategoricalUnivariateDistribution::hasRequiredElements();

  if (isSetProb() == false)
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
DistribBernoulliDistribution::writeElements(XMLOutputStream& stream) const
{
  DistribCategoricalUnivariateDistribution::writeElements(stream);

  if (isSetProb() == true)
  {
    mProb->write(stream);
  }

  SBase::writeExtensionElements(stream);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Accepts the given SBMLVisitor
 */
bool
DistribBernoulliDistribution::accept(SBMLVisitor& v) const
{
  v.visit(*this);

  if (mProb != NULL)
  {
    mProb->accept(v);
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
DistribBernoulliDistribution::setSBMLDocument(SBMLDocument* d)
{
  DistribCategoricalUnivariateDistribution::setSBMLDocument(d);

  if (mProb != NULL)
  {
    mProb->setSBMLDocument(d);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Connects to child elements
 */
void
DistribBernoulliDistribution::connectToChild()
{
  DistribCategoricalUnivariateDistribution::connectToChild();

  if (mProb != NULL)
  {
    mProb->connectToParent(this);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Enables/disables the given package with this element
 */
void
DistribBernoulliDistribution::enablePackageInternal(const std::string& pkgURI,
                                                    const std::string&
                                                      pkgPrefix,
                                                    bool flag)
{
  DistribCategoricalUnivariateDistribution::enablePackageInternal(pkgURI,
    pkgPrefix, flag);

  if (isSetProb())
  {
    mProb->enablePackageInternal(pkgURI, pkgPrefix, flag);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Updates the namespaces when setLevelVersion is used
 */
void
DistribBernoulliDistribution::updateSBMLNamespace(const std::string& package,
                                                  unsigned int level,
                                                  unsigned int version)
{
  DistribCategoricalUnivariateDistribution::updateSBMLNamespace(package, level,
    version);

  if (mProb != NULL)
  {
    mProb->updateSBMLNamespace(package, level, version);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::getAttribute(const std::string& attributeName,
                                           bool& value) const
{
  int return_value =
    DistribCategoricalUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::getAttribute(const std::string& attributeName,
                                           int& value) const
{
  int return_value =
    DistribCategoricalUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::getAttribute(const std::string& attributeName,
                                           double& value) const
{
  int return_value =
    DistribCategoricalUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::getAttribute(const std::string& attributeName,
                                           unsigned int& value) const
{
  int return_value =
    DistribCategoricalUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::getAttribute(const std::string& attributeName,
                                           std::string& value) const
{
  int return_value =
    DistribCategoricalUnivariateDistribution::getAttribute(attributeName, value);

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
 * Predicate returning @c true if this DistribBernoulliDistribution's attribute
 * "attributeName" is set.
 */
bool
DistribBernoulliDistribution::isSetAttribute(const std::string& attributeName)
  const
{
  bool value =
    DistribCategoricalUnivariateDistribution::isSetAttribute(attributeName);

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
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setAttribute(const std::string& attributeName,
                                           bool value)
{
  int return_value =
    DistribCategoricalUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setAttribute(const std::string& attributeName,
                                           int value)
{
  int return_value =
    DistribCategoricalUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setAttribute(const std::string& attributeName,
                                           double value)
{
  int return_value =
    DistribCategoricalUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setAttribute(const std::string& attributeName,
                                           unsigned int value)
{
  int return_value =
    DistribCategoricalUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::setAttribute(const std::string& attributeName,
                                           const std::string& value)
{
  int return_value =
    DistribCategoricalUnivariateDistribution::setAttribute(attributeName, value);

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
 * DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::unsetAttribute(const std::string& attributeName)
{
  int value =
    DistribCategoricalUnivariateDistribution::unsetAttribute(attributeName);

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
 * DistribBernoulliDistribution.
 */
SBase*
DistribBernoulliDistribution::createChildObject(const std::string& elementName)
{
  DistribCategoricalUnivariateDistribution* obj = NULL;

  if (elementName == "prob")
  {
    return createProb();
  }

  return obj;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Adds a new "elementName" object to this DistribBernoulliDistribution.
 */
int
DistribBernoulliDistribution::addChildObject(const std::string& elementName,
                                             const SBase* element)
{
  if (elementName == "prob" && element->getTypeCode() ==
    SBML_DISTRIB_UNCERTVALUE)
  {
    return setProb((const DistribUncertValue*)(element));
  }

  return LIBSBML_OPERATION_FAILED;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Removes and returns the new "elementName" object with the given id in this
 * DistribBernoulliDistribution.
 */
SBase*
DistribBernoulliDistribution::removeChildObject(const std::string& elementName,
                                                const std::string& id)
{
  if (elementName == "prob")
  {
    DistribUncertValue * obj = getProb();
    if (unsetProb() == LIBSBML_OPERATION_SUCCESS) return obj;
  }

  return NULL;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the number of "elementName" in this DistribBernoulliDistribution.
 */
unsigned int
DistribBernoulliDistribution::getNumObjects(const std::string& elementName)
{
  unsigned int n = 0;

  if (elementName == "prob")
  {
    if (isSetProb())
    {
      return 1;
    }
  }

  return n;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the nth object of "objectName" in this DistribBernoulliDistribution.
 */
SBase*
DistribBernoulliDistribution::getObject(const std::string& elementName,
                                        unsigned int index)
{
  SBase* obj = NULL;

  if (elementName == "prob")
  {
    return getProb();
  }

  return obj;
}

/** @endcond */


/*
 * Returns the first child element that has the given @p id in the model-wide
 * SId namespace, or @c NULL if no such object is found.
 */
SBase*
DistribBernoulliDistribution::getElementBySId(const std::string& id)
{
  if (id.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mProb != NULL)
  {
    if (mProb->getId() == id)
    {
      return mProb;
    }

    obj = mProb->getElementBySId(id);
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
DistribBernoulliDistribution::getElementByMetaId(const std::string& metaid)
{
  if (metaid.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mProb != NULL)
  {
    if (mProb->getMetaId() == metaid)
    {
      return mProb;
    }

    obj = mProb->getElementByMetaId(metaid);
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
DistribBernoulliDistribution::getAllElements(ElementFilter* filter)
{
  List* ret = new List();
  List* sublist = NULL;

  ADD_FILTERED_POINTER(ret, sublist, mProb, filter);


  ADD_FILTERED_FROM_PLUGIN(ret, sublist, filter);

  return ret;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Creates a new object from the next XMLToken on the XMLInputStream
 */
SBase*
DistribBernoulliDistribution::createObject(XMLInputStream& stream)
{
  SBase* obj = DistribCategoricalUnivariateDistribution::createObject(stream);

  const std::string& name = stream.peek().getName();

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());

  if (name == "prob")
  {
    if (isSetProb())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribBernoulliDistributionAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mProb;
    mProb = new DistribUncertValue(distribns);
    mProb->setElementName(name);
    obj = mProb;
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
DistribBernoulliDistribution::addExpectedAttributes(ExpectedAttributes&
  attributes)
{
  DistribCategoricalUnivariateDistribution::addExpectedAttributes(attributes);

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
DistribBernoulliDistribution::readAttributes(const XMLAttributes& attributes,
                                             const ExpectedAttributes&
                                               expectedAttributes)
{
  unsigned int level = getLevel();
  unsigned int version = getVersion();
  unsigned int pkgVersion = getPackageVersion();
  unsigned int numErrs;
  bool assigned = false;
  SBMLErrorLog* log = getErrorLog();

  DistribCategoricalUnivariateDistribution::readAttributes(attributes,
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
          DistribDistribBernoulliDistributionAllowedAttributes, pkgVersion,
            level, version, details);
      }
      else if (log->getError(n)->getErrorId() == UnknownCoreAttribute)
      {
        const std::string details = log->getError(n)->getMessage();
        log->remove(UnknownCoreAttribute);
        log->logPackageError("distrib",
          DistribDistribBernoulliDistributionAllowedCoreAttributes, pkgVersion,
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
DistribBernoulliDistribution::readL3V1V1Attributes(const XMLAttributes&
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
      logEmptyString(mId, level, version, "<DistribBernoulliDistribution>");
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
      logEmptyString(mName, level, version, "<DistribBernoulliDistribution>");
    }
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Reads the expected attributes into the member data variables
 */
void
DistribBernoulliDistribution::readL3V2V1Attributes(const XMLAttributes&
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
      logEmptyString(mId, level, version, "<DistribBernoulliDistribution>");
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
DistribBernoulliDistribution::writeAttributes(XMLOutputStream& stream) const
{
  DistribCategoricalUnivariateDistribution::writeAttributes(stream);

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
DistribBernoulliDistribution::writeL3V1V1Attributes(XMLOutputStream& stream)
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
DistribBernoulliDistribution::writeL3V2V1Attributes(XMLOutputStream& stream)
  const
{
}

/** @endcond */




#endif /* __cplusplus */


/*
 * Creates a new DistribBernoulliDistribution_t using the given SBML Level,
 * Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribBernoulliDistribution_t *
DistribBernoulliDistribution_create(unsigned int level,
                                    unsigned int version,
                                    unsigned int pkgVersion)
{
  return new DistribBernoulliDistribution(level, version, pkgVersion);
}


/*
 * Creates and returns a deep copy of this DistribBernoulliDistribution_t
 * object.
 */
LIBSBML_EXTERN
DistribBernoulliDistribution_t*
DistribBernoulliDistribution_clone(const DistribBernoulliDistribution_t* dbd)
{
  if (dbd != NULL)
  {
    return static_cast<DistribBernoulliDistribution_t*>(dbd->clone());
  }
  else
  {
    return NULL;
  }
}


/*
 * Frees this DistribBernoulliDistribution_t object.
 */
LIBSBML_EXTERN
void
DistribBernoulliDistribution_free(DistribBernoulliDistribution_t* dbd)
{
  if (dbd != NULL)
  {
    delete dbd;
  }
}


/*
 * Returns the value of the "id" attribute of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
char *
DistribBernoulliDistribution_getId(const DistribBernoulliDistribution_t * dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return dbd->getId().empty() ? NULL : safe_strdup(dbd->getId().c_str());
}


/*
 * Returns the value of the "name" attribute of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
char *
DistribBernoulliDistribution_getName(const DistribBernoulliDistribution_t *
  dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return dbd->getName().empty() ? NULL : safe_strdup(dbd->getName().c_str());
}


/*
 * Predicate returning @c 1 (true) if this DistribBernoulliDistribution_t's
 * "id" attribute is set.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_isSetId(const DistribBernoulliDistribution_t *
  dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetId()) : 0;
}


/*
 * Predicate returning @c 1 (true) if this DistribBernoulliDistribution_t's
 * "name" attribute is set.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_isSetName(const DistribBernoulliDistribution_t *
  dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetName()) : 0;
}


/*
 * Sets the value of the "id" attribute of this DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_setId(DistribBernoulliDistribution_t * dbd,
                                   const char * id)
{
  return (dbd != NULL) ? dbd->setId(id) : LIBSBML_INVALID_OBJECT;
}


/*
 * Sets the value of the "name" attribute of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_setName(DistribBernoulliDistribution_t * dbd,
                                     const char * name)
{
  return (dbd != NULL) ? dbd->setName(name) : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "id" attribute of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_unsetId(DistribBernoulliDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetId() : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "name" attribute of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_unsetName(DistribBernoulliDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetName() : LIBSBML_INVALID_OBJECT;
}


/*
 * Returns the value of the "prob" element of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
const DistribUncertValue_t*
DistribBernoulliDistribution_getProb(const DistribBernoulliDistribution_t *
  dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dbd->getProb());
}


/*
 * Predicate returning @c 1 (true) if this DistribBernoulliDistribution_t's
 * "prob" element is set.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_isSetProb(const DistribBernoulliDistribution_t *
  dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetProb()) : 0;
}


/*
 * Sets the value of the "prob" element of this DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_setProb(DistribBernoulliDistribution_t * dbd,
                                     const DistribUncertValue_t* prob)
{
  return (dbd != NULL) ? dbd->setProb(prob) : LIBSBML_INVALID_OBJECT;
}


/*
 * Creates a new DistribUncertValue_t object, adds it to this
 * DistribBernoulliDistribution_t object and returns the DistribUncertValue_t
 * object created.
 */
LIBSBML_EXTERN
DistribUncertValue_t*
DistribBernoulliDistribution_createProb(DistribBernoulliDistribution_t* dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dbd->createProb());
}


/*
 * Unsets the value of the "prob" element of this
 * DistribBernoulliDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_unsetProb(DistribBernoulliDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetProb() : LIBSBML_INVALID_OBJECT;
}


/*
 * Predicate returning @c 1 (true) if all the required attributes for this
 * DistribBernoulliDistribution_t object have been set.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_hasRequiredAttributes(const
  DistribBernoulliDistribution_t * dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->hasRequiredAttributes()) : 0;
}


/*
 * Predicate returning @c 1 (true) if all the required elements for this
 * DistribBernoulliDistribution_t object have been set.
 */
LIBSBML_EXTERN
int
DistribBernoulliDistribution_hasRequiredElements(const
  DistribBernoulliDistribution_t * dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->hasRequiredElements()) : 0;
}




LIBSBML_CPP_NAMESPACE_END



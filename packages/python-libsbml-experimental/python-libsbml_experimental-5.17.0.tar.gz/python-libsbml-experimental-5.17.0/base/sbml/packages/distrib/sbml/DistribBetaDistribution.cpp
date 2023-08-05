/**
 * @file DistribBetaDistribution.cpp
 * @brief Implementation of the DistribBetaDistribution class.
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
#include <sbml/packages/distrib/sbml/DistribBetaDistribution.h>
#include <sbml/packages/distrib/validator/DistribSBMLError.h>
#include <sbml/util/ElementFilter.h>


using namespace std;



LIBSBML_CPP_NAMESPACE_BEGIN




#ifdef __cplusplus


/*
 * Creates a new DistribBetaDistribution using the given SBML Level, Version
 * and &ldquo;distrib&rdquo; package version.
 */
DistribBetaDistribution::DistribBetaDistribution(unsigned int level,
                                                 unsigned int version,
                                                 unsigned int pkgVersion)
  : DistribContinuousUnivariateDistribution(level, version)
  , mAlpha (NULL)
  , mBeta (NULL)
{
  setSBMLNamespacesAndOwn(new DistribPkgNamespaces(level, version,
    pkgVersion));
  connectToChild();
}


/*
 * Creates a new DistribBetaDistribution using the given DistribPkgNamespaces
 * object.
 */
DistribBetaDistribution::DistribBetaDistribution(DistribPkgNamespaces
  *distribns)
  : DistribContinuousUnivariateDistribution(distribns)
  , mAlpha (NULL)
  , mBeta (NULL)
{
  setElementNamespace(distribns->getURI());
  connectToChild();
  loadPlugins(distribns);
}


/*
 * Copy constructor for DistribBetaDistribution.
 */
DistribBetaDistribution::DistribBetaDistribution(const DistribBetaDistribution&
  orig)
  : DistribContinuousUnivariateDistribution( orig )
  , mAlpha ( NULL )
  , mBeta ( NULL )
{
  if (orig.mAlpha != NULL)
  {
    mAlpha = orig.mAlpha->clone();
  }

  if (orig.mBeta != NULL)
  {
    mBeta = orig.mBeta->clone();
  }

  connectToChild();
}


/*
 * Assignment operator for DistribBetaDistribution.
 */
DistribBetaDistribution&
DistribBetaDistribution::operator=(const DistribBetaDistribution& rhs)
{
  if (&rhs != this)
  {
    DistribContinuousUnivariateDistribution::operator=(rhs);
    delete mAlpha;
    if (rhs.mAlpha != NULL)
    {
      mAlpha = rhs.mAlpha->clone();
    }
    else
    {
      mAlpha = NULL;
    }

    delete mBeta;
    if (rhs.mBeta != NULL)
    {
      mBeta = rhs.mBeta->clone();
    }
    else
    {
      mBeta = NULL;
    }

    connectToChild();
  }

  return *this;
}


/*
 * Creates and returns a deep copy of this DistribBetaDistribution object.
 */
DistribBetaDistribution*
DistribBetaDistribution::clone() const
{
  return new DistribBetaDistribution(*this);
}


/*
 * Destructor for DistribBetaDistribution.
 */
DistribBetaDistribution::~DistribBetaDistribution()
{
  delete mAlpha;
  mAlpha = NULL;
  delete mBeta;
  mBeta = NULL;
}


/*
 * Returns the value of the "id" attribute of this DistribBetaDistribution.
 */
const std::string&
DistribBetaDistribution::getId() const
{
  return mId;
}


/*
 * Returns the value of the "name" attribute of this DistribBetaDistribution.
 */
const std::string&
DistribBetaDistribution::getName() const
{
  return mName;
}


/*
 * Predicate returning @c true if this DistribBetaDistribution's "id" attribute
 * is set.
 */
bool
DistribBetaDistribution::isSetId() const
{
  return (mId.empty() == false);
}


/*
 * Predicate returning @c true if this DistribBetaDistribution's "name"
 * attribute is set.
 */
bool
DistribBetaDistribution::isSetName() const
{
  return (mName.empty() == false);
}


/*
 * Sets the value of the "id" attribute of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::setId(const std::string& id)
{
  return SyntaxChecker::checkAndSetSId(id, mId);
}


/*
 * Sets the value of the "name" attribute of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::setName(const std::string& name)
{
  mName = name;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Unsets the value of the "id" attribute of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::unsetId()
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
 * Unsets the value of the "name" attribute of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::unsetName()
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
 * Returns the value of the "alpha" element of this DistribBetaDistribution.
 */
const DistribUncertValue*
DistribBetaDistribution::getAlpha() const
{
  return mAlpha;
}


/*
 * Returns the value of the "alpha" element of this DistribBetaDistribution.
 */
DistribUncertValue*
DistribBetaDistribution::getAlpha()
{
  return mAlpha;
}


/*
 * Returns the value of the "beta" element of this DistribBetaDistribution.
 */
const DistribUncertValue*
DistribBetaDistribution::getBeta() const
{
  return mBeta;
}


/*
 * Returns the value of the "beta" element of this DistribBetaDistribution.
 */
DistribUncertValue*
DistribBetaDistribution::getBeta()
{
  return mBeta;
}


/*
 * Predicate returning @c true if this DistribBetaDistribution's "alpha"
 * element is set.
 */
bool
DistribBetaDistribution::isSetAlpha() const
{
  return (mAlpha != NULL);
}


/*
 * Predicate returning @c true if this DistribBetaDistribution's "beta" element
 * is set.
 */
bool
DistribBetaDistribution::isSetBeta() const
{
  return (mBeta != NULL);
}


/*
 * Sets the value of the "alpha" element of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::setAlpha(const DistribUncertValue* alpha)
{
  if (alpha == NULL)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else if (alpha->hasRequiredElements() == false)
  {
    return LIBSBML_INVALID_OBJECT;
  }
  else if (getLevel() != alpha->getLevel())
  {
    return LIBSBML_LEVEL_MISMATCH;
  }
  else if (getVersion() != alpha->getVersion())
  {
    return LIBSBML_VERSION_MISMATCH;
  }
  else if (getPackageVersion() != alpha->getPackageVersion())
  {
    return LIBSBML_PKG_VERSION_MISMATCH;
  }
  else
  {
    delete mAlpha;
    mAlpha = (alpha != NULL) ? static_cast<DistribUncertValue*>(alpha->clone())
      : NULL;
    if (mAlpha != NULL) mAlpha->connectToParent(this);
    return LIBSBML_OPERATION_SUCCESS;
  }
}


/*
 * Sets the value of the "beta" element of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::setBeta(const DistribUncertValue* beta)
{
  if (beta == NULL)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else if (beta->hasRequiredElements() == false)
  {
    return LIBSBML_INVALID_OBJECT;
  }
  else if (getLevel() != beta->getLevel())
  {
    return LIBSBML_LEVEL_MISMATCH;
  }
  else if (getVersion() != beta->getVersion())
  {
    return LIBSBML_VERSION_MISMATCH;
  }
  else if (getPackageVersion() != beta->getPackageVersion())
  {
    return LIBSBML_PKG_VERSION_MISMATCH;
  }
  else
  {
    delete mBeta;
    mBeta = (beta != NULL) ? static_cast<DistribUncertValue*>(beta->clone()) :
      NULL;
    if (mBeta != NULL) mBeta->connectToParent(this);
    return LIBSBML_OPERATION_SUCCESS;
  }
}


/*
 * Creates a new DistribUncertValue object, adds it to this
 * DistribBetaDistribution object and returns the DistribUncertValue object
 * created.
 */
DistribUncertValue*
DistribBetaDistribution::createAlpha()
{
  if (mAlpha != NULL)
  {
    delete mAlpha;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mAlpha = new DistribUncertValue(distribns);

  mAlpha->setElementName("alpha");

  delete distribns;

  connectToChild();

  return mAlpha;
}


/*
 * Creates a new DistribUncertValue object, adds it to this
 * DistribBetaDistribution object and returns the DistribUncertValue object
 * created.
 */
DistribUncertValue*
DistribBetaDistribution::createBeta()
{
  if (mBeta != NULL)
  {
    delete mBeta;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mBeta = new DistribUncertValue(distribns);

  mBeta->setElementName("beta");

  delete distribns;

  connectToChild();

  return mBeta;
}


/*
 * Unsets the value of the "alpha" element of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::unsetAlpha()
{
  delete mAlpha;
  mAlpha = NULL;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Unsets the value of the "beta" element of this DistribBetaDistribution.
 */
int
DistribBetaDistribution::unsetBeta()
{
  delete mBeta;
  mBeta = NULL;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Returns the XML element name of this DistribBetaDistribution object.
 */
const std::string&
DistribBetaDistribution::getElementName() const
{
  static const string name = "betaDistribution";
  return name;
}


/*
 * Returns the libSBML type code for this DistribBetaDistribution object.
 */
int
DistribBetaDistribution::getTypeCode() const
{
  return SBML_DISTRIB_BETADISTRIBUTION;
}


/*
 * Predicate returning @c true if all the required attributes for this
 * DistribBetaDistribution object have been set.
 */
bool
DistribBetaDistribution::hasRequiredAttributes() const
{
  bool allPresent =
    DistribContinuousUnivariateDistribution::hasRequiredAttributes();

  return allPresent;
}


/*
 * Predicate returning @c true if all the required elements for this
 * DistribBetaDistribution object have been set.
 */
bool
DistribBetaDistribution::hasRequiredElements() const
{
  bool allPresent =
    DistribContinuousUnivariateDistribution::hasRequiredElements();

  if (isSetAlpha() == false)
  {
    allPresent = false;
  }

  if (isSetBeta() == false)
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
DistribBetaDistribution::writeElements(XMLOutputStream& stream) const
{
  DistribContinuousUnivariateDistribution::writeElements(stream);

  if (isSetAlpha() == true)
  {
    mAlpha->write(stream);
  }

  if (isSetBeta() == true)
  {
    mBeta->write(stream);
  }

  SBase::writeExtensionElements(stream);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Accepts the given SBMLVisitor
 */
bool
DistribBetaDistribution::accept(SBMLVisitor& v) const
{
  v.visit(*this);

  if (mAlpha != NULL)
  {
    mAlpha->accept(v);
  }

  if (mBeta != NULL)
  {
    mBeta->accept(v);
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
DistribBetaDistribution::setSBMLDocument(SBMLDocument* d)
{
  DistribContinuousUnivariateDistribution::setSBMLDocument(d);

  if (mAlpha != NULL)
  {
    mAlpha->setSBMLDocument(d);
  }

  if (mBeta != NULL)
  {
    mBeta->setSBMLDocument(d);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Connects to child elements
 */
void
DistribBetaDistribution::connectToChild()
{
  DistribContinuousUnivariateDistribution::connectToChild();

  if (mAlpha != NULL)
  {
    mAlpha->connectToParent(this);
  }

  if (mBeta != NULL)
  {
    mBeta->connectToParent(this);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Enables/disables the given package with this element
 */
void
DistribBetaDistribution::enablePackageInternal(const std::string& pkgURI,
                                               const std::string& pkgPrefix,
                                               bool flag)
{
  DistribContinuousUnivariateDistribution::enablePackageInternal(pkgURI,
    pkgPrefix, flag);

  if (isSetAlpha())
  {
    mAlpha->enablePackageInternal(pkgURI, pkgPrefix, flag);
  }

  if (isSetBeta())
  {
    mBeta->enablePackageInternal(pkgURI, pkgPrefix, flag);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Updates the namespaces when setLevelVersion is used
 */
void
DistribBetaDistribution::updateSBMLNamespace(const std::string& package,
                                             unsigned int level,
                                             unsigned int version)
{
  DistribContinuousUnivariateDistribution::updateSBMLNamespace(package, level,
    version);

  if (mAlpha != NULL)
  {
    mAlpha->updateSBMLNamespace(package, level, version);
  }

  if (mBeta != NULL)
  {
    mBeta->updateSBMLNamespace(package, level, version);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::getAttribute(const std::string& attributeName,
                                      bool& value) const
{
  int return_value =
    DistribContinuousUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::getAttribute(const std::string& attributeName,
                                      int& value) const
{
  int return_value =
    DistribContinuousUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::getAttribute(const std::string& attributeName,
                                      double& value) const
{
  int return_value =
    DistribContinuousUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::getAttribute(const std::string& attributeName,
                                      unsigned int& value) const
{
  int return_value =
    DistribContinuousUnivariateDistribution::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::getAttribute(const std::string& attributeName,
                                      std::string& value) const
{
  int return_value =
    DistribContinuousUnivariateDistribution::getAttribute(attributeName, value);

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
 * Predicate returning @c true if this DistribBetaDistribution's attribute
 * "attributeName" is set.
 */
bool
DistribBetaDistribution::isSetAttribute(const std::string& attributeName) const
{
  bool value =
    DistribContinuousUnivariateDistribution::isSetAttribute(attributeName);

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
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::setAttribute(const std::string& attributeName,
                                      bool value)
{
  int return_value =
    DistribContinuousUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::setAttribute(const std::string& attributeName,
                                      int value)
{
  int return_value =
    DistribContinuousUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::setAttribute(const std::string& attributeName,
                                      double value)
{
  int return_value =
    DistribContinuousUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::setAttribute(const std::string& attributeName,
                                      unsigned int value)
{
  int return_value =
    DistribContinuousUnivariateDistribution::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::setAttribute(const std::string& attributeName,
                                      const std::string& value)
{
  int return_value =
    DistribContinuousUnivariateDistribution::setAttribute(attributeName, value);

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
 * DistribBetaDistribution.
 */
int
DistribBetaDistribution::unsetAttribute(const std::string& attributeName)
{
  int value =
    DistribContinuousUnivariateDistribution::unsetAttribute(attributeName);

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
 * DistribBetaDistribution.
 */
SBase*
DistribBetaDistribution::createChildObject(const std::string& elementName)
{
  DistribContinuousUnivariateDistribution* obj = NULL;

  if (elementName == "alpha")
  {
    return createAlpha();
  }
  else if (elementName == "beta")
  {
    return createBeta();
  }

  return obj;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Adds a new "elementName" object to this DistribBetaDistribution.
 */
int
DistribBetaDistribution::addChildObject(const std::string& elementName,
                                        const SBase* element)
{
  if (elementName == "alpha" && element->getTypeCode() ==
    SBML_DISTRIB_UNCERTVALUE)
  {
    return setAlpha((const DistribUncertValue*)(element));
  }
  else if (elementName == "beta" && element->getTypeCode() ==
    SBML_DISTRIB_UNCERTVALUE)
  {
    return setBeta((const DistribUncertValue*)(element));
  }

  return LIBSBML_OPERATION_FAILED;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Removes and returns the new "elementName" object with the given id in this
 * DistribBetaDistribution.
 */
SBase*
DistribBetaDistribution::removeChildObject(const std::string& elementName,
                                           const std::string& id)
{
  if (elementName == "alpha")
  {
    DistribUncertValue * obj = getAlpha();
    if (unsetAlpha() == LIBSBML_OPERATION_SUCCESS) return obj;
  }
  else if (elementName == "beta")
  {
    DistribUncertValue * obj = getBeta();
    if (unsetBeta() == LIBSBML_OPERATION_SUCCESS) return obj;
  }

  return NULL;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the number of "elementName" in this DistribBetaDistribution.
 */
unsigned int
DistribBetaDistribution::getNumObjects(const std::string& elementName)
{
  unsigned int n = 0;

  if (elementName == "alpha")
  {
    if (isSetAlpha())
    {
      return 1;
    }
  }
  else if (elementName == "beta")
  {
    if (isSetBeta())
    {
      return 1;
    }
  }

  return n;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the nth object of "objectName" in this DistribBetaDistribution.
 */
SBase*
DistribBetaDistribution::getObject(const std::string& elementName,
                                   unsigned int index)
{
  SBase* obj = NULL;

  if (elementName == "alpha")
  {
    return getAlpha();
  }
  else if (elementName == "beta")
  {
    return getBeta();
  }

  return obj;
}

/** @endcond */


/*
 * Returns the first child element that has the given @p id in the model-wide
 * SId namespace, or @c NULL if no such object is found.
 */
SBase*
DistribBetaDistribution::getElementBySId(const std::string& id)
{
  if (id.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mAlpha != NULL)
  {
    if (mAlpha->getId() == id)
    {
      return mAlpha;
    }

    obj = mAlpha->getElementBySId(id);
    if (obj != NULL)
    {
      return obj;
    }
  }

  if (mBeta != NULL)
  {
    if (mBeta->getId() == id)
    {
      return mBeta;
    }

    obj = mBeta->getElementBySId(id);
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
DistribBetaDistribution::getElementByMetaId(const std::string& metaid)
{
  if (metaid.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mAlpha != NULL)
  {
    if (mAlpha->getMetaId() == metaid)
    {
      return mAlpha;
    }

    obj = mAlpha->getElementByMetaId(metaid);
    if (obj != NULL)
    {
      return obj;
    }
  }

  if (mBeta != NULL)
  {
    if (mBeta->getMetaId() == metaid)
    {
      return mBeta;
    }

    obj = mBeta->getElementByMetaId(metaid);
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
DistribBetaDistribution::getAllElements(ElementFilter* filter)
{
  List* ret = new List();
  List* sublist = NULL;

  ADD_FILTERED_POINTER(ret, sublist, mAlpha, filter);
  ADD_FILTERED_POINTER(ret, sublist, mBeta, filter);


  ADD_FILTERED_FROM_PLUGIN(ret, sublist, filter);

  return ret;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Creates a new object from the next XMLToken on the XMLInputStream
 */
SBase*
DistribBetaDistribution::createObject(XMLInputStream& stream)
{
  SBase* obj = DistribContinuousUnivariateDistribution::createObject(stream);

  const std::string& name = stream.peek().getName();

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());

  if (name == "alpha")
  {
    if (isSetAlpha())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribBetaDistributionAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mAlpha;
    mAlpha = new DistribUncertValue(distribns);
    mAlpha->setElementName(name);
    obj = mAlpha;
  }
  else if (name == "beta")
  {
    if (isSetBeta())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribBetaDistributionAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mBeta;
    mBeta = new DistribUncertValue(distribns);
    mBeta->setElementName(name);
    obj = mBeta;
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
DistribBetaDistribution::addExpectedAttributes(ExpectedAttributes& attributes)
{
  DistribContinuousUnivariateDistribution::addExpectedAttributes(attributes);

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
DistribBetaDistribution::readAttributes(const XMLAttributes& attributes,
                                        const ExpectedAttributes&
                                          expectedAttributes)
{
  unsigned int level = getLevel();
  unsigned int version = getVersion();
  unsigned int pkgVersion = getPackageVersion();
  unsigned int numErrs;
  bool assigned = false;
  SBMLErrorLog* log = getErrorLog();

  DistribContinuousUnivariateDistribution::readAttributes(attributes,
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
          DistribDistribBetaDistributionAllowedAttributes, pkgVersion, level,
            version, details);
      }
      else if (log->getError(n)->getErrorId() == UnknownCoreAttribute)
      {
        const std::string details = log->getError(n)->getMessage();
        log->remove(UnknownCoreAttribute);
        log->logPackageError("distrib",
          DistribDistribBetaDistributionAllowedCoreAttributes, pkgVersion, level,
            version, details);
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
DistribBetaDistribution::readL3V1V1Attributes(const XMLAttributes& attributes)
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
      logEmptyString(mId, level, version, "<DistribBetaDistribution>");
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
      logEmptyString(mName, level, version, "<DistribBetaDistribution>");
    }
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Reads the expected attributes into the member data variables
 */
void
DistribBetaDistribution::readL3V2V1Attributes(const XMLAttributes& attributes)
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
      logEmptyString(mId, level, version, "<DistribBetaDistribution>");
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
DistribBetaDistribution::writeAttributes(XMLOutputStream& stream) const
{
  DistribContinuousUnivariateDistribution::writeAttributes(stream);

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
DistribBetaDistribution::writeL3V1V1Attributes(XMLOutputStream& stream) const
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
DistribBetaDistribution::writeL3V2V1Attributes(XMLOutputStream& stream) const
{
}

/** @endcond */




#endif /* __cplusplus */


/*
 * Creates a new DistribBetaDistribution_t using the given SBML Level, Version
 * and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribBetaDistribution_t *
DistribBetaDistribution_create(unsigned int level,
                               unsigned int version,
                               unsigned int pkgVersion)
{
  return new DistribBetaDistribution(level, version, pkgVersion);
}


/*
 * Creates and returns a deep copy of this DistribBetaDistribution_t object.
 */
LIBSBML_EXTERN
DistribBetaDistribution_t*
DistribBetaDistribution_clone(const DistribBetaDistribution_t* dbd)
{
  if (dbd != NULL)
  {
    return static_cast<DistribBetaDistribution_t*>(dbd->clone());
  }
  else
  {
    return NULL;
  }
}


/*
 * Frees this DistribBetaDistribution_t object.
 */
LIBSBML_EXTERN
void
DistribBetaDistribution_free(DistribBetaDistribution_t* dbd)
{
  if (dbd != NULL)
  {
    delete dbd;
  }
}


/*
 * Returns the value of the "id" attribute of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
char *
DistribBetaDistribution_getId(const DistribBetaDistribution_t * dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return dbd->getId().empty() ? NULL : safe_strdup(dbd->getId().c_str());
}


/*
 * Returns the value of the "name" attribute of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
char *
DistribBetaDistribution_getName(const DistribBetaDistribution_t * dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return dbd->getName().empty() ? NULL : safe_strdup(dbd->getName().c_str());
}


/*
 * Predicate returning @c 1 (true) if this DistribBetaDistribution_t's "id"
 * attribute is set.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_isSetId(const DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetId()) : 0;
}


/*
 * Predicate returning @c 1 (true) if this DistribBetaDistribution_t's "name"
 * attribute is set.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_isSetName(const DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetName()) : 0;
}


/*
 * Sets the value of the "id" attribute of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_setId(DistribBetaDistribution_t * dbd,
                              const char * id)
{
  return (dbd != NULL) ? dbd->setId(id) : LIBSBML_INVALID_OBJECT;
}


/*
 * Sets the value of the "name" attribute of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_setName(DistribBetaDistribution_t * dbd,
                                const char * name)
{
  return (dbd != NULL) ? dbd->setName(name) : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "id" attribute of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_unsetId(DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetId() : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "name" attribute of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_unsetName(DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetName() : LIBSBML_INVALID_OBJECT;
}


/*
 * Returns the value of the "alpha" element of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
const DistribUncertValue_t*
DistribBetaDistribution_getAlpha(const DistribBetaDistribution_t * dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dbd->getAlpha());
}


/*
 * Returns the value of the "beta" element of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
const DistribUncertValue_t*
DistribBetaDistribution_getBeta(const DistribBetaDistribution_t * dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dbd->getBeta());
}


/*
 * Predicate returning @c 1 (true) if this DistribBetaDistribution_t's "alpha"
 * element is set.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_isSetAlpha(const DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetAlpha()) : 0;
}


/*
 * Predicate returning @c 1 (true) if this DistribBetaDistribution_t's "beta"
 * element is set.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_isSetBeta(const DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->isSetBeta()) : 0;
}


/*
 * Sets the value of the "alpha" element of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_setAlpha(DistribBetaDistribution_t * dbd,
                                 const DistribUncertValue_t* alpha)
{
  return (dbd != NULL) ? dbd->setAlpha(alpha) : LIBSBML_INVALID_OBJECT;
}


/*
 * Sets the value of the "beta" element of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_setBeta(DistribBetaDistribution_t * dbd,
                                const DistribUncertValue_t* beta)
{
  return (dbd != NULL) ? dbd->setBeta(beta) : LIBSBML_INVALID_OBJECT;
}


/*
 * Creates a new DistribUncertValue_t object, adds it to this
 * DistribBetaDistribution_t object and returns the DistribUncertValue_t object
 * created.
 */
LIBSBML_EXTERN
DistribUncertValue_t*
DistribBetaDistribution_createAlpha(DistribBetaDistribution_t* dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dbd->createAlpha());
}


/*
 * Creates a new DistribUncertValue_t object, adds it to this
 * DistribBetaDistribution_t object and returns the DistribUncertValue_t object
 * created.
 */
LIBSBML_EXTERN
DistribUncertValue_t*
DistribBetaDistribution_createBeta(DistribBetaDistribution_t* dbd)
{
  if (dbd == NULL)
  {
    return NULL;
  }

  return (DistribUncertValue_t*)(dbd->createBeta());
}


/*
 * Unsets the value of the "alpha" element of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_unsetAlpha(DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetAlpha() : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "beta" element of this DistribBetaDistribution_t.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_unsetBeta(DistribBetaDistribution_t * dbd)
{
  return (dbd != NULL) ? dbd->unsetBeta() : LIBSBML_INVALID_OBJECT;
}


/*
 * Predicate returning @c 1 (true) if all the required attributes for this
 * DistribBetaDistribution_t object have been set.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_hasRequiredAttributes(const DistribBetaDistribution_t *
  dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->hasRequiredAttributes()) : 0;
}


/*
 * Predicate returning @c 1 (true) if all the required elements for this
 * DistribBetaDistribution_t object have been set.
 */
LIBSBML_EXTERN
int
DistribBetaDistribution_hasRequiredElements(const DistribBetaDistribution_t *
  dbd)
{
  return (dbd != NULL) ? static_cast<int>(dbd->hasRequiredElements()) : 0;
}




LIBSBML_CPP_NAMESPACE_END



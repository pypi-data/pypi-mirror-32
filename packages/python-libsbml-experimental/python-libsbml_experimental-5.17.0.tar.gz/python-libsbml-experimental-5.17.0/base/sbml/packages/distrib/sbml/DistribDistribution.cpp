/**
 * @file DistribDistribution.cpp
 * @brief Implementation of the DistribDistribution class.
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
#include <sbml/packages/distrib/sbml/DistribDistribution.h>
#include <sbml/packages/distrib/validator/DistribSBMLError.h>

#include <sbml/packages/distrib/sbml/DistribBetaDistribution.h>
#include <sbml/packages/distrib/sbml/DistribCauchyDistribution.h>
#include <sbml/packages/distrib/sbml/DistribChiSquareDistribution.h>
#include <sbml/packages/distrib/sbml/DistribExponentialDistribution.h>
#include <sbml/packages/distrib/sbml/DistribFDistribution.h>
#include <sbml/packages/distrib/sbml/DistribGammaDistribution.h>
#include <sbml/packages/distrib/sbml/DistribInverseGammaDistribution.h>
#include <sbml/packages/distrib/sbml/DistribLaPlaceDistribution.h>
#include <sbml/packages/distrib/sbml/DistribLogNormalDistribution.h>
#include <sbml/packages/distrib/sbml/DistribLogisticDistribution.h>
#include <sbml/packages/distrib/sbml/DistribNormalDistribution.h>
#include <sbml/packages/distrib/sbml/DistribParetoDistribution.h>
#include <sbml/packages/distrib/sbml/DistribRayleighDistribution.h>
#include <sbml/packages/distrib/sbml/DistribStudentTDistribution.h>
#include <sbml/packages/distrib/sbml/DistribUniformDistribution.h>
#include <sbml/packages/distrib/sbml/DistribWeibullDistribution.h>
#include <sbml/packages/distrib/sbml/DistribBinomialDistribution.h>
#include <sbml/packages/distrib/sbml/DistribGeometricDistribution.h>
#include <sbml/packages/distrib/sbml/DistribHypergeometricDistribution.h>
#include <sbml/packages/distrib/sbml/DistribNegativeBinomialDistribution.h>
#include <sbml/packages/distrib/sbml/DistribPoissonDistribution.h>
#include <sbml/packages/distrib/sbml/DistribBernoulliDistribution.h>
#include <sbml/packages/distrib/sbml/DistribCategoricalDistribution.h>
#include <sbml/packages/distrib/sbml/DistribMultivariateDistribution.h>
#include <sbml/packages/distrib/sbml/DistribExternalDistribution.h>


using namespace std;



LIBSBML_CPP_NAMESPACE_BEGIN




#ifdef __cplusplus


/*
 * Creates a new DistribDistribution using the given SBML Level, Version and
 * &ldquo;distrib&rdquo; package version.
 */
DistribDistribution::DistribDistribution(unsigned int level,
                                         unsigned int version,
                                         unsigned int pkgVersion)
  : SBase(level, version)
  , mElementName("distribution")
{
  setSBMLNamespacesAndOwn(new DistribPkgNamespaces(level, version,
    pkgVersion));
}


/*
 * Creates a new DistribDistribution using the given DistribPkgNamespaces
 * object.
 */
DistribDistribution::DistribDistribution(DistribPkgNamespaces *distribns)
  : SBase(distribns)
  , mElementName("distribution")
{
  setElementNamespace(distribns->getURI());
  loadPlugins(distribns);
}


/*
 * Copy constructor for DistribDistribution.
 */
DistribDistribution::DistribDistribution(const DistribDistribution& orig)
  : SBase( orig )
  , mElementName ( orig.mElementName )
{
}


/*
 * Assignment operator for DistribDistribution.
 */
DistribDistribution&
DistribDistribution::operator=(const DistribDistribution& rhs)
{
  if (&rhs != this)
  {
    SBase::operator=(rhs);
    mElementName = rhs.mElementName;
  }

  return *this;
}


/*
 * Creates and returns a deep copy of this DistribDistribution object.
 */
DistribDistribution*
DistribDistribution::clone() const
{
  return new DistribDistribution(*this);
}


/*
 * Destructor for DistribDistribution.
 */
DistribDistribution::~DistribDistribution()
{
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribBetaDistribution
 */
bool
DistribDistribution::isDistribBetaDistribution() const
{
  return dynamic_cast<const DistribBetaDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribCauchyDistribution
 */
bool
DistribDistribution::isDistribCauchyDistribution() const
{
  return dynamic_cast<const DistribCauchyDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribChiSquareDistribution
 */
bool
DistribDistribution::isDistribChiSquareDistribution() const
{
  return dynamic_cast<const DistribChiSquareDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribExponentialDistribution
 */
bool
DistribDistribution::isDistribExponentialDistribution() const
{
  return dynamic_cast<const DistribExponentialDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribFDistribution
 */
bool
DistribDistribution::isDistribFDistribution() const
{
  return dynamic_cast<const DistribFDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribGammaDistribution
 */
bool
DistribDistribution::isDistribGammaDistribution() const
{
  return dynamic_cast<const DistribGammaDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribInverseGammaDistribution
 */
bool
DistribDistribution::isDistribInverseGammaDistribution() const
{
  return dynamic_cast<const DistribInverseGammaDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribLaPlaceDistribution
 */
bool
DistribDistribution::isDistribLaPlaceDistribution() const
{
  return dynamic_cast<const DistribLaPlaceDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribLogNormalDistribution
 */
bool
DistribDistribution::isDistribLogNormalDistribution() const
{
  return dynamic_cast<const DistribLogNormalDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribLogisticDistribution
 */
bool
DistribDistribution::isDistribLogisticDistribution() const
{
  return dynamic_cast<const DistribLogisticDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribNormalDistribution
 */
bool
DistribDistribution::isDistribNormalDistribution() const
{
  return dynamic_cast<const DistribNormalDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribParetoDistribution
 */
bool
DistribDistribution::isDistribParetoDistribution() const
{
  return dynamic_cast<const DistribParetoDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribRayleighDistribution
 */
bool
DistribDistribution::isDistribRayleighDistribution() const
{
  return dynamic_cast<const DistribRayleighDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribStudentTDistribution
 */
bool
DistribDistribution::isDistribStudentTDistribution() const
{
  return dynamic_cast<const DistribStudentTDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribUniformDistribution
 */
bool
DistribDistribution::isDistribUniformDistribution() const
{
  return dynamic_cast<const DistribUniformDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribWeibullDistribution
 */
bool
DistribDistribution::isDistribWeibullDistribution() const
{
  return dynamic_cast<const DistribWeibullDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribBinomialDistribution
 */
bool
DistribDistribution::isDistribBinomialDistribution() const
{
  return dynamic_cast<const DistribBinomialDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribGeometricDistribution
 */
bool
DistribDistribution::isDistribGeometricDistribution() const
{
  return dynamic_cast<const DistribGeometricDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribHypergeometricDistribution
 */
bool
DistribDistribution::isDistribHypergeometricDistribution() const
{
  return dynamic_cast<const DistribHypergeometricDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribNegativeBinomialDistribution
 */
bool
DistribDistribution::isDistribNegativeBinomialDistribution() const
{
  return dynamic_cast<const DistribNegativeBinomialDistribution*>(this) !=
    NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribPoissonDistribution
 */
bool
DistribDistribution::isDistribPoissonDistribution() const
{
  return dynamic_cast<const DistribPoissonDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribBernoulliDistribution
 */
bool
DistribDistribution::isDistribBernoulliDistribution() const
{
  return dynamic_cast<const DistribBernoulliDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribCategoricalDistribution
 */
bool
DistribDistribution::isDistribCategoricalDistribution() const
{
  return dynamic_cast<const DistribCategoricalDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribMultivariateDistribution
 */
bool
DistribDistribution::isDistribMultivariateDistribution() const
{
  return dynamic_cast<const DistribMultivariateDistribution*>(this) != NULL;
}


/*
 * Predicate returning @c true if this abstract "DistribDistribution" is of
 * type DistribExternalDistribution
 */
bool
DistribDistribution::isDistribExternalDistribution() const
{
  return dynamic_cast<const DistribExternalDistribution*>(this) != NULL;
}


/*
 * Returns the XML element name of this DistribDistribution object.
 */
const std::string&
DistribDistribution::getElementName() const
{
  return mElementName;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the XML name of this DistribDistribution object.
 */
void
DistribDistribution::setElementName(const std::string& name)
{
  mElementName = name;
}

/** @endcond */


/*
 * Returns the libSBML type code for this DistribDistribution object.
 */
int
DistribDistribution::getTypeCode() const
{
  return SBML_DISTRIB_DISTRIBUTION;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Write any contained elements
 */
void
DistribDistribution::writeElements(XMLOutputStream& stream) const
{
  SBase::writeElements(stream);

  SBase::writeExtensionElements(stream);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Accepts the given SBMLVisitor
 */
bool
DistribDistribution::accept(SBMLVisitor& v) const
{
  return v.visit(*this);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the parent SBMLDocument
 */
void
DistribDistribution::setSBMLDocument(SBMLDocument* d)
{
  SBase::setSBMLDocument(d);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Enables/disables the given package with this element
 */
void
DistribDistribution::enablePackageInternal(const std::string& pkgURI,
                                           const std::string& pkgPrefix,
                                           bool flag)
{
  SBase::enablePackageInternal(pkgURI, pkgPrefix, flag);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::getAttribute(const std::string& attributeName,
                                  bool& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::getAttribute(const std::string& attributeName,
                                  int& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::getAttribute(const std::string& attributeName,
                                  double& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::getAttribute(const std::string& attributeName,
                                  unsigned int& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::getAttribute(const std::string& attributeName,
                                  std::string& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Predicate returning @c true if this DistribDistribution's attribute
 * "attributeName" is set.
 */
bool
DistribDistribution::isSetAttribute(const std::string& attributeName) const
{
  bool value = SBase::isSetAttribute(attributeName);

  return value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::setAttribute(const std::string& attributeName,
                                  bool value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::setAttribute(const std::string& attributeName, int value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::setAttribute(const std::string& attributeName,
                                  double value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::setAttribute(const std::string& attributeName,
                                  unsigned int value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribDistribution.
 */
int
DistribDistribution::setAttribute(const std::string& attributeName,
                                  const std::string& value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Unsets the value of the "attributeName" attribute of this
 * DistribDistribution.
 */
int
DistribDistribution::unsetAttribute(const std::string& attributeName)
{
  int value = SBase::unsetAttribute(attributeName);

  return value;
}

/** @endcond */




#endif /* __cplusplus */


/*
 * Creates a new DistribBetaDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribBetaDistribution(unsigned int level,
                                                  unsigned int version,
                                                  unsigned int pkgVersion)
{
  return new DistribBetaDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribCauchyDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribCauchyDistribution(unsigned int level,
                                                    unsigned int version,
                                                    unsigned int pkgVersion)
{
  return new DistribCauchyDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribChiSquareDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribChiSquareDistribution(unsigned int level,
                                                       unsigned int version,
                                                       unsigned int pkgVersion)
{
  return new DistribChiSquareDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribExponentialDistribution (DistribDistribution_t) using
 * the given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribExponentialDistribution(unsigned int level,
                                                         unsigned int version,
                                                         unsigned int
                                                           pkgVersion)
{
  return new DistribExponentialDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribFDistribution (DistribDistribution_t) using the given
 * SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribFDistribution(unsigned int level,
                                               unsigned int version,
                                               unsigned int pkgVersion)
{
  return new DistribFDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribGammaDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribGammaDistribution(unsigned int level,
                                                   unsigned int version,
                                                   unsigned int pkgVersion)
{
  return new DistribGammaDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribInverseGammaDistribution (DistribDistribution_t) using
 * the given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribInverseGammaDistribution(unsigned int level,
                                                          unsigned int version,
                                                          unsigned int
                                                            pkgVersion)
{
  return new DistribInverseGammaDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribLaPlaceDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribLaPlaceDistribution(unsigned int level,
                                                     unsigned int version,
                                                     unsigned int pkgVersion)
{
  return new DistribLaPlaceDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribLogNormalDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribLogNormalDistribution(unsigned int level,
                                                       unsigned int version,
                                                       unsigned int pkgVersion)
{
  return new DistribLogNormalDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribLogisticDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribLogisticDistribution(unsigned int level,
                                                      unsigned int version,
                                                      unsigned int pkgVersion)
{
  return new DistribLogisticDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribNormalDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribNormalDistribution(unsigned int level,
                                                    unsigned int version,
                                                    unsigned int pkgVersion)
{
  return new DistribNormalDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribParetoDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribParetoDistribution(unsigned int level,
                                                    unsigned int version,
                                                    unsigned int pkgVersion)
{
  return new DistribParetoDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribRayleighDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribRayleighDistribution(unsigned int level,
                                                      unsigned int version,
                                                      unsigned int pkgVersion)
{
  return new DistribRayleighDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribStudentTDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribStudentTDistribution(unsigned int level,
                                                      unsigned int version,
                                                      unsigned int pkgVersion)
{
  return new DistribStudentTDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribUniformDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribUniformDistribution(unsigned int level,
                                                     unsigned int version,
                                                     unsigned int pkgVersion)
{
  return new DistribUniformDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribWeibullDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribWeibullDistribution(unsigned int level,
                                                     unsigned int version,
                                                     unsigned int pkgVersion)
{
  return new DistribWeibullDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribBinomialDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribBinomialDistribution(unsigned int level,
                                                      unsigned int version,
                                                      unsigned int pkgVersion)
{
  return new DistribBinomialDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribGeometricDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribGeometricDistribution(unsigned int level,
                                                       unsigned int version,
                                                       unsigned int pkgVersion)
{
  return new DistribGeometricDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribHypergeometricDistribution (DistribDistribution_t)
 * using the given SBML Level, Version and &ldquo;distrib&rdquo; package
 * version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribHypergeometricDistribution(unsigned int level,
                                                            unsigned int
                                                              version,
                                                            unsigned int
                                                              pkgVersion)
{
  return new DistribHypergeometricDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribNegativeBinomialDistribution (DistribDistribution_t)
 * using the given SBML Level, Version and &ldquo;distrib&rdquo; package
 * version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribNegativeBinomialDistribution(
                                                              unsigned int
                                                                level,
                                                              unsigned int
                                                                version,
                                                              unsigned int
                                                                pkgVersion)
{
  return new DistribNegativeBinomialDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribPoissonDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribPoissonDistribution(unsigned int level,
                                                     unsigned int version,
                                                     unsigned int pkgVersion)
{
  return new DistribPoissonDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribBernoulliDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribBernoulliDistribution(unsigned int level,
                                                       unsigned int version,
                                                       unsigned int pkgVersion)
{
  return new DistribBernoulliDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribCategoricalDistribution (DistribDistribution_t) using
 * the given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribCategoricalDistribution(unsigned int level,
                                                         unsigned int version,
                                                         unsigned int
                                                           pkgVersion)
{
  return new DistribCategoricalDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribMultivariateDistribution (DistribDistribution_t) using
 * the given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribMultivariateDistribution(unsigned int level,
                                                          unsigned int version,
                                                          unsigned int
                                                            pkgVersion)
{
  return new DistribMultivariateDistribution(level, version, pkgVersion);
}


/*
 * Creates a new DistribExternalDistribution (DistribDistribution_t) using the
 * given SBML Level, Version and &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribDistribution_t *
DistribDistribution_createDistribExternalDistribution(unsigned int level,
                                                      unsigned int version,
                                                      unsigned int pkgVersion)
{
  return new DistribExternalDistribution(level, version, pkgVersion);
}


/*
 * Creates and returns a deep copy of this DistribDistribution_t object.
 */
LIBSBML_EXTERN
DistribDistribution_t*
DistribDistribution_clone(const DistribDistribution_t* dd)
{
  if (dd != NULL)
  {
    return static_cast<DistribDistribution_t*>(dd->clone());
  }
  else
  {
    return NULL;
  }
}


/*
 * Frees this DistribDistribution_t object.
 */
LIBSBML_EXTERN
void
DistribDistribution_free(DistribDistribution_t* dd)
{
  if (dd != NULL)
  {
    delete dd;
  }
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribBetaDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribBetaDistribution(const DistribDistribution_t * dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribBetaDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribCauchyDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribCauchyDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribCauchyDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribChiSquareDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribChiSquareDistribution(const DistribDistribution_t
  * dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribChiSquareDistribution())
    : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribExponentialDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribExponentialDistribution(const
  DistribDistribution_t * dd)
{
  return (dd != NULL) ?
    static_cast<int>(dd->isDistribExponentialDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribFDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribFDistribution(const DistribDistribution_t * dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribFDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribGammaDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribGammaDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribGammaDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribInverseGammaDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribInverseGammaDistribution(const
  DistribDistribution_t * dd)
{
  return (dd != NULL) ?
    static_cast<int>(dd->isDistribInverseGammaDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribLaPlaceDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribLaPlaceDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribLaPlaceDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribLogNormalDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribLogNormalDistribution(const DistribDistribution_t
  * dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribLogNormalDistribution())
    : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribLogisticDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribLogisticDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribLogisticDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribNormalDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribNormalDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribNormalDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribParetoDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribParetoDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribParetoDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribRayleighDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribRayleighDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribRayleighDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribStudentTDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribStudentTDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribStudentTDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribUniformDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribUniformDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribUniformDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribWeibullDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribWeibullDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribWeibullDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribBinomialDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribBinomialDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribBinomialDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribGeometricDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribGeometricDistribution(const DistribDistribution_t
  * dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribGeometricDistribution())
    : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribHypergeometricDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribHypergeometricDistribution(const
  DistribDistribution_t * dd)
{
  return (dd != NULL) ?
    static_cast<int>(dd->isDistribHypergeometricDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribNegativeBinomialDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribNegativeBinomialDistribution(const
  DistribDistribution_t * dd)
{
  return (dd != NULL) ?
    static_cast<int>(dd->isDistribNegativeBinomialDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribPoissonDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribPoissonDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribPoissonDistribution()) :
    0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribBernoulliDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribBernoulliDistribution(const DistribDistribution_t
  * dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribBernoulliDistribution())
    : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribCategoricalDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribCategoricalDistribution(const
  DistribDistribution_t * dd)
{
  return (dd != NULL) ?
    static_cast<int>(dd->isDistribCategoricalDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribMultivariateDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribMultivariateDistribution(const
  DistribDistribution_t * dd)
{
  return (dd != NULL) ?
    static_cast<int>(dd->isDistribMultivariateDistribution()) : 0;
}


/*
 * Predicate returning @c 1 if this DistribDistribution_t is of type
 * DistribExternalDistribution_t
 */
LIBSBML_EXTERN
int
DistribDistribution_isDistribExternalDistribution(const DistribDistribution_t *
  dd)
{
  return (dd != NULL) ? static_cast<int>(dd->isDistribExternalDistribution()) :
    0;
}




LIBSBML_CPP_NAMESPACE_END



/**
 * @file DistribUncertainty.cpp
 * @brief Implementation of the DistribUncertainty class.
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
#include <sbml/packages/distrib/sbml/DistribUncertainty.h>
#include <sbml/packages/distrib/validator/DistribSBMLError.h>
#include <sbml/util/ElementFilter.h>

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
 * Creates a new DistribUncertainty using the given SBML Level, Version and
 * &ldquo;distrib&rdquo; package version.
 */
DistribUncertainty::DistribUncertainty(unsigned int level,
                                       unsigned int version,
                                       unsigned int pkgVersion)
  : SBase(level, version)
  , mUncertStatistics (NULL)
  , mDistribution (NULL)
{
  setSBMLNamespacesAndOwn(new DistribPkgNamespaces(level, version,
    pkgVersion));
  connectToChild();
}


/*
 * Creates a new DistribUncertainty using the given DistribPkgNamespaces
 * object.
 */
DistribUncertainty::DistribUncertainty(DistribPkgNamespaces *distribns)
  : SBase(distribns)
  , mUncertStatistics (NULL)
  , mDistribution (NULL)
{
  setElementNamespace(distribns->getURI());
  connectToChild();
  loadPlugins(distribns);
}


/*
 * Copy constructor for DistribUncertainty.
 */
DistribUncertainty::DistribUncertainty(const DistribUncertainty& orig)
  : SBase( orig )
  , mUncertStatistics ( NULL )
  , mDistribution ( NULL )
{
  if (orig.mUncertStatistics != NULL)
  {
    mUncertStatistics = orig.mUncertStatistics->clone();
  }

  if (orig.mDistribution != NULL)
  {
    mDistribution = orig.mDistribution->clone();
  }

  connectToChild();
}


/*
 * Assignment operator for DistribUncertainty.
 */
DistribUncertainty&
DistribUncertainty::operator=(const DistribUncertainty& rhs)
{
  if (&rhs != this)
  {
    SBase::operator=(rhs);
    delete mUncertStatistics;
    if (rhs.mUncertStatistics != NULL)
    {
      mUncertStatistics = rhs.mUncertStatistics->clone();
    }
    else
    {
      mUncertStatistics = NULL;
    }

    delete mDistribution;
    if (rhs.mDistribution != NULL)
    {
      mDistribution = rhs.mDistribution->clone();
    }
    else
    {
      mDistribution = NULL;
    }

    connectToChild();
  }

  return *this;
}


/*
 * Creates and returns a deep copy of this DistribUncertainty object.
 */
DistribUncertainty*
DistribUncertainty::clone() const
{
  return new DistribUncertainty(*this);
}


/*
 * Destructor for DistribUncertainty.
 */
DistribUncertainty::~DistribUncertainty()
{
  delete mUncertStatistics;
  mUncertStatistics = NULL;
  delete mDistribution;
  mDistribution = NULL;
}


/*
 * Returns the value of the "id" attribute of this DistribUncertainty.
 */
const std::string&
DistribUncertainty::getId() const
{
  return mId;
}


/*
 * Returns the value of the "name" attribute of this DistribUncertainty.
 */
const std::string&
DistribUncertainty::getName() const
{
  return mName;
}


/*
 * Predicate returning @c true if this DistribUncertainty's "id" attribute is
 * set.
 */
bool
DistribUncertainty::isSetId() const
{
  return (mId.empty() == false);
}


/*
 * Predicate returning @c true if this DistribUncertainty's "name" attribute is
 * set.
 */
bool
DistribUncertainty::isSetName() const
{
  return (mName.empty() == false);
}


/*
 * Sets the value of the "id" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setId(const std::string& id)
{
  return SyntaxChecker::checkAndSetSId(id, mId);
}


/*
 * Sets the value of the "name" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setName(const std::string& name)
{
  mName = name;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Unsets the value of the "id" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::unsetId()
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
 * Unsets the value of the "name" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::unsetName()
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
 * Returns the value of the "uncertStatistics" element of this
 * DistribUncertainty.
 */
const DistribUncertStatistics*
DistribUncertainty::getUncertStatistics() const
{
  return mUncertStatistics;
}


/*
 * Returns the value of the "uncertStatistics" element of this
 * DistribUncertainty.
 */
DistribUncertStatistics*
DistribUncertainty::getUncertStatistics()
{
  return mUncertStatistics;
}


/*
 * Returns the value of the "distribution" element of this DistribUncertainty.
 */
const DistribDistribution*
DistribUncertainty::getDistribution() const
{
  return mDistribution;
}


/*
 * Returns the value of the "distribution" element of this DistribUncertainty.
 */
DistribDistribution*
DistribUncertainty::getDistribution()
{
  return mDistribution;
}


/*
 * Predicate returning @c true if this DistribUncertainty's "uncertStatistics"
 * element is set.
 */
bool
DistribUncertainty::isSetUncertStatistics() const
{
  return (mUncertStatistics != NULL);
}


/*
 * Predicate returning @c true if this DistribUncertainty's "distribution"
 * element is set.
 */
bool
DistribUncertainty::isSetDistribution() const
{
  return (mDistribution != NULL);
}


/*
 * Sets the value of the "uncertStatistics" element of this DistribUncertainty.
 */
int
DistribUncertainty::setUncertStatistics(const DistribUncertStatistics*
  uncertStatistics)
{
  if (uncertStatistics == NULL)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else if (uncertStatistics->hasRequiredElements() == false)
  {
    return LIBSBML_INVALID_OBJECT;
  }
  else if (getLevel() != uncertStatistics->getLevel())
  {
    return LIBSBML_LEVEL_MISMATCH;
  }
  else if (getVersion() != uncertStatistics->getVersion())
  {
    return LIBSBML_VERSION_MISMATCH;
  }
  else if (getPackageVersion() != uncertStatistics->getPackageVersion())
  {
    return LIBSBML_PKG_VERSION_MISMATCH;
  }
  else
  {
    delete mUncertStatistics;
    mUncertStatistics = (uncertStatistics != NULL) ?
      static_cast<DistribUncertStatistics*>(uncertStatistics->clone()) : NULL;
    if (mUncertStatistics != NULL) mUncertStatistics->connectToParent(this);
    return LIBSBML_OPERATION_SUCCESS;
  }
}


/*
 * Sets the value of the "distribution" element of this DistribUncertainty.
 */
int
DistribUncertainty::setDistribution(const DistribDistribution* distribution)
{
  if (distribution == NULL)
  {
    return LIBSBML_OPERATION_SUCCESS;
  }
  else if (distribution->hasRequiredElements() == false)
  {
    return LIBSBML_INVALID_OBJECT;
  }
  else if (getLevel() != distribution->getLevel())
  {
    return LIBSBML_LEVEL_MISMATCH;
  }
  else if (getVersion() != distribution->getVersion())
  {
    return LIBSBML_VERSION_MISMATCH;
  }
  else if (getPackageVersion() != distribution->getPackageVersion())
  {
    return LIBSBML_PKG_VERSION_MISMATCH;
  }
  else
  {
    delete mDistribution;
    mDistribution = (distribution != NULL) ?
      static_cast<DistribDistribution*>(distribution->clone()) : NULL;
    if (mDistribution != NULL) mDistribution->connectToParent(this);
    return LIBSBML_OPERATION_SUCCESS;
  }
}


/*
 * Creates a new DistribUncertStatistics object, adds it to this
 * DistribUncertainty object and returns the DistribUncertStatistics object
 * created.
 */
DistribUncertStatistics*
DistribUncertainty::createUncertStatistics()
{
  if (mUncertStatistics != NULL)
  {
    delete mUncertStatistics;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mUncertStatistics = new DistribUncertStatistics(distribns);

  mUncertStatistics->setElementName("uncertStatistics");

  delete distribns;

  connectToChild();

  return mUncertStatistics;
}


/*
 * Creates a new DistribBetaDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribBetaDistribution object
 * created.
 */
DistribBetaDistribution*
DistribUncertainty::createDistribBetaDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribBetaDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribBetaDistribution*>(mDistribution);
}


/*
 * Creates a new DistribCauchyDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribCauchyDistribution object
 * created.
 */
DistribCauchyDistribution*
DistribUncertainty::createDistribCauchyDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribCauchyDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribCauchyDistribution*>(mDistribution);
}


/*
 * Creates a new DistribChiSquareDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribChiSquareDistribution
 * object created.
 */
DistribChiSquareDistribution*
DistribUncertainty::createDistribChiSquareDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribChiSquareDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribChiSquareDistribution*>(mDistribution);
}


/*
 * Creates a new DistribExponentialDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribExponentialDistribution
 * object created.
 */
DistribExponentialDistribution*
DistribUncertainty::createDistribExponentialDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribExponentialDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribExponentialDistribution*>(mDistribution);
}


/*
 * Creates a new DistribFDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribFDistribution object
 * created.
 */
DistribFDistribution*
DistribUncertainty::createDistribFDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribFDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribFDistribution*>(mDistribution);
}


/*
 * Creates a new DistribGammaDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribGammaDistribution object
 * created.
 */
DistribGammaDistribution*
DistribUncertainty::createDistribGammaDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribGammaDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribGammaDistribution*>(mDistribution);
}


/*
 * Creates a new DistribInverseGammaDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribInverseGammaDistribution
 * object created.
 */
DistribInverseGammaDistribution*
DistribUncertainty::createDistribInverseGammaDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribInverseGammaDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribInverseGammaDistribution*>(mDistribution);
}


/*
 * Creates a new DistribLaPlaceDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribLaPlaceDistribution object
 * created.
 */
DistribLaPlaceDistribution*
DistribUncertainty::createDistribLaPlaceDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribLaPlaceDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribLaPlaceDistribution*>(mDistribution);
}


/*
 * Creates a new DistribLogNormalDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribLogNormalDistribution
 * object created.
 */
DistribLogNormalDistribution*
DistribUncertainty::createDistribLogNormalDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribLogNormalDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribLogNormalDistribution*>(mDistribution);
}


/*
 * Creates a new DistribLogisticDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribLogisticDistribution object
 * created.
 */
DistribLogisticDistribution*
DistribUncertainty::createDistribLogisticDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribLogisticDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribLogisticDistribution*>(mDistribution);
}


/*
 * Creates a new DistribNormalDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribNormalDistribution object
 * created.
 */
DistribNormalDistribution*
DistribUncertainty::createDistribNormalDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribNormalDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribNormalDistribution*>(mDistribution);
}


/*
 * Creates a new DistribParetoDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribParetoDistribution object
 * created.
 */
DistribParetoDistribution*
DistribUncertainty::createDistribParetoDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribParetoDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribParetoDistribution*>(mDistribution);
}


/*
 * Creates a new DistribRayleighDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribRayleighDistribution object
 * created.
 */
DistribRayleighDistribution*
DistribUncertainty::createDistribRayleighDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribRayleighDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribRayleighDistribution*>(mDistribution);
}


/*
 * Creates a new DistribStudentTDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribStudentTDistribution object
 * created.
 */
DistribStudentTDistribution*
DistribUncertainty::createDistribStudentTDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribStudentTDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribStudentTDistribution*>(mDistribution);
}


/*
 * Creates a new DistribUniformDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribUniformDistribution object
 * created.
 */
DistribUniformDistribution*
DistribUncertainty::createDistribUniformDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribUniformDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribUniformDistribution*>(mDistribution);
}


/*
 * Creates a new DistribWeibullDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribWeibullDistribution object
 * created.
 */
DistribWeibullDistribution*
DistribUncertainty::createDistribWeibullDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribWeibullDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribWeibullDistribution*>(mDistribution);
}


/*
 * Creates a new DistribBinomialDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribBinomialDistribution object
 * created.
 */
DistribBinomialDistribution*
DistribUncertainty::createDistribBinomialDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribBinomialDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribBinomialDistribution*>(mDistribution);
}


/*
 * Creates a new DistribGeometricDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribGeometricDistribution
 * object created.
 */
DistribGeometricDistribution*
DistribUncertainty::createDistribGeometricDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribGeometricDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribGeometricDistribution*>(mDistribution);
}


/*
 * Creates a new DistribHypergeometricDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribHypergeometricDistribution
 * object created.
 */
DistribHypergeometricDistribution*
DistribUncertainty::createDistribHypergeometricDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribHypergeometricDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribHypergeometricDistribution*>(mDistribution);
}


/*
 * Creates a new DistribNegativeBinomialDistribution object, adds it to this
 * DistribUncertainty object and returns the
 * DistribNegativeBinomialDistribution object created.
 */
DistribNegativeBinomialDistribution*
DistribUncertainty::createDistribNegativeBinomialDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribNegativeBinomialDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribNegativeBinomialDistribution*>(mDistribution);
}


/*
 * Creates a new DistribPoissonDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribPoissonDistribution object
 * created.
 */
DistribPoissonDistribution*
DistribUncertainty::createDistribPoissonDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribPoissonDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribPoissonDistribution*>(mDistribution);
}


/*
 * Creates a new DistribBernoulliDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribBernoulliDistribution
 * object created.
 */
DistribBernoulliDistribution*
DistribUncertainty::createDistribBernoulliDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribBernoulliDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribBernoulliDistribution*>(mDistribution);
}


/*
 * Creates a new DistribCategoricalDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribCategoricalDistribution
 * object created.
 */
DistribCategoricalDistribution*
DistribUncertainty::createDistribCategoricalDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribCategoricalDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribCategoricalDistribution*>(mDistribution);
}


/*
 * Creates a new DistribMultivariateDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribMultivariateDistribution
 * object created.
 */
DistribMultivariateDistribution*
DistribUncertainty::createDistribMultivariateDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribMultivariateDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribMultivariateDistribution*>(mDistribution);
}


/*
 * Creates a new DistribExternalDistribution object, adds it to this
 * DistribUncertainty object and returns the DistribExternalDistribution object
 * created.
 */
DistribExternalDistribution*
DistribUncertainty::createDistribExternalDistribution()
{
  if (mDistribution != NULL)
  {
    delete mDistribution;
  }

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());
  mDistribution = new DistribExternalDistribution(distribns);

  delete distribns;

  connectToChild();

  return static_cast<DistribExternalDistribution*>(mDistribution);
}


/*
 * Unsets the value of the "uncertStatistics" element of this
 * DistribUncertainty.
 */
int
DistribUncertainty::unsetUncertStatistics()
{
  delete mUncertStatistics;
  mUncertStatistics = NULL;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Unsets the value of the "distribution" element of this DistribUncertainty.
 */
int
DistribUncertainty::unsetDistribution()
{
  delete mDistribution;
  mDistribution = NULL;
  return LIBSBML_OPERATION_SUCCESS;
}


/*
 * Returns the XML element name of this DistribUncertainty object.
 */
const std::string&
DistribUncertainty::getElementName() const
{
  static const string name = "uncertainty";
  return name;
}


/*
 * Returns the libSBML type code for this DistribUncertainty object.
 */
int
DistribUncertainty::getTypeCode() const
{
  return SBML_DISTRIB_UNCERTAINTY;
}


/*
 * Predicate returning @c true if all the required attributes for this
 * DistribUncertainty object have been set.
 */
bool
DistribUncertainty::hasRequiredAttributes() const
{
  bool allPresent = true;

  return allPresent;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Write any contained elements
 */
void
DistribUncertainty::writeElements(XMLOutputStream& stream) const
{
  SBase::writeElements(stream);

  if (isSetUncertStatistics() == true)
  {
    mUncertStatistics->write(stream);
  }

  if (isSetDistribution() == true)
  {
    mDistribution->write(stream);
  }

  SBase::writeExtensionElements(stream);
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Accepts the given SBMLVisitor
 */
bool
DistribUncertainty::accept(SBMLVisitor& v) const
{
  v.visit(*this);

  if (mUncertStatistics != NULL)
  {
    mUncertStatistics->accept(v);
  }

  if (mDistribution != NULL)
  {
    mDistribution->accept(v);
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
DistribUncertainty::setSBMLDocument(SBMLDocument* d)
{
  SBase::setSBMLDocument(d);

  if (mUncertStatistics != NULL)
  {
    mUncertStatistics->setSBMLDocument(d);
  }

  if (mDistribution != NULL)
  {
    mDistribution->setSBMLDocument(d);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Connects to child elements
 */
void
DistribUncertainty::connectToChild()
{
  SBase::connectToChild();

  if (mUncertStatistics != NULL)
  {
    mUncertStatistics->connectToParent(this);
  }

  if (mDistribution != NULL)
  {
    mDistribution->connectToParent(this);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Enables/disables the given package with this element
 */
void
DistribUncertainty::enablePackageInternal(const std::string& pkgURI,
                                          const std::string& pkgPrefix,
                                          bool flag)
{
  SBase::enablePackageInternal(pkgURI, pkgPrefix, flag);

  if (isSetUncertStatistics())
  {
    mUncertStatistics->enablePackageInternal(pkgURI, pkgPrefix, flag);
  }

  if (isSetDistribution())
  {
    mDistribution->enablePackageInternal(pkgURI, pkgPrefix, flag);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Updates the namespaces when setLevelVersion is used
 */
void
DistribUncertainty::updateSBMLNamespace(const std::string& package,
                                        unsigned int level,
                                        unsigned int version)
{
  SBase::updateSBMLNamespace(package, level, version);

  if (mUncertStatistics != NULL)
  {
    mUncertStatistics->updateSBMLNamespace(package, level, version);
  }

  if (mDistribution != NULL)
  {
    mDistribution->updateSBMLNamespace(package, level, version);
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::getAttribute(const std::string& attributeName,
                                 bool& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::getAttribute(const std::string& attributeName,
                                 int& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::getAttribute(const std::string& attributeName,
                                 double& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::getAttribute(const std::string& attributeName,
                                 unsigned int& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Gets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::getAttribute(const std::string& attributeName,
                                 std::string& value) const
{
  int return_value = SBase::getAttribute(attributeName, value);

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
 * Predicate returning @c true if this DistribUncertainty's attribute
 * "attributeName" is set.
 */
bool
DistribUncertainty::isSetAttribute(const std::string& attributeName) const
{
  bool value = SBase::isSetAttribute(attributeName);

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
 * Sets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setAttribute(const std::string& attributeName, bool value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setAttribute(const std::string& attributeName, int value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setAttribute(const std::string& attributeName,
                                 double value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setAttribute(const std::string& attributeName,
                                 unsigned int value)
{
  int return_value = SBase::setAttribute(attributeName, value);

  return return_value;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Sets the value of the "attributeName" attribute of this DistribUncertainty.
 */
int
DistribUncertainty::setAttribute(const std::string& attributeName,
                                 const std::string& value)
{
  int return_value = SBase::setAttribute(attributeName, value);

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
 * DistribUncertainty.
 */
int
DistribUncertainty::unsetAttribute(const std::string& attributeName)
{
  int value = SBase::unsetAttribute(attributeName);

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
 * Creates and returns an new "elementName" object in this DistribUncertainty.
 */
SBase*
DistribUncertainty::createChildObject(const std::string& elementName)
{
  SBase* obj = NULL;

  // TO DO

  return obj;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Adds a new "elementName" object to this DistribUncertainty.
 */
int
DistribUncertainty::addChildObject(const std::string& elementName,
                                   const SBase* element)
{
  // TO DO

  return -1;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Removes and returns the new "elementName" object with the given id in this
 * DistribUncertainty.
 */
SBase*
DistribUncertainty::removeChildObject(const std::string& elementName,
                                      const std::string& id)
{
  // TO DO

  return NULL;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the number of "elementName" in this DistribUncertainty.
 */
unsigned int
DistribUncertainty::getNumObjects(const std::string& elementName)
{
  // TO DO

  return 0;
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Returns the nth object of "objectName" in this DistribUncertainty.
 */
SBase*
DistribUncertainty::getObject(const std::string& elementName,
                              unsigned int index)
{
  // TO DO

  return NULL;
}

/** @endcond */


/*
 * Returns the first child element that has the given @p id in the model-wide
 * SId namespace, or @c NULL if no such object is found.
 */
SBase*
DistribUncertainty::getElementBySId(const std::string& id)
{
  if (id.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mUncertStatistics != NULL)
  {
    if (mUncertStatistics->getId() == id)
    {
      return mUncertStatistics;
    }

    obj = mUncertStatistics->getElementBySId(id);
    if (obj != NULL)
    {
      return obj;
    }
  }

  if (mDistribution != NULL)
  {
    if (mDistribution->getId() == id)
    {
      return mDistribution;
    }

    obj = mDistribution->getElementBySId(id);
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
DistribUncertainty::getElementByMetaId(const std::string& metaid)
{
  if (metaid.empty())
  {
    return NULL;
  }

  SBase* obj = NULL;

  if (mUncertStatistics != NULL)
  {
    if (mUncertStatistics->getMetaId() == metaid)
    {
      return mUncertStatistics;
    }

    obj = mUncertStatistics->getElementByMetaId(metaid);
    if (obj != NULL)
    {
      return obj;
    }
  }

  if (mDistribution != NULL)
  {
    if (mDistribution->getMetaId() == metaid)
    {
      return mDistribution;
    }

    obj = mDistribution->getElementByMetaId(metaid);
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
DistribUncertainty::getAllElements(ElementFilter* filter)
{
  List* ret = new List();
  List* sublist = NULL;

  ADD_FILTERED_POINTER(ret, sublist, mUncertStatistics, filter);
  ADD_FILTERED_POINTER(ret, sublist, mDistribution, filter);


  ADD_FILTERED_FROM_PLUGIN(ret, sublist, filter);

  return ret;
}



/** @cond doxygenLibsbmlInternal */

/*
 * Creates a new object from the next XMLToken on the XMLInputStream
 */
SBase*
DistribUncertainty::createObject(XMLInputStream& stream)
{
  SBase* obj = NULL;

  const std::string& name = stream.peek().getName();

  DISTRIB_CREATE_NS(distribns, getSBMLNamespaces());

  if (name == "betaDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribBetaDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "cauchyDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribCauchyDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "chiSquareDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribChiSquareDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "exponentialDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribExponentialDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "fDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribFDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "gammaDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribGammaDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "inverseGammaDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribInverseGammaDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "laPlaceDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribLaPlaceDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "logNormalDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribLogNormalDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "logisticDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribLogisticDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "normalDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribNormalDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "paretoDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribParetoDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "rayleighDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribRayleighDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "studentTDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribStudentTDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "uniformDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribUniformDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "weibullDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribWeibullDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "binomialDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribBinomialDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "geometricDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribGeometricDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "hypergeometricDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribHypergeometricDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "negativeBinomialDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribNegativeBinomialDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "poissonDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribPoissonDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "bernoulliDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribBernoulliDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "categoricalDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribCategoricalDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "multivariateDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribMultivariateDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "externalDistribution")
  {
    if (isSetDistribution())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mDistribution;
    mDistribution = new DistribExternalDistribution(distribns);
    obj = mDistribution;
  }
  else if (name == "uncertStatistics")
  {
    if (isSetUncertStatistics())
    {
      getErrorLog()->logPackageError("distrib",
        DistribDistribUncertaintyAllowedElements, getPackageVersion(),
          getLevel(), getVersion());
    }

    delete mUncertStatistics;
    mUncertStatistics = new DistribUncertStatistics(distribns);
    mUncertStatistics->setElementName(name);
    obj = mUncertStatistics;
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
DistribUncertainty::addExpectedAttributes(ExpectedAttributes& attributes)
{
  SBase::addExpectedAttributes(attributes);

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
DistribUncertainty::readAttributes(const XMLAttributes& attributes,
                                   const ExpectedAttributes&
                                     expectedAttributes)
{
  unsigned int level = getLevel();
  unsigned int version = getVersion();
  unsigned int pkgVersion = getPackageVersion();
  unsigned int numErrs;
  bool assigned = false;
  SBMLErrorLog* log = getErrorLog();

  SBase::readAttributes(attributes, expectedAttributes);

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
          DistribDistribUncertaintyAllowedAttributes, pkgVersion, level, version,
            details);
      }
      else if (log->getError(n)->getErrorId() == UnknownCoreAttribute)
      {
        const std::string details = log->getError(n)->getMessage();
        log->remove(UnknownCoreAttribute);
        log->logPackageError("distrib",
          DistribDistribUncertaintyAllowedCoreAttributes, pkgVersion, level,
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
DistribUncertainty::readL3V1V1Attributes(const XMLAttributes& attributes)
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
      logEmptyString(mId, level, version, "<DistribUncertainty>");
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
      logEmptyString(mName, level, version, "<DistribUncertainty>");
    }
  }
}

/** @endcond */



/** @cond doxygenLibsbmlInternal */

/*
 * Reads the expected attributes into the member data variables
 */
void
DistribUncertainty::readL3V2V1Attributes(const XMLAttributes& attributes)
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
      logEmptyString(mId, level, version, "<DistribUncertainty>");
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
DistribUncertainty::writeAttributes(XMLOutputStream& stream) const
{
  SBase::writeAttributes(stream);

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
DistribUncertainty::writeL3V1V1Attributes(XMLOutputStream& stream) const
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
DistribUncertainty::writeL3V2V1Attributes(XMLOutputStream& stream) const
{
}

/** @endcond */




#endif /* __cplusplus */


/*
 * Creates a new DistribUncertainty_t using the given SBML Level, Version and
 * &ldquo;distrib&rdquo; package version.
 */
LIBSBML_EXTERN
DistribUncertainty_t *
DistribUncertainty_create(unsigned int level,
                          unsigned int version,
                          unsigned int pkgVersion)
{
  return new DistribUncertainty(level, version, pkgVersion);
}


/*
 * Creates and returns a deep copy of this DistribUncertainty_t object.
 */
LIBSBML_EXTERN
DistribUncertainty_t*
DistribUncertainty_clone(const DistribUncertainty_t* du)
{
  if (du != NULL)
  {
    return static_cast<DistribUncertainty_t*>(du->clone());
  }
  else
  {
    return NULL;
  }
}


/*
 * Frees this DistribUncertainty_t object.
 */
LIBSBML_EXTERN
void
DistribUncertainty_free(DistribUncertainty_t* du)
{
  if (du != NULL)
  {
    delete du;
  }
}


/*
 * Returns the value of the "id" attribute of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
char *
DistribUncertainty_getId(const DistribUncertainty_t * du)
{
  if (du == NULL)
  {
    return NULL;
  }

  return du->getId().empty() ? NULL : safe_strdup(du->getId().c_str());
}


/*
 * Returns the value of the "name" attribute of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
char *
DistribUncertainty_getName(const DistribUncertainty_t * du)
{
  if (du == NULL)
  {
    return NULL;
  }

  return du->getName().empty() ? NULL : safe_strdup(du->getName().c_str());
}


/*
 * Predicate returning @c 1 (true) if this DistribUncertainty_t's "id"
 * attribute is set.
 */
LIBSBML_EXTERN
int
DistribUncertainty_isSetId(const DistribUncertainty_t * du)
{
  return (du != NULL) ? static_cast<int>(du->isSetId()) : 0;
}


/*
 * Predicate returning @c 1 (true) if this DistribUncertainty_t's "name"
 * attribute is set.
 */
LIBSBML_EXTERN
int
DistribUncertainty_isSetName(const DistribUncertainty_t * du)
{
  return (du != NULL) ? static_cast<int>(du->isSetName()) : 0;
}


/*
 * Sets the value of the "id" attribute of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_setId(DistribUncertainty_t * du, const char * id)
{
  return (du != NULL) ? du->setId(id) : LIBSBML_INVALID_OBJECT;
}


/*
 * Sets the value of the "name" attribute of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_setName(DistribUncertainty_t * du, const char * name)
{
  return (du != NULL) ? du->setName(name) : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "id" attribute of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_unsetId(DistribUncertainty_t * du)
{
  return (du != NULL) ? du->unsetId() : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "name" attribute of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_unsetName(DistribUncertainty_t * du)
{
  return (du != NULL) ? du->unsetName() : LIBSBML_INVALID_OBJECT;
}


/*
 * Returns the value of the "uncertStatistics" element of this
 * DistribUncertainty_t.
 */
LIBSBML_EXTERN
const DistribUncertStatistics_t*
DistribUncertainty_getUncertStatistics(const DistribUncertainty_t * du)
{
  if (du == NULL)
  {
    return NULL;
  }

  return (DistribUncertStatistics_t*)(du->getUncertStatistics());
}


/*
 * Returns the value of the "distribution" element of this
 * DistribUncertainty_t.
 */
LIBSBML_EXTERN
const DistribDistribution_t*
DistribUncertainty_getDistribution(const DistribUncertainty_t * du)
{
  if (du == NULL)
  {
    return NULL;
  }

  return (DistribDistribution_t*)(du->getDistribution());
}


/*
 * Predicate returning @c 1 (true) if this DistribUncertainty_t's
 * "uncertStatistics" element is set.
 */
LIBSBML_EXTERN
int
DistribUncertainty_isSetUncertStatistics(const DistribUncertainty_t * du)
{
  return (du != NULL) ? static_cast<int>(du->isSetUncertStatistics()) : 0;
}


/*
 * Predicate returning @c 1 (true) if this DistribUncertainty_t's
 * "distribution" element is set.
 */
LIBSBML_EXTERN
int
DistribUncertainty_isSetDistribution(const DistribUncertainty_t * du)
{
  return (du != NULL) ? static_cast<int>(du->isSetDistribution()) : 0;
}


/*
 * Sets the value of the "uncertStatistics" element of this
 * DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_setUncertStatistics(DistribUncertainty_t * du,
                                       const DistribUncertStatistics_t*
                                         uncertStatistics)
{
  return (du != NULL) ? du->setUncertStatistics(uncertStatistics) :
    LIBSBML_INVALID_OBJECT;
}


/*
 * Sets the value of the "distribution" element of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_setDistribution(DistribUncertainty_t * du,
                                   const DistribDistribution_t* distribution)
{
  return (du != NULL) ? du->setDistribution(distribution) :
    LIBSBML_INVALID_OBJECT;
}


/*
 * Creates a new DistribUncertStatistics_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribUncertStatistics_t object
 * created.
 */
LIBSBML_EXTERN
DistribUncertStatistics_t*
DistribUncertainty_createUncertStatistics(DistribUncertainty_t* du)
{
  if (du == NULL)
  {
    return NULL;
  }

  return (DistribUncertStatistics_t*)(du->createUncertStatistics());
}


/*
 * Creates a new DistribBetaDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribBetaDistribution_t object
 * created.
 */
LIBSBML_EXTERN
DistribBetaDistribution_t*
DistribUncertainty_createDistribBetaDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribBetaDistribution() : NULL;
}


/*
 * Creates a new DistribCauchyDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribCauchyDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribCauchyDistribution_t*
DistribUncertainty_createDistribCauchyDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribCauchyDistribution() : NULL;
}


/*
 * Creates a new DistribChiSquareDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribChiSquareDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribChiSquareDistribution_t*
DistribUncertainty_createDistribChiSquareDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribChiSquareDistribution() : NULL;
}


/*
 * Creates a new DistribExponentialDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribExponentialDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribExponentialDistribution_t*
DistribUncertainty_createDistribExponentialDistribution(DistribUncertainty_t*
  du)
{
  return (du != NULL) ? du->createDistribExponentialDistribution() : NULL;
}


/*
 * Creates a new DistribFDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribFDistribution_t object
 * created.
 */
LIBSBML_EXTERN
DistribFDistribution_t*
DistribUncertainty_createDistribFDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribFDistribution() : NULL;
}


/*
 * Creates a new DistribGammaDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribGammaDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribGammaDistribution_t*
DistribUncertainty_createDistribGammaDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribGammaDistribution() : NULL;
}


/*
 * Creates a new DistribInverseGammaDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the
 * DistribInverseGammaDistribution_t object created.
 */
LIBSBML_EXTERN
DistribInverseGammaDistribution_t*
DistribUncertainty_createDistribInverseGammaDistribution(DistribUncertainty_t*
  du)
{
  return (du != NULL) ? du->createDistribInverseGammaDistribution() : NULL;
}


/*
 * Creates a new DistribLaPlaceDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribLaPlaceDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribLaPlaceDistribution_t*
DistribUncertainty_createDistribLaPlaceDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribLaPlaceDistribution() : NULL;
}


/*
 * Creates a new DistribLogNormalDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribLogNormalDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribLogNormalDistribution_t*
DistribUncertainty_createDistribLogNormalDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribLogNormalDistribution() : NULL;
}


/*
 * Creates a new DistribLogisticDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribLogisticDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribLogisticDistribution_t*
DistribUncertainty_createDistribLogisticDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribLogisticDistribution() : NULL;
}


/*
 * Creates a new DistribNormalDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribNormalDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribNormalDistribution_t*
DistribUncertainty_createDistribNormalDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribNormalDistribution() : NULL;
}


/*
 * Creates a new DistribParetoDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribParetoDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribParetoDistribution_t*
DistribUncertainty_createDistribParetoDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribParetoDistribution() : NULL;
}


/*
 * Creates a new DistribRayleighDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribRayleighDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribRayleighDistribution_t*
DistribUncertainty_createDistribRayleighDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribRayleighDistribution() : NULL;
}


/*
 * Creates a new DistribStudentTDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribStudentTDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribStudentTDistribution_t*
DistribUncertainty_createDistribStudentTDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribStudentTDistribution() : NULL;
}


/*
 * Creates a new DistribUniformDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribUniformDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribUniformDistribution_t*
DistribUncertainty_createDistribUniformDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribUniformDistribution() : NULL;
}


/*
 * Creates a new DistribWeibullDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribWeibullDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribWeibullDistribution_t*
DistribUncertainty_createDistribWeibullDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribWeibullDistribution() : NULL;
}


/*
 * Creates a new DistribBinomialDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribBinomialDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribBinomialDistribution_t*
DistribUncertainty_createDistribBinomialDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribBinomialDistribution() : NULL;
}


/*
 * Creates a new DistribGeometricDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribGeometricDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribGeometricDistribution_t*
DistribUncertainty_createDistribGeometricDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribGeometricDistribution() : NULL;
}


/*
 * Creates a new DistribHypergeometricDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the
 * DistribHypergeometricDistribution_t object created.
 */
LIBSBML_EXTERN
DistribHypergeometricDistribution_t*
DistribUncertainty_createDistribHypergeometricDistribution(DistribUncertainty_t*
  du)
{
  return (du != NULL) ? du->createDistribHypergeometricDistribution() : NULL;
}


/*
 * Creates a new DistribNegativeBinomialDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the
 * DistribNegativeBinomialDistribution_t object created.
 */
LIBSBML_EXTERN
DistribNegativeBinomialDistribution_t*
DistribUncertainty_createDistribNegativeBinomialDistribution(DistribUncertainty_t*
  du)
{
  return (du != NULL) ? du->createDistribNegativeBinomialDistribution() : NULL;
}


/*
 * Creates a new DistribPoissonDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribPoissonDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribPoissonDistribution_t*
DistribUncertainty_createDistribPoissonDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribPoissonDistribution() : NULL;
}


/*
 * Creates a new DistribBernoulliDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribBernoulliDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribBernoulliDistribution_t*
DistribUncertainty_createDistribBernoulliDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribBernoulliDistribution() : NULL;
}


/*
 * Creates a new DistribCategoricalDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribCategoricalDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribCategoricalDistribution_t*
DistribUncertainty_createDistribCategoricalDistribution(DistribUncertainty_t*
  du)
{
  return (du != NULL) ? du->createDistribCategoricalDistribution() : NULL;
}


/*
 * Creates a new DistribMultivariateDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the
 * DistribMultivariateDistribution_t object created.
 */
LIBSBML_EXTERN
DistribMultivariateDistribution_t*
DistribUncertainty_createDistribMultivariateDistribution(DistribUncertainty_t*
  du)
{
  return (du != NULL) ? du->createDistribMultivariateDistribution() : NULL;
}


/*
 * Creates a new DistribExternalDistribution_t object, adds it to this
 * DistribUncertainty_t object and returns the DistribExternalDistribution_t
 * object created.
 */
LIBSBML_EXTERN
DistribExternalDistribution_t*
DistribUncertainty_createDistribExternalDistribution(DistribUncertainty_t* du)
{
  return (du != NULL) ? du->createDistribExternalDistribution() : NULL;
}


/*
 * Unsets the value of the "uncertStatistics" element of this
 * DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_unsetUncertStatistics(DistribUncertainty_t * du)
{
  return (du != NULL) ? du->unsetUncertStatistics() : LIBSBML_INVALID_OBJECT;
}


/*
 * Unsets the value of the "distribution" element of this DistribUncertainty_t.
 */
LIBSBML_EXTERN
int
DistribUncertainty_unsetDistribution(DistribUncertainty_t * du)
{
  return (du != NULL) ? du->unsetDistribution() : LIBSBML_INVALID_OBJECT;
}


/*
 * Predicate returning @c 1 (true) if all the required attributes for this
 * DistribUncertainty_t object have been set.
 */
LIBSBML_EXTERN
int
DistribUncertainty_hasRequiredAttributes(const DistribUncertainty_t * du)
{
  return (du != NULL) ? static_cast<int>(du->hasRequiredAttributes()) : 0;
}




LIBSBML_CPP_NAMESPACE_END



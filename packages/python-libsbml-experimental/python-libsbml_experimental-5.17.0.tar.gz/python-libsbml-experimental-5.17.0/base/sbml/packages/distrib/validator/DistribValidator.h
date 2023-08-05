/**
 * @file DistribValidator.h
 * @brief Definition of DistribValidator.
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
 *
 * @class DistribValidator
 * @sbmlbrief{distrib} Entry point for &ldquo;distrib&rdquo; package
 * validation.
 *
 * @htmlinclude not-sbml-warning.html
 *
 * @copydetails doc_common_intro_package_validators
 *
 * The DistribValidator class extends the Validator class from core libSBML to
 * apply validation to the constructs introduced by the SBML Level&nbsp;3
 * Distributions package. This class then acts as a base class for any
 * validators that apply rules to the &ldquo;distrib&rdquo; package
 * specification constructs or to entire models that use the
 * &ldquo;distrib&rdquo; package, and may therefore be subject to other global
 * restrictions introduced.
 *
 * @copydetails doc_section_package_validators_general_info
 */


#ifndef DistribValidator_H__
#define DistribValidator_H__




#ifdef __cplusplus



/** @cond doxygenLibsbmlInternal */

#include <list>
#include <string>

/** @endcond */

#include <sbml/SBMLError.h>
#include <sbml/validator/Validator.h>
#include <sbml/SBMLReader.h>


LIBSBML_CPP_NAMESPACE_BEGIN


class VConstraint;
struct DistribValidatorConstraints;
class SBMLDocument;


class DistribValidator : public Validator
{
public:

  /**
   * Creates a new DistribValidator object for the given category of
   * validation.
   *
   * @param category code indicating the type of validation that this validator
   * will perform.
   */
  DistribValidator(SBMLErrorCategory_t category = LIBSBML_CAT_SBML);


  /**
   * Destroys this DistribValidator object.
   */
  virtual ~DistribValidator();


  /**
   * Initializes this DistribValidator object.
   *
   * When creating a subclass of DistribValidator, override this method to add
   * your own validation code.
   */
  virtual void init() = 0;


  /**
   * Adds the given VConstraint object to this DistribValidator.
   *
   * @param c the VConstraint ("validator constraint") object to add.
   */
  virtual void addConstraint(VConstraint* c);


  /**
   * Validates the given SBMLDocument
   *
   * @param d the SBMLDocument object to be validated.
   *
   * @return the number of validation failures that occurred. The objects
   * describing the actual failures can be retrieved using getFailures().
   */
  virtual unsigned int validate(const SBMLDocument& d);


  /**
   * Validates the SBMLDocument located at the given filename
   *
   * @param filename the path to the file to be read and validated.
   *
   * @return the number of validation failures that occurred. The objects
   * describing the actual failures can be retrieved using getFailures().
   */
  virtual unsigned int validate(const std::string& filename);


protected:

  /** @cond doxygenLibsbmlInternal */

  DistribValidatorConstraints* mDistribConstraints;

  friend class DistribValidatingVisitor;

  /** @endcond */

};



LIBSBML_CPP_NAMESPACE_END




#endif /* __cplusplus */




#endif /* !DistribValidator_H__ */



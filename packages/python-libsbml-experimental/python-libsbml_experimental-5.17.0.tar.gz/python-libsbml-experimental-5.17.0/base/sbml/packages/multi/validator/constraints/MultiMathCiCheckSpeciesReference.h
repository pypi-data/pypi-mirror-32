/**
 * @cond doxygenLibsbmlInternal
 *
 * @file    MultiMathCiCheckSpeciesReference.h
 * @brief   checks &lt;ci&gt; references is a valid SpeciesReference Id
 * @author  Fengkai Zhang
 * @author  Sarah Keating
 * 
 * <!--------------------------------------------------------------------------
 * This file is part of libSBML.  Please visit http://sbml.org for more
 * information about SBML, and the latest version of libSBML.
 *
 * Copyright (C) 2013-2018 jointly by the following organizations:
 *     1. California Institute of Technology, Pasadena, CA, USA
 *     2. EMBL European Bioinformatics Institute (EMBL-EBI), Hinxton, UK
 *     3. University of Heidelberg, Heidelberg, Germany
 * 
 * Copyright (C) 2009-2013 jointly by the following organizations: 
 *     1. California Institute of Technology, Pasadena, CA, USA
 *     2. EMBL European Bioinformatics Institute (EMBL-EBI), Hinxton, UK
 *  
 * Copyright (C) 2006-2008 by the California Institute of Technology,
 *     Pasadena, CA, USA 
 *  
 * Copyright (C) 2002-2005 jointly by the following organizations: 
 *     1. California Institute of Technology, Pasadena, CA, USA
 *     2. Japan Science and Technology Agency, Japan
 * 
 * This library is free software; you can redistribute it and/or modify it
 * under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation.  A copy of the license agreement is provided
 * in the file named "LICENSE.txt" included with this software distribution
 * and also available online as http://sbml.org/software/libsbml/license.html
 * ---------------------------------------------------------------------- -->*/

#ifndef MultiMathCiCheckSpeciesReference_h
#define MultiMathCiCheckSpeciesReference_h


#ifdef __cplusplus


#include <string>
#include <sstream>
#include <math.h>

#include <sbml/validator/VConstraint.h>

#include "../../../../validator/constraints/MathMLBase.h"

LIBSBML_CPP_NAMESPACE_BEGIN

class ASTNode;
class MultiValidator;

class MultiMathCiCheckSpeciesReference: public MathMLBase
{
public:

  /**
   * Creates a new Constraint with the given @p id.
   */
  MultiMathCiCheckSpeciesReference (unsigned int id, MultiValidator& v);

  /**
   * Destroys this Constraint.
   */
  virtual ~MultiMathCiCheckSpeciesReference ();


protected:

  /**
   * Checks the MathML of the ASTnode 
   * is appropriate for the function being performed
   *
   * If an inconsistency is found, an error message is logged.
   */
  virtual void checkMath (const Model& m, const ASTNode& node, const SBase & sb);
  


  /**
   * Returns the error message to use when logging constraint violations.
   * This method is called by logFailure.
   *
   * @return the error message to use when logging constraint violations.
   */
  virtual const std::string
  getMessage (const ASTNode& node, const SBase& object);

private:
  void checkCiSpeciesReference(const Model& m, const ASTNode& node, const SBase & sb);

};

LIBSBML_CPP_NAMESPACE_END

#endif  /* __cplusplus */
#endif  /* MultiMathCiCheckSpeciesReference_h */
/** @endcond */


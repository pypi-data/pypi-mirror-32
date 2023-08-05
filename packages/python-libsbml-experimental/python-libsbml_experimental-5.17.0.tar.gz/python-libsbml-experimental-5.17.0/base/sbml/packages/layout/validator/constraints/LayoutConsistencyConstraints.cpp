/** @cond doxygenLibsbmlInternal */
/**
 * @file   LayoutConsistencyConstraints.cpp
 * @brief  Implementation of the LayoutConsistencyConstraints class
 * @author Generated by autocreate code
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
 * ------------------------------------------------------------------------ -->
 */

#ifndef  AddingConstraintsToValidator

#include <sbml/validator/VConstraint.h>

#include <sbml/packages/layout/validator/LayoutSBMLError.h>
#include <sbml/packages/layout/common/LayoutExtensionTypes.h>
#include <sbml/packages/layout/extension/LayoutSBMLDocumentPlugin.h>
#include <sbml/util/ElementFilter.h>

#endif  /* AddingConstrainstToValidator */

#include <sbml/common/libsbml-namespace.h>
#include <sbml/validator/ConstraintMacros.h>

/** @cond doxygenIgnored */
using namespace std;
LIBSBML_CPP_NAMESPACE_USE
/** @endcond */

class GraphicalObjectFilter : public ElementFilter
{
public:
  GraphicalObjectFilter (): ElementFilter ()
  {
  };


  virtual bool filter(const SBase* element)
  {
    // return in case we don't have a valid element with an id
    if (element == NULL || element->isSetId() == false)
    {
        return false;
    }

    // otherwise we have an id set and want to keep the element
    // if it is derived from graphical object
    int tc = element->getTypeCode();
    bool keep = false;
    switch (tc)
    {
    case SBML_LAYOUT_COMPARTMENTGLYPH:
    case SBML_LAYOUT_GRAPHICALOBJECT:
    case SBML_LAYOUT_REACTIONGLYPH:
    case SBML_LAYOUT_SPECIESGLYPH:
    case SBML_LAYOUT_SPECIESREFERENCEGLYPH:
    case SBML_LAYOUT_TEXTGLYPH:
    case SBML_LAYOUT_REFERENCEGLYPH:
    case SBML_LAYOUT_GENERALGLYPH:
      keep = true;
      break;
    default:
      break;
    }

    return keep;      
  };

};

//20315
START_CONSTRAINT (LayoutLayoutMustHaveDimensions, Layout, l)
{
  bool fail = false;

  if (l.getDimensionsExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20406
START_CONSTRAINT (LayoutGOMetaIdRefMustReferenceObject, GraphicalObject, go)
{
  pre(go.isSetMetaIdRef() == true);

  bool fail = false;
  msg = "The <" + go.getElementName() + "> ";
  if (go.isSetId()) {
    msg += "with the id '" + go.getId() + "' ";
  }
  msg += "has a metaidRef '" + go.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (go.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(go.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT


//20407
START_CONSTRAINT (LayoutGOMustContainBoundingBox, GraphicalObject, go)
{
  bool fail = false;

  if (go.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20503
START_CONSTRAINT (LayoutCGAllowedElements, CompartmentGlyph, glyph)
{
  bool fail = false;

  if (glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20506
START_CONSTRAINT (LayoutCGMetaIdRefMustReferenceObject, CompartmentGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20508
START_CONSTRAINT (LayoutCGCompartmentMustRefComp, CompartmentGlyph, glyph)
{
  pre(glyph.isSetCompartmentId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a compartment '" + glyph.getCompartmentId() +
    "' which is not the id of any <compartment> in the model.";

  if (m.getCompartment(glyph.getCompartmentId()) == NULL)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20509
START_CONSTRAINT (LayoutCGNoDuplicateReferences, CompartmentGlyph, glyph)
{
  pre(glyph.isSetCompartmentId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  const Compartment * c = m.getCompartment(glyph.getCompartmentId());

  // bail if we have not found a match
  pre ( c != NULL);

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (c->isSetMetaId() == false || c->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20603
START_CONSTRAINT (LayoutSGAllowedElements, SpeciesGlyph, glyph)
{
  bool fail = false;

  if (glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20606
START_CONSTRAINT (LayoutSGMetaIdRefMustReferenceObject, SpeciesGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20608
START_CONSTRAINT (LayoutSGSpeciesMustRefSpecies, SpeciesGlyph, glyph)
{
  pre(glyph.isSetSpeciesId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a species '" + glyph.getSpeciesId() + 
    "' which is not the id of any <species> in the model.";

  if (m.getSpecies(glyph.getSpeciesId()) == NULL)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20609
START_CONSTRAINT (LayoutSGNoDuplicateReferences, SpeciesGlyph, glyph)
{
  pre(glyph.isSetSpeciesId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  const Species * obj = m.getSpecies(glyph.getSpeciesId());

  // bail if we have not found a match
  pre ( obj != NULL);

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (obj->isSetMetaId() == false || obj->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20703
START_CONSTRAINT (LayoutRGAllowedElements, ReactionGlyph, glyph)
{
  bool fail = false;

  if (glyph.getCurveExplicitlySet() == false
    && glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  if (glyph.getNumSpeciesReferenceGlyphs() == 0)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20706
START_CONSTRAINT (LayoutRGMetaIdRefMustReferenceObject, ReactionGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20708
START_CONSTRAINT (LayoutRGReactionMustRefReaction, ReactionGlyph, glyph)
{
  pre(glyph.isSetReactionId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a reaction '" + glyph.getReactionId() +
    "' which is not the id of any <reaction> in the model.";

  if (m.getReaction(glyph.getReactionId()) == NULL)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20709
START_CONSTRAINT (LayoutRGNoDuplicateReferences, ReactionGlyph, glyph)
{
  pre(glyph.isSetReactionId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  const Reaction * obj = m.getReaction(glyph.getReactionId());

  // bail if we have not found a match
  pre ( obj != NULL);

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (obj->isSetMetaId() == false || obj->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20803
START_CONSTRAINT (LayoutGGAllowedElements, GeneralGlyph, glyph)
{
  bool fail = false;

  if (glyph.getCurveExplicitlySet() == false
    && glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20806
START_CONSTRAINT (LayoutGGMetaIdRefMustReferenceObject, GeneralGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT


//20808
START_CONSTRAINT (LayoutGGReferenceMustRefObject, GeneralGlyph, glyph)
{
  pre(glyph.isSetReferenceId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a reference '" + glyph.getReferenceId() + 
    "' which is not the id of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getIdList().contains(glyph.getReferenceId()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20809
START_CONSTRAINT (LayoutGGNoDuplicateReferences, GeneralGlyph, glyph)
{
  pre(glyph.isSetReferenceId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  List * elements = plug->getListElementsWithId();
  unsigned int i = 0;
  SBase * obj = NULL;
  while (i < elements->getSize())
  {
    obj = (SBase*)(elements->get(i));
    if (obj->getId() == glyph.getReferenceId())
    {
      break;
    }
    i++;
  }

  // bail if we have not found a match
  pre ( i < elements->getSize());

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (obj == NULL || obj->isSetMetaId() == false || obj->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20903
START_CONSTRAINT (LayoutTGAllowedElements, TextGlyph, glyph)
{
  bool fail = false;

  if (glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20906
START_CONSTRAINT (LayoutTGMetaIdRefMustReferenceObject, TextGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20908
START_CONSTRAINT (LayoutTGOriginOfTextMustRefObject, TextGlyph, glyph)
{
  pre(glyph.isSetOriginOfTextId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has an originOfText '" + glyph.getOriginOfTextId() + 
    "' which is not the id of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getIdList().contains(glyph.getOriginOfTextId()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20909
START_CONSTRAINT (LayoutTGNoDuplicateReferences, TextGlyph, glyph)
{
  pre(glyph.isSetOriginOfTextId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  std::string origin = glyph.getOriginOfTextId();

  bool fail = false;

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  List * elements = plug->getListElementsWithId();
  unsigned int i = 0;
  SBase * obj = NULL;
  while (i < elements->getSize())
  {
    obj = (SBase*)(elements->get(i));
    if (obj->getId() == origin)
    {
      break;
    }
    i++;
  }

  // bail if we have not found a match
  pre ( i < elements->getSize());

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (obj == NULL || obj->isSetMetaId() == false || obj->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20911
START_CONSTRAINT (LayoutTGGraphicalObjectMustRefObject, TextGlyph, glyph)
{
  pre(glyph.isSetGraphicalObjectId() == true);

  std::string goRef = glyph.getGraphicalObjectId();
  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a graphicalObject '" + goRef +
    "' which is not the id of any <graphicalObject> in the model.";

  const Layout* layout = static_cast<const Layout *>
                        (glyph.getAncestorOfType(SBML_LAYOUT_LAYOUT, "layout"));


  GraphicalObjectFilter filter;

  List * allGO = const_cast<Layout *>(layout)->getAllElements(&filter);

  unsigned int i = 0;
  bool match = false;
  
  while(match == false && i < allGO->getSize())
  {
    if (static_cast<SBase*>(allGO->get(i))->getId() == goRef)
    {
      match = true;
    }

    i++;
  }
  
  delete allGO;

  if (match == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21003
START_CONSTRAINT (LayoutSRGAllowedElements, SpeciesReferenceGlyph, glyph)
{
  bool fail = false;

  if (glyph.getCurveExplicitlySet() == false
    && glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21006
START_CONSTRAINT (LayoutSRGMetaIdRefMustReferenceObject, SpeciesReferenceGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21008
START_CONSTRAINT (LayoutSRGSpeciesRefMustRefObject, SpeciesReferenceGlyph, glyph)
{
  pre(glyph.isSetSpeciesReferenceId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a speciesReference '" + glyph.getSpeciesReferenceId() + 
    "' which is not the id of any <speciesReference> in the model.";

  if (m.getSpeciesReference(glyph.getSpeciesReferenceId()) == NULL
    && m.getModifierSpeciesReference(glyph.getSpeciesReferenceId()) == NULL)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21009
START_CONSTRAINT (LayoutSRGNoDuplicateReferences, SpeciesReferenceGlyph, glyph)
{
  pre(glyph.isSetSpeciesReferenceId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  std::string sref = glyph.getSpeciesReferenceId();

  bool fail = false;

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  List * elements = plug->getListElementsWithId();
  unsigned int i = 0;
  SBase * obj = NULL;
  while (i < elements->getSize())
  {
    obj = (SBase*)(elements->get(i));
    if (obj->getId() == sref)
    {
      break;
    }
    i++;
  }

  // bail if we have not found a match
  pre ( i < elements->getSize());

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (obj == NULL || obj->isSetMetaId() == false || obj->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//20911
START_CONSTRAINT (LayoutSRGSpeciesGlyphMustRefObject, SpeciesReferenceGlyph, glyph)
{
  pre(glyph.isSetSpeciesGlyphId() == true);

  std::string sGlyph = glyph.getSpeciesGlyphId();
  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a graphicalObject '" + sGlyph + 
    "' which is not the id of any <graphicalObject> in the model.";

  const Layout* layout = static_cast<const Layout *>
                        (glyph.getAncestorOfType(SBML_LAYOUT_LAYOUT, "layout"));


  unsigned int i = 0;
  bool match = false;
  
  while(match == false && i < layout->getNumSpeciesGlyphs())
  {
    if (layout->getSpeciesGlyph(i)->getId() == sGlyph)
    {
      match = true;
    }

    i++;
  }

  if (match == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21103
START_CONSTRAINT (LayoutREFGAllowedElements, ReferenceGlyph, glyph)
{
  bool fail = false;

  if (glyph.getCurveExplicitlySet() == false
    && glyph.getBoundingBoxExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21106
START_CONSTRAINT (LayoutREFGMetaIdRefMustReferenceObject, ReferenceGlyph, glyph)
{
  pre(glyph.isSetMetaIdRef() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a metaidRef '" + glyph.getMetaIdRef() + 
    "' which is not the metaid of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getMetaidList().contains(glyph.getMetaIdRef()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21108
START_CONSTRAINT (LayoutREFGReferenceMustRefObject, ReferenceGlyph, glyph)
{
  pre(glyph.isSetReferenceId() == true);

  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a reference '" + glyph.getReferenceId() + 
    "' which is not the id of any element in the model.";

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  if (plug->getIdList().contains(glyph.getReferenceId()) == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21109
START_CONSTRAINT (LayoutREFGNoDuplicateReferences, ReferenceGlyph, glyph)
{
  pre(glyph.isSetReferenceId() == true);
  pre(glyph.isSetMetaIdRef() == true);

  std::string origin = glyph.getReferenceId();

  bool fail = false;

  const LayoutSBMLDocumentPlugin * plug = 
                            static_cast<const LayoutSBMLDocumentPlugin*>
                            (glyph.getSBMLDocument()->getPlugin("layout"));

  List * elements = plug->getListElementsWithId();
  unsigned int i = 0;
  SBase * obj = NULL;
  while (i < elements->getSize())
  {
    obj = (SBase*)(elements->get(i));
    if (obj->getId() == origin)
    {
      break;
    }
    i++;
  }

  // bail if we have not found a match
  pre ( i < elements->getSize());

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "references multiple objects.";

  if (obj == NULL || obj->isSetMetaId() == false || obj->getMetaId() != glyph.getMetaIdRef())
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21111
START_CONSTRAINT (LayoutREFGGlyphMustRefObject, ReferenceGlyph, glyph)
{
  pre(glyph.isSetGlyphId() == true);

  std::string goRef = glyph.getGlyphId();
  bool fail = false;

  msg = "The <" + glyph.getElementName() + "> ";
  if (glyph.isSetId()) {
    msg += "with the id '" + glyph.getId() + "' ";
  }
  msg += "has a glyph '" + goRef + 
    "' which is not the id of any <graphicalObject> in the model.";

  const Layout* layout = static_cast<const Layout *>
                        (glyph.getAncestorOfType(SBML_LAYOUT_LAYOUT, "layout"));


  GraphicalObjectFilter filter;

  List * allGO = const_cast<Layout *>(layout)->getAllElements(&filter);

  unsigned int i = 0;
  bool match = false;
  
  while(match == false && i < allGO->getSize())
  {
    if (static_cast<SBase*>(allGO->get(i))->getId() == goRef)
    {
      match = true;
    }

    i++;
  }

  delete allGO;

  if (match == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21303
START_CONSTRAINT (LayoutBBoxAllowedElements, BoundingBox, bb)
{
  bool fail = false;

  if (bb.getPositionExplicitlySet() == false
    || bb.getDimensionsExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21305
START_CONSTRAINT (LayoutBBoxConsistent3DDefinition, BoundingBox, bb)
{
  pre (bb.getPositionExplicitlySet() == true);
  pre (bb.getDimensionsExplicitlySet() == true);

  bool fail = false;

  if (bb.getPosition()->getZOffsetExplicitlySet() == false)
  {
    if (bb.getDimensions()->getDExplicitlySet() == true)
    {
      fail = true;
    }
  }

  inv(fail == false);
}
END_CONSTRAINT


//21403
START_CONSTRAINT (LayoutCurveAllowedElements, Curve, curve)
{
  bool fail = false;

  if (curve.getNumCurveSegments() == 0)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT

//21503
START_CONSTRAINT (LayoutLSegAllowedElements, LineSegment, line)
{
  bool fail = false;

  if (line.getStartExplicitlySet() == false
    || line.getEndExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT


//21603
START_CONSTRAINT (LayoutCBezAllowedElements, CubicBezier, line)
{
  bool fail = false;

  if (line.getStartExplicitlySet() == false
    || line.getEndExplicitlySet() == false
    || line.getBasePt1ExplicitlySet() == false
    || line.getBasePt2ExplicitlySet() == false)
  {
    fail = true;
  }

  inv(fail == false);
}
END_CONSTRAINT
/** @endcond */



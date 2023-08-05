/**
 * @cond doxygenLibsbmlInternal
 *
 * @file    IdentifierConsistencyConstraintsDeclared.cxx
 * @brief   Declarations of constraints
 * @author  SBML Team
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

addConstraint(new UniqueIdsInModel(10301, *this));

addConstraint(new UniqueIdsForUnitDefinitions(10302, *this));

addConstraint(new UniqueIdsInKineticLaw(10303, *this));

addConstraint(new UniqueVarsInRules(10304, *this));

addConstraint(new UniqueVarsInEventAssignments(10305, *this));

addConstraint(new UniqueVarsInEventsAndRules(10306, *this));

addConstraint(new UniqueMetaId(10307, *this));

addConstraint(new ModelUnitsDangling(10313, *this));

addConstraint(new VConstraintParameter10313(*this));

addConstraint(new VConstraintSpecies10313(*this));

addConstraint(new VConstraintCompartment10313(*this));

addConstraint(new VConstraintLocalParameter10313(*this));

addConstraint(new VConstraintParameter99303(*this));

addConstraint(new VConstraintLocalParameter99303(*this));

addConstraint(new VConstraintSpecies99303(*this));

addConstraint(new VConstraintCompartment99303(*this));

addConstraint(new VConstraintKineticLaw99303(*this));

addConstraint(new VConstraintEvent99303(*this));

addConstraint(new VConstraintModel99303(*this));
/** @endcond */


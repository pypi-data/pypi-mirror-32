/**
 * @cond doxygenLibsbmlInternal
 *
 * @file    FbcConsistencyValidator.cpp
 * @brief   Checks an SBML model for structural consistency
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
 * This library is free software; you can redistribute it and/or modify it
 * under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation.  A copy of the license agreement is provided
 * in the file named "LICENSE.txt" included with this software distribution and
 * also available online as http://sbml.org/software/libsbml/license.html
 * ---------------------------------------------------------------------- -->*/

#ifndef doxygen_ignore

#include <sbml/packages/fbc/validator/FbcConsistencyValidator.h>


/*
 * Compile ConsistencyConstraints
 */
#include "constraints/FbcConsistencyConstraints.cpp"

LIBSBML_CPP_NAMESPACE_BEGIN
#ifdef __cplusplus

/*
 * Initializes this Validator with a set of Constraints.
 */
void
FbcConsistencyValidator::init ()
{
#define  AddingConstraintsToValidator 1
#include "constraints/FbcConsistencyConstraintsDeclared.cxx"
}

#endif /* __cplusplus */
LIBSBML_CPP_NAMESPACE_END

#endif
/** @endcond */

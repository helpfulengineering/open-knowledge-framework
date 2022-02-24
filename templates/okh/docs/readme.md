## Design Files
- Fieldname: design-files
- Purpose: Storage for design files such as CAD and vector files.

design-files:
  - path: [value]         # Absolute or relative path
    title: [value]        # text
    format: [string]        # file extension and required software (if any) to access the files

Rules: Recommended. May either refer to a location in which files are stored or can list and link to design files individually. In the former case, the title field is not required.


Schematics

Fieldname: schematics

Purpose: Links to schematics. Includes all types of engineering drawings.

Format: Absolute or relative path

schematics:
  - path: [value]     # Absolute or relative path
    title: [value]    # text

Bill of Materials

Fieldname: bom

Purpose: Links to the bill of materials.

Format: Absolute or relative path.

Rules: Recommended. Should be provided as a single file.
List of Tools

Fieldname: tool-list

Purpose: Links to a list of tools required to make the thing.

Format: Absolute or relative path.

Rules: Recommended. Should be provided as a single file.
Assembly Instructions

Fieldname: making-instructions

Purpose: Links to the making instructions, e.g. assembly instructions.

Format: 

making-instructions:
  - path: [value]      # Absolute or relative path
    title: [value]     # text

Rules: Recommended. May include more than one assembly instructions, e.g. where the instructions are provided in different formats.
Manufacturing files

Fieldname: manufacturing-files

Purpose: Links to the manufacturing files, such as 3D printing files.

Format: 

manufacturing-files:
 - path: [value]       # Absolute or relative path
   title: [value]      # text

Rules: Recommended where applicable. 
Risk Assessment

Fieldname: risk-assessment

Purpose: Links to the risk assessment.

Format: 

risk-assessment:
  - path: [value]     # Absolute or relative path
    title: [value]    # text

Rules: Recommended. May include more than one, e.g. where different assessments cover different aspects or making, operation, maintenance and disposal.
Tool settings and documentation

Fieldname: tool-settings

Purpose: Links to settings for tools and machines.

Format: 

tool-settings: 
  - path: [value]     # Absolute or relative path 
    title: [value]    # text

Rules: Recommended where applicable. May include more than one.
Quality control instructions

Fieldname: quality-instructions

Purpose: Links to instructions for testing and/or quality management.

Format: 

quality-instructions:
  - path: [value]     # Absolute or relative path
    title: [value]    # text

Rules: Recommended. May include more than one.
Operating instructions

Fieldname: operating-instructions

Purpose: Links to instructions for operating the thing.

Format: 

operating-instructions:
  - path: [value]    # Absolute or relative path
    title: [value]   # text

Rules: Recommended.
Maintenance instructions

Fieldname: maintenance-instructions

Purpose: Links to instructions for maintaining the thing.

Format: 

maintenance-instructions:
  - path: [value]    # Absolute or relative path
    title: [value]   # text

Rules: Recommended.
Disposal instructions

Fieldname: disposal-instructions

Purpose: Links to instructions for disposing the thing at the end of its life.

Format: 

disposal-instructions:
  - path: [value]    # Absolute or relative path
    title: [value]   # text

Rules: Recommended.
Software

Fieldname: software

Purpose: Links to source code repository or software executables used by the thing.

Format: 

software:
 - path: [value]     # Absolute or relative path
   title: [value]    # text

Rules: Recommended where applicable. May include more than one.
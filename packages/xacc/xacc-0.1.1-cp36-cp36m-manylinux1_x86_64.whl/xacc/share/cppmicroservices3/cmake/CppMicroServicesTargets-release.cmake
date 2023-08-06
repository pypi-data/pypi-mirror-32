#----------------------------------------------------------------
# Generated CMake target import file for configuration "RELEASE".
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "usResourceCompiler" for configuration "RELEASE"
set_property(TARGET usResourceCompiler APPEND PROPERTY IMPORTED_CONFIGURATIONS RELEASE)
set_target_properties(usResourceCompiler PROPERTIES
  IMPORTED_LOCATION_RELEASE "${_IMPORT_PREFIX}/bin/usResourceCompiler3"
  )

list(APPEND _IMPORT_CHECK_TARGETS usResourceCompiler )
list(APPEND _IMPORT_CHECK_FILES_FOR_usResourceCompiler "${_IMPORT_PREFIX}/bin/usResourceCompiler3" )

# Import target "usShell" for configuration "RELEASE"
set_property(TARGET usShell APPEND PROPERTY IMPORTED_CONFIGURATIONS RELEASE)
set_target_properties(usShell PROPERTIES
  IMPORTED_LOCATION_RELEASE "${_IMPORT_PREFIX}/bin/usShell3"
  )

list(APPEND _IMPORT_CHECK_TARGETS usShell )
list(APPEND _IMPORT_CHECK_FILES_FOR_usShell "${_IMPORT_PREFIX}/bin/usShell3" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)

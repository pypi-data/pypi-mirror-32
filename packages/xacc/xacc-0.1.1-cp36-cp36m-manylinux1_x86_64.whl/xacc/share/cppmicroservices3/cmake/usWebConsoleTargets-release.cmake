#----------------------------------------------------------------
# Generated CMake target import file for configuration "RELEASE".
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "usWebConsole" for configuration "RELEASE"
set_property(TARGET usWebConsole APPEND PROPERTY IMPORTED_CONFIGURATIONS RELEASE)
set_target_properties(usWebConsole PROPERTIES
  IMPORTED_LOCATION_RELEASE "${_IMPORT_PREFIX}/lib/libusWebConsole.so.0.1.0"
  IMPORTED_SONAME_RELEASE "libusWebConsole.so.0.1.0"
  )

list(APPEND _IMPORT_CHECK_TARGETS usWebConsole )
list(APPEND _IMPORT_CHECK_FILES_FOR_usWebConsole "${_IMPORT_PREFIX}/lib/libusWebConsole.so.0.1.0" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)

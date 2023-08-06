#----------------------------------------------------------------
# Generated CMake target import file for configuration "RELEASE".
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "usHttpService" for configuration "RELEASE"
set_property(TARGET usHttpService APPEND PROPERTY IMPORTED_CONFIGURATIONS RELEASE)
set_target_properties(usHttpService PROPERTIES
  IMPORTED_LOCATION_RELEASE "${_IMPORT_PREFIX}/lib/libusHttpService.so.0.1.0"
  IMPORTED_SONAME_RELEASE "libusHttpService.so.0.1.0"
  )

list(APPEND _IMPORT_CHECK_TARGETS usHttpService )
list(APPEND _IMPORT_CHECK_FILES_FOR_usHttpService "${_IMPORT_PREFIX}/lib/libusHttpService.so.0.1.0" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)

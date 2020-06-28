cmake_minimum_required(VERSION 3.0)
project(gimp-plugin LANGUAGES NONE)

find_program(GOLANG_COMPILER go)
if (GOLANG_COMPILER STREQUAL "GOLANG_COMPILER-NOTFOUND")
    message(FATAL_ERROR "Could not find golang compiler")
else()
    message(STATUS "Found golang compiler: ${GOLANG_COMPILER}")
endif()

find_library(GIMP_LIB gimp-2.0)
if (GIMP_LIB STREQUAL "GIMP_LIB-NOTFOUND")
    if (NOT CMAKE_SYSTEM_NAME STREQUAL "Windows")
        message(FATAL_ERROR "Could not find GIMP")
    endif()
else()
    message(STATUS "Found GIMP: ${GIMP_LIB}")
endif()

find_program(GIMPTOOL NAMES gimptool gimptool-2.0)
if (GIMPTOOL STREQUAL "GIMPTOOL-NOTFOUND")
    if (CMAKE_SYSTEM_NAME STREQUAL "Windows")
        get_filename_component(GIMP_INSTALLDIR "[HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\GIMP-2_is1;InstallLocation]" ABSOLUTE CACHE)
        if (GIMP_INSTALLDIR STREQUAL "")
            message(FATAL_ERROR "Could not determine GIMP plugin directory")
        else()
            set(INSTALL_PATH "${GIMP_INSTALLDIR}/lib/gimp/2.0/plug-ins")
        endif()
    else()
        message(FATAL_ERROR "Could find gimptool")
    endif()
else()
    message(STATUS "Found gimptool: ${GIMPTOOL}")
    execute_process(
        COMMAND "${GIMPTOOL}" --gimpplugindir
        RESULT_VARIABLE GIMPTOOL_RETVAL
        OUTPUT_VARIABLE GIMPTOOL_PLUGINDIR
    )
    if (NOT GIMPTOOL_RETVAL EQUAL 0)
        message(FATAL_ERROR "Could not determine GIMP plugin directory")
    else()
        string(REGEX REPLACE "\n$" "" GIMPTOOL_PLUGINDIR_STRIP "${GIMPTOOL_PLUGINDIR}")
        set(INSTALL_PATH "${GIMPTOOL_PLUGINDIR_STRIP}/plug-ins")
    endif()
endif()
message(STATUS "Found install dir: ${INSTALL_PATH}")

if (CMAKE_SYSTEM_NAME STREQUAL "Windows")
    set(OUTPUT_FILENAME "radio-tracker.exe")
else()
    set(OUTPUT_FILENAME "radio-tracker")
endif()
set(PLUGIN_INSTALL_DIR "${INSTALL_PATH}/radio-tracker")

set(SOURCES
    plugin/eeprom/main.go
    plugin/eeprom/procedure.go
    plugin/gimp/bufferedImage.go
    plugin/gimp/image.go
    plugin/gimp/layer.go
    plugin/debug.go
    plugin/exports.go
    plugin/fileHandlers.go
    plugin/main.go
    plugin/param.go
    plugin/procedure.go
    go.mod
    go.sum
    main.go
)

add_custom_command(
    OUTPUT "${CMAKE_CURRENT_BINARY_DIR}/${OUTPUT_FILENAME}"
    COMMAND go build -o "${CMAKE_CURRENT_BINARY_DIR}/${OUTPUT_FILENAME}"
    DEPENDS ${SOURCES}
    WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
)

add_custom_target(
    build_plugin ALL
    DEPENDS "${CMAKE_CURRENT_BINARY_DIR}/${OUTPUT_FILENAME}"
)

install(
    FILES "${CMAKE_CURRENT_BINARY_DIR}/${OUTPUT_FILENAME}"
    DESTINATION "${PLUGIN_INSTALL_DIR}"
)

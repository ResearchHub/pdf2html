cmake_minimum_required(VERSION 3.23.1 FATAL_ERROR)

set(CMAKE_BUILD_TYPE Release CACHE STRING "Build configuration (Debug, Release, RelWithDebInfo, MinSizeRel)")
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED on)

# debug build flags (overwrite default cmake debug flags)
set(CMAKE_C_FLAGS_DEBUG "-ggdb -pg")
set(CMAKE_CXX_FLAGS_DEBUG "-ggdb -pg")
set(CMAKE_EXE_LINKER_FLAGS_DEBUG "-pg")
# release build flags (overwrite default cmake release flags)
set(CMAKE_C_FLAGS_RELEASE "-O2 -DNDEBUG")
set(CMAKE_CXX_FLAGS_RELEASE "-O2 -DNDEBUG")
# generic flags
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Woverloaded-virtual")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++14 -pthread")
# clang compiler need c++11 flag
if ("${CMAKE_CXX_COMPILER_ID}" STREQUAL "Clang")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -stdlib=libc++")
endif()
# CYGWIN or GCC 4.5.x bug
if(CYGWIN)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++14")
endif()


project(pdf2html)
include(ExternalProject)
find_package(PkgConfig REQUIRED)

set(ENABLE_SVG 1)
include(${CMAKE_SOURCE_DIR}/src/css_class_names.cmakelists.txt)
configure_file (${CMAKE_SOURCE_DIR}/src/util/css_const.h.in ${CMAKE_SOURCE_DIR}/src/util/css_const.h)
configure_file (${CMAKE_SOURCE_DIR}/share/base.css.in ${CMAKE_SOURCE_DIR}/share/base.css)
configure_file(${CMAKE_SOURCE_DIR}/src/pdf2htmlEX-config.h.in ${CMAKE_SOURCE_DIR}/src/pdf2htmlEX-config.h)

set(PROJECT_SRC ${PROJECT_SRC}
    src/Param.h
    src/pdf2htmlEX.cc
    src/pdf2htmlEX-config.h
    src/HTMLRenderer/HTMLRenderer.h
    src/HTMLRenderer/draw.cc
    src/HTMLRenderer/general.cc
    src/HTMLRenderer/image.cc
    src/HTMLRenderer/font.cc
    src/HTMLRenderer/form.cc
    src/HTMLRenderer/link.cc
    src/HTMLRenderer/outline.cc
    src/HTMLRenderer/state.cc
    src/HTMLRenderer/text.cc
    src/BackgroundRenderer/BackgroundRenderer.h
    src/BackgroundRenderer/BackgroundRenderer.cc
    src/BackgroundRenderer/SplashBackgroundRenderer.h
    src/BackgroundRenderer/SplashBackgroundRenderer.cc
    src/BackgroundRenderer/CairoBackgroundRenderer.h
    src/BackgroundRenderer/CairoBackgroundRenderer.cc
    src/util/const.h
    src/util/const.cc
    src/util/css_const.h
    src/util/encoding.h
    src/util/encoding.cc
    src/util/ffw.h
    src/util/ffw.c
    src/util/math.h
    src/util/math.cc
    src/util/misc.h
    src/util/misc.cc
    src/util/namespace.h
    src/util/path.h
    src/util/path.cc
    src/util/unicode.h
    src/util/unicode.cc
    src/util/mingw.h
    src/util/mingw.cc
    src/util/SignalHandler.h
    src/util/SignalHandler.cc
    src/ArgParser.h
    src/ArgParser.cc
    src/Base64Stream.h
    src/Base64Stream.cc
    src/Color.h
    src/Color.cc
    src/CoveredTextDetector.h
    src/CoveredTextDetector.cc
    src/DrawingTracer.h
    src/DrawingTracer.cc
    src/HTMLState.h
    src/HTMLTextLine.h
    src/HTMLTextLine.cc
    src/HTMLTextPage.h
    src/HTMLTextPage.cc
    src/Preprocessor.h
    src/Preprocessor.cc
    src/StringFormatter.h
    src/StringFormatter.cc
    src/TmpFiles.h
    src/TmpFiles.cc
)

include_directories(
    ${CMAKE_SOURCE_DIR}/libs/src/poppler-build/poppler
    ${CMAKE_SOURCE_DIR}/libs/src/poppler-build
    ${CMAKE_SOURCE_DIR}/libs/src/poppler/poppler
    ${CMAKE_SOURCE_DIR}/libs/src/poppler
    ${CMAKE_SOURCE_DIR}/libs/src/fontforge-build/inc
    ${CMAKE_SOURCE_DIR}/libs/src/fontforge/inc
    ${CMAKE_SOURCE_DIR}/libs/src/fontforge/fontforge
    ${CMAKE_SOURCE_DIR}/libs/src/fontforge
    ${CMAKE_SOURCE_DIR}/src
)

add_executable(${PROJECT_NAME} ${PROJECT_SRC})

# If we are using Alpine Linux then we need to add -lintl
if (EXISTS /usr/lib/libintl.so )
    target_link_libraries(${PROJECT_NAME} -lintl)
endif()

# include/link glib-2.0
pkg_search_module(GLIB REQUIRED glib-2.0)
target_include_directories(${PROJECT_NAME} PRIVATE ${GLIB_INCLUDE_DIRS})
target_link_libraries(${PROJECT_NAME} ${GLIB_LDFLAGS})

pkg_search_module(CAIRO REQUIRED cairo>=1.10.0)
message(NOTICE ${CAIRO_INCLUDE_DIRS})
include_directories(${CAIRO_INCLUDE_DIRS})
link_directories(${CAIRO_LIBRARY_DIRS})
target_link_libraries(${PROJECT_NAME} ${CAIRO_LIBRARIES})

find_package(Freetype REQUIRED)
message(NOTICE ${FREETYPE_INCLUDE_DIRS})
include_directories(${FREETYPE_INCLUDE_DIRS})
link_directories(${FREETYPE_LIBRARY_DIRS})
target_link_libraries(${PROJECT_NAME} ${FREETYPE_LIBRARIES})

target_link_libraries(${PROJECT_NAME}
    # link poppler
    ${CMAKE_SOURCE_DIR}/libs/src/poppler-build/glib/libpoppler-glib.a
    ${CMAKE_SOURCE_DIR}/libs/src/poppler-build/libpoppler.a
    #link fontforge
    ${CMAKE_SOURCE_DIR}/libs/src/fontforge-build/lib/libfontforge.so
    #link dependency libs
    -ljpeg
    -lpng
    -lfontconfig
    -lfreetype
    -lxml2
    -lglib-2.0
    -lgio-2.0
    -pthread
    -lz
    -lm
)

set(PDF2HTML_RESOURCES
    ${CMAKE_SOURCE_DIR}/share/base.css
    ${CMAKE_SOURCE_DIR}/share/manifest
)

install (TARGETS ${PROJECT_NAME} DESTINATION bin)
install (FILES ${PDF2HTML_RESOURCES} DESTINATION share/pdf2html)

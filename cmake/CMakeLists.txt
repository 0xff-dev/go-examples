CMAKE_MINIMUM_REQUIRED (VERSION 2.8)

project(main)

configure_file(
    "${PROJECT_SOURCE_DIR}/config.h.in"
    "${PROJECT_BINARY_DIR}/config.h"
    )

option (USE_MYMATH
        "Use provided math implementation" ON)

if (USE_MYMATH)
    include_directories("${PROJECT_SOURCE_DIR}/math")
    add_subdirectory(math)
    set (EXTRA_LIBS ${EXTRA_LIBS} MathFunctions)
endif (USE_MYMATH)


aux_source_directory(. base_dir)
add_executable(main ${base_dir})
target_link_libraries(main ${EXTRA_LIBS})

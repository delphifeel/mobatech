cmake_minimum_required(VERSION 3.10)
project(Mobatech LANGUAGES CXX VERSION 0.0.0.1)

# set(CMAKE_BUILD_TYPE "Release")
# set(CMAKE_CXX_FLAGS_RELEASE "-Of -march=native -funroll-loops -DNDEBUG")

# cmake -DCMAKE_CXX_FLAGS_RELEASE="-Of -march=native -funroll-loops -DNDEBUG" -DCMAKE_BUILD_TYPE=Release ..

# add_compile_options(-DCORE_SET_DEBUG_COLORED)
set(LOCAL_COMPLILE_OPTIONS

    # -std=c99 -pedantic
    # -std=c++20
    $<$<NOT:$<CONFIG:RELEASE>>:
    -Wall -Wextra
    -Wno-missing-braces -Wno-missing-field-initializers -Wformat=2
    -Wswitch-default -Wswitch-enum -Wcast-align -Wpointer-arith
    -Wbad-function-cast -Wstrict-overflow=5 -Wstrict-prototypes -Winline
    -Wundef -Wnested-externs -Wcast-qual -Wshadow -Wunreachable-code
    -Wlogical-op -Wfloat-equal -Wstrict-aliasing=2 -Wredundant-decls
    -Wold-style-definition
    -ggdb3
    -O0
    -fno-omit-frame-pointer -ffloat-store -fno-common -fstrict-aliasing
    >
    $<$<CONFIG:RELEASE>:
    -O3
    -funroll-loops
    -DNDEBUG
    >
)

add_executable(${PROJECT_NAME}
    main.cpp
)

add_subdirectory(external/json)
target_include_directories(${PROJECT_NAME} PUBLIC ${PROJECT_SOURCE_DIR}/external/json/include)
target_link_libraries(${PROJECT_NAME} PRIVATE
    nlohmann_json::nlohmann_json
)

# target_compile_options(${PROJECT_NAME} PRIVATE ${LOCAL_COMPLILE_OPTIONS})
# target_include_directories(Mobatech PUBLIC include)

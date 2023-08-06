
set(COPY_DIR ${PROJECT_BINARY_DIR}/examples/TomoPy)
set(FILES
    basic.py
)

foreach(_FILE ${FILES})
    configure_file(${_FILE} ${COPY_DIR}/${_FILE} COPYONLY)
endforeach(_FILE ${FILES})

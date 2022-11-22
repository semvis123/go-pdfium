package implementation_webassembly

import (
	"unsafe"

	"github.com/google/uuid"
	"github.com/klippa-app/go-pdfium/references"
)

func (p *PdfiumImplementation) registerFormHandle(formHandle *uint64, formInfo unsafe.Pointer) *FormHandleHandle {
	ref := uuid.New()
	handle := &FormHandleHandle{
		handle:           formHandle,
		nativeRef:        references.FPDF_FORMHANDLE(ref.String()),
		formInfo:         formInfo,
		pagePointers:     map[unsafe.Pointer]references.FPDF_PAGE{},
		documentPointers: map[unsafe.Pointer]references.FPDF_DOCUMENT{},
	}

	p.formHandleRefs[handle.nativeRef] = handle

	return handle
}

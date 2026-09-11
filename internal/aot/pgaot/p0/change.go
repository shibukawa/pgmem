package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_changeDependencyFor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l2
	if l2 == int32(2613) {
		v31 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l4
	v33 = int32(0)
	if l2 == int32(2613) {
		v46 = v33
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	if base.Ui32(int32(11999)) < base.Ui32(l3) {
		v31 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = (base.B2i32(l2 != int32(2615)) | base.B2i32(l3 != int32(2200))) & base.B2i32(l2 != int32(1262))
	goto L2
L5:
	;
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	goto L5
L7:
	;
	if base.Ui32(int32(11999)) < base.Ui32(l4) {
		v46 = v33
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v46 = (base.B2i32(l2 != int32(2615)) | base.B2i32(l4 != int32(2200))) & base.B2i32(l2 != int32(1262))
	goto L6
L9:
	;
	m.G0 = v13 + int32(112)
	return v155
L10:
	;
	v47 = int32(1)
	if v46 != 0 {
		v155 = v47
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v64 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	F_recordMultipleDependencies(m, v13+int32(16), v13+int32(4), int32(1), int32(110))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v155 = v47
	goto L9
L16:
	;
	F_ScanKeyInit(m, v13+int32(16), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_ScanKeyInit(m, v13-int32(-64), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v86 = F_systable_beginscan(m, v64, int32(2673), int32(1), int32(0), int32(2), v13+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v88 = F_systable_getnext(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if v88 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v96 = v88
	v98 = v6
	goto L24
L22:
	;
	v140 = v6
	goto L23
L23:
	;
	F_systable_endscan(m, v86)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L39
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	v102 = v100 + v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	if v103 != l2 {
		v129 = v98
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v140 = v129
	goto L23
L26:
	;
	v130 = F_systable_getnext(m, v86)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L14
	} else {
		goto L37
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v105 != l3 {
		v129 = v98
		goto L26
	} else {
		goto L28
	}
L28:
	;
	if v46 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v129 = v98 + int32(1)
	goto L26
L30:
	;
	F_CatalogTupleDelete(m, v64, v96+int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L14
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v111 = F_heap_copytuple(m, v96)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L34
	}
L33:
	;
	goto L29
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v113+v114)+16)) = l4
	F_CatalogTupleUpdate(m, v64, v111+int32(4), v111)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v111)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	if v130 != 0 {
		v96 = v130
		v98 = v129
		goto L24
	} else {
		goto L38
	}
L38:
	;
	goto L25
L39:
	;
	F_sequence_close(m, v64, int32(3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v155 = v140
	goto L9
}

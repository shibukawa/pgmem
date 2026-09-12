package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyToBinaryStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v7, int32(1616816), int32(11))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_appendBinaryStringInfo(m, v14, v5+int32(8), int32(4))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_appendBinaryStringInfo(m, v22, v5+int32(12), int32(4))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
func F_CopyToTextOneRow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CopySendTextLikeEndOfRow(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L28
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v20 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v47 <= v46 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v22<<(uint(int32(2))%32))))
	v36 = F_OutputFunctionCall(m, v18+v22*int32(28), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v42 = F_strlen(m, v41)
	mBase = m.M
	F_appendBinaryStringInfo(m, v40, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	F_CopyAttributeOutText(m, l0, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	goto L4
L12:
	;
	v52 = v46
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v62 = int32(2)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v52<<(uint(v62)%32))))
	v66 = int32(1)
	v67 = v65 - v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v67<<(uint(v62)%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v82 <= v79+v66 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L1
L15:
	;
	if v70&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	F_appendStringInfoChar(m, v78, v77)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86+v79))) = uint8(v77)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v92 = v90 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v92))) = uint8(v96)
	goto L15
L19:
	;
	goto L15
L20:
	;
	v116 = v52 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v116 < v117 {
		v52 = v116
		goto L13
	} else {
		goto L27
	}
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v104 = F_strlen(m, v103)
	mBase = m.M
	F_appendBinaryStringInfo(m, v102, v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v110 = F_OutputFunctionCall(m, v18+v67*int32(28), v75)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	F_CopyAttributeOutText(m, l0, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L14
L28:
	;
	return
}
func F_to_tsquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1175), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}

package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XidInMVCCSnapshot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v7))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v147
L2:
	;
	if v19 != 0 {
		v147 = v3
		goto L1
	} else {
		goto L6
	}
L3:
	;
	v19 = base.B2i32(base.Ui32(l0) < base.Ui32(v7))
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = int32(base.Ui32(l0-v7) >> (uint(int32(31)) % 32))
	goto L2
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v21))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33 != 0 {
		v147 = int32(1)
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v33 = base.B2i32(base.Ui32(v21) <= base.Ui32(l0))
	goto L7
L9:
	;
	goto L10
L10:
	;
	v33 = base.B2i32(int32(0) <= l0-v21)
	goto L7
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+29)))
	if v35 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v34&int32(1) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	if v34&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v86 == v85 {
		v147 = v85
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v42 == int32(0) {
		v79 = l0
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v62 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v50 = v3
	goto L20
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45+v50<<(uint(int32(2))%32))))
	if l0 == v56 {
		v147 = int32(1)
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v79 = l0
	goto L15
L22:
	;
	v59 = v50 + int32(1)
	if v42 != v59 {
		v50 = v59
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	return int32(0)
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v66))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v62)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v78 != 0 {
		v147 = int32(0)
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v78 = base.B2i32(base.Ui32(v62) < base.Ui32(v66))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v78 = int32(base.Ui32(v62-v66) >> (uint(int32(31)) % 32))
	goto L26
L30:
	;
	v79 = v62
	goto L15
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v92 = int32(0)
	goto L32
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89+v92<<(uint(int32(2))%32))))
	v101 = base.B2i32(v79 == v100)
	if v79 == v100 {
		v147 = v101
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v147 = v101
	goto L1
L34:
	;
	v103 = v92 + int32(1)
	if v103 != v86 {
		v92 = v103
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v108 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L39
	}
L37:
	;
	v123 = l0
	goto L38
L38:
	;
	v125 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v126 == v125 {
		v147 = v125
		goto L1
	} else {
		goto L45
	}
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v110))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v108)) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v122 != 0 {
		v147 = int32(0)
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v122 = base.B2i32(base.Ui32(v108) < base.Ui32(v110))
	goto L40
L42:
	;
	goto L43
L43:
	;
	v122 = int32(base.Ui32(v108-v110) >> (uint(int32(31)) % 32))
	goto L40
L44:
	;
	v123 = v108
	goto L38
L45:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v132 = int32(0)
	goto L46
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129+v132<<(uint(int32(2))%32))))
	v141 = base.B2i32(v123 == v140)
	if v123 == v140 {
		v147 = v141
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v147 = v141
	goto L1
L48:
	;
	v143 = v132 + int32(1)
	if v143 != v126 {
		v132 = v143
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
}
func F_xmin_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = int32(48)
	v6 = l0 - v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v9 = l1 - v5
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v10))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v7)) == int32(0) {
		v22 = base.B2i32(base.Ui32(v7) < base.Ui32(v10))
	} else {
		v22 = int32(base.Ui32(v7-v10) >> (uint(int32(31)) % 32))
	}
	if v22 != 0 {
		v39 = int32(1)
	} else {
		v23 = int32(0)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v25))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v24)) == v23 {
			v37 = base.B2i32(base.Ui32(v25) < base.Ui32(v24))
		} else {
			v37 = base.B2i32(int32(0) < v24-v25)
		}
		v39 = v23 - v37
	}
	return v39
}
func F_xmlconcat(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(363269), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(578929), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497496), int32(609), int32(112449))
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_xmltotext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}

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
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v3 = int32(0)
	v7 = int32(3)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(l0) < base.Ui32(v7))|base.B2i32(base.Ui32(v9) < base.Ui32(v7)) == v3 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v155
L2:
	;
	v132 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v133 == v132 {
		v155 = v132
		goto L1
	} else {
		goto L47
	}
L3:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v107 == v106 {
		v155 = v106
		goto L1
	} else {
		goto L42
	}
L4:
	;
	v155 = int32(0)
	goto L1
L5:
	;
	v19 = int32(3)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(l0) < base.Ui32(v19))|base.B2i32(base.Ui32(v21) < base.Ui32(v19)) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if l0-v9 < int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(l0) < base.Ui32(v9) {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	goto L5
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+29)))
	if v36 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	if l0-v21 < int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(l0) < base.Ui32(v21) {
		goto L11
	} else {
		goto L16
	}
L15:
	;
	return int32(1)
L16:
	;
	return int32(1)
L17:
	;
	if v35&int32(1) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if v35&int32(1) == int32(0) {
		v103 = l0
		goto L3
	} else {
		goto L35
	}
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v43 == int32(0) {
		v126 = l0
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v63 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v51 = int32(0)
	goto L24
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46+v51<<(uint(int32(2))%32))))
	if l0 == v58 {
		v155 = int32(1)
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v126 = l0
	goto L2
L26:
	;
	v61 = v51 + int32(1)
	if v43 != v61 {
		v51 = v61
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(0)
L29:
	;
	v67 = int32(3)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(v63) < base.Ui32(v67))|base.B2i32(base.Ui32(v69) < base.Ui32(v67)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v75 = int32(0)
	if v63-v69 < v75 {
		v155 = v75
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(v69) <= base.Ui32(v63) {
		v126 = v63
		goto L2
	} else {
		goto L34
	}
L33:
	;
	v126 = v63
	goto L2
L34:
	;
	goto L4
L35:
	;
	v84 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v86 = int32(3)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(v84) < base.Ui32(v86))|base.B2i32(base.Ui32(v88) < base.Ui32(v86)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v94 = int32(0)
	if v84-v88 < v94 {
		v155 = v94
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v88) <= base.Ui32(v84) {
		v103 = v84
		goto L3
	} else {
		goto L41
	}
L40:
	;
	v103 = v84
	goto L3
L41:
	;
	goto L4
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v113 = int32(0)
	goto L43
L43:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v110+v113<<(uint(int32(2))%32))))
	v122 = base.B2i32(v103 == v121)
	if v103 == v121 {
		v155 = v122
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v155 = v122
	goto L1
L45:
	;
	v124 = v113 + int32(1)
	if v124 != v107 {
		v113 = v124
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v139 = int32(0)
	goto L48
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v136+v139<<(uint(int32(2))%32))))
	v148 = base.B2i32(v126 == v147)
	if v126 == v147 {
		v155 = v148
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v155 = v148
	goto L1
L50:
	;
	v150 = v139 + int32(1)
	if v150 != v133 {
		v139 = v150
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
}
func F_xmin_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v28 int32
	_ = v28
	v4 = int32(0)
	v5 = int32(48)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0-v5)))
	v8 = int32(3)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1-v5)))
	if base.B2i32(base.Ui32(v7) < base.Ui32(v8))|base.B2i32(base.Ui32(v12) < base.Ui32(v8)) == v4 {
		if int32(0) <= v7-v12 {
			v28 = base.B2i32(v7 != v12)
			return v4 - v28
		} else {
			return int32(1)
		}
	} else {
		if base.Ui32(v7) < base.Ui32(v12) {
			return int32(1)
		} else {
			v28 = base.B2i32(base.Ui32(v12) < base.Ui32(v7))
			return v4 - v28
		}
	}
}
func F_xmlconcat(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14412(m, int32(_a_F_xmlconcat_0), int32(631))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_xmltotext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}

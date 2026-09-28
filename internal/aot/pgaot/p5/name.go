package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NameOfDatum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		v9 = v3
		return v9
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v5 = F_NameListToString(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = v5
			return v9
		}
	}
}
func F_choose_plan_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v193
L2:
	;
	v184 = F_pstrdup(m, l1)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L21
	} else {
		goto L48
	}
L3:
	;
	v79 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	v86 = F_psprintf(m, int32(_a_F_choose_plan_name_0), v12+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v20 = int32(0)
	if v20 < v17 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v24 = v17
	goto L9
L8:
	;
	v24 = v20
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v28 = v20
	goto L10
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25+v28<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v41 == int32(0))|base.B2i32(v41 != v44) != 0 {
		v62 = v41
		v63 = v44
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L2
L12:
	;
	if v62-v63 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L13:
	;
	goto L12
L14:
	;
	v47 = v38
	v48 = l1
	goto L15
L15:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v52
		v63 = v51
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v62 = v52
	v63 = v51
	goto L13
L17:
	;
	v55 = int32(1)
	if v52 == v51 {
		v47 = v47 + v55
		v48 = v48 + v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v68 = v28 + int32(1)
	if v24 != v68 {
		v28 = v68
		goto L10
	} else {
		goto L20
	}
L20:
	;
	goto L11
L21:
	;
	return int32(0)
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v90 == int32(0) {
		v166 = v86
		v167 = v90
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v172 = F_lappend(m, v167, v166)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L21
	} else {
		goto L47
	}
L24:
	;
	v96 = v86
	v97 = v90
	v99 = v79
	goto L25
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v102 <= int32(0) {
		v166 = v96
		v167 = v97
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v166 = v160
	v167 = v162
	goto L23
L27:
	;
	v105 = int32(0)
	if v105 < v102 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v108 = v102
	goto L30
L29:
	;
	v108 = v105
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v113 = int32(0)
	goto L31
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v109+v113<<(uint(int32(2))%32))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if base.B2i32(v126 == int32(0))|base.B2i32(v126 != v129) != 0 {
		v147 = v126
		v148 = v129
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_pfree(m, v96)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L21
	} else {
		goto L44
	}
L33:
	;
	if v147-v148 != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	goto L33
L35:
	;
	v132 = v123
	v133 = v96
	goto L36
L36:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v137 == int32(0) {
		v147 = v137
		v148 = v136
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v147 = v137
	v148 = v136
	goto L34
L38:
	;
	v140 = int32(1)
	if v137 == v136 {
		v132 = v132 + v140
		v133 = v133 + v140
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v151 = v113 + int32(1)
	if v108 != v151 {
		v113 = v151
		goto L31
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L32
L43:
	;
	v166 = v96
	v167 = v97
	goto L23
L44:
	;
	v156 = v99 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	v160 = F_psprintf(m, int32(_a_F_choose_plan_name_0), v12)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v162 != 0 {
		v96 = v160
		v97 = v162
		v99 = v156
		goto L25
	} else {
		goto L46
	}
L46:
	;
	goto L26
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v172
	v193 = v166
	goto L1
L48:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v187 = F_lappend(m, v186, v184)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v187
	v193 = v184
	goto L1
}

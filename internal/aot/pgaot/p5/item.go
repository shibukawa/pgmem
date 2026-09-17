package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ItemPointerEquals(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v5 = int32(16)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v3|v4<<(uint(v5)%32) == v8|v9<<(uint(v5)%32) {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v15 == v16 {
			v19 = int32(1)
		} else {
			v19 = int32(0)
		}
	} else {
		v19 = int32(0)
	}
	return v19
}
func F_makeItemLikeRegex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(192)
	m.G0 = v12
	v15 = F_palloc(m, int32(24))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_makeItemLikeRegex[0]))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(42)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v28
	if l2 == v29 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v12 + int32(192)
	return v205
L8:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v164 = F_palloc(m, v159<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L45
	}
L9:
	;
	v153 = int32(67)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(0)
	v47 = v6
	goto L15
L13:
	;
	v108 = v6
	goto L14
L14:
	;
	if v108&int32(1) != 0 {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v39))))
	switch v51 - int32(105) {
	case 0:
		v93 = int32(1)
		goto L17
	default:
		goto L19
	case 4:
		goto L22
	case 8:
		goto L20
	case 10:
		goto L18
	case 15:
		goto L21
	}
L16:
	;
	v108 = v94
	goto L14
L17:
	;
	v94 = v93 | v47
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v94
	v97 = v39 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v97 < v98 {
		v39 = v97
		v47 = v94
		goto L15
	} else {
		goto L30
	}
L18:
	;
	v93 = int32(2)
	goto L17
L19:
	;
	v57 = int32(0)
	v58 = F_errsave_start(m, l4)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	v93 = int32(16)
	goto L17
L21:
	;
	v93 = int32(8)
	goto L17
L22:
	;
	v93 = int32(4)
	goto L17
L23:
	;
	if v58 == int32(0) {
		v205 = v57
		goto L7
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(_a_F_makeItemLikeRegex_0)
	F_errmsg(m, int32(_a_F_makeItemLikeRegex_1), v12+int32(32))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v76 = F_pg_mblen_range(m, v72+v39, v72+v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v39 + v78
	F_errdetail(m, int32(_a_F_makeItemLikeRegex_2), v12+int32(16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, l4, int32(_a_F_makeItemLikeRegex_3), int32(603), int32(_a_F_makeItemLikeRegex_4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v205 = v57
	goto L7
L30:
	;
	goto L16
L31:
	;
	v113 = int32(11)
	goto L33
L32:
	;
	v113 = int32(3)
	goto L33
L33:
	;
	if v108&int32(16) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v153 = v113&int32(8) | int32(4)
	goto L8
L35:
	;
	goto L36
L36:
	;
	if v108&int32(8) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v153 = v113 | v108<<(uint(int32(5))%32)&int32(192) ^ int32(64)
	goto L8
L38:
	;
	goto L39
L39:
	;
	v131 = int32(0)
	v132 = F_errsave_start(m, l4)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v132 == int32(0) {
		v205 = v131
		goto L7
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_makeItemLikeRegex_5), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, l4, int32(_a_F_makeItemLikeRegex_3), int32(680), int32(_a_F_makeItemLikeRegex_6))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v205 = v131
	goto L7
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v168 = F_pg_mb2wchar_with_len(m, v166, v164, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v171 = F_pg_regcomp(m, v12+int32(160), v164, v168, v153, int32(100))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v171 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v174 = v12 + int32(48)
	F_pg_regerror(m, v171, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_pg_regfree(m, v12+int32(160))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L57
	}
L51:
	;
	v177 = int32(0)
	v178 = F_errsave_start(m, l4)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v178 == int32(0) {
		v205 = v177
		goto L7
	} else {
		goto L53
	}
L53:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v174
	F_errmsg(m, int32(_a_F_makeItemLikeRegex_7), v12)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errsave_finish(m, l4, int32(_a_F_makeItemLikeRegex_3), int32(632), int32(_a_F_makeItemLikeRegex_4))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v205 = v177
	goto L7
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
	v205 = int32(1)
	goto L7
}

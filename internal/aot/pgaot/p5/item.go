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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v107 int32
	_ = v107
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
	var v175 int32
	_ = v175
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
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
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
	v20 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
	return v207
L8:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v164 = F_palloc(m, v159<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L46
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
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v36 <= v35 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v107&int32(1) != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	v107 = v6
	goto L12
L14:
	;
	goto L15
L15:
	;
	v39 = v35
	v46 = v6
	goto L16
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v39))))
	switch v51 - int32(105) {
	case 0:
		v93 = int32(1)
		goto L18
	default:
		goto L20
	case 4:
		goto L23
	case 8:
		goto L21
	case 10:
		goto L19
	case 15:
		goto L22
	}
L17:
	;
	v107 = v94
	goto L12
L18:
	;
	v94 = v93 | v46
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v94
	v97 = v39 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v97 < v98 {
		v39 = v97
		v46 = v94
		goto L16
	} else {
		goto L31
	}
L19:
	;
	v93 = int32(2)
	goto L18
L20:
	;
	v57 = int32(0)
	v58 = F_errsave_start(m, l4)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v93 = int32(16)
	goto L18
L22:
	;
	v93 = int32(8)
	goto L18
L23:
	;
	v93 = int32(4)
	goto L18
L24:
	;
	if v58 == int32(0) {
		v207 = v57
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(320382)
	F_errmsg(m, int32(188787), v12+int32(32))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v76 = F_pg_mblen_range(m, v72+v39, v72+v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v39 + v78
	F_errdetail(m, int32(614522), v12+int32(16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errsave_finish(m, l4, int32(26956), int32(603), int32(27917))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v207 = v57
	goto L7
L31:
	;
	goto L17
L32:
	;
	v113 = int32(11)
	goto L34
L33:
	;
	v113 = int32(3)
	goto L34
L34:
	;
	if v107&int32(16) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v153 = v113&int32(8) | int32(4)
	goto L8
L36:
	;
	goto L37
L37:
	;
	if v107&int32(8) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v153 = v113 | v107<<(uint(int32(5))%32)&int32(192) ^ int32(64)
	goto L8
L39:
	;
	goto L40
L40:
	;
	v131 = int32(0)
	v132 = F_errsave_start(m, l4)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v132 == int32(0) {
		v207 = v131
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(444478), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, l4, int32(26956), int32(680), int32(156532))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v207 = v131
	goto L7
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v168 = F_pg_mb2wchar_with_len(m, v166, v164, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v171 = F_pg_regcomp(m, v12+int32(160), v164, v168, v153, int32(100))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v171 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v175 = F_pg_regerror(m, v171, v12+int32(48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_pg_regfree(m, v12+int32(160))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L58
	}
L52:
	;
	v177 = int32(0)
	v178 = F_errsave_start(m, l4)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v178 == int32(0) {
		v207 = v177
		goto L7
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(48)
	F_errmsg(m, int32(201563), v12)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errsave_finish(m, l4, int32(26956), int32(632), int32(27917))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v207 = v177
	goto L7
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
	v207 = int32(1)
	goto L7
}

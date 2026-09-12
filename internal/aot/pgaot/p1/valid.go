package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_valid_polymorphic_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 <= int32(3830) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v149
L2:
	;
	v137 = F_format_type_be(m, l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L48
	} else {
		goto L49
	}
L3:
	;
	if base.Ui32(int32(2)) < base.Ui32(l0-int32(5077)) {
		v149 = v4
		goto L1
	} else {
		goto L41
	}
L4:
	;
	switch l0 - int32(2277) {
	case 0, 6:
		goto L7
	case 1, 2, 3, 4, 5:
		goto L3
	default:
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	switch l0 - int32(4537) {
	case 0:
		goto L25
	case 1:
		goto L24
	default:
		goto L26
	}
L7:
	;
	v21 = int32(655265)
	if l2 <= int32(0) {
		v135 = v21
		goto L2
	} else {
		goto L11
	}
L8:
	;
	if l0 == int32(2776) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l0 != int32(3500) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v28 = v4
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1+v28<<(uint(int32(2))%32))))
	if v35 <= int32(3499) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v135 = v21
	goto L2
L14:
	;
	v49 = v28 + int32(1)
	if v49 != l2 {
		v28 = v49
		goto L12
	} else {
		goto L23
	}
L15:
	;
	switch v35 - int32(2277) {
	case 0, 6:
		v149 = v4
		goto L1
	case 1, 2, 3, 4, 5:
		goto L14
	default:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v35 == int32(3500) {
		v149 = v4
		goto L1
	} else {
		goto L20
	}
L18:
	;
	if v35 != int32(2776) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v149 = v4
	goto L1
L20:
	;
	if v35 == int32(3831) {
		v149 = v4
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v35 == int32(4537) {
		v149 = v4
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	goto L13
L24:
	;
	v79 = int32(655393)
	if l2 <= int32(0) {
		v135 = v79
		goto L2
	} else {
		goto L35
	}
L25:
	;
	v57 = int32(655182)
	if l2 <= int32(0) {
		v135 = v57
		goto L2
	} else {
		goto L29
	}
L26:
	;
	if l0 == int32(5080) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if l0 != int32(3831) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v63 = v4
	goto L30
L30:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1+v63<<(uint(int32(2))%32))))
	if v71 == int32(3831) {
		v149 = v4
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v135 = v57
	goto L2
L32:
	;
	if v71 == int32(4537) {
		v149 = v4
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v77 = v63 + int32(1)
	if v77 != l2 {
		v63 = v77
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v85 = v4
	goto L36
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1+v85<<(uint(int32(2))%32))))
	if v93 == int32(4538) {
		v149 = v4
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v135 = v79
	goto L2
L38:
	;
	if v93 == int32(5080) {
		v149 = v4
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v99 = v85 + int32(1)
	if v99 != l2 {
		v85 = v99
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v105 = int32(655496)
	if l2 <= int32(0) {
		v135 = v105
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v111 = v4
	goto L43
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1+v111<<(uint(int32(2))%32))))
	if base.Ui32(v119-int32(5077)) < base.Ui32(int32(4)) {
		v149 = v4
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v135 = v105
	goto L2
L45:
	;
	if v119 == int32(4538) {
		v149 = v4
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v127 = v111 + int32(1)
	if v127 != l2 {
		v111 = v127
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	return int32(0)
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v137
	v142 = F_psprintf(m, v135, v11)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v149 = v142
	goto L1
}
func F_is_valid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return (v3 ^ int32(-1)) & int32(1)
}
func F_make_valid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v6 = F_Int64GetDatum(m, v3&int64(-2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_validOperatorName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v2 = int32(0)
	v5 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(v5+int32(-64)) < base.Ui32(int32(-63)) {
		v242 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v242
L2:
	;
	v10 = int32(559153)
	v14 = m.G0
	v16 = v14 - int32(32)
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v17
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[328])))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v93 != v5 {
		v242 = v2
		goto L1
	} else {
		goto L24
	}
L4:
	;
	v93 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[329])))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = l0
	goto L10
L8:
	;
	goto L9
L9:
	;
	v43 = v10
	v44 = v25
	goto L13
L10:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v39 == v25 {
		v33 = v33 + int32(1)
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v93 = v33 - l0
	goto L3
L12:
	;
	goto L11
L13:
	;
	v51 = v16 + int32(base.Ui32(v44)>>(uint(int32(3))%32))&int32(28)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v52 | v53<<(uint(v44)%32)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v57 != 0 {
		v43 = v43 + v53
		v44 = v57
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v60 == int32(0) {
		v85 = l0
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v93 = v85 - l0
	goto L3
L17:
	;
	v64 = l0
	v65 = v60
	goto L18
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v65)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v73)>>(uint(v65)%32))&int32(1) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v85 = v81
	goto L16
L20:
	;
	v85 = v64
	goto L16
L21:
	;
	goto L22
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v81 = v64 + int32(1)
	if v79 != 0 {
		v64 = v81
		v65 = v79
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v96 = F_strstr(m, l0, int32(684082))
	mBase = m.M
	if v96 != 0 {
		v242 = v2
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v98 = F_strstr(m, l0, int32(683983))
	mBase = m.M
	if v98 != 0 {
		v242 = v2
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v5) < base.Ui32(int32(2)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v230 = int32(1)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v231 != int32(33) {
		v242 = v230
		goto L1
	} else {
		goto L60
	}
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v5-int32(1)))))
	switch v104 - int32(43) {
	case 0, 2:
		goto L29
	default:
		goto L27
	}
L29:
	;
	v110 = v5 - int32(2)
	goto L30
L30:
	;
	v113 = int32(703908)
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v110))))
	v116 = int32(11)
	goto L35
L31:
	;
	v242 = v2
	goto L1
L32:
	;
	if v219 != 0 {
		goto L27
	} else {
		goto L58
	}
L33:
	;
	v219 = int32(0)
	goto L32
L34:
	;
	v197 = v190
	v199 = v192
	goto L52
L35:
	;
	goto L43
L43:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[330])))
	if v153 == v115&int32(255) {
		v183 = v113
		v185 = v116
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v185 == int32(0) {
		goto L33
	} else {
		goto L51
	}
L45:
	;
	goto L46
L46:
	;
	v163 = v113
	v165 = v116
	goto L47
L47:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v170 = v169 ^ v115&int32(255)*int32(16843009)
	v173 = int32(-2139062144)
	if (int32(16843008)-v170|v170)&v173 != v173 {
		v190 = v163
		v192 = v165
		goto L34
	} else {
		goto L49
	}
L48:
	;
	v183 = v178
	v185 = v180
	goto L44
L49:
	;
	v177 = int32(4)
	v178 = v163 + v177
	v180 = v165 - v177
	if base.Ui32(int32(3)) < base.Ui32(v180) {
		v163 = v178
		v165 = v180
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v190 = v183
	v192 = v185
	goto L34
L52:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v115&int32(255) == v202 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L33
L54:
	;
	v219 = v197
	goto L32
L55:
	;
	goto L56
L56:
	;
	v204 = int32(1)
	v207 = v199 - v204
	if v207 != 0 {
		v197 = v197 + v204
		v199 = v207
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	v220 = int32(0)
	if base.B2i32(v110 <= v220) == v220 {
		v110 = v110 - int32(1)
		goto L30
	} else {
		goto L59
	}
L59:
	;
	goto L31
L60:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v234 != int32(61) {
		v242 = v230
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v242 = base.B2i32(v237 != int32(0))
	goto L1
}

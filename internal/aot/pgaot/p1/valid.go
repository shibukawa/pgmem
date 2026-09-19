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
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
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
	return v156
L2:
	;
	v142 = F_format_type_be(m, l0)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L43
	} else {
		goto L44
	}
L3:
	;
	if base.Ui32(int32(2)) < base.Ui32(l0-int32(_a_F_check_valid_polymorphic_signature_0)) {
		v156 = v4
		goto L1
	} else {
		goto L37
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
	switch l0 - int32(_a_F_check_valid_polymorphic_signature_1) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L24
	}
L7:
	;
	v21 = int32(_a_F_check_valid_polymorphic_signature_2)
	if l2 <= int32(0) {
		v139 = v21
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
	v139 = v21
	goto L2
L14:
	;
	v51 = v28 + int32(1)
	if v51 != l2 {
		v28 = v51
		goto L12
	} else {
		goto L21
	}
L15:
	;
	switch v35 - int32(2277) {
	case 0, 6:
		v156 = v4
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
	if base.B2i32(v35 == int32(3500))|base.B2i32(v35 == int32(3831))|base.B2i32(v35 == int32(_a_F_check_valid_polymorphic_signature_1)) != 0 {
		v156 = v4
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
	v156 = v4
	goto L1
L20:
	;
	goto L14
L21:
	;
	goto L13
L22:
	;
	v82 = int32(_a_F_check_valid_polymorphic_signature_3)
	if l2 <= int32(0) {
		v139 = v82
		goto L2
	} else {
		goto L32
	}
L23:
	;
	v59 = int32(_a_F_check_valid_polymorphic_signature_4)
	if l2 <= int32(0) {
		v139 = v59
		goto L2
	} else {
		goto L27
	}
L24:
	;
	if l0 == int32(_a_F_check_valid_polymorphic_signature_5) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if l0 != int32(3831) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v65 = v4
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1+v65<<(uint(int32(2))%32))))
	if base.B2i32(v73 == int32(3831))|base.B2i32(v73 == int32(_a_F_check_valid_polymorphic_signature_1)) != 0 {
		v156 = v4
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v139 = v59
	goto L2
L30:
	;
	v80 = v65 + int32(1)
	if v80 != l2 {
		v65 = v80
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v88 = v4
	goto L33
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1+v88<<(uint(int32(2))%32))))
	if base.B2i32(v96 == int32(_a_F_check_valid_polymorphic_signature_6))|base.B2i32(v96 == int32(_a_F_check_valid_polymorphic_signature_5)) != 0 {
		v156 = v4
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v139 = v82
	goto L2
L35:
	;
	v103 = v88 + int32(1)
	if v103 != l2 {
		v88 = v103
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v109 = int32(_a_F_check_valid_polymorphic_signature_7)
	if l2 <= int32(0) {
		v139 = v109
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v115 = v4
	goto L39
L39:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v115<<(uint(int32(2))%32))))
	if base.B2i32(base.Ui32(v123-int32(_a_F_check_valid_polymorphic_signature_0)) < base.Ui32(int32(4)))|base.B2i32(v123 == int32(_a_F_check_valid_polymorphic_signature_6)) != 0 {
		v156 = v4
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v139 = v109
	goto L2
L41:
	;
	v132 = v115 + int32(1)
	if v132 != l2 {
		v115 = v132
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	return int32(0)
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v142
	v147 = F_psprintf(m, v139, v11)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v156 = v147
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
	var v83 int32
	_ = v83
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v2 = int32(0)
	v5 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(v5+int32(-64)) < base.Ui32(int32(-63)) {
		v244 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v244
L2:
	;
	v10 = int32(_a_F_validOperatorName_0)
	v14 = m.G0
	v16 = v14 - int32(32)
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v17
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validOperatorName[0])))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v93 != v5 {
		v244 = v2
		goto L1
	} else {
		goto L22
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
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validOperatorName[1])))
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
		v83 = l0
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v93 = v83 - l0
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
		v83 = v64
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v83 = v81
	goto L16
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v81 = v64 + int32(1)
	if v79 != 0 {
		v64 = v81
		v65 = v79
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v96 = F_strstr(m, l0, int32(_a_F_validOperatorName_1))
	mBase = m.M
	if v96 != 0 {
		v244 = v2
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v98 = F_strstr(m, l0, int32(_a_F_validOperatorName_2))
	mBase = m.M
	if v98 != 0 {
		v244 = v2
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v5) < base.Ui32(int32(2)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v232 = int32(1)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v233 != int32(33) {
		v244 = v232
		goto L1
	} else {
		goto L57
	}
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v5-int32(1)))))
	switch v104 - int32(43) {
	case 0, 2:
		goto L27
	default:
		goto L25
	}
L27:
	;
	v110 = v5 - int32(2)
	goto L28
L28:
	;
	v113 = int32(_a_F_validOperatorName_3)
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v110))))
	v116 = int32(11)
	goto L33
L29:
	;
	v244 = v2
	goto L1
L30:
	;
	if v221 != 0 {
		goto L25
	} else {
		goto L55
	}
L31:
	;
	v221 = int32(0)
	goto L30
L32:
	;
	v199 = v192
	v201 = v194
	goto L49
L33:
	;
	goto L40
L40:
	;
	v155 = v115 & int32(255)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validOperatorName[2])))
	if base.B2i32(v155 == v156)|int32(0) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v165 = v113
	v167 = v116
	goto L44
L42:
	;
	v185 = v113
	v187 = v116
	goto L43
L43:
	;
	if v187 == int32(0) {
		goto L31
	} else {
		goto L48
	}
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v172 = v171 ^ v155*int32(16843009)
	v175 = int32(-2139062144)
	if (int32(16843008)-v172|v172)&v175 != v175 {
		v192 = v165
		v194 = v167
		goto L32
	} else {
		goto L46
	}
L45:
	;
	v185 = v180
	v187 = v182
	goto L43
L46:
	;
	v179 = int32(4)
	v180 = v165 + v179
	v182 = v167 - v179
	if base.Ui32(int32(3)) < base.Ui32(v182) {
		v165 = v180
		v167 = v182
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v192 = v185
	v194 = v187
	goto L32
L49:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v115&int32(255) == v204 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L31
L51:
	;
	v221 = v199
	goto L30
L52:
	;
	goto L53
L53:
	;
	v206 = int32(1)
	v209 = v201 - v206
	if v209 != 0 {
		v199 = v199 + v206
		v201 = v209
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v222 = int32(0)
	if base.B2i32(v110 <= v222) == v222 {
		v110 = v110 - int32(1)
		goto L28
	} else {
		goto L56
	}
L56:
	;
	goto L29
L57:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v236 != int32(61) {
		v244 = v232
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v244 = base.B2i32(v239 != int32(0))
	goto L1
}

package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_relation_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = F_try_relation_open(m, v5, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v22 = int32(274282)
	v23 = F_text_to_cstring(m, v7)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v160 = F_calculate_relation_size(m, v12, v20, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L50
	}
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v28 == int32(0) {
		v47 = v27
		v48 = v28
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v48-v47 == int32(0) {
		v159 = int32(0)
		goto L7
	} else {
		goto L17
	}
L10:
	;
	goto L9
L11:
	;
	if v27 != v28 {
		v47 = v27
		v48 = v28
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = v22
	v33 = v23
	goto L13
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v37 == int32(0) {
		v47 = v36
		v48 = v37
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v47 = v36
	v48 = v37
	goto L10
L15:
	;
	v40 = int32(1)
	if v36 == v37 {
		v32 = v32 + v40
		v33 = v33 + v40
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v53 = int32(283400)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v57 == int32(0) {
		v76 = v56
		v77 = v57
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v77-v76 == int32(0) {
		v159 = int32(1)
		goto L7
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	if v56 != v57 {
		v76 = v56
		v77 = v57
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v61 = v53
	v62 = v23
	goto L22
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v65
		v77 = v66
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v76 = v65
	v77 = v66
	goto L19
L24:
	;
	v69 = int32(1)
	if v65 == v66 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v82 = int32(281414)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1071])))
	if v86 == int32(0) {
		v105 = v85
		v106 = v86
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v106-v105 == int32(0) {
		v159 = int32(2)
		goto L7
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	if v85 != v86 {
		v105 = v85
		v106 = v86
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v90 = v82
	v91 = v23
	goto L31
L31:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v95 == int32(0) {
		v105 = v94
		v106 = v95
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v105 = v94
	v106 = v95
	goto L28
L33:
	;
	v98 = int32(1)
	if v94 == v95 {
		v90 = v90 + v98
		v91 = v91 + v98
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v111 = int32(99007)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _consts[981])))
	if v115 == int32(0) {
		v134 = v114
		v135 = v115
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v135-v134 == int32(0) {
		v159 = int32(3)
		goto L7
	} else {
		goto L44
	}
L37:
	;
	goto L36
L38:
	;
	if v114 != v115 {
		v134 = v114
		v135 = v115
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v119 = v111
	v120 = v23
	goto L40
L40:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v124 == int32(0) {
		v134 = v123
		v135 = v124
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v134 = v123
	v135 = v124
	goto L37
L42:
	;
	v127 = int32(1)
	if v123 == v124 {
		v119 = v119 + v127
		v120 = v120 + v127
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(375690), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errhint(m, int32(639547), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(490209), int32(63), int32(224930))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_relation_close(m, v12, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v165 = F_Int64GetDatum(m, v160)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	return v165
}

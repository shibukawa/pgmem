package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_define_custom_variable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[956]))
	v19 = F_hash_search(m, v14, v9+int32(8), v2, v2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L85
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L81
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return
L4:
	;
	return
L5:
	;
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_InitializeOneGUCOption(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+21)))
	if v36&int32(2) == int32(0) {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[956]))
	v30 = F_hash_search(m, v26, l0, int32(3), v9+int32(15))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v30 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = l0
	goto L3
L12:
	;
	F_InitializeOneGUCOption(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	goto L16
L15:
	;
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v58 = int32(4534316)
	goto L22
L18:
	;
	goto L19
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+28)))
	if v67&int32(4) != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	goto L20
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v61 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v65
	goto L21
L24:
	;
	if v61 != v35+int32(72) {
		v58 = v61
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v74 = int32(4534308)
	goto L31
L27:
	;
	goto L28
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v35)+112))
	if v83 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	goto L29
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v77 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v81
	goto L30
L33:
	;
	if v77 != v35+int32(76) {
		v74 = v77
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v93 = F_set_config_with_handle(m, v84, v85, v83, v86, v87, v88, v85, int32(1), int32(19), v85)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	F_reapply_stacked_values(m, l0, v35, v95, v97, v98, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v35)+84))
	if v103 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v131 = v35 + int32(56)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v134
	if v133 == v134 {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v113 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(12)
	goto L44
L43:
	;
	v114 = int32(15)
	goto L44
L44:
	;
	v115 = F_find_option(m, v107, int32(1), int32(0), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v115 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v119 = F_guc_strdup(m, v114, v103)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+84))
	if v121 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_pfree(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+88)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v115)+84)) = v119
	goto L40
L51:
	;
	goto L50
L52:
	;
	v171 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v171
	if v167 == v171 {
		goto L66
	} else {
		goto L67
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v35)+112))
	v167 = v164
	goto L52
L54:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v133 == v139 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v35)+112))
	if v133 == v141 {
		v167 = v141
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
	if v133 == v143 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v147 = v131
	goto L58
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v151 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	F_pfree(m, v133)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L65
	}
L60:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+32))
	if v133 == v152 {
		goto L53
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+48))
	if v133 != v154 {
		v147 = v151
		goto L58
	} else {
		goto L64
	}
L64:
	;
	goto L53
L65:
	;
	goto L53
L66:
	;
	F_pfree(m, v35)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L80
	}
L67:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v167 == v176 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
	if v167 == v178 {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v180 = v131
	goto L71
L70:
	;
	if v35 == int32(0) {
		goto L3
	} else {
		goto L79
	}
L71:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v186 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	F_pfree(m, v167)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L78
	}
L73:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+32))
	if v167 == v187 {
		goto L70
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L72
L76:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+48))
	if v167 != v189 {
		v180 = v186
		goto L71
	} else {
		goto L77
	}
L77:
	;
	goto L70
L78:
	;
	goto L70
L79:
	;
	goto L66
L80:
	;
	goto L3
L81:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(13961), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(512044), int32(1060), int32(404921))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v235
	F_errmsg(m, int32(714996), v9)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(512044), int32(4972), int32(404761))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}

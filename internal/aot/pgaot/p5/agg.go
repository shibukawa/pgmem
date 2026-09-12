package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_finalize_agg_primnode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v4 == int32(9) {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v8 = F_finalize_primnode(m, v7, l1)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v13 = F_finalize_primnode(m, v12, l1)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			v18 = F_expression_tree_walker_impl(m, l0, int32(849), l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = v18
				return v20
			}
		}
	} else {
		v20 = int32(0)
		return v20
	}
}
func F_lookup_agg_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v27 = F_func_get_detail(m, l0, v6, v6, l1, l2, v6, v6, v6, v9+int32(76), l4, v9+int32(75), v9+int32(68), v9-int32(-64), v9+int32(60), v6)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L44
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L39
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L34
	}
L4:
	;
	return int32(0)
L5:
	;
	if v27 != int32(2) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	if v33 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+75)))
	if v36 == int32(1) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if l3 == int32(2276) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	if v41 != int32(2276) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v47 = F_enforce_generic_type_consistency(m, l2, v44, l1, v45, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v47
	v50 = int32(0)
	if l1 <= v50 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v102 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v104 = F_object_aclcheck(m, int32(1255), v100, v102, int64(128))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L28
	}
L15:
	;
	v57 = v50
	goto L16
L16:
	;
	v60 = v57 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2+v60)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63+v60)))
	v66 = F_IsBinaryCoercible(m, v62, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L23
	}
L18:
	;
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = v57 + int32(1)
	if l1 != v69 {
		v57 = v69
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L17
L22:
	;
	goto L14
L23:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v80 = F_func_signature_string(m, l0, l1, int32(0), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v80
	F_errmsg(m, int32(271492), v9+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(495341), int32(906), int32(249900))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	if v104 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v108 = F_get_func_name(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	m.G0 = v9 + int32(80)
	return v112
L32:
	;
	F_aclcheck_error(m, v104, int32(19), v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v125 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v125
	F_errmsg(m, int32(69423), v9+int32(48))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(495341), int32(861), int32(249900))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v146 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v146
	F_errmsg(m, int32(106275), v9)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(495341), int32(867), int32(249900))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v165 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v165
	F_errmsg(m, int32(352834), v9+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(495341), int32(882), int32(249900))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}

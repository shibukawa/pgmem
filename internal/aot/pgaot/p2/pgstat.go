package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_reset_all_cb(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v10 = v8 + int32(24)
	v12 = F_LWLockAcquire(m, v10, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = l0
	F_LWLockRelease(m, v10)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L10
L8:
	;
	goto L7
L9:
	;
	if v24&int32(1) != 0 {
		goto L3
	} else {
		goto L13
	}
L10:
	;
	v30 = F__emscripten_memcpy_bulkmem(m, v8+int32(184), v8+int32(48), int32(136))
	mBase = m.M
	goto L12
L12:
	;
	goto L9
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v24 != v34 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L4
L15:
	;
	return
}
func F_pgstat_bestart_final(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v5 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	if base.Ui32(int32(6)) < base.Ui32(v8) {
		v20 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(4449876)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v24 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v23 + v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v27 + v24
	v32 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v5)+52)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v5)+208)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v27 + int32(2)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v43 - v24
	v48 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	goto L4
L2:
	;
	v11 = int32(0)
	if int32(1)<<(uint(v8)%32)&int32(98) == v11 {
		v20 = v11
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1004]))
	v20 = v19
	goto L1
L4:
	;
	if int32(base.Ui32(int32(115186))>>(uint(v48)%32))&base.B2i32(base.Ui32(v48) < base.Ui32(int32(17))) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v55 = int32(0)
	v57 = int64(*(*int32)(unsafe.Add(mBase, _consts[743])))
	v59 = F_pgstat_get_entry_ref_locked(m, int32(6), v55, v57, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[687]))
	if v96 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	return
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v67 = F__emscripten_memset_bulkmem(m, v61+int32(24), base.I32_extend8_s(int32(0)), int32(2920))
	mBase = m.M
	goto L10
L10:
	;
	F_pgstat_unlock_entry(m, v59)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v74 = F__emscripten_memset_bulkmem(m, int32(4431896), base.I32_extend8_s(int32(0)), int32(2880))
	mBase = m.M
	goto L12
L12:
	;
	v77 = *(*int64)(unsafe.Add(mBase, _consts[43]))
	*(*int64)(unsafe.Add(mBase, _consts[1005])) = v77
	v81 = *(*int64)(unsafe.Add(mBase, _consts[41]))
	*(*int64)(unsafe.Add(mBase, _consts[1006])) = v81
	v85 = *(*int64)(unsafe.Add(mBase, _consts[39]))
	*(*int64)(unsafe.Add(mBase, _consts[1007])) = v85
	v89 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, _consts[1008])) = v89
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[245])) = uint8(v92)
	goto L7
L13:
	;
	return
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v100 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v96&int32(3) == int32(0) {
		v126 = v96
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v161 = F_pg_mbcliplen(m, v96, v159, int32(63))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L33
	}
L17:
	;
	v159 = v151 - v96
	goto L16
L18:
	;
	v130 = v126
	goto L27
L19:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v110 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v159 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v115 = v96
	goto L23
L23:
	;
	v119 = v115 + int32(1)
	if v119&int32(3) == int32(0) {
		v126 = v119
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v151 = v119
	goto L17
L25:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v124 != 0 {
		v115 = v119
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v139 = int32(-2139062144)
	if (int32(16843008)-v136|v136)&v139 == v139 {
		v130 = v130 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v145 = v130
	goto L30
L29:
	;
	goto L28
L30:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 != 0 {
		v145 = v145 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v151 = v145
	goto L17
L32:
	;
	goto L31
L33:
	;
	v163 = int32(4449876)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v166 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v165 + v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v169 + v166
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v100)+212))
	if v161 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v100)+212))
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v161+v176))) = uint8(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v181 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v180 + v181
	v184 = int32(4449876)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v186 - v181
	goto L13
L35:
	;
	v174 = F__emscripten_memcpy_bulkmem(m, v173, v96, v161)
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L34
}
func F_pgstat_checkpointer_snapshot_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	v16 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	goto L1
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+424))
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v47 = v16 + int32(408)
	v49 = F_LWLockAcquire(m, v47, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L14
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L9
L6:
	;
	return
L7:
	;
	goto L5
L8:
	;
	if v34&int32(1) != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v40 = F__emscripten_memcpy_bulkmem(m, int32(4379320), v16+int32(432), int32(88))
	mBase = m.M
	goto L11
L11:
	;
	goto L8
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)+424))
	if v34 != v44 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v16)+592))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v16)+584))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v16)+576))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v16)+568))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v16)+560))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v16)+552))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v16)+544))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v16)+536))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v16)+528))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v16)+520))
	F_LWLockRelease(m, v47)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v63 = int32(4379320)
	v65 = *(*int64)(unsafe.Add(mBase, _consts[1033]))
	*(*int64)(unsafe.Add(mBase, _consts[1033])) = v65 - v60
	v68 = int32(4379328)
	v70 = *(*int64)(unsafe.Add(mBase, _consts[1034]))
	*(*int64)(unsafe.Add(mBase, _consts[1034])) = v70 - v59
	v73 = int32(4379336)
	v75 = *(*int64)(unsafe.Add(mBase, _consts[1035]))
	*(*int64)(unsafe.Add(mBase, _consts[1035])) = v75 - v58
	v78 = int32(4379344)
	v80 = *(*int64)(unsafe.Add(mBase, _consts[1036]))
	*(*int64)(unsafe.Add(mBase, _consts[1036])) = v80 - v57
	v83 = int32(4379352)
	v85 = *(*int64)(unsafe.Add(mBase, _consts[1037]))
	*(*int64)(unsafe.Add(mBase, _consts[1037])) = v85 - v56
	v88 = int32(4379360)
	v90 = *(*int64)(unsafe.Add(mBase, _consts[1038]))
	*(*int64)(unsafe.Add(mBase, _consts[1038])) = v90 - v55
	v93 = int32(4379368)
	v95 = *(*int64)(unsafe.Add(mBase, _consts[1039]))
	*(*int64)(unsafe.Add(mBase, _consts[1039])) = v95 - v54
	v98 = int32(4379376)
	v100 = *(*int64)(unsafe.Add(mBase, _consts[1040]))
	*(*int64)(unsafe.Add(mBase, _consts[1040])) = v100 - v53
	v103 = int32(4379384)
	v105 = *(*int64)(unsafe.Add(mBase, _consts[1041]))
	*(*int64)(unsafe.Add(mBase, _consts[1041])) = v105 - v52
	v108 = int32(4379392)
	v110 = *(*int64)(unsafe.Add(mBase, _consts[1042]))
	*(*int64)(unsafe.Add(mBase, _consts[1042])) = v110 - v51
	return
}
func F_pgstat_count_heap_delete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v5 == int32(0) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
		if v8 != int32(1) {
			return
		} else {
			F_pgstat_assoc_relation(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v14 = v13
				v16 = *(*int32)(unsafe.Add(mBase, _consts[61]))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
					if v19 == v17 {
						v36 = v18
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
						return
					} else {
						v21 = F_pgstat_get_xact_stack_level(m, v17)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
							v26 = F_MemoryContextAllocZero(m, v24, int32(72))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
								*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
								*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
								*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
								v36 = v26
								v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
								return
							}
						}
					}
				} else {
					v21 = F_pgstat_get_xact_stack_level(m, v17)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
						v26 = F_MemoryContextAllocZero(m, v24, int32(72))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
							*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
							v36 = v26
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
							return
						}
					}
				}
			}
		}
	} else {
		v14 = v5
		v16 = *(*int32)(unsafe.Add(mBase, _consts[61]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
			if v19 == v17 {
				v36 = v18
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
				return
			} else {
				v21 = F_pgstat_get_xact_stack_level(m, v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
					v26 = F_MemoryContextAllocZero(m, v24, int32(72))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
						*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
						v36 = v26
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
						return
					}
				}
			}
		} else {
			v21 = F_pgstat_get_xact_stack_level(m, v17)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
				v26 = F_MemoryContextAllocZero(m, v24, int32(72))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v17
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v26
					v36 = v26
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v39 + int64(1)
					return
				}
			}
		}
	}
}
func F_pgstat_count_slru_page_hit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[241])) = uint8(v3)
	*(*uint8)(unsafe.Add(mBase, _consts[1053])) = uint8(v3)
	v9 = l0 << (uint(int32(6)) % 32)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1054])))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1054]))) = v12 + int64(1)
	return
}
func F_pgstat_create_transactional(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pgstat_get_entry_ref(m, l0, l1, l2, v4, v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v16 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 != 0 {
					if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
						v46 = l0*int32(72) + int32(1618272)
					} else {
						if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
							v44 = int32(0)
						} else {
							v32 = int32(0)
							v34 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
							if v34 == v32 {
								v44 = v32
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v34+l0<<(uint(int32(2))%32)-int32(96))))
								v44 = v42
							}
						}
						v46 = v44
					}
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v47
					F_errmsg(m, int32(36238), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(476714), int32(368), int32(301668))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_pgstat_reset(m, l0, l1, l2)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								F_create_drop_transactional_internal(m, l0, l1, l2, int32(1))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_pgstat_reset(m, l0, l1, l2)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_create_drop_transactional_internal(m, l0, l1, l2, int32(1))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			F_create_drop_transactional_internal(m, l0, l1, l2, int32(1))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_pgstat_drop_all_entries(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v81 int64
	_ = v81
	var v86 int64
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v225 int64
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	v7 = int64(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = v10 + int32(-44)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
	v18 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)) = uint8(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = v7
	goto L1
L1:
	;
	v29 = F_dshash_seq_next(m, v10+int32(-44))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v12 - int32(-64)
	return
L3:
	;
	return
L4:
	;
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_dshash_seq_term(m, v10+int32(-44))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v39 = v29
	v45 = v7
	goto L9
L8:
	;
	goto L2
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+16)))
	if v46 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_dshash_seq_term(m, v10+int32(-44))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L49
	}
L11:
	;
	goto L10
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
	if v50 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v215 = F_dshash_seq_next(m, v10+int32(-44))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L47
	}
L15:
	;
	v203 = F_pgstat_drop_entry_internal(m, v39, v10+int32(-44))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L44
	}
L16:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v59 = int64(23)
	v62 = int64(2388976653695081527)
	v63 = (v55 ^ int64(base.Ui64(v55)>>(uint(v59)%64))) * v62
	v64 = int64(47)
	v69 = int64(-8645972361240307355)
	v75 = (v53 ^ int64(base.Ui64(v53)>>(uint(v59)%64))) * v62
	v81 = ((v63^int64(base.Ui64(v63)>>(uint(v64)%64))^int64(-9208349263878056368))*v69 ^ int64(base.Ui64(v75)>>(uint(v64)%64)) ^ v75) * v69
	v86 = (int64(base.Ui64(v81)>>(uint(v59)%64)) ^ v81) * v62
	v94 = v58 & base.I32_wrap_i64(int64(base.Ui64(v86)>>(uint(v64)%64))^v86-int64(base.Ui64(v86)>>(uint(int64(32))%64)))
	v97 = v57 + v94*int32(24)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+16)))
	if v98 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v102 = v97
	v104 = v94
	goto L18
L18:
	;
	v111 = v10 + int32(-16)
	v112 = int32(16)
	goto L23
L19:
	;
	if v102 == int32(0) {
		goto L15
	} else {
		goto L42
	}
L20:
	;
	if v174 != 0 {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v174 = int32(0)
	goto L20
L22:
	;
	v148 = v143
	v149 = v144
	v150 = v145
	goto L32
L23:
	;
	if (v102|v111)&int32(3) != 0 {
		v143 = v102
		v144 = v111
		v145 = v112
		goto L22
	} else {
		goto L26
	}
L25:
	;
	if v133 == int32(0) {
		goto L21
	} else {
		goto L31
	}
L26:
	;
	v120 = v102
	v121 = v111
	v122 = v112
	goto L27
L27:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v125 != v126 {
		v143 = v120
		v144 = v121
		v145 = v122
		goto L22
	} else {
		goto L29
	}
L28:
	;
	goto L25
L29:
	;
	v128 = int32(4)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v143 = v131
	v144 = v129
	v145 = v133
	goto L22
L32:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 == v154 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v174 = v153 - v154
	goto L20
L34:
	;
	v156 = int32(1)
	v161 = v150 - v156
	if v161 != 0 {
		v148 = v148 + v156
		v149 = v149 + v156
		v150 = v161
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L21
L38:
	;
	v177 = (v104 + int32(1)) & v58
	v180 = v57 + v177*int32(24)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+16)))
	if v181 != 0 {
		v102 = v180
		v104 = v177
		goto L18
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L19
L41:
	;
	goto L15
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v102)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v185
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v187
	F_pgstat_release_entry_ref(m, v12, v184, int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	goto L15
L44:
	;
	v208 = v45 + base.I64_extend_i32_u(v203^int32(1))
	v211 = F_dshash_seq_next(m, v10+int32(-44))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	if v211 != 0 {
		v39 = v211
		v45 = v208
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v225 = v208
	goto L11
L47:
	;
	if v215 != 0 {
		v39 = v215
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v225 = v45
	goto L11
L49:
	;
	if v225 == int64(0) {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v233)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v233)+16)) = v234 + int64(1)
	goto L2
}
func F_pgstat_drop_entry(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v44 int64
	_ = v44
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int64
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int64
	_ = v380
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v182 = int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
	v188 = F_dshash_find(m, v184, v16+int32(32), v182)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L29
	} else {
		goto L32
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	v28 = int64(23)
	v31 = int64(2388976653695081527)
	v32 = (int64(base.Ui64(v27)>>(uint(v28)%64)) ^ v27) * v31
	v33 = int64(47)
	v38 = int64(-8645972361240307355)
	v44 = (int64(base.Ui64(l2)>>(uint(v28)%64)) ^ l2) * v31
	v50 = ((v32^int64(base.Ui64(v32)>>(uint(v33)%64))^int64(-9208349263878056368))*v38 ^ int64(base.Ui64(v44)>>(uint(v33)%64)) ^ v44) * v38
	v55 = (int64(base.Ui64(v50)>>(uint(v28)%64)) ^ v50) * v31
	v63 = v26 & base.I32_wrap_i64(int64(base.Ui64(v55)>>(uint(v33)%64))^v55-int64(base.Ui64(v55)>>(uint(int64(32))%64)))
	v66 = v25 + v63*int32(24)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+16)))
	if v67 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v70 = v63
	v71 = v66
	goto L4
L4:
	;
	v84 = v16 + int32(32)
	v85 = int32(16)
	goto L9
L5:
	;
	if v71 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L6:
	;
	if v147 != 0 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v147 = int32(0)
	goto L6
L8:
	;
	v121 = v116
	v122 = v117
	v123 = v118
	goto L18
L9:
	;
	if (v71|v84)&int32(3) != 0 {
		v116 = v71
		v117 = v84
		v118 = v85
		goto L8
	} else {
		goto L12
	}
L11:
	;
	if v106 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L12:
	;
	v93 = v71
	v94 = v84
	v95 = v85
	goto L13
L13:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v98 != v99 {
		v116 = v93
		v117 = v94
		v118 = v95
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L11
L15:
	;
	v101 = int32(4)
	v102 = v94 + v101
	v104 = v93 + v101
	v106 = v95 - v101
	if base.Ui32(int32(3)) < base.Ui32(v106) {
		v93 = v104
		v94 = v102
		v95 = v106
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v116 = v104
	v117 = v102
	v118 = v106
	goto L8
L18:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v126 == v127 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v147 = v126 - v127
	goto L6
L20:
	;
	v129 = int32(1)
	v134 = v123 - v129
	if v134 != 0 {
		v121 = v121 + v129
		v122 = v122 + v129
		v123 = v134
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	goto L7
L24:
	;
	v150 = (v70 + int32(1)) & v26
	v153 = v25 + v150*int32(24)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+16)))
	if v154 != 0 {
		v70 = v150
		v71 = v153
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L5
L27:
	;
	goto L1
L28:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v158
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v160
	F_pgstat_release_entry_ref(m, v16+int32(16), v157, int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L1
L31:
	;
	m.G0 = v16 + int32(80)
	return v397
L32:
	;
	if v188 == int32(0) {
		v397 = v182
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v193 = F_pgstat_drop_entry_internal(m, v188, int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v195 != int32(1) {
		v397 = v193
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
	if v200 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v319 = v16 + int32(52)
	v321 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
	v322 = int32(1)
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v319)+4)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v321
	*(*uint8)(unsafe.Add(mBase, uint32(v319)+24)) = uint8(v322)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v319)+12)) = v323
	goto L53
L37:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	if v204 == int64(0) {
		v237 = int32(-1)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v247 = v200
	v248 = v237
	v253 = int32(0)
	goto L44
L39:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v200)+20))
	v213 = int32(0)
	goto L40
L40:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v213*int32(24))+16)))
	if v225 != int32(1) {
		v237 = v213
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v237 = int32(-1)
	goto L38
L42:
	;
	v229 = v213 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v229)) < base.Ui64(v204) {
		v213 = v229
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v261 = v248
	v266 = v253
	v267 = v253
	goto L46
L46:
	;
	if v267&int32(1) != 0 {
		goto L36
	} else {
		goto L48
	}
L47:
	;
	if v287 == int32(0) {
		goto L36
	} else {
		goto L50
	}
L48:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v276 = int32(1)
	v277 = v261 - v276
	v281 = base.B2i32(v275&(v277^v237) == int32(0))
	v282 = v281 | v266
	v285 = v275 & v277
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	v287 = v261*int32(24) + v286
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+16)))
	if v288 != v276 {
		v261 = v285
		v266 = v282
		v267 = v281
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v293 != v198 {
		v248 = v285
		v253 = v282
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v287)+20))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v287)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v296
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v287)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v298
	F_pgstat_release_entry_ref(m, v16, v295, int32(1))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
	v247 = v304
	v248 = v285
	v253 = v282
	goto L44
L53:
	;
	v333 = F_dshash_seq_next(m, v16+int32(52))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	if v333 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v337 = v333
	v338 = int64(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_dshash_seq_term(m, v16+int32(52))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L29
	} else {
		goto L71
	}
L58:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+16)))
	if v349 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	F_dshash_seq_term(m, v16+int32(52))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L29
	} else {
		goto L69
	}
L60:
	;
	goto L59
L61:
	;
	v368 = F_dshash_seq_next(m, v16+int32(52))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L29
	} else {
		goto L67
	}
L62:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v352 != v198 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v356 = F_pgstat_drop_entry_internal(m, v337, v16+int32(52))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L29
	} else {
		goto L64
	}
L64:
	;
	v361 = v338 + base.I64_extend_i32_u(v356^int32(1))
	v364 = F_dshash_seq_next(m, v16+int32(52))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L29
	} else {
		goto L65
	}
L65:
	;
	if v364 != 0 {
		v337 = v364
		v338 = v361
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v371 = v361
	goto L60
L67:
	;
	if v368 != 0 {
		v337 = v368
		goto L58
	} else {
		goto L68
	}
L68:
	;
	v371 = v338
	goto L60
L69:
	;
	if v371 == int64(0) {
		v397 = v193
		goto L31
	} else {
		goto L70
	}
L70:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v379)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v379)+16)) = v380 + int64(1)
	v397 = v193
	goto L31
L71:
	;
	v397 = v193
	goto L31
}
func F_pgstat_end_function_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		F___clock_gettime(m, int32(1), v11)
		mBase = m.M
		v16 = int32(4434992)
		v17 = *(*int64)(unsafe.Add(mBase, _consts[428]))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v26 = v20 + v21*int64(1000000000) - v25
		*(*int64)(unsafe.Add(mBase, _consts[428])) = v19 + v26
		v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if l1 != 0 {
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
			*(*int64)(unsafe.Add(mBase, uint32(v13))) = v31 + int64(1)
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v29 + v26
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v36 + (v26 - v17 + v19)
	} else {
	}
	m.G0 = v11 + int32(16)
	return
}
func F_pgstat_execute_transactional_drops(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	v3 = int32(0)
	if l0 <= v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = v3
	v11 = v3
	goto L3
L3:
	;
	v15 = l1 + v11<<(uint(int32(4))%32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v19 = F_pgstat_drop_entry(m, v16, v17, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 <= int32(0) {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	v21 = int32(1)
	v23 = v10 + (v19 ^ v21)
	v25 = v11 + v21
	if v25 != l0 {
		v10 = v23
		v11 = v25
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v32 + int64(1)
	goto L9
L9:
	;
	goto L1
}
func F_pgstat_fetch_stat_bgwriter(m *base.Module) int32 {
	var v5 int32
	_ = v5
	F_pgstat_snapshot_fixed(m, int32(8))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(4379288)
	}
}
func F_pgstat_fetch_stat_dbentry(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_pgstat_fetch_entry(m, int32(1), l0, int64(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pgstat_fetch_stat_funcentry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v6 = F_pgstat_fetch_entry(m, int32(3), v4, base.I64_extend_i32_u(l0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_pgstat_get_slru_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	v2 = int32(227837)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1046])))
	if v6 == int32(0) {
		v25 = v5
		v26 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	goto L1
L3:
	;
	if v5 != v6 {
		v25 = v5
		v26 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v10 = v2
	v11 = l0
	goto L5
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 == int32(0) {
		v25 = v14
		v26 = v15
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v25 = v14
	v26 = v15
	goto L2
L7:
	;
	v18 = int32(1)
	if v14 == v15 {
		v10 = v10 + v18
		v11 = v11 + v18
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = int32(220359)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1047])))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v56-v55 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = v32
	v41 = l0
	goto L16
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v55 = v44
	v56 = v45
	goto L13
L18:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(1)
L21:
	;
	goto L22
L22:
	;
	v62 = int32(100134)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1048])))
	if v66 == int32(0) {
		v85 = v65
		v86 = v66
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v86-v85 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	goto L23
L25:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v70 = v62
	v71 = l0
	goto L27
L27:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v85 = v74
	v86 = v75
	goto L24
L29:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return int32(2)
L32:
	;
	goto L33
L33:
	;
	v92 = int32(19597)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1049])))
	if v96 == int32(0) {
		v115 = v95
		v116 = v96
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v116-v115 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	goto L34
L36:
	;
	if v95 != v96 {
		v115 = v95
		v116 = v96
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v100 = v92
	v101 = l0
	goto L38
L38:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v104
		v116 = v105
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v115 = v104
	v116 = v105
	goto L35
L40:
	;
	v108 = int32(1)
	if v104 == v105 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	return int32(3)
L43:
	;
	goto L44
L44:
	;
	v122 = int32(377389)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1050])))
	if v126 == int32(0) {
		v145 = v125
		v146 = v126
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v146-v145 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	goto L45
L47:
	;
	if v125 != v126 {
		v145 = v125
		v146 = v126
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v130 = v122
	v131 = l0
	goto L49
L49:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v135 == int32(0) {
		v145 = v134
		v146 = v135
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v145 = v134
	v146 = v135
	goto L46
L51:
	;
	v138 = int32(1)
	if v134 == v135 {
		v130 = v130 + v138
		v131 = v131 + v138
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	return int32(4)
L54:
	;
	goto L55
L55:
	;
	v152 = int32(246855)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1051])))
	if v156 == int32(0) {
		v175 = v155
		v176 = v156
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v176-v175 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	goto L56
L58:
	;
	if v155 != v156 {
		v175 = v155
		v176 = v156
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v160 = v152
	v161 = l0
	goto L60
L60:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	if v165 == int32(0) {
		v175 = v164
		v176 = v165
		goto L57
	} else {
		goto L62
	}
L61:
	;
	v175 = v164
	v176 = v165
	goto L57
L62:
	;
	v168 = int32(1)
	if v164 == v165 {
		v160 = v160 + v168
		v161 = v161 + v168
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	return int32(5)
L65:
	;
	goto L66
L66:
	;
	v184 = int32(248800)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1052])))
	if v188 == int32(0) {
		v207 = v187
		v208 = v188
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v208-v207 != 0 {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	goto L67
L69:
	;
	if v187 != v188 {
		v207 = v187
		v208 = v188
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v192 = v184
	v193 = l0
	goto L71
L71:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	if v197 == int32(0) {
		v207 = v196
		v208 = v197
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v207 = v196
	v208 = v197
	goto L68
L73:
	;
	v200 = int32(1)
	if v196 == v197 {
		v192 = v192 + v200
		v193 = v193 + v200
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v210 = int32(7)
	goto L77
L76:
	;
	v210 = int32(6)
	goto L77
L77:
	;
	return v210
}
func F_pgstat_io_reset_all_cb(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v8 = v6 + int32(608)
	v10 = F_LWLockAcquire(m, v8, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+896)) = l0
	v20 = F__emscripten_memset_bulkmem(m, v6+int32(904), base.I32_extend8_s(int32(0)), int32(2880))
	mBase = m.M
	goto L3
L3:
	;
	F_LWLockRelease(m, v8)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(1)
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v34 = v29 + v25<<(uint(int32(4))%32) + int32(608)
	v36 = F_LWLockAcquire(m, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v38 = int32(2880)
	v46 = F__emscripten_memset_bulkmem(m, v29+v25*v38+int32(904), base.I32_extend8_s(int32(0)), v38)
	mBase = m.M
	goto L8
L8:
	;
	F_LWLockRelease(m, v34)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v50 = v25 + int32(1)
	if v50 != int32(18) {
		v25 = v50
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
}
func F_pgstat_progress_start_command(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
		if v9 != int32(1) {
		} else {
			v12 = int32(4449876)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			v15 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v14 + v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v18 + v15
			*(*int32)(unsafe.Add(mBase, uint32(v5)+220)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v5)+224)) = l1
			v25 = v5 + int32(232)
			if v25&int32(3) == int32(0) {
				v31 = v5 + int32(392)
				if base.Ui32(v31) <= base.Ui32(v25) {
				} else {
					v35 = v5 + int32(236)
					if base.Ui32(v35) < base.Ui32(v31) {
						v37 = v31
					} else {
						v37 = v35
					}
					v46 = F__emscripten_memset_bulkmem(m, v25, base.I32_extend8_s(int32(0)), (v37-v5-int32(233))&int32(-4)+int32(4))
					mBase = m.M
				}
			} else {
				v50 = F__emscripten_memset_bulkmem(m, v25, base.I32_extend8_s(int32(0)), int32(160))
				mBase = m.M
			}
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v54 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v53 + v54
			v57 = int32(4449876)
			v59 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v59 - v54
		}
	}
	return
}
func F_pgstat_relation_delete_pending_cb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+128))
	if v4 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+272))
		if v7 == int32(0) {
		} else {
			v10 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v4)+272)) = v10
		}
	}
	return
}
func F_pgstat_report_checkpointer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1010]))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1011]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[1012]))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1013]))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1014]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1015]))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1016]))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1018]))
	if v5|(v7|(v9|(v11|(v13|(v15|(v17|v19)))))) != 0 {
		v64 = int32(4449876)
		v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		v67 = int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
		*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
		v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
		v75 = int32(4434848)
		v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
		v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
		v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
		v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
		v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
		v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
		v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
		v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
		v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
		v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
		v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
		v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
		v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
		v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
		v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
		v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
		v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
		v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
		*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
		v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
		v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
		mBase = m.M
		F_pgstat_flush_io(m, int32(0))
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return
		} else {
			return
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _consts[1025]))
		v30 = *(*int32)(unsafe.Add(mBase, _consts[1022]))
		v32 = *(*int32)(unsafe.Add(mBase, _consts[1026]))
		v34 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
		v36 = *(*int32)(unsafe.Add(mBase, _consts[1027]))
		v38 = *(*int32)(unsafe.Add(mBase, _consts[1020]))
		v40 = *(*int32)(unsafe.Add(mBase, _consts[1028]))
		v42 = *(*int32)(unsafe.Add(mBase, _consts[1019]))
		if v28|(v30|(v32|(v34|(v36|(v38|(v40|v42)))))) != 0 {
			v64 = int32(4449876)
			v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			v67 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
			*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
			v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
			v75 = int32(4434848)
			v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
			v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
			v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
			v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
			v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
			v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
			v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
			v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
			v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
			v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
			v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
			v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
			v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
			v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
			v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
			v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
			v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
			v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
			v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
			*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
			*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
			v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
			v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
			mBase = m.M
			F_pgstat_flush_io(m, int32(0))
			mBase = m.M
			v141 = m.ExcPending
			if v141 != 0 {
				return
			} else {
				return
			}
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, _consts[1023]))
			if v51 != 0 {
				v64 = int32(4449876)
				v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				v67 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
				*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
				v75 = int32(4434848)
				v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
				v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
				v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
				v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
				v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
				v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
				v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
				v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
				v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
				v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
				v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
				v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
				v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
				v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
				v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
				v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
				*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
				*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
				v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
				v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
				mBase = m.M
				F_pgstat_flush_io(m, int32(0))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return
				} else {
					return
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, _consts[1029]))
				if v53 != 0 {
					v64 = int32(4449876)
					v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
					v67 = int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
					*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
					v75 = int32(4434848)
					v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
					v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
					v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
					v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
					v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
					v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
					v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
					v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
					v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
					v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
					v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
					v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
					v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
					v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
					*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
					*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
					v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
					*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
					v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
					mBase = m.M
					F_pgstat_flush_io(m, int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						return
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _consts[1024]))
					if v55 != 0 {
						v64 = int32(4449876)
						v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
						v67 = int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
						*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
						v75 = int32(4434848)
						v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
						v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
						v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
						v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
						v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
						v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
						v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
						v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
						v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
						v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
						v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
						v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
						*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
						*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
						v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
						*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
						v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
						mBase = m.M
						F_pgstat_flush_io(m, int32(0))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							return
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, _consts[1030]))
						if v57 != 0 {
							v64 = int32(4449876)
							v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							v67 = int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
							*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
							v75 = int32(4434848)
							v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
							v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
							v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
							v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
							v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
							v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
							v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
							v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
							v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
							v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
							v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
							*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
							*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
							v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
							*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
							v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
							mBase = m.M
							F_pgstat_flush_io(m, int32(0))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								return
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _consts[1031]))
							if v59 != 0 {
								v64 = int32(4449876)
								v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
								v67 = int32(1)
								*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
								*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
								v75 = int32(4434848)
								v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
								v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
								v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
								v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
								v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
								v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
								v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
								v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
								v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
								v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
								v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
								v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
								v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
								v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
								v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
								v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
								*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
								*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
								v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
								*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
								v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
								mBase = m.M
								F_pgstat_flush_io(m, int32(0))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return
								} else {
									return
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, _consts[1032]))
								if v61 == int32(0) {
									return
								} else {
									v64 = int32(4449876)
									v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									v67 = int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[7])) = v66 + v67
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v3)+424))
									*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + v67
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v3)+432))
									v75 = int32(4434848)
									v76 = *(*int64)(unsafe.Add(mBase, _consts[1018]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+432)) = v74 + v76
									v79 = *(*int64)(unsafe.Add(mBase, uint32(v3)+440))
									v81 = *(*int64)(unsafe.Add(mBase, _consts[1016]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+440)) = v79 + v81
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v3)+448))
									v86 = *(*int64)(unsafe.Add(mBase, _consts[1014]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+448)) = v84 + v86
									v89 = *(*int64)(unsafe.Add(mBase, uint32(v3)+456))
									v91 = *(*int64)(unsafe.Add(mBase, _consts[1012]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+456)) = v89 + v91
									v94 = *(*int64)(unsafe.Add(mBase, uint32(v3)+464))
									v96 = *(*int64)(unsafe.Add(mBase, _consts[1019]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+464)) = v94 + v96
									v99 = *(*int64)(unsafe.Add(mBase, uint32(v3)+472))
									v101 = *(*int64)(unsafe.Add(mBase, _consts[1020]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+472)) = v99 + v101
									v104 = *(*int64)(unsafe.Add(mBase, uint32(v3)+480))
									v106 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+480)) = v104 + v106
									v109 = *(*int64)(unsafe.Add(mBase, uint32(v3)+488))
									v111 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+488)) = v109 + v111
									v114 = *(*int64)(unsafe.Add(mBase, uint32(v3)+496))
									v116 = *(*int64)(unsafe.Add(mBase, _consts[1023]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+496)) = v114 + v116
									v119 = *(*int64)(unsafe.Add(mBase, uint32(v3)+504))
									v121 = *(*int64)(unsafe.Add(mBase, _consts[1024]))
									*(*int64)(unsafe.Add(mBase, uint32(v3)+504)) = v119 + v121
									*(*int32)(unsafe.Add(mBase, uint32(v3)+424)) = v70 + int32(2)
									v130 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									*(*int32)(unsafe.Add(mBase, _consts[7])) = v130 - v67
									v138 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), int32(88))
									mBase = m.M
									F_pgstat_flush_io(m, int32(0))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstat_reset_matching_entries(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = v11 + int32(4)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = v18
	goto L1
L1:
	;
	v28 = F_dshash_seq_next(m, v11+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = v28
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_dshash_seq_term(m, v11+int32(4))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L44
	}
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+16)))
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v158 = F_dshash_seq_next(m, v11+int32(4))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L42
	}
L10:
	;
	v39 = m.T0[l0].(func(*base.Module, int32, int32) int32)(m, v33, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v46 = F_dsa_get_address(m, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v49 = v46 + int32(4)
	v51 = F_LWLockAcquire(m, v49, int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if base.Ui32(v53-int32(1)) <= base.Ui32(int32(11)) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if base.Ui32(v53-int32(1)) <= base.Ui32(int32(11)) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v82 = v53*int32(72) + int32(1618272)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(int32(8)) < base.Ui32(v53-int32(24)) {
		v80 = int32(0)
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v82 = v80
	goto L15
L20:
	;
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
	if v70 == v68 {
		v80 = v68
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+v53<<(uint(int32(2))%32)-int32(96))))
	v80 = v78
	goto L19
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	if base.Ui32(v53-int32(1)) <= base.Ui32(int32(11)) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v111 = v53*int32(72) + int32(1618272)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(int32(8)) < base.Ui32(v53-int32(24)) {
		v109 = int32(0)
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v111 = v109
	goto L22
L27:
	;
	v97 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
	if v99 == v97 {
		v109 = v97
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v99+v53<<(uint(int32(2))%32)-int32(96))))
	v109 = v107
	goto L26
L29:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	v146 = F__emscripten_memset_bulkmem(m, v46+v112, base.I32_extend8_s(int32(0)), v144)
	mBase = m.M
	goto L36
L30:
	;
	v143 = v53*int32(72) + int32(1618272)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(8)) < base.Ui32(v53-int32(24)) {
		v141 = int32(0)
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v143 = v141
	goto L29
L34:
	;
	v129 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
	if v131 == v129 {
		v141 = v129
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v131+v53<<(uint(int32(2))%32)-int32(96))))
	v141 = v139
	goto L33
L36:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v147 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	m.T0[v147].(func(*base.Module, int32, int64))(m, v46, l2)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_LWLockRelease(m, v49)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L9
L42:
	;
	if v158 != 0 {
		v33 = v158
		goto L7
	} else {
		goto L43
	}
L43:
	;
	goto L8
L44:
	;
	m.G0 = v11 + int32(32)
	return
}
func F_pgstat_reset_of_kind(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		v29 = l0*int32(72) + int32(1618272)
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
			v29 = int32(0)
		} else {
			v17 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[1009]))
			if v19 == v17 {
				v29 = v17
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v19+l0<<(uint(int32(2))%32)-int32(96))))
				v29 = v27
			}
		}
	}
	v33 = m.G0
	v34 = int32(16)
	v35 = v33 - v34
	m.G0 = v35
	F___gettimeofday(m, v35)
	mBase = m.M
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	m.G0 = v35 + v34
	v47 = v39 + v38*int64(1000000) - int64(946684800000000)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v48&int32(1) != 0 {
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
		m.T0[v51].(func(*base.Module, int64))(m, v47)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			return
		}
	} else {
		F_pgstat_reset_matching_entries(m, int32(1248), l0, v47)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgstat_snapshot_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v46 int64
	_ = v46
	var v52 int64
	_ = v52
	var v57 int64
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v106 int64
	_ = v106
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v133 int64
	_ = v133
	var v141 int32
	_ = v141
	var v148 float64
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v202 int64
	_ = v202
	var v208 int64
	_ = v208
	var v213 int64
	_ = v213
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v273 int64
	_ = v273
	var v279 int64
	_ = v279
	var v284 int64
	_ = v284
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v511 int64
	_ = v511
	var v512 int64
	_ = v512
	var v513 int64
	_ = v513
	var v518 int64
	_ = v518
	var v520 int64
	_ = v520
	var v525 int64
	_ = v525
	var v531 int64
	_ = v531
	var v536 int64
	_ = v536
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int64
	_ = v580
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int64
	_ = v635
	var v637 int64
	_ = v637
	var v639 int64
	_ = v639
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int64
	_ = v663
	var v665 int64
	_ = v665
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v675 int64
	_ = v675
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int64
	_ = v720
	var v722 int64
	_ = v722
	var v734 int32
	_ = v734
	var v752 int32
	_ = v752
	var v761 int32
	_ = v761
	var v773 int32
	_ = v773
	var v784 int32
	_ = v784
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(8))))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v24
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v28
	v30 = int64(23)
	v33 = int64(2388976653695081527)
	v34 = (v25 ^ int64(base.Ui64(v25)>>(uint(v30)%64))) * v33
	v35 = int64(47)
	v40 = int64(-8645972361240307355)
	v46 = (v24 ^ int64(base.Ui64(v24)>>(uint(v30)%64))) * v33
	v52 = ((v34^int64(base.Ui64(v34)>>(uint(v35)%64))^int64(-9208349263878056368))*v40 ^ int64(base.Ui64(v46)>>(uint(v35)%64)) ^ v46) * v40
	v57 = (int64(base.Ui64(v52)>>(uint(v30)%64)) ^ v52) * v33
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = v65
	v73 = v66
	goto L1
L1:
	;
	if base.Ui32(v72) <= base.Ui32(v73) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v784 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v784
	v72 = v784
	v73 = v773
	goto L1
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v761)
	m.G0 = v20 + int32(16)
	return v752
L5:
	;
	v752 = v734
	v761 = int32(0)
	goto L4
L6:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v717 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v716 + v717
	v720 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v707)+8)) = v720
	v722 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v707))) = v722
	*(*uint8)(unsafe.Add(mBase, uint32(v707)+16)) = uint8(v717)
	v734 = v707
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L23
	} else {
		goto L113
	}
L8:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v85 == int64(4294967296) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v412 = int32(0)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v415 = v414 & base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(v35)%64))^v57-int64(base.Ui64(v57)>>(uint(int64(32))%64)))
	v418 = v413 + v415*int32(24)
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+16)))
	if v419 == v412 {
		v707 = v418
		goto L6
	} else {
		goto L64
	}
L11:
	;
	v88 = int32(0)
	v90 = int64(2)
	v92 = v85 << (uint(int64(1)) % 64)
	if base.Ui64(v92) <= base.Ui64(v90) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L23
	} else {
		goto L61
	}
L14:
	;
	v95 = v90
	goto L16
L15:
	;
	v95 = v92
	goto L16
L16:
	;
	v96 = int64(1)
	if v95&(v95-v96) == int64(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v106 = v95
	goto L19
L18:
	;
	v106 = v96 << (uint(int64(64)-base.I64_clz(v95)) % 64)
	goto L19
L19:
	;
	if base.Ui64(v106*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v118 = F_MemoryContextAllocExtended(m, v113, base.I32_wrap_i64(v106)*int32(24), int32(5))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L23
	} else {
		goto L58
	}
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v118
	v123 = int64(1)
	if v106&(v106-v123) == int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v133 = v106
	goto L27
L26:
	;
	v133 = v123 << (uint(int64(64)-base.I64_clz(v106)) % 64)
	goto L27
L27:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v133*int64(24)) {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v133
	v141 = base.I32_wrap_i64(v133) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v141
	v148 = base.F64_mul(base.F64_convert_i64_u(v133), float64(0.9))
	if base.F64_lt(v148, float64(4.294967296e+09))&base.F64_ge(v148, float64(0)) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v133 == int64(4294967296) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v154 = base.I32_trunc_f64_u(v148)
	v156 = v154
	goto L29
L31:
	;
	goto L32
L32:
	;
	v156 = int32(0)
	goto L29
L33:
	;
	v157 = int32(-85899346)
	goto L35
L34:
	;
	v157 = v156
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v157
	if v112 != int64(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v166 = v88
	goto L40
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v111)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L23
	} else {
		goto L57
	}
L39:
	;
	v235 = v229
	v240 = v88
	goto L45
L40:
	;
	v180 = v111 + v166*int32(24)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+16)))
	if v181 != int32(1) {
		v229 = v166
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v229 = int32(0)
	goto L39
L42:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v180)))
	v185 = int64(23)
	v188 = int64(2388976653695081527)
	v189 = (int64(base.Ui64(v184)>>(uint(v185)%64)) ^ v184) * v188
	v190 = int64(47)
	v195 = int64(-8645972361240307355)
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	v202 = (int64(base.Ui64(v197)>>(uint(v185)%64)) ^ v197) * v188
	v208 = ((v189^int64(base.Ui64(v189)>>(uint(v190)%64))^int64(-9208349263878056368))*v195 ^ int64(base.Ui64(v202)>>(uint(v190)%64)) ^ v202) * v195
	v213 = (int64(base.Ui64(v208)>>(uint(v185)%64)) ^ v208) * v188
	if v141&base.I32_wrap_i64(int64(base.Ui64(v213)>>(uint(v190)%64))^v213-int64(base.Ui64(v213)>>(uint(int64(32))%64))) == v166 {
		v229 = v166
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v224 = v166 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v224)) < base.Ui64(v112) {
		v166 = v224
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v249 = v111 + v235*int32(24)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+16)))
	if v250 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L38
L47:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v249)))
	v254 = int64(23)
	v257 = int64(2388976653695081527)
	v258 = (int64(base.Ui64(v253)>>(uint(v254)%64)) ^ v253) * v257
	v259 = int64(47)
	v264 = int64(-8645972361240307355)
	v267 = v249 + int32(8)
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	v273 = (int64(base.Ui64(v268)>>(uint(v254)%64)) ^ v268) * v257
	v279 = ((v258^int64(base.Ui64(v258)>>(uint(v259)%64))^int64(-9208349263878056368))*v264 ^ int64(base.Ui64(v273)>>(uint(v259)%64)) ^ v273) * v264
	v284 = (int64(base.Ui64(v279)>>(uint(v254)%64)) ^ v279) * v257
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v299 = base.I32_wrap_i64(int64(base.Ui64(v284)>>(uint(v259)%64)) ^ v284 - int64(base.Ui64(v284)>>(uint(int64(32))%64)))
	goto L50
L48:
	;
	goto L49
L49:
	;
	v341 = v235 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v341)) < base.Ui64(v112) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v310 = v299 & v292
	v315 = v118 + v310*int32(24)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+16)))
	if v316 != 0 {
		v299 = v310 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v249)))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v317
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v249)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v315)+16)) = v319
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	*(*int64)(unsafe.Add(mBase, uint32(v315)+8)) = v321
	goto L49
L52:
	;
	goto L51
L53:
	;
	v345 = v341
	goto L55
L54:
	;
	v345 = int32(0)
	goto L55
L55:
	;
	v347 = v240 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v347)) < base.Ui64(v112) {
		v235 = v345
		v240 = v347
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L46
L57:
	;
	goto L12
L58:
	;
	F_errmsg_internal(m, int32(386506), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(314692), int32(327), int32(328665))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errmsg_internal(m, int32(386506), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L23
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(314692), int32(327), int32(328665))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v427 = v415
	v430 = v418
	v431 = v412
	goto L65
L65:
	;
	v440 = int32(16)
	goto L70
L66:
	;
	v707 = v684
	goto L6
L67:
	;
	if v502 == int32(0) {
		v752 = v430
		v761 = int32(1)
		goto L4
	} else {
		goto L85
	}
L68:
	;
	v502 = int32(0)
	goto L67
L69:
	;
	v476 = v471
	v477 = v472
	v478 = v473
	goto L79
L70:
	;
	if (v430|v20)&int32(3) != 0 {
		v471 = v430
		v472 = v20
		v473 = v440
		goto L69
	} else {
		goto L73
	}
L72:
	;
	if v461 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L73:
	;
	v448 = v430
	v449 = v20
	v450 = v440
	goto L74
L74:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	if v453 != v454 {
		v471 = v448
		v472 = v449
		v473 = v450
		goto L69
	} else {
		goto L76
	}
L75:
	;
	goto L72
L76:
	;
	v456 = int32(4)
	v457 = v449 + v456
	v459 = v448 + v456
	v461 = v450 - v456
	if base.Ui32(int32(3)) < base.Ui32(v461) {
		v448 = v459
		v449 = v457
		v450 = v461
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v471 = v459
	v472 = v457
	v473 = v461
	goto L69
L79:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v481 == v482 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v502 = v481 - v482
	goto L67
L81:
	;
	v484 = int32(1)
	v489 = v478 - v484
	if v489 != 0 {
		v476 = v476 + v484
		v477 = v477 + v484
		v478 = v489
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	goto L68
L85:
	;
	v506 = v427 + int32(1)
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v430)))
	v508 = int64(23)
	v511 = int64(2388976653695081527)
	v512 = (int64(base.Ui64(v507)>>(uint(v508)%64)) ^ v507) * v511
	v513 = int64(47)
	v518 = int64(-8645972361240307355)
	v520 = *(*int64)(unsafe.Add(mBase, uint32(v430)+8))
	v525 = (int64(base.Ui64(v520)>>(uint(v508)%64)) ^ v520) * v511
	v531 = ((v512^int64(base.Ui64(v512)>>(uint(v513)%64))^int64(-9208349263878056368))*v518 ^ int64(base.Ui64(v525)>>(uint(v513)%64)) ^ v525) * v518
	v536 = (int64(base.Ui64(v531)>>(uint(v508)%64)) ^ v531) * v511
	v544 = v414 & base.I32_wrap_i64(int64(base.Ui64(v536)>>(uint(v513)%64))^v536-int64(base.Ui64(v536)>>(uint(int64(32))%64)))
	if base.Ui32(v427) < base.Ui32(v544) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v548 = v427 + v546
	goto L88
L87:
	;
	v548 = v427
	goto L88
L88:
	;
	if base.Ui32(v548-v544) < base.Ui32(v431) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v552 = v414 & v506
	v555 = v413 + v552*int32(24)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+16)))
	if v556 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v670 = v431 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v670) {
		goto L108
	} else {
		goto L109
	}
L92:
	;
	v566 = v552
	v568 = int32(0)
	goto L95
L93:
	;
	v599 = v555
	v602 = v552
	goto L94
L94:
	;
	if v427 != v602 {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	v575 = v568 + int32(1)
	if int32(151) <= v575 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v599 = v591
	v602 = v588
	goto L94
L97:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v580 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v578), base.F64_convert_i64_u(v580)), float64(0.1)) != 0 {
		v773 = v578
		goto L3
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v588 = (v566 + int32(1)) & v414
	v591 = v413 + v588*int32(24)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+16)))
	if v592 != 0 {
		v566 = v588
		v568 = v575
		goto L95
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	goto L96
L102:
	;
	v617 = v599
	v620 = v602
	goto L105
L103:
	;
	goto L104
L104:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v660 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v659 + v660
	v663 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v430)+8)) = v663
	v665 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v430))) = v665
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+16)) = uint8(v660)
	v734 = v430
	goto L5
L105:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v631 = v628 & (v620 - int32(1))
	v634 = v413 + v631*int32(24)
	v635 = *(*int64)(unsafe.Add(mBase, uint32(v634)))
	*(*int64)(unsafe.Add(mBase, uint32(v617))) = v635
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v634)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v617)+16)) = v637
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v634)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v617)+8)) = v639
	if v427 != v631 {
		v617 = v634
		v620 = v631
		goto L105
	} else {
		goto L107
	}
L106:
	;
	goto L104
L107:
	;
	goto L106
L108:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v675 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v673), base.F64_convert_i64_u(v675)), float64(0.1)) != 0 {
		v773 = v673
		goto L3
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v681 = v414 & v506
	v684 = v413 + v681*int32(24)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684)+16)))
	if v685 != 0 {
		v427 = v681
		v430 = v684
		v431 = v670
		goto L65
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	goto L66
L113:
	;
	F_errmsg_internal(m, int32(446888), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L23
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(314692), int32(630), int32(300337))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L23
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_subscription_flush_cb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = F_pgstat_lock_entry(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v12 + v13
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v16 + v17
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)+40))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v20 + v21
			v25 = v6 + int32(48)
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v25))) = v26 + v27
			v31 = v6 + int32(56)
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v7)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v31))) = v32 + v33
			v36 = int32(-64)
			v37 = v6 - v36
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
			v39 = *(*int64)(unsafe.Add(mBase, uint32(v7)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v37))) = v38 + v39
			v43 = v6 + int32(72)
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
			*(*int64)(unsafe.Add(mBase, uint32(v43))) = v44 + v45
			v49 = v6 + int32(80)
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
			v51 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50 + v51
			v55 = v6 + int32(88)
			v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v7-v36)))
			*(*int64)(unsafe.Add(mBase, uint32(v55))) = v56 + v59
			F_pgstat_unlock_entry(m, l0)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				return v8
			}
		} else {
			return v8
		}
	}
}
func F_pgstat_wal_init_backend_cb(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	v3 = *(*int64)(unsafe.Add(mBase, _consts[43]))
	*(*int64)(unsafe.Add(mBase, _consts[1055])) = v3
	v7 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, _consts[1056])) = v7
	v11 = *(*int64)(unsafe.Add(mBase, _consts[39]))
	*(*int64)(unsafe.Add(mBase, _consts[1057])) = v11
	v15 = *(*int64)(unsafe.Add(mBase, _consts[41]))
	*(*int64)(unsafe.Add(mBase, _consts[1058])) = v15
	return
}

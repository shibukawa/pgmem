package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CleanUpLock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_CleanUpLock[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v28 = F_hash_search_with_hash_value(m, v21, l1, v22<<(uint(int32(4))%32)^l3, int32(2), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			if v28 == int32(0) {
				F_errstart_cold(m, int32(24), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_CleanUpLock_0), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1799), int32(_a_F_CleanUpLock_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if v33 == int32(0) {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_CleanUpLock[1]))
					v40 = F_hash_search_with_hash_value(m, v37, l0, l3, int32(2), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if v40 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(24), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_CleanUpLock_3), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1815), int32(_a_F_CleanUpLock_2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					if l4 == int32(0) {
						return
					} else {
						F_ProcLockWakeup(m, l2, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v33 == int32(0) {
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_CleanUpLock[1]))
			v40 = F_hash_search_with_hash_value(m, v37, l0, l3, int32(2), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				if v40 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(24), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_CleanUpLock_3), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1815), int32(_a_F_CleanUpLock_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			if l4 == int32(0) {
				return
			} else {
				F_ProcLockWakeup(m, l2, l0)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_CleanupInvalidationState(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[0]))
	v16 = F_LWLockAcquire(m, v12+int32(768), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[1]))
	v20 = base.I32_wrap_i64(l1)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[2]))
	v25 = v20 + v22<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_CleanupInvalidationState[3]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_CleanupInvalidationState[4]))) = v19
	v35 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_CleanupInvalidationState[5]))) = uint16(v35)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_CleanupInvalidationState[6])))
	v39 = v37 - int32(1)
	if v39 < v35 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L13
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_CleanupInvalidationState[7])))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[2]))
	if v46 == v48 {
		v76 = v39
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_CleanupInvalidationState[6]))) = v76
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[0]))
	F_LWLockRelease(m, v86+int32(768))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v50 = v39
	goto L7
L7:
	;
	if v50 <= int32(0) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	if v37 == v50 {
		v76 = v39
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v63 = v50 - int32(1)
	v66 = v42 + v63<<(uint(int32(2))%32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v67 != v48 {
		v50 = v63
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v46
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_CleanupInvalidationState[6])))
	v76 = v71 - int32(1)
	goto L5
L12:
	;
	return
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CleanupInvalidationState_0), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_CleanupInvalidationState_1), int32(361), int32(_a_F_CleanupInvalidationState_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CloneRowTriggersToPartition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int64
	_ = v252
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = v17 + int32(40)
	v24 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v20, int32(2), int32(3), int32(184), v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(1)
	v35 = F_systable_beginscan(m, v29, int32(2701), v32, int32(0), v32, v20)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	v43 = F_AllocSetContextCreateInternal(m, v38, int32(_a_F_CloneRowTriggersToPartition_0), int32(0), int32(1024), int32(_a_F_CloneRowTriggersToPartition_1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v45 = F_systable_getnext(m, v35)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L61
	}
L7:
	;
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = v45
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_MemoryContextDelete(m, v43)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L58
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+22)))
	v63 = v61 + v62
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+80)))
	if v64&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v309 = F_systable_getnext(m, v35)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L56
	}
L14:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+83)))
	if v69 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	switch v64 & int32(66) {
	case 0, 2:
		goto L16
	default:
		goto L17
	}
L16:
	;
	v87 = int32(0)
	v88 = int32(_a_F_CloneRowTriggersToPartition_2)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v43
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v96 = F_heap_getattr_6(m, v51, int32(17), v93, v17+int32(39))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v63 + int32(12)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_3), v17)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_5), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v98 = int32(0)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+39)))
	if v99 == v98 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v103 = F_text_to_cstring(m, base.I32_wrap_i64(v96))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v113 = v98
	goto L24
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v63)+116))
	if int32(0) < v114 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v105 = F_stringToNode(m, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v108 = F_map_partition_varattnos(m, v105, int32(1), l1, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v111 = F_map_partition_varattnos(m, v108, int32(2), l1, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v113 = v111
	goto L24
L29:
	;
	v122 = int32(0)
	v127 = v87
	goto L32
L30:
	;
	v165 = v87
	goto L31
L31:
	;
	v172 = int32(0)
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+98)))
	if v173 <= v172 {
		v235 = v172
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63+int32(124)+v122<<(uint(int32(1))%32)))))
	v148 = F_pstrdup(m, v134+v135<<(uint(int32(3))%32)+v142*int32(100)-int32(68))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v165 = v152
	goto L31
L34:
	;
	v150 = F_makeString(m, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v152 = F_lappend(m, v127, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v155 = v122 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v63)+116))
	if v155 < v156 {
		v122 = v155
		v127 = v152
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v242 = F_palloc0(m, int32(52))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L53
	}
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v180 = F_heap_getattr_6(m, v51, int32(16), v177, v17+int32(39))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+39)))
	if v182 == int32(1) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v186 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v180))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+98)))
	if v188 <= int32(0) {
		v235 = v172
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v191 = int32(1)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v193&v191 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v196 = v191
	goto L46
L45:
	;
	v196 = int32(4)
	goto L46
L46:
	;
	v201 = v186 + v196
	v203 = int32(0)
	v207 = v172
	goto L47
L47:
	;
	v213 = F_pstrdup(m, v201)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v235 = v217
	goto L38
L49:
	;
	v215 = F_makeString(m, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v217 = F_lappend(m, v207, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v219 = F_strlen(m, v201)
	mBase = m.M
	v221 = int32(1)
	v224 = v203 + v221
	v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+98)))
	if v224 < v225 {
		v201 = v219 + v201 + v221
		v203 = v224
		v207 = v217
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v244 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)) = uint8(v244)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = int32(181)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v63)+92))
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+24)) = uint8(v249)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v235
	v252 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v63 + int32(12)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+5)) = uint8(base.B2i32(v248 != v244))
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+80)))
	v262 = v260 & int32(66)
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+26)) = uint16(v262)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v242)+36)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v242)+32)) = v165
	v269 = v264 & int32(60)
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+28)) = uint16(v269)
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+44)) = uint8(v271)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+97)))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+48)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+45)) = uint8(v273)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v63)+84))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v288 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63)+82)))
	F_CreateTriggerFiringOn(m, v17+int32(24), v242, v244, v280, v281, v244, v244, v284, v285, v113, v244, v249, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v89
	F_MemoryContextReset(m, v43)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L13
L56:
	;
	if v309 != 0 {
		v51 = v309
		goto L11
	} else {
		goto L57
	}
L57:
	;
	goto L12
L58:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_relation_close(m, v29, int32(3))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v17 + int32(96)
	return
L61:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v63 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v339 + int32(4)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_7), v17+int32(16))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_8), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clamp_row_est(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	var v11 float64
	_ = v11
	var v15 float64
	_ = v15
	v2 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)))|base.F64_gt(l0, v2) != 0 {
		v15 = v2
	} else {
		v11 = float64(1)
		if base.F64_le(l0, v11) != 0 {
			v15 = v11
		} else {
			v15 = base.F64_nearest(l0)
		}
	}
	return v15
}
func F_clamp_width_est(m *base.Module, l0 int64) int32 {
	var v2 int64
	_ = v2
	var v5 int64
	_ = v5
	v2 = int64(1073741823)
	if v2 <= l0 {
		v5 = v2
	} else {
		v5 = l0
	}
	return base.I32_wrap_i64(v5)
}
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_clauselist_selectivity_ext(m, l0, l1, l2, l3, l4, int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return float64(0)
	} else {
		return v7
	}
}
func F_clauselist_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v173 float64
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v212 float64
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 float64
	_ = v221
	var v222 float64
	_ = v222
	var v225 float64
	_ = v225
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v244 float64
	_ = v244
	var v245 float64
	_ = v245
	var v246 float64
	_ = v246
	var v247 float64
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 float64
	_ = v253
	var v268 float64
	_ = v268
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v7
	if l1 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v268
L2:
	;
	v37 = float64(1)
	v40 = F_find_single_rel_for_clauses(m, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L8
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = F_clause_selectivity_ext(m, l0, v32, l2, l3, l4, l5)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return float64(0)
L6:
	;
	v268 = v33
	goto L1
L7:
	;
	if l1 == int32(0) {
		v268 = v54
		goto L1
	} else {
		goto L13
	}
L8:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(v40 == int32(0)) != 0 {
		v54 = v37
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+84))
	if v45 != 0 {
		v54 = v37
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	if v46 == int32(0) {
		v54 = v37
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v52 = F_statext_clauselist_selectivity(m, l0, l1, l2, l3, l4, v40, v20+int32(12), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v54 = v52
	goto L7
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v57 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v7
	v72 = int32(-1)
	v75 = v54
	goto L17
L15:
	;
	v192 = v54
	goto L16
L16:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v195 == int32(0) {
		v268 = v192
		goto L1
	} else {
		goto L56
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v71<<(uint(int32(2))%32))))
	v84 = v72 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v86 = F_bms_is_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	v192 = v173
	goto L16
L19:
	;
	v175 = v71 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v175 < v176 {
		v71 = v175
		v72 = v84
		v75 = v173
		goto L17
	} else {
		goto L55
	}
L20:
	;
	if v86 != 0 {
		v173 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = F_clause_selectivity_ext(m, l0, v82, l2, l3, l4, l5)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v91 != int32(320) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v173 = base.F64_mul(v75, v88)
	goto L19
L24:
	;
	if v103 != int32(17) {
		goto L23
	} else {
		goto L30
	}
L25:
	;
	v101 = v82
	v102 = int32(0)
	v103 = v91
	goto L24
L26:
	;
	goto L27
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+10)))
	if v94 == int32(1) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v97 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v101 = v97
	v102 = v82
	v103 = v100
	goto L24
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	if v106 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 != int32(2) {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	if v102 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v147 = F_get_oprrest(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L52
	}
L34:
	;
	v145 = int32(0)
	goto L33
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	if v112 != int32(1) {
		goto L23
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v127 = F_NumRelids(m, l0, v101)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L43
	}
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v102)+48))
	v119 = F_is_pseudo_constant_clause_relids(m, v117, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v119 != 0 {
		v145 = int32(1)
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v102)+44))
	v125 = F_is_pseudo_constant_clause_relids(m, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	if v125 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L23
L43:
	;
	if v127 != int32(1) {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v135 = F_is_pseudo_constant_clause(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	if v135 != 0 {
		v145 = int32(1)
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = F_is_pseudo_constant_clause(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	if v140 == int32(0) {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	goto L34
L49:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(0), v88)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L54
	}
L50:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(1), v88)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	switch v147 - int32(336) {
	case 0:
		goto L50
	case 1:
		goto L49
	default:
		goto L23
	}
L52:
	;
	switch v147 - int32(103) {
	case 0:
		goto L50
	case 1:
		goto L49
	default:
		goto L51
	}
L53:
	;
	v173 = v75
	goto L19
L54:
	;
	v173 = v75
	goto L19
L55:
	;
	goto L18
L56:
	;
	v201 = v195
	v212 = v192
	goto L57
L57:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+8)))
	if v215 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v268 = v253
	goto L1
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	F_pfree(m, v201)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L73
	}
L60:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+9)))
	if v218 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v201)+24))
	v247 = v246
	goto L59
L63:
	;
	v221 = float64(0.005)
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v201)+24))
	if base.F64_eq(v222, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L59
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v201)+16))
	v247 = v245
	goto L59
L66:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v201)+16))
	if base.F64_eq(v225, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L59
	} else {
		goto L67
	}
L67:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v233 = F_nulltestsel(m, l0, int32(0), v232, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v235 = base.F64_add(base.F64_add(base.F64_add(v222, v225), float64(-1)), v233)
	if base.F64_le(v235, float64(0)) == int32(0) {
		v247 = v235
		goto L59
	} else {
		goto L69
	}
L69:
	;
	if base.F64_lt(v235, float64(-0.01)) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v244 = float64(0.005)
	goto L72
L71:
	;
	v244 = float64(1e-10)
	goto L72
L72:
	;
	v247 = v244
	goto L59
L73:
	;
	v253 = base.F64_mul(v212, v247)
	if v250 != 0 {
		v201 = v250
		v212 = v253
		goto L57
	} else {
		goto L74
	}
L74:
	;
	goto L58
}
func F_cleartraverse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = m.T0[v6].(func(*base.Module) int32)(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(101)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v15 = v13
	goto L8
L7:
	;
	v15 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v15
	return
L9:
	;
	return
L10:
	;
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v22 == v20 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = v22
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	F_cleartraverse(m, l0, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v30 != 0 {
		v26 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_clonesuccessorstates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v457 int32
	_ = v457
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v516 int32
	_ = v516
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = m.T0[v18].(func(*base.Module) int32)(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(101)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if l5 != 0 {
		v51 = l5
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v27 = v25
	goto L8
L7:
	;
	v27 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v27
	return
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v52))) = uint8(v54)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v56 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L10:
	;
	v30 = F_palloc_extended(m, l7, int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v30 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(101)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if l6 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v40 = v38
	goto L17
L16:
	;
	v40 = int32(12)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v40
	return
L18:
	;
	if l7 == int32(0) {
		v51 = v30
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l7 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	base.MemoryCopy(m, v30, l6, l7)
	v51 = v30
	goto L9
L22:
	;
	base.MemoryFill(m, v30, int32(0), l7)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v47))) = uint8(v49)
	v51 = v30
	goto L9
L25:
	;
	if l5 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L26:
	;
	v69 = v56
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	if v74 != 0 {
		goto L25
	} else {
		goto L29
	}
L28:
	;
	goto L25
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v76 - int32(76) {
	case 0, 18, 21, 38:
		goto L33
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L32
	default:
		goto L34
	}
L30:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v457 != 0 {
		v69 = v457
		goto L27
	} else {
		goto L148
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v120))))
	if v122 != 0 {
		goto L30
	} else {
		goto L44
	}
L32:
	;
	F_cparc(m, l0, v69, l2, v75)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v81 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	if v76 != int32(36) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v85 = v81
	goto L37
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v98 - int32(76) {
	case 0, 18, 21, 38:
		goto L31
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L39
	default:
		goto L40
	}
L38:
	;
	goto L32
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v103 != 0 {
		v85 = v103
		goto L37
	} else {
		goto L42
	}
L40:
	;
	if v98 == int32(36) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L38
L43:
	;
	goto L30
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v123 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if l4 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	v154 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v128 = v123
	goto L49
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	if v142 == v75 {
		v154 = v141
		goto L45
	} else {
		goto L51
	}
L50:
	;
	v154 = int32(0)
	goto L45
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v144 != 0 {
		v128 = v144
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v436 = F_newstate(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L145
	}
L54:
	;
	if v154 != 0 {
		goto L71
	} else {
		goto L72
	}
L55:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v167 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v76 != v162 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v164 == v165 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v169 = v167
	v179 = l2
	goto L62
L60:
	;
	goto L61
L61:
	;
	if v154 == int32(0) {
		goto L53
	} else {
		goto L69
	}
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v182 != int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L61
L64:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	if v191 != 0 {
		v169 = v191
		v179 = v190
		goto L62
	} else {
		goto L68
	}
L65:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v76 != v185 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if v187 == v188 {
		goto L54
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L63
L69:
	;
	F_cparc(m, l0, v69, l2, v154)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L30
L71:
	;
	goto L74
L72:
	;
	goto L73
L73:
	;
	F_clonesuccessorstates(m, l0, v75, l2, l3, l4, v51, l6, l7)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L144
	}
L74:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v238 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L105
L76:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+4)))
	if v245 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L74
L80:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	if v280 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v250 = v248 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v250))|base.B2i32(int32(1)<<(uint(v250)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v260 != 0 {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	if v261 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v273 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v245*int32(24))+12)) = v269
	v273 = v269
	goto L84
L86:
	;
	goto L87
L87:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+32)) = v271
	v273 = v271
	goto L84
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+36)) = v261
	goto L90
L89:
	;
	goto L90
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = int64(0)
	goto L80
L91:
	;
	if v279 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+20)) = v279
	goto L91
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v279
	goto L91
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v280
	goto L97
L96:
	;
	goto L97
L97:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v286 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v238)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v238)+28))
	if v291 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v290 != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v290
	goto L98
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v290
	goto L98
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v291
	goto L104
L103:
	;
	goto L104
L104:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v297 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = int32(0)
	v304 = v238 + int32(8)
	v305 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v304)+16)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304)+8)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v238
	goto L79
L105:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	if v328 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(-1)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v154)+32))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v409 != 0 {
		goto L137
	} else {
		goto L138
	}
L107:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328)+4)))
	if v335 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	goto L105
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v328)+20))
	if v370 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L112:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v340 = v338 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v340))|base.B2i32(int32(1)<<(uint(v340)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0)) != 0 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v350 != 0 {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v328)+36))
	if v351 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v363 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v355+v335*int32(24))+12)) = v359
	v363 = v359
	goto L115
L117:
	;
	goto L118
L118:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+32)) = v361
	v363 = v361
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+36)) = v351
	goto L121
L120:
	;
	goto L121
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v328)+32)) = int64(0)
	goto L111
L122:
	;
	if v369 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+20)) = v369
	goto L122
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+16)) = v369
	goto L122
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+20)) = v370
	goto L128
L127:
	;
	goto L128
L128:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+12)) = v376 - int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	if v381 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v380 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+16)) = v380
	goto L129
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+24)) = v380
	goto L129
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = v381
	goto L135
L134:
	;
	goto L135
L135:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = v387 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(0)
	v394 = v328 + int32(8)
	v395 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v394)+16)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = v395
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+16)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v328
	goto L110
L136:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v408 != 0 {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+32)) = v408
	goto L136
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v408
	goto L136
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v154
	goto L73
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+28)) = v412
	goto L140
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v412
	goto L140
L144:
	;
	goto L30
L145:
	;
	if v436 == int32(0) {
		goto L25
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = v75
	F_cparc(m, l0, v69, l2, v436)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L30
L148:
	;
	goto L28
L149:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v474 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	return
L152:
	;
	F_pfree(m, v51)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L162
	}
L153:
	;
	v478 = v474
	goto L154
L154:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+12))
	if v492 != 0 {
		goto L152
	} else {
		goto L156
	}
L155:
	;
	goto L152
L156:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	if v494 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v493)+24)) = v495
	F_clonesuccessorstates(m, l0, v494, v493, l3, l4, v495, v51, l7)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v478)+16))
	if v500 != 0 {
		v478 = v500
		goto L154
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	goto L155
L162:
	;
	goto L151
}

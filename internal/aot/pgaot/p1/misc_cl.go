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
				F_errstart_cold(m, int32(23), int32(0))
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
						F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1758), int32(_a_F_CleanUpLock_2))
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
							F_errstart_cold(m, int32(23), int32(0))
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
									F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1774), int32(_a_F_CleanUpLock_2))
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
					F_errstart_cold(m, int32(23), int32(0))
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
							F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1774), int32(_a_F_CleanUpLock_2))
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
func F_CleanupInvalidationState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[0]))
	v15 = F_LWLockAcquire(m, v11+int32(768), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[1]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[2]))
	v23 = l1 + v20<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_CleanupInvalidationState[3]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_CleanupInvalidationState[4]))) = v18
	v33 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_CleanupInvalidationState[5]))) = uint16(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[6])))
	v37 = v35 - int32(1)
	if v37 < v33 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L13
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[7])))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[2]))
	if v44 == v46 {
		v73 = v37
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[6]))) = v73
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[0]))
	F_LWLockRelease(m, v82+int32(768))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v48 = v37
	goto L7
L7:
	;
	if v48 <= int32(0) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	if v48 == v35 {
		v73 = v37
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v60 = v48 - int32(1)
	v63 = v40 + v60<<(uint(int32(2))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 != v46 {
		v48 = v60
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v44
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[6])))
	v73 = v68 - int32(1)
	goto L5
L12:
	;
	return
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CleanupInvalidationState_0), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_CleanupInvalidationState_1), int32(359), int32(_a_F_CleanupInvalidationState_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
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
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v17+int32(48), int32(2), int32(3), int32(184), v24)
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
	v37 = F_systable_beginscan(m, v29, int32(2701), v32, int32(0), v32, v17+int32(48))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	v45 = F_AllocSetContextCreateInternal(m, v40, int32(_a_F_CloneRowTriggersToPartition_0), int32(0), int32(1024), int32(_a_F_CloneRowTriggersToPartition_1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = F_systable_getnext(m, v37)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L62
	}
L7:
	;
	if v47 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v47
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_MemoryContextDelete(m, v45)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L59
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+80)))
	if v66&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v311 = F_systable_getnext(m, v37)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L57
	}
L14:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+83)))
	if v71 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	switch v66 & int32(66) {
	case 0, 2:
		goto L16
	default:
		goto L17
	}
L16:
	;
	v89 = int32(0)
	v90 = int32(_a_F_CloneRowTriggersToPartition_2)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v45
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v98 = F_heap_getattr_6(m, v57, int32(17), v95, v17+int32(47))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v65 + int32(12)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_3), v17)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_5), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	v100 = int32(0)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
	if v101 == v100 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v104 = F_text_to_cstring(m, v98)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v114 = v100
	goto L24
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v65)+116))
	if int32(0) < v115 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v106 = F_stringToNode(m, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v109 = F_map_partition_varattnos(m, v106, int32(1), l1, l0)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v112 = F_map_partition_varattnos(m, v109, int32(2), l1, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v114 = v112
	goto L24
L29:
	;
	v123 = int32(0)
	v128 = v89
	goto L32
L30:
	;
	v166 = v89
	goto L31
L31:
	;
	v173 = int32(0)
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+98)))
	if v174 <= v173 {
		v235 = v173
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65+int32(124)+v123<<(uint(int32(1))%32)))))
	v149 = F_pstrdup(m, v135+v136<<(uint(int32(4))%32)+v143*int32(100)-int32(76))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v166 = v153
	goto L31
L34:
	;
	v151 = F_makeString(m, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v153 = F_lappend(m, v128, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v156 = v123 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v65)+116))
	if v156 < v157 {
		v123 = v156
		v128 = v153
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v244 = F_palloc0(m, int32(52))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L54
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v181 = F_heap_getattr_6(m, v57, int32(16), v178, v17+int32(47))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
	if v183 == int32(1) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v186 = F_pg_detoast_datum_packed(m, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v189 = F_pg_detoast_datum_packed(m, v181)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+98)))
	if v191 <= int32(0) {
		v235 = v173
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v194 = int32(1)
	if v188&v194 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v198 = v194
	goto L47
L46:
	;
	v198 = int32(4)
	goto L47
L47:
	;
	v203 = v189 + v198
	v206 = int32(0)
	v207 = v173
	goto L48
L48:
	;
	v215 = F_pstrdup(m, v203)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v235 = v219
	goto L38
L50:
	;
	v217 = F_makeString(m, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v219 = F_lappend(m, v207, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v221 = F_strlen(m, v203)
	mBase = m.M
	v223 = int32(1)
	v226 = v206 + v223
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+98)))
	if v226 < v227 {
		v203 = v221 + v203 + v223
		v206 = v226
		v207 = v219
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v246 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+4)) = uint8(v246)
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = int32(181)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v65)+92))
	v251 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+24)) = uint8(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v244)+20)) = v235
	v254 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v244)+12)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v65 + int32(12)
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+5)) = uint8(base.B2i32(v250 != v246))
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+80)))
	v264 = v262 & int32(66)
	*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v264)
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v244)+36)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v244)+32)) = v166
	v271 = v266 & int32(60)
	*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v271)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+44)) = uint8(v273)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+97)))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+48)) = v246
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+45)) = uint8(v275)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v290 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+82)))
	F_CreateTriggerFiringOn(m, v17+int32(32), v244, v246, v282, v283, v246, v246, v286, v287, v114, v246, v251, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v91
	F_MemoryContextReset(m, v45)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L13
L57:
	;
	if v311 != 0 {
		v57 = v311
		goto L11
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_sequence_close(m, v29, int32(3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	m.G0 = v17 + int32(96)
	return
L62:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v65 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v341 + int32(4)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_7), v17+int32(16))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_8), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clamp_cardinality_to_long(m *base.Module, l0 float64) int32 {
	var v13 float64
	_ = v13
	var v16 float64
	_ = v16
	var v20 int32
	_ = v20
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
		return int32(2147483647)
	} else {
		if base.F64_le(l0, float64(0)) != 0 {
			return int32(0)
		} else {
			v13 = float64(2.147483647e+09)
			if base.F64_lt(l0, v13) != 0 {
				v16 = l0
			} else {
				v16 = v13
			}
			if base.F64_lt(base.F64_abs(v16), float64(2.147483648e+09)) != 0 {
				v20 = base.I32_trunc_f64_s(v16)
				return v20
			} else {
				return int32(-2147483648)
			}
		}
	}
}
func F_clamp_row_est(m *base.Module, l0 float64) float64 {
	var v3 float64
	_ = v3
	var v11 float64
	_ = v11
	var v15 float64
	_ = v15
	v3 = float64(1e+100)
	if base.F64_gt(l0, v3) != 0 {
		v15 = v3
	} else {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v15 = v3
		} else {
			v11 = float64(1)
			if base.F64_le(l0, v11) != 0 {
				v15 = v11
			} else {
				v15 = base.F64_nearest(l0)
			}
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
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
	var v205 int32
	_ = v205
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
	v38 = F_find_single_rel_for_clauses(m, l0, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
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
	if l5 == int32(0) {
		v53 = v37
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 == int32(0) {
		v268 = v53
		goto L1
	} else {
		goto L14
	}
L9:
	;
	if v38 == int32(0) {
		v53 = v37
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	if v44 != 0 {
		v53 = v37
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+112))
	if v45 == int32(0) {
		v53 = v37
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = F_statext_clauselist_selectivity(m, l0, l1, l2, l3, l4, v38, v20+int32(12), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v53 = v51
	goto L8
L14:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v56 < v57 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = v56
	v73 = int32(-1)
	v75 = v53
	goto L18
L16:
	;
	v192 = v53
	goto L17
L17:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v195 == int32(0) {
		v268 = v192
		goto L1
	} else {
		goto L57
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v67<<(uint(int32(2))%32))))
	v84 = v73 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v86 = F_bms_is_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v192 = v173
	goto L17
L20:
	;
	v175 = v67 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v175 < v176 {
		v67 = v175
		v73 = v84
		v75 = v173
		goto L18
	} else {
		goto L56
	}
L21:
	;
	if v86 != 0 {
		v173 = v75
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v88 = F_clause_selectivity_ext(m, l0, v82, l2, l3, l4, l5)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v91 != int32(318) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v173 = base.F64_mul(v75, v88)
	goto L20
L25:
	;
	if v103 != int32(17) {
		goto L24
	} else {
		goto L31
	}
L26:
	;
	v101 = v82
	v102 = int32(0)
	v103 = v91
	goto L25
L27:
	;
	goto L28
L28:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+10)))
	if v94 == int32(1) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v97 == int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v101 = v97
	v102 = v82
	v103 = v100
	goto L25
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	if v106 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 != int32(2) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v147 = F_get_oprrest(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L53
	}
L35:
	;
	v145 = int32(0)
	goto L34
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	if v112 != int32(1) {
		goto L24
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v127 = F_NumRelids(m, l0, v101)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L44
	}
L39:
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
		goto L40
	}
L40:
	;
	if v119 != 0 {
		v145 = int32(1)
		goto L34
	} else {
		goto L41
	}
L41:
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
		goto L42
	}
L42:
	;
	if v125 != 0 {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L24
L44:
	;
	if v127 != int32(1) {
		goto L24
	} else {
		goto L45
	}
L45:
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
		goto L46
	}
L46:
	;
	if v135 != 0 {
		v145 = int32(1)
		goto L34
	} else {
		goto L47
	}
L47:
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
		goto L48
	}
L48:
	;
	if v140 == int32(0) {
		goto L24
	} else {
		goto L49
	}
L49:
	;
	goto L35
L50:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(0), v88)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L55
	}
L51:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(1), v88)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L54
	}
L52:
	;
	switch v147 - int32(336) {
	case 0:
		goto L51
	case 1:
		goto L50
	default:
		goto L24
	}
L53:
	;
	switch v147 - int32(103) {
	case 0:
		goto L51
	case 1:
		goto L50
	default:
		goto L52
	}
L54:
	;
	v173 = v75
	goto L20
L55:
	;
	v173 = v75
	goto L20
L56:
	;
	goto L19
L57:
	;
	v205 = v195
	v212 = v192
	goto L58
L58:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
	if v215 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v268 = v253
	goto L1
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	F_pfree(m, v205)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L74
	}
L61:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+9)))
	if v218 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v205)+24))
	v247 = v246
	goto L60
L64:
	;
	v221 = float64(0.005)
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v205)+24))
	if base.F64_eq(v222, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L60
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v205)+16))
	v247 = v245
	goto L60
L67:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v205)+16))
	if base.F64_eq(v225, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L60
	} else {
		goto L68
	}
L68:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v233 = F_nulltestsel(m, l0, int32(0), v232, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v235 = base.F64_add(base.F64_add(base.F64_add(v222, v225), float64(-1)), v233)
	if base.F64_le(v235, float64(0)) == int32(0) {
		v247 = v235
		goto L60
	} else {
		goto L70
	}
L70:
	;
	if base.F64_lt(v235, float64(-0.01)) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v244 = float64(0.005)
	goto L73
L72:
	;
	v244 = float64(1e-10)
	goto L73
L73:
	;
	v247 = v244
	goto L60
L74:
	;
	v253 = base.F64_mul(v212, v247)
	if v250 != 0 {
		v205 = v250
		v212 = v253
		goto L58
	} else {
		goto L75
	}
L75:
	;
	goto L59
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
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
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
	var v178 int32
	_ = v178
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
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
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
	var v299 int32
	_ = v299
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
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
	var v389 int32
	_ = v389
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
		goto L26
	} else {
		goto L27
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
	if l7 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v46 = F__emscripten_memset_bulkmem(m, v30, base.I32_extend8_s(int32(0)), l7)
	mBase = m.M
	goto L25
L21:
	;
	v51 = v30
	goto L9
L22:
	;
	v42 = F__emscripten_memcpy_bulkmem(m, v30, l6, l7)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))) = uint8(v49)
	v51 = v30
	goto L9
L26:
	;
	if l5 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L27:
	;
	v68 = v56
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	if v74 != 0 {
		goto L26
	} else {
		goto L30
	}
L29:
	;
	goto L26
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	switch v76 - int32(76) {
	case 0, 18, 21, 38:
		goto L34
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L33
	default:
		goto L35
	}
L31:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v457 != 0 {
		v68 = v457
		goto L28
	} else {
		goto L151
	}
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v120))))
	if v122 != 0 {
		goto L31
	} else {
		goto L45
	}
L33:
	;
	F_cparc(m, l0, v68, l2, v75)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v81 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	if v76 != int32(36) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v85 = v81
	goto L38
L38:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v98 - int32(76) {
	case 0, 18, 21, 38:
		goto L32
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L40
	default:
		goto L41
	}
L39:
	;
	goto L33
L40:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v103 != 0 {
		v85 = v103
		goto L38
	} else {
		goto L43
	}
L41:
	;
	if v98 == int32(36) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L39
L44:
	;
	goto L31
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v123 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if l4 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	v154 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v128 = v123
	goto L50
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	if v142 == v75 {
		v154 = v141
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v154 = int32(0)
	goto L46
L52:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v144 != 0 {
		v128 = v144
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v436 = F_newstate(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L148
	}
L55:
	;
	if v154 != 0 {
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v167 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v76 != v162 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v164 == v165 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	v169 = v167
	v178 = l2
	goto L63
L61:
	;
	goto L62
L62:
	;
	if v154 == int32(0) {
		goto L54
	} else {
		goto L70
	}
L63:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	if v182 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L62
L65:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	if v191 != 0 {
		v169 = v191
		v178 = v190
		goto L63
	} else {
		goto L69
	}
L66:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v76 != v185 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if v187 == v188 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	goto L64
L70:
	;
	F_cparc(m, l0, v68, l2, v154)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L31
L72:
	;
	goto L75
L73:
	;
	goto L74
L74:
	;
	F_clonesuccessorstates(m, l0, v75, l2, l3, l4, v51, l6, l7)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L147
	}
L75:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v238 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L107
L77:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+4)))
	if v246 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	goto L75
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	if v280 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v251 = v249 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v251) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if int32(1)<<(uint(v251)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v260 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	if v261 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v273 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v246*int32(24))+12)) = v269
	v273 = v269
	goto L86
L88:
	;
	goto L89
L89:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+32)) = v271
	v273 = v271
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+36)) = v261
	goto L92
L91:
	;
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = int64(0)
	goto L81
L93:
	;
	if v279 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+20)) = v279
	goto L93
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v279
	goto L93
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v280
	goto L99
L98:
	;
	goto L99
L99:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+12)) = v286 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v238)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v238)+28))
	if v291 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v297 = v238 + int32(8)
	if v290 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v290
	goto L100
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v290
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v291
	goto L106
L105:
	;
	goto L106
L106:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v299 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = int32(0)
	v305 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v305
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v238
	goto L80
L107:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	if v328 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(-1)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v154)+32))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v409 != 0 {
		goto L140
	} else {
		goto L141
	}
L109:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v336 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328)+4)))
	if v336 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	goto L107
L113:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v328)+20))
	if v370 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v341 = v339 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v341) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if int32(1)<<(uint(v341)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v350 != 0 {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v328)+36))
	if v351 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v363 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v355+v336*int32(24))+12)) = v359
	v363 = v359
	goto L118
L120:
	;
	goto L121
L121:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+32)) = v361
	v363 = v361
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+36)) = v351
	goto L124
L123:
	;
	goto L124
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v328)+32)) = int64(0)
	goto L113
L125:
	;
	if v369 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+20)) = v369
	goto L125
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+16)) = v369
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+20)) = v370
	goto L131
L130:
	;
	goto L131
L131:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+12)) = v376 - int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	if v381 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v387 = v328 + int32(8)
	if v380 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+16)) = v380
	goto L132
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+24)) = v380
	goto L132
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = v381
	goto L138
L137:
	;
	goto L138
L138:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+8)) = v389 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(0)
	v395 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v387)+16)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v387))) = v395
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+16)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v328
	goto L112
L139:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v408 != 0 {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+32)) = v408
	goto L139
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v408
	goto L139
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v154
	goto L74
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+28)) = v412
	goto L143
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v412
	goto L143
L147:
	;
	goto L31
L148:
	;
	if v436 == int32(0) {
		goto L26
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = v75
	F_cparc(m, l0, v68, l2, v436)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L31
L151:
	;
	goto L29
L152:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v474 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	return
L155:
	;
	F_pfree(m, v51)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L165
	}
L156:
	;
	v478 = v474
	goto L157
L157:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+12))
	if v492 != 0 {
		goto L155
	} else {
		goto L159
	}
L158:
	;
	goto L155
L159:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	if v494 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v493)+24)) = v495
	F_clonesuccessorstates(m, l0, v494, v493, l3, l4, v495, v51, l7)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v478)+16))
	if v500 != 0 {
		v478 = v500
		goto L157
	} else {
		goto L164
	}
L163:
	;
	goto L162
L164:
	;
	goto L158
L165:
	;
	goto L154
}

package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_FindLogicalRepLocalIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v248 int32
	_ = v248
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
	if v13 == int32(112) {
		v248 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v248
L2:
	;
	v16 = F_RelationGetReplicaIndex(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v16 != 0 {
		v248 = v16
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = F_RelationGetPrimaryKeyIndex(m, l0, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v21 != 0 {
		v248 = v21
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v23 != int32(102) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v248 = int32(0)
	goto L1
L9:
	;
	v27 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v27 == int32(0) {
		v248 = int32(0)
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v31 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v35<<(uint(int32(2))%32))))
	v52 = F_index_open(m, v50, int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L8
L15:
	;
	v229 = v35 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v229 < v230 {
		v35 = v229
		goto L13
	} else {
		goto L52
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+192))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+18)))
	if v55 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_relation_close(m, v52, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+196))
	v66 = F_heap_attisnull(m, v63, int32(21), v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L23
	}
L20:
	;
	goto L15
L21:
	;
	F_relation_close(m, v52, int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L50
	}
L22:
	;
	v213 = v199
	goto L21
L23:
	;
	if v66 == int32(0) {
		v199 = v61
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52)+196))
	v73 = F_SysCacheGetAttrNotNull(m, int32(34), v71, int32(18))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v52)+192))
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+10)))
	if int32(0) < v76 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v61
	goto L29
L27:
	;
	v117 = v75
	goto L28
L28:
	;
	v123 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v123 < v125 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v73)+int32(24)+v85<<(uint(int32(2))%32))))
	v97 = F_get_opclass_family(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L31
	}
L30:
	;
	v117 = v109
	goto L28
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+84))
	v103 = F_IndexAmTranslateCompareType(m, int32(3), v101, v97, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v103 == int32(0) {
		v199 = v61
		goto L22
	} else {
		goto L33
	}
L33:
	;
	v108 = v85 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v52)+192))
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v109)+10)))
	if v108 < v110 {
		v85 = v108
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v131 = v123
	v134 = v124
	v137 = v125
	goto L39
L36:
	;
	v168 = v117
	goto L37
L37:
	;
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v168)+48)))
	if v169 == int32(0) {
		v199 = v61
		goto L22
	} else {
		goto L46
	}
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v52)+192))
	v168 = v156
	goto L37
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134+v137<<(uint(int32(3))%32)+v131*int32(100))+96))
	v147 = F_lookup_type_cache(m, v145, int32(32))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L41
	}
L40:
	;
	v213 = int32(0)
	goto L21
L41:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	if v149 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v151 = v131 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v151 < v153 {
		v131 = v151
		v134 = v152
		v137 = v153
		goto L39
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L40
L45:
	;
	goto L38
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v172 < v169 {
		v199 = v61
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v174+v169<<(uint(int32(1))%32)-int32(2)))))
	if v180 < int32(0) {
		v199 = v61
		goto L22
	} else {
		goto L48
	}
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+84))
	v186 = F_GetIndexAmRoutineByAmId(m, v184, int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186)+100))
	v199 = base.B2i32(v188 != int32(0))
	goto L22
L50:
	;
	if v213 != 0 {
		v248 = v50
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L15
L52:
	;
	goto L14
}
func F_LogicalDecodingCtlShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_LogicalDecodingCtlShmemRequest_0), int64(3), int32(_a_F_LogicalDecodingCtlShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_LogicalIncreaseRestartDecodingForSlot(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v2 int64
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v1 = l0
	v2 = l1
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalIncreaseRestartDecodingForSlot[0]))
	v16 = base.AtomicRmwXchg32(m, v13, int32(0), int32(1))
	if v16 != 0 {
		F_s_lock(m, v13, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
			if base.Ui64(v2) <= base.Ui64(v20) {
				v22 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v22))
				m.G0 = v10 - int32(-64)
				return
			} else {
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
				if base.Ui64(v1) <= base.Ui64(v25) {
					*(*int64)(unsafe.Add(mBase, uint32(v13)+256)) = v2
					*(*int64)(unsafe.Add(mBase, uint32(v13)+248)) = v1
					v29 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v29))
					F_LogicalConfirmReceivedLocation(m, v25)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						m.G0 = v10 - int32(-64)
						return
					}
				} else {
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v13)+248))
					if v34 == int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v13)+256)) = v2
						*(*int64)(unsafe.Add(mBase, uint32(v13)+248)) = v1
						v39 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v39))
						v44 = F_errstart(m, int32(14), v39)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							if v44 == int32(0) {
								m.G0 = v10 - int32(-64)
								return
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+12)) = uint32(v1)
								v49 = int64(32)
								v50 = int64(base.Ui64(v1) >> (uint(v49) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+8)) = uint32(v50)
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v2)
								v54 = int64(base.Ui64(v2) >> (uint(v49) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v54)
								F_errmsg_internal(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_1), v10)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_2), int32(1872), int32(_a_F_LogicalIncreaseRestartDecodingForSlot_3))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										m.G0 = v10 - int32(-64)
										return
									}
								}
							}
						}
					} else {
						v64 = *(*int64)(unsafe.Add(mBase, uint32(v13)+256))
						v65 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v65))
						v70 = F_errstart(m, int32(14), v65)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 == int32(0) {
								m.G0 = v10 - int32(-64)
								return
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+52)) = uint32(v25)
								v75 = int64(32)
								v76 = int64(base.Ui64(v25) >> (uint(v75) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+48)) = uint32(v76)
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+44)) = uint32(v34)
								v80 = int64(base.Ui64(v34) >> (uint(v75) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+40)) = uint32(v80)
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v64)
								v84 = int64(base.Ui64(v64) >> (uint(v75) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v84)
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v1)
								v88 = int64(base.Ui64(v1) >> (uint(v75) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v88)
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v2)
								v92 = int64(base.Ui64(v2) >> (uint(v75) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v92)
								F_errmsg_internal(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_4), v8+int32(-48))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_2), int32(1890), int32(_a_F_LogicalIncreaseRestartDecodingForSlot_3))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										m.G0 = v10 - int32(-64)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
		if base.Ui64(v2) <= base.Ui64(v20) {
			v22 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v22))
			m.G0 = v10 - int32(-64)
			return
		} else {
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
			if base.Ui64(v1) <= base.Ui64(v25) {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+256)) = v2
				*(*int64)(unsafe.Add(mBase, uint32(v13)+248)) = v1
				v29 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v29))
				F_LogicalConfirmReceivedLocation(m, v25)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v10 - int32(-64)
					return
				}
			} else {
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v13)+248))
				if v34 == int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v13)+256)) = v2
					*(*int64)(unsafe.Add(mBase, uint32(v13)+248)) = v1
					v39 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v39))
					v44 = F_errstart(m, int32(14), v39)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						if v44 == int32(0) {
							m.G0 = v10 - int32(-64)
							return
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+12)) = uint32(v1)
							v49 = int64(32)
							v50 = int64(base.Ui64(v1) >> (uint(v49) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+8)) = uint32(v50)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v2)
							v54 = int64(base.Ui64(v2) >> (uint(v49) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v54)
							F_errmsg_internal(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_1), v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_2), int32(1872), int32(_a_F_LogicalIncreaseRestartDecodingForSlot_3))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v10 - int32(-64)
									return
								}
							}
						}
					}
				} else {
					v64 = *(*int64)(unsafe.Add(mBase, uint32(v13)+256))
					v65 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v65))
					v70 = F_errstart(m, int32(14), v65)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						if v70 == int32(0) {
							m.G0 = v10 - int32(-64)
							return
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+52)) = uint32(v25)
							v75 = int64(32)
							v76 = int64(base.Ui64(v25) >> (uint(v75) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+48)) = uint32(v76)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+44)) = uint32(v34)
							v80 = int64(base.Ui64(v34) >> (uint(v75) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+40)) = uint32(v80)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v64)
							v84 = int64(base.Ui64(v64) >> (uint(v75) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v84)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v1)
							v88 = int64(base.Ui64(v1) >> (uint(v75) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v88)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v2)
							v92 = int64(base.Ui64(v2) >> (uint(v75) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v92)
							F_errmsg_internal(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_4), v8+int32(-48))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_LogicalIncreaseRestartDecodingForSlot_2), int32(1890), int32(_a_F_LogicalIncreaseRestartDecodingForSlot_3))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									m.G0 = v10 - int32(-64)
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
func F_LogicalOutputPrepareWrite(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v7
	return
}
func F_LogicalRepWorkersWakeupAtCommit(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = int32(_a_F_LogicalRepWorkersWakeupAtCommit_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalRepWorkersWakeupAtCommit[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalRepWorkersWakeupAtCommit[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalRepWorkersWakeupAtCommit[0])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalRepWorkersWakeupAtCommit[2]))
	v11 = F_list_append_unique_oid(m, v10, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_LogicalRepWorkersWakeupAtCommit[0])) = v4
		*(*int32)(unsafe.Add(mBase, _c_F_LogicalRepWorkersWakeupAtCommit[2])) = v11
		return
	}
}
func F_LogicalTapeWrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_palloc(m, int32(_a_F_LogicalTapeWrite_0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v18 == int64(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(_a_F_LogicalTapeWrite_0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v12
	goto L3
L6:
	;
	v21 = F_ltsGetBlock(m, v7, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v21
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_LogicalTapeWrite[0]))) = int64(-1)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L32
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v31 = l1
	v32 = l2
	v33 = v29
	goto L14
L12:
	;
	goto L13
L13:
	;
	return
L14:
	;
	if int32(_a_F_LogicalTapeWrite_1) <= v33 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	v58 = v33
	goto L18
L18:
	;
	v61 = int32(_a_F_LogicalTapeWrite_1) - v58
	if base.Ui32(v61) < base.Ui32(v32) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = F_ltsGetBlock(m, v41, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_LogicalTapeWrite[1]))) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ltsWriteBlock(m, v46, v47, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_LogicalTapeWrite[0]))) = v52
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v42
	v58 = int32(0)
	goto L18
L22:
	;
	v63 = v61
	goto L24
L23:
	;
	v63 = v32
	goto L24
L24:
	;
	if v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	base.MemoryCopy(m, v64+v58, v31, v63)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v70 = v69 + v63
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v72 < v70 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v70
	goto L30
L29:
	;
	goto L30
L30:
	;
	v76 = v32 - v63
	if v76 != 0 {
		v31 = v31 + v63
		v32 = v76
		v33 = v70
		goto L14
	} else {
		goto L31
	}
L31:
	;
	goto L15
L32:
	;
	F_errmsg_internal(m, int32(_a_F_LogicalTapeWrite_2), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_LogicalTapeWrite_3), int32(797), int32(_a_F_LogicalTapeWrite_4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logical_rewrite_log_mapping(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v10 = m.G0
	v12 = v10 - int32(1088)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1084)) = l1
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v23 = F_hash_search(m, v17, v12+int32(1084), int32(1), v12+int32(1083))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1083)))
		if v25 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_logical_rewrite_log_mapping[0]))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+117)))
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1084))
			v35 = F_GetCurrentTransactionId(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v34
				*(*uint32)(unsafe.Add(mBase, uint32(v12)+32)) = uint32(v33)
				v41 = int64(base.Ui64(v33) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v12)+28)) = uint32(v41)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v16
				if v32 != 0 {
					v45 = int32(0)
				} else {
					v45 = v29
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_logical_rewrite_log_mapping_0)
				v50 = v12 + int32(48)
				v55 = F_pg_snprintf(m, v50, int32(1024), int32(_a_F_logical_rewrite_log_mapping_1), v12+int32(16))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
					v60 = v23 + int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v60
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v60
					*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = int64(0)
					base.MemoryCopy(m, v23+int32(28), v50, int32(1024))
					v70 = F_PathNameOpenFile(m, v50, int32(193))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v70
						if v70 < int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(48)
									F_errmsg(m, int32(_a_F_logical_rewrite_log_mapping_2), v12)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_logical_rewrite_log_mapping_3), int32(980), int32(_a_F_logical_rewrite_log_mapping_4))
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v82 = F_MemoryContextAlloc(m, v80, int32(44))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v84
								v86 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v82)+24)) = v86
								v88 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v82)+16)) = v88
								v90 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v90
								v92 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
								*(*int64)(unsafe.Add(mBase, uint32(v82))) = v92
								v95 = v23 + int32(16)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
								if v96 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v95
									*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v95
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v95
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v104
								v107 = v82 + int32(36)
								*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v107
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
								v111 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v110 + v111
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v116 = v114 + v111
								*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v116
								if base.Ui32(int32(1000)) <= base.Ui32(v116) {
									F_logical_heap_rewrite_flush_mappings(m, l0)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										m.G0 = v12 + int32(1088)
										return
									}
								} else {
									m.G0 = v12 + int32(1088)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v82 = F_MemoryContextAlloc(m, v80, int32(44))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v84
				v86 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v82)+24)) = v86
				v88 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v82)+16)) = v88
				v90 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v90
				v92 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v82))) = v92
				v95 = v23 + int32(16)
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
				if v96 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v95
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v95
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v104
				v107 = v82 + int32(36)
				*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v107
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v107
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
				v111 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v110 + v111
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				v116 = v114 + v111
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v116
				if base.Ui32(int32(1000)) <= base.Ui32(v116) {
					F_logical_heap_rewrite_flush_mappings(m, l0)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						m.G0 = v12 + int32(1088)
						return
					}
				} else {
					m.G0 = v12 + int32(1088)
					return
				}
			}
		}
	}
}

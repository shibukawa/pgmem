package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LockErrorCleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(_a_F_LockErrorCleanup_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[0])) = v15 + int32(1)
	F_AbortStrongLockAcquire(m)
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
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[1]))
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(2)
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	F_disable_timeouts(m, v11)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v147 = int32(_a_F_LockErrorCleanup_0)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[0])) = v149 - int32(1)
	m.G0 = v11 + int32(16)
	return
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[2]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v42 = v34 + v35&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockErrorCleanup_1)
	v44 = F_LWLockAcquire(m, v42, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[3]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[1])) = int32(0)
	F_LWLockRelease(m, v42)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	F_RemoveFromWaitQueue(m, v47, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v52 != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[4]))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_LockErrorCleanup[1]))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+32)) = v57 + int64(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	if int32(0) < v62 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L8
L15:
	;
	v68 = int32(0)
	goto L18
L16:
	;
	v94 = int32(0)
	goto L17
L17:
	;
	v97 = v61 + v94<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v54
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v101 + int32(1)
	if v54 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v75 = v61 + v68<<(uint(int32(4))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v54 == v76 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v94 = v62
	goto L17
L20:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v78 + int64(1)
	goto L14
L21:
	;
	goto L22
L22:
	;
	v83 = v68 + int32(1)
	if v83 != v62 {
		v68 = v83
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+18)))
	if base.Ui32(v106) <= base.Ui32(int32(15)) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	goto L14
L27:
	;
	goto L26
L28:
	;
	if v106 != int32(15) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54+v106<<(uint(int32(2))%32))+292)) = v56
	goto L33
L32:
	;
	goto L33
L33:
	;
	v116 = v106 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+18)) = uint8(v116)
	goto L30
L34:
	;
	goto L5
}
func F_LockRefindAndRelease(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[0]))
	v20 = F_get_hash_value(m, v19, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[1]))
		v30 = v23 + v20&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockRefindAndRelease_0)
		v32 = F_LWLockAcquire(m, v30, int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[0]))
			v36 = int32(0)
			v38 = F_hash_search_with_hash_value(m, v35, l2, v20, v36, v36)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l1
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[2]))
					v49 = int32(0)
					v51 = F_hash_search_with_hash_value(m, v43, v16+int32(8), v20^l1<<(uint(int32(4))%32), v49, v49)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						if v51 == int32(0) {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_LockRefindAndRelease_1), int32(0))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_LockRefindAndRelease_2), int32(3296), int32(_a_F_LockRefindAndRelease_3))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v56 = int32(1) << (uint(l3) % 32)
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
							if v56&v57 == int32(0) {
								F_LWLockRelease(m, v30)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v65 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										if v65 == int32(0) {
											m.G0 = v16 + int32(16)
											return
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+l3<<(uint(int32(2))%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v73
											F_errmsg_internal(m, int32(_a_F_LockRefindAndRelease_4), v16)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_LockRefindAndRelease_2), int32(3307), int32(_a_F_LockRefindAndRelease_3))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													m.G0 = v16 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
								v84 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v83 - v84
								v88 = l3 << (uint(int32(2)) % 32)
								v89 = v38 + v88
								v91 = v89 + int32(44)
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
								*(*int32)(unsafe.Add(mBase, uint32(v91))) = v92 - v84
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v38)+128))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = v96 - v84
								v101 = v89 + int32(88)
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
								v104 = v102 - v84
								*(*int32)(unsafe.Add(mBase, uint32(v101))) = v104
								v107 = v56 ^ int32(-1)
								if v104 == int32(0) {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v110 & v107
								} else {
								}
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v113+v88)))
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v117 & v107
								F_CleanUpLock(m, v38, v51, l0, v20, base.B2i32(v116&v115 != int32(0)))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return
								} else {
									F_LWLockRelease(m, v30)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										if l4 == int32(0) {
											m.G0 = v16 + int32(16)
											return
										} else {
											v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
											if v129 != int32(1) {
												m.G0 = v16 + int32(16)
												return
											} else {
												v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)))
												if v132|base.B2i32(l3 < int32(5)) != 0 {
													m.G0 = v16 + int32(16)
													return
												} else {
													v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													if v136 == int32(0) {
														m.G0 = v16 + int32(16)
														return
													} else {
														v140 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[3]))
														v143 = base.AtomicRmwXchg32(m, v140, int32(0), int32(1))
														if v143 != 0 {
															v145 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[3]))
															F_s_lock(m, v145, int32(_a_F_LockRefindAndRelease_2), int32(3330), int32(_a_F_LockRefindAndRelease_3))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return
															} else {
																v152 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[3]))
																v157 = v152 + v20&int32(1023)<<(uint(int32(2))%32)
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v158 - int32(1)
																v162 = int32(0)
																atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v152))), uint32(v162))
																m.G0 = v16 + int32(16)
																return
															}
														} else {
															v152 = *(*int32)(unsafe.Add(mBase, _c_F_LockRefindAndRelease[3]))
															v157 = v152 + v20&int32(1023)<<(uint(int32(2))%32)
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v158 - int32(1)
															v162 = int32(0)
															atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v152))), uint32(v162))
															m.G0 = v16 + int32(16)
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
				} else {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_LockRefindAndRelease_5), int32(0))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_LockRefindAndRelease_2), int32(3280), int32(_a_F_LockRefindAndRelease_3))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
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
		}
	}
}
func F_LookupExplicitNamespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_LookupExplicitNamespace_0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LookupExplicitNamespace[0])))
	if base.B2i32(v13 == int32(0))|base.B2i32(v13 != v16) != 0 {
		v34 = v13
		v35 = v16
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L30
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v73
L3:
	;
	if v34-v35 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	v19 = l0
	v20 = v10
	goto L6
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v24
		v35 = v23
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v34 = v24
	v35 = v23
	goto L4
L8:
	;
	v27 = int32(1)
	if v24 == v23 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_LookupExplicitNamespace[1]))
	if v40 != 0 {
		v73 = v40
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v43 = int32(0)
	v46 = F_GetSysCacheOid(m, int32(37), l0, v43, v43, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	if l1|v46 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if l1 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v53 = int32(0)
	if v46 == v53 {
		v73 = v53
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_LookupExplicitNamespace[2]))
	v61 = F_object_aclcheck(m, int32(2615), v46, v59, int64(256))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v61 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_aclcheck_error(m, v61, int32(36), l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_LookupExplicitNamespace[3]))
	if v67 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v69 = F_RunNamespaceSearchHook(m, v46, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L14
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v73 = v46
	goto L2
L29:
	;
	goto L28
L30:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(_a_F_LookupExplicitNamespace_1), v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_LookupExplicitNamespace_2), int32(3547), int32(_a_F_LookupExplicitNamespace_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___letf2(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32 {
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	v8 = int32(1)
	v12 = l1 & int64(9223372036854775807)
	v13 = int64(9223090561878065152)
	if v12 == v13 {
		v17 = base.B2i32(l0 != int64(0))
	} else {
		v17 = base.B2i32(base.Ui64(v13) < base.Ui64(v12))
	}
	if v17 != 0 {
		v60 = v8
		return v60
	} else {
		v21 = l3 & int64(9223372036854775807)
		v22 = int64(9223090561878065152)
		if v21 == v22 {
			v26 = base.B2i32(l2 != int64(0))
		} else {
			v26 = base.B2i32(base.Ui64(v22) < base.Ui64(v21))
		}
		if v26 != 0 {
			v60 = v8
			return v60
		} else {
			if l0|l2|(v12|v21) == int64(0) {
				return int32(0)
			} else {
				if int64(0) <= l1&l3 {
					if l1 == l3 {
						v40 = base.B2i32(base.Ui64(l0) < base.Ui64(l2))
					} else {
						v40 = base.B2i32(l1 < l3)
					}
					if v40 != 0 {
						return int32(-1)
					} else {
						return base.B2i32(l0^l2|(l1^l3) != int64(0))
					}
				} else {
					if l1 == l3 {
						v52 = base.B2i32(base.Ui64(l2) < base.Ui64(l0))
					} else {
						v52 = base.B2i32(l3 < l1)
					}
					if v52 != 0 {
						return int32(-1)
					} else {
						v60 = base.B2i32(l0^l2|(l1^l3) != int64(0))
						return v60
					}
				}
			}
		}
	}
}
func F__lt_q_regex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v19 = F_ArrayGetNItemsSafe(m, v16, v13+int32(16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v24 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L9:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if int32(0) < v19 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v73 != v8 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v72 = int32(0)
	goto L11
L15:
	;
	v34 = v15
	goto L17
L16:
	;
	v34 = (v16<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L17
L17:
	;
	v38 = v13 + v34
	v39 = v19
	goto L18
L18:
	;
	v45 = F_array_iterator(m, v8, int32(_a_F__lt_q_regex_0), v38, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	if v45 != 0 {
		v72 = int32(1)
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v55 = int32(1)
	if v55 < v39 {
		v38 = v38 + (int32(base.Ui32(v47)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v39 = v39 - v55
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	F_pfree(m, v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v77 != v13 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_pfree(m, v13)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	return v72
L30:
	;
	goto L29
L31:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F__lt_q_regex_1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F__lt_q_regex_2), int32(146), int32(_a_F__lt_q_regex_3))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F__lt_q_regex_4), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F__lt_q_regex_2), int32(150), int32(_a_F__lt_q_regex_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_l1_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 float32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 float32
	_ = v55
	var v57 float32
	_ = v57
	var v61 float32
	_ = v61
	var v63 float32
	_ = v63
	var v67 float32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v86 float32
	_ = v86
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v92 float32
	_ = v92
	var v105 float32
	_ = v105
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v10 = float32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
			if v23 == v24 {
				v26 = base.I32_extend16_s(v23)
				if v26 <= int32(0) {
					v117 = float64(0)
				} else {
					v30 = int32(8)
					v31 = v21 + v30
					v33 = v16 + v30
					if v26 == int32(1) {
						v77 = int32(0)
						v86 = v10
						v88 = v77 << (uint(int32(2)) % 32)
						v90 = *(*float32)(unsafe.Add(mBase, uint32(v33+v88)))
						v92 = *(*float32)(unsafe.Add(mBase, uint32(v88+v31)))
						v105 = base.F32_add(base.F32_abs(base.F32_sub(v90, v92)), v86)
					} else {
						v40 = int32(0)
						v48 = int32(0)
						v49 = v10
						for {
							v50 = int32(2)
							v51 = v40 << (uint(v50) % 32)
							v53 = v51 | int32(4)
							v55 = *(*float32)(unsafe.Add(mBase, uint32(v33+v53)))
							v57 = *(*float32)(unsafe.Add(mBase, uint32(v31+v53)))
							v61 = *(*float32)(unsafe.Add(mBase, uint32(v33+v51)))
							v63 = *(*float32)(unsafe.Add(mBase, uint32(v31+v51)))
							v67 = base.F32_add(base.F32_abs(base.F32_sub(v55, v57)), base.F32_add(base.F32_abs(base.F32_sub(v61, v63)), v49))
							v69 = v40 + v50
							v71 = v48 + v50
							if v71 != v26&int32(_a_F_l1_distance_0) {
								v40 = v69
								v48 = v71
								v49 = v67
								continue
							} else {
								break
							}
							break
						}
						if v26&int32(1) == int32(0) {
							v105 = v67
						} else {
							v77 = v69
							v86 = v67
							v88 = v77 << (uint(int32(2)) % 32)
							v90 = *(*float32)(unsafe.Add(mBase, uint32(v33+v88)))
							v92 = *(*float32)(unsafe.Add(mBase, uint32(v88+v31)))
							v105 = base.F32_add(base.F32_abs(base.F32_sub(v90, v92)), v86)
						}
					}
					v117 = base.F64_promote_f32(v105)
				}
				v118 = F_Float8GetDatum(m, v117)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(16)
					return v118
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
						v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v132
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v131
						F_errmsg(m, int32(_a_F_l1_distance_1), v13)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_l1_distance_2), int32(76), int32(_a_F_l1_distance_3))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_lastval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[0]))
	if v11 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v14 = int32(0)
		v17 = F_SearchSysCacheExists(m, int32(57), v13, v14, v14, v14)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_lastval_0), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_lastval_1), int32(911), int32(_a_F_lastval_2))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[1]))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[0]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				if v25 != v28 {
					v30 = int32(_a_F_lastval_3)
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[2]))
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[3]))
					*(*int32)(unsafe.Add(mBase, _c_F_lastval[2])) = v34
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					F_LockRelationOid(m, v36, int32(3))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_lastval[2])) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v25
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v46 = F_sequence_open(m, v44, int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[0]))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[4]))
							v54 = F_pg_class_aclcheck(m, v50, v52, int64(258))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v107 + int32(4)
											F_errmsg(m, int32(_a_F_lastval_4), v8)
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_lastval_1), int32(923), int32(_a_F_lastval_2))
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[0]))
									v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+16))
									F_relation_close(m, v46, int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = F_Int64GetDatum(m, v58)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 + int32(16)
											return v62
										}
									}
								}
							}
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v46 = F_sequence_open(m, v44, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[0]))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[4]))
						v54 = F_pg_class_aclcheck(m, v50, v52, int64(258))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v54 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v107 + int32(4)
										F_errmsg(m, int32(_a_F_lastval_4), v8)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_lastval_1), int32(923), int32(_a_F_lastval_2))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, _c_F_lastval[0]))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+16))
								F_relation_close(m, v46, int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = F_Int64GetDatum(m, v58)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return v62
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_lastval_0), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_lastval_1), int32(905), int32(_a_F_lastval_2))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_latin1_to_mic(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13957(m, l0, int32(8), int32(129))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_latin2mic_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v4 = l3
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v67)
	return v64 - l0
L2:
	;
	v58 = l1
	v64 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v14 = l1
	v15 = l2
	v20 = l0
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		v58 = v14
		v64 = v20
		goto L1
	} else {
		goto L20
	}
L7:
	;
	if l6 != 0 {
		v58 = v14
		v64 = v20
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v30 = base.I32_extend8_s(v23)
	if int32(0) <= v30 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_report_invalid_encoding(m, l4, v20, v15)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	goto L6
L14:
	;
	v48 = int32(1)
	v49 = v20 + v48
	if v48 < v15 {
		v14 = v47
		v15 = v15 - v48
		v20 = v49
		goto L5
	} else {
		goto L19
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v30)
	v47 = v14 + int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v23-int32(128)))))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v39)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v4)
	v47 = v14 + int32(2)
	goto L14
L19:
	;
	v58 = v47
	v64 = v49
	goto L1
L20:
	;
	F_report_untranslatable_char(m, l4, int32(7), v20, v15)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_latin3_to_mic(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13957(m, l0, int32(10), int32(131))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lazy_check_wraparound_failsafe(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int64
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[0])))
	if v17 != 0 {
		v314 = int32(1)
		m.G0 = v13 + int32(48)
		return v314
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[1]))
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[2]))
		v24 = l0 + int32(28)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		v28 = F_ReadNextFullTransactionId(m)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v33 = base.F64_convert_i32_s(v22)
			v36 = base.F64_mul(base.F64_convert_i32_s(v20), float64(1.05))
			if base.F64_gt(v33, v36) != 0 {
				v38 = v33
			} else {
				v38 = v36
			}
			v40 = base.I32_wrap_i64(v28) - base.I32_trunc_sat_f64_s(v38)
			if base.Ui32(v40) <= base.Ui32(int32(3)) {
				v43 = int32(3)
			} else {
				v43 = v40
			}
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v43))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v26)) == int32(0) {
				v55 = base.B2i32(base.Ui32(v26) < base.Ui32(v43))
			} else {
				v55 = int32(base.Ui32(v26-v43) >> (uint(int32(31)) % 32))
			}
			if v55 != 0 {
				v82 = int32(1)
				if v82 == int32(0) {
					v314 = int32(0)
					m.G0 = v13 + int32(48)
					return v314
				} else {
					v85 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[0])) = uint8(v85)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = int64(38654705672)
					v91 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v91
					*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v91
					v95 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v95)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+23)) = uint16(v95)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v95
					v115 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[3]))
					if v115 == int32(0) {
					} else {
						v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[4])))
						if v119&int32(1) == int32(0) {
						} else {
							v124 = int32(_a_F_lazy_check_wraparound_failsafe_0)
							v126 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5]))
							v127 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5])) = v126 + v127
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
							*(*int32)(unsafe.Add(mBase, uint32(v115))) = v130 + v127
							v225 = int32(0)
							v228 = v95
							for {
								v237 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(40)+v228<<(uint(int32(2))%32))))
								v238 = int32(3)
								v244 = *(*int64)(unsafe.Add(mBase, uint32(v13+int32(16)+v228<<(uint(v238)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v115+int32(232)+v237<<(uint(v238)%32)))) = v244
								v246 = int32(1)
								v249 = v225 + v246
								if v249 != int32(2) {
									v225 = v249
									v228 = v228 + v246
									continue
								} else {
									break
								}
								break
							}
							v260 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
							v261 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v115))) = v260 + v261
							v264 = int32(_a_F_lazy_check_wraparound_failsafe_0)
							v266 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5]))
							*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5])) = v266 - v261
						}
					}
					v281 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v282 = m.ExcPending
					if v282 != 0 {
						return int32(0)
					} else {
						if v281 != 0 {
							v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)+68))
							v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v285
							*(*int64)(unsafe.Add(mBase, uint32(v13))) = v284
							*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v283
							F_errmsg(m, int32(_a_F_lazy_check_wraparound_failsafe_1), v13)
							mBase = m.M
							v291 = m.ExcPending
							if v291 != 0 {
								return int32(0)
							} else {
								F_errdetail(m, int32(_a_F_lazy_check_wraparound_failsafe_2), int32(0))
								mBase = m.M
								v295 = m.ExcPending
								if v295 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_lazy_check_wraparound_failsafe_3), int32(0))
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_lazy_check_wraparound_failsafe_4), int32(2987), int32(_a_F_lazy_check_wraparound_failsafe_5))
										mBase = m.M
										v304 = m.ExcPending
										if v304 != 0 {
											return int32(0)
										} else {
											v308 = int32(0)
											*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[6])) = v308
											*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[7])) = uint8(v308)
											v314 = v85
											m.G0 = v13 + int32(48)
											return v314
										}
									}
								}
							}
						} else {
							v308 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[6])) = v308
							*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[7])) = uint8(v308)
							v314 = v85
							m.G0 = v13 + int32(48)
							return v314
						}
					}
				}
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[8]))
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[9]))
				v62 = F_ReadNextMultiXactId(m)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v64 = base.F64_convert_i32_s(v60)
					v67 = base.F64_mul(base.F64_convert_i32_s(v58), float64(1.05))
					if base.F64_gt(v64, v67) != 0 {
						v69 = v64
					} else {
						v69 = v67
					}
					v70 = base.I32_trunc_sat_f64_s(v69)
					if v70 == v62 {
						v73 = int32(1)
					} else {
						v73 = v62 - v70
					}
					v82 = int32(base.Ui32(v25-v73) >> (uint(int32(31)) % 32))
					if v82 == int32(0) {
						v314 = int32(0)
						m.G0 = v13 + int32(48)
						return v314
					} else {
						v85 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[0])) = uint8(v85)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = int64(38654705672)
						v91 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v91
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v91
						v95 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v95)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+23)) = uint16(v95)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v95
						v115 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[3]))
						if v115 == int32(0) {
						} else {
							v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[4])))
							if v119&int32(1) == int32(0) {
							} else {
								v124 = int32(_a_F_lazy_check_wraparound_failsafe_0)
								v126 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5]))
								v127 = int32(1)
								*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5])) = v126 + v127
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
								*(*int32)(unsafe.Add(mBase, uint32(v115))) = v130 + v127
								v225 = int32(0)
								v228 = v95
								for {
									v237 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(40)+v228<<(uint(int32(2))%32))))
									v238 = int32(3)
									v244 = *(*int64)(unsafe.Add(mBase, uint32(v13+int32(16)+v228<<(uint(v238)%32))))
									*(*int64)(unsafe.Add(mBase, uint32(v115+int32(232)+v237<<(uint(v238)%32)))) = v244
									v246 = int32(1)
									v249 = v225 + v246
									if v249 != int32(2) {
										v225 = v249
										v228 = v228 + v246
										continue
									} else {
										break
									}
									break
								}
								v260 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
								v261 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v115))) = v260 + v261
								v264 = int32(_a_F_lazy_check_wraparound_failsafe_0)
								v266 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5]))
								*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[5])) = v266 - v261
							}
						}
						v281 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v282 = m.ExcPending
						if v282 != 0 {
							return int32(0)
						} else {
							if v281 != 0 {
								v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
								v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)+68))
								v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v285
								*(*int64)(unsafe.Add(mBase, uint32(v13))) = v284
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v283
								F_errmsg(m, int32(_a_F_lazy_check_wraparound_failsafe_1), v13)
								mBase = m.M
								v291 = m.ExcPending
								if v291 != 0 {
									return int32(0)
								} else {
									F_errdetail(m, int32(_a_F_lazy_check_wraparound_failsafe_2), int32(0))
									mBase = m.M
									v295 = m.ExcPending
									if v295 != 0 {
										return int32(0)
									} else {
										F_errhint(m, int32(_a_F_lazy_check_wraparound_failsafe_3), int32(0))
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_lazy_check_wraparound_failsafe_4), int32(2987), int32(_a_F_lazy_check_wraparound_failsafe_5))
											mBase = m.M
											v304 = m.ExcPending
											if v304 != 0 {
												return int32(0)
											} else {
												v308 = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[6])) = v308
												*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[7])) = uint8(v308)
												v314 = v85
												m.G0 = v13 + int32(48)
												return v314
											}
										}
									}
								}
							} else {
								v308 = int32(0)
								*(*int32)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[6])) = v308
								*(*uint8)(unsafe.Add(mBase, _c_F_lazy_check_wraparound_failsafe[7])) = uint8(v308)
								v314 = v85
								m.G0 = v13 + int32(48)
								return v314
							}
						}
					}
				}
			}
		}
	}
}
func F_leftmostvalue_text(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_cstring_to_text_with_len(m, int32(_a_F_leftmostvalue_text_0), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_leftmostvalue_timetz(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(-86400)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(0)
		return v3
	}
}
func F_levenshtein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v22 = v20 & int32(1)
			if v22 != 0 {
				v23 = v13
			} else {
				v23 = v8 + int32(4)
			}
			if v20 == int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v29 == int32(18) {
					v32 = int32(16)
				} else {
					v32 = int32(0)
				}
				if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v39 = int32(4)
				} else {
					v39 = v32
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v22 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v52 = v15 + v51
			if v17&v51 != 0 {
				v57 = v52
			} else {
				v57 = v15 + int32(4)
			}
			if v17 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v86 = v73
			} else {
				v74 = int32(1)
				if v17&v74 != 0 {
					v86 = int32(base.Ui32(v17)>>(uint(v74)%32)) - v74
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = int32(1)
			v90 = F_varstr_levenshtein(m, v23, v50, v57, v86, v87, v87, v87)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				return v90
			}
		}
	}
}
func F_lexdigits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = v11
	v19 = v5
	v20 = v5
	goto L3
L1:
	;
	if base.Ui32(v54) < base.Ui32(l2) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
	v54 = v19
	v55 = v20
	goto L1
L3:
	;
	if base.Ui32(v12) <= base.Ui32(v17) {
		v54 = v19
		v55 = v20
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v54 = l3
	v55 = v46
	goto L1
L5:
	;
	v25 = v17 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v29 = v27 - int32(48)
	if base.B2i32(base.Ui32(int32(54)) < base.Ui32(v29))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v29))%64)))&int32(1) == int32(0)) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_c_F_lexdigits[0])))
	if base.Ui32(l1) <= base.Ui32(v43) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v46 = v43 + l1*v20
	v48 = v19 + int32(1)
	if v48 != l3 {
		v17 = v25
		v19 = v48
		v20 = v46
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	return v55
L12:
	;
	v62 = v60
	goto L14
L13:
	;
	v62 = int32(5)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v62
	goto L11
}
func F_lexeme_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
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
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v4 - int32(1636608432)
	if v3&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v4) {
			v119 = v3
			v120 = v4
			v121 = v10
			v122 = v10
			v123 = v10
			for {
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
				v126 = v125 + v122
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
				v130 = v129 + v123
				v132 = int32(4)
				v134 = v127 + v121 - v130 ^ base.I32_rotl(v130, v132)
				v138 = v126 - v134 ^ base.I32_rotl(v134, int32(6))
				v139 = v130 + v126
				v140 = v134 + v139
				v141 = v138 + v140
				v145 = v139 - v138 ^ base.I32_rotl(v138, int32(8))
				v149 = v140 - v145 ^ base.I32_rotl(v145, int32(16))
				v153 = v141 - v149 ^ base.I32_rotl(v149, int32(19))
				v154 = v145 + v141
				v155 = v149 + v154
				v156 = v153 + v155
				v160 = v154 - v153 ^ base.I32_rotl(v153, v132)
				v161 = int32(12)
				v162 = v119 + v161
				v164 = v120 - v161
				if base.Ui32(int32(11)) < base.Ui32(v164) {
					v119 = v162
					v120 = v164
					v121 = v155
					v122 = v156
					v123 = v160
					continue
				} else {
					break
				}
				break
			}
			v167 = v162
			v168 = v164
			v169 = v155
			v170 = v156
			v171 = v160
		} else {
			v167 = v3
			v168 = v4
			v169 = v10
			v170 = v10
			v171 = v10
		}
		switch v168 - int32(1) {
		case 0:
			v230 = v169
			v231 = v170
			v232 = v171
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 1:
			v223 = v169
			v224 = v170
			v225 = v171
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 2:
			v216 = v169
			v217 = v170
			v218 = v171
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 3:
			v210 = v170
			v211 = v171
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 4:
			v206 = v170
			v207 = v171
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 5:
			v200 = v170
			v201 = v171
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 6:
			v194 = v170
			v195 = v171
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 7:
			v189 = v171
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 8:
			v184 = v171
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 9:
			v179 = v171
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+10)))
			v179 = v175<<(uint(int32(24))%32) + v171
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		default:
			v237 = v169
			v238 = v170
			v239 = v171
		}
	} else {
		if base.Ui32(v4) < base.Ui32(int32(12)) {
			v65 = v3
			v66 = v4
			v67 = v10
			v68 = v10
			v69 = v10
		} else {
			v17 = v3
			v18 = v4
			v19 = v10
			v20 = v10
			v21 = v10
			for {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v24 = v23 + v20
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				v28 = v27 + v21
				v30 = int32(4)
				v32 = v25 + v19 - v28 ^ base.I32_rotl(v28, v30)
				v36 = v24 - v32 ^ base.I32_rotl(v32, int32(6))
				v37 = v28 + v24
				v38 = v32 + v37
				v39 = v36 + v38
				v43 = v37 - v36 ^ base.I32_rotl(v36, int32(8))
				v47 = v38 - v43 ^ base.I32_rotl(v43, int32(16))
				v51 = v39 - v47 ^ base.I32_rotl(v47, int32(19))
				v52 = v43 + v39
				v53 = v47 + v52
				v54 = v51 + v53
				v58 = v52 - v51 ^ base.I32_rotl(v51, v30)
				v59 = int32(12)
				v60 = v17 + v59
				v62 = v18 - v59
				if base.Ui32(int32(11)) < base.Ui32(v62) {
					v17 = v60
					v18 = v62
					v19 = v53
					v20 = v54
					v21 = v58
					continue
				} else {
					break
				}
				break
			}
			v65 = v60
			v66 = v62
			v67 = v53
			v68 = v54
			v69 = v58
		}
		switch v66 - int32(1) {
		case 0:
			v116 = v67
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
			v237 = v116 + v117
			v238 = v68
			v239 = v69
		case 1:
			v111 = v67
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
			v237 = v116 + v117
			v238 = v68
			v239 = v69
		case 2:
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
			v111 = v107<<(uint(int32(16))%32) + v67
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
			v237 = v116 + v117
			v238 = v68
			v239 = v69
		case 3:
			v104 = v68
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 4:
			v101 = v68
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 5:
			v96 = v68
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 6:
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+6)))
			v96 = v92<<(uint(int32(16))%32) + v68
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 7:
			v87 = v69
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		case 8:
			v82 = v69
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		case 9:
			v77 = v69
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		case 10:
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+10)))
			v77 = v73<<(uint(int32(24))%32) + v69
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		default:
			v237 = v67
			v238 = v68
			v239 = v69
		}
	}
	v242 = int32(14)
	v244 = v238 ^ v239 - base.I32_rotl(v238, v242)
	v248 = v244 ^ v237 - base.I32_rotl(v244, int32(11))
	v252 = v248 ^ v238 - base.I32_rotl(v248, int32(25))
	v256 = v252 ^ v244 - base.I32_rotl(v252, int32(16))
	v260 = v256 ^ v248 - base.I32_rotl(v256, int32(4))
	v264 = v260 ^ v252 - base.I32_rotl(v260, v242)
	return v264 ^ v256 - base.I32_rotl(v264, int32(24))
}
func F_like_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v276 float64
	_ = v276
	var v279 int32
	_ = v279
	var v293 float64
	_ = v293
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v311 int32
	_ = v311
	var v328 float64
	_ = v328
	var v329 float64
	_ = v329
	var v332 float64
	_ = v332
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	v6 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_like_fixed_prefix[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20*int32(28))+uint32(_c_F_like_fixed_prefix[1])))
	goto L1
L1:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L12
	} else {
		goto L103
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L99
	}
L4:
	;
	v93 = F_palloc(m, v87+int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L35
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = F_pg_detoast_datum_packed(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L17
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = F_text_to_cstring(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L15
	}
L7:
	;
	if v17 == int32(17) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v17 == int32(17) {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v30 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v36 = v30
	goto L6
L14:
	;
	v36 = v6
	goto L6
L15:
	;
	v40 = F_strlen(m, v38)
	mBase = m.M
	v87 = v40
	v88 = v38
	v89 = v36
	v90 = v6
	goto L4
L16:
	;
	v74 = F_palloc(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L28
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v44 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v50 == int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v61 = int32(1)
	if v44&v61 != 0 {
		v73 = int32(base.Ui32(v44)>>(uint(v61)%32)) - v61
		goto L16
	} else {
		goto L27
	}
L21:
	;
	v53 = int32(16)
	goto L23
L22:
	;
	v53 = int32(0)
	goto L23
L23:
	;
	if base.Ui32((v50-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = int32(4)
	goto L26
L25:
	;
	v60 = v53
	goto L26
L26:
	;
	v73 = v60
	goto L16
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v73 = int32(base.Ui32(v67)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L28:
	;
	if v73 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v77 = int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v79&v77 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v87 = v73
	v88 = v74
	v89 = v6
	v90 = int32(1)
	goto L4
L32:
	;
	v82 = v77
	goto L34
L33:
	;
	v82 = int32(4)
	goto L34
L34:
	;
	base.MemoryCopy(m, v74, v42+v82, v73)
	goto L31
L35:
	;
	v95 = int32(0)
	if v87 <= v95 {
		v183 = v95
		v185 = v95
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v93))) = uint8(v200)
	if v90 == v200 {
		goto L61
	} else {
		goto L62
	}
L37:
	;
	v101 = v95
	v103 = v95
	goto L38
L38:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v88))))
	switch v118 - int32(92) {
	case 0:
		goto L41
	case 1, 2:
		v126 = v101
		goto L40
	case 3:
		v183 = v101
		v185 = v103
		goto L36
	default:
		goto L42
	}
L39:
	;
	v183 = v181
	v185 = v179
	goto L36
L40:
	;
	if l1 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v124 = v101 + int32(1)
	if v87 <= v124 {
		v183 = v124
		v185 = v103
		goto L36
	} else {
		goto L44
	}
L42:
	;
	if v118 != int32(37) {
		v126 = v101
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v183 = v101
	v185 = v103
	goto L36
L44:
	;
	v126 = v124
	goto L40
L45:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v88))))
	*(*uint8)(unsafe.Add(mBase, uint32(v103+v93))) = uint8(v176)
	v178 = int32(1)
	v179 = v103 + v178
	v181 = v126 + v178
	if v181 < v87 {
		v101 = v181
		v103 = v179
		goto L38
	} else {
		goto L59
	}
L46:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v88))))
	v131 = base.I32_extend8_s(v130)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+3)))
	if v132 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if base.Ui32((v131&int32(223)-int32(91))&int32(255)) < base.Ui32(int32(230)) {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v143 = int32(0)
	v144 = base.B2i32(v131 < v143)
	if v144&base.B2i32(base.B2i32(v25 < int32(2)) == v143) != 0 {
		v183 = v126
		v185 = v103
		goto L36
	} else {
		goto L51
	}
L50:
	;
	v183 = v126
	v185 = v103
	goto L36
L51:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v148 != int32(99) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if base.B2i32(base.Ui32(int32(229)) < base.Ui32((v131-int32(91))&int32(255)))|v144 != 0 {
		v183 = v126
		v185 = v103
		goto L36
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L57
L55:
	;
	if base.Ui32((v131-int32(123))&int32(255)) < base.Ui32(int32(230)) {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	v183 = v126
	v185 = v103
	goto L36
L57:
	;
	if base.Ui32(v130|int32(32)-int32(97)) < base.Ui32(int32(26)) {
		v183 = v126
		v185 = v103
		goto L36
	} else {
		goto L58
	}
L58:
	;
	goto L45
L59:
	;
	goto L39
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v226
	if l4 != 0 {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	v204 = F_string_to_const(m, v93, v17)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v207 = v185 + int32(4)
	v208 = F_palloc(m, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L65
	}
L64:
	;
	v226 = v204
	goto L60
L65:
	;
	if v185 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	base.MemoryCopy(m, v208+int32(4), v93, v185)
	goto L68
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v207 << (uint(int32(2)) % 32)
	v217 = int32(-1)
	v218 = int32(0)
	v222 = F_makeConst(m, int32(17), v217, v218, v217, v208, v218, v218)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	v226 = v222
	goto L60
L70:
	;
	v228 = v183 + v88
	v229 = int32(0)
	v230 = v87 - v183
	if v230 <= v229 {
		v261 = v229
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	F_pfree(m, v88)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L94
	}
L73:
	;
	v329 = float64(1)
	if base.F64_gt(v328, v329) != 0 {
		goto L91
	} else {
		goto L92
	}
L74:
	;
	v276 = float64(1)
	if v230 <= v261 {
		v328 = v276
		goto L73
	} else {
		goto L80
	}
L75:
	;
	v234 = v229
	goto L76
L76:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v228))))
	if base.B2i32(v250 != int32(95))&base.B2i32(v250 != int32(37)) != 0 {
		v261 = v234
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v328 = float64(1)
	goto L73
L78:
	;
	v257 = v234 + int32(1)
	if v257 != v230 {
		v234 = v257
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v279 = v261
	v293 = v276
	goto L81
L81:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v228))))
	switch v295 - int32(92) {
	case 0:
		goto L85
	case 1, 2:
		v305 = v279
		goto L84
	case 3:
		goto L86
	default:
		goto L87
	}
L82:
	;
	v328 = v309
	goto L73
L83:
	;
	v309 = base.F64_mul(v293, v308)
	v311 = v307 + int32(1)
	if v311 < v230 {
		v279 = v311
		v293 = v309
		goto L81
	} else {
		goto L90
	}
L84:
	;
	v307 = v305
	v308 = float64(0.2)
	goto L83
L85:
	;
	v303 = v279 + int32(1)
	if v230 <= v303 {
		v328 = v293
		goto L73
	} else {
		goto L89
	}
L86:
	;
	v307 = v279
	v308 = float64(0.9)
	goto L83
L87:
	;
	if v295 != int32(37) {
		v305 = v279
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v307 = v279
	v308 = float64(5)
	goto L83
L89:
	;
	v305 = v303
	goto L84
L90:
	;
	goto L82
L91:
	;
	v332 = v329
	goto L93
L92:
	;
	v332 = v328
	goto L93
L93:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v332
	goto L72
L94:
	;
	F_pfree(m, v93)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	if v183 == v87 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v358 = int32(2)
	goto L98
L97:
	;
	v358 = base.B2i32(int32(0) < v185)
	goto L98
L98:
	;
	return v358
L99:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_like_fixed_prefix_0), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_like_fixed_prefix_1), int32(1009), int32(_a_F_like_fixed_prefix_2))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(_a_F_like_fixed_prefix_3), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	F_errhint(m, int32(_a_F_like_fixed_prefix_4), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_like_fixed_prefix_1), int32(1020), int32(_a_F_like_fixed_prefix_2))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lo_get_fragment_internal(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v9 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_lo_get_fragment_internal[0])) = uint8(v9)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_lo_get_fragment_internal[1]))
	v14 = F_inv_open(m, l0, int32(_a_F_lo_get_fragment_internal_0), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_inv_seek(m, v14, int64(0), int32(2))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 <= l1 {
				v52 = int64(0)
				v53 = base.I32_wrap_i64(v52)
				v55 = v53 + int32(4)
				v56 = F_palloc(m, v55)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v59 = F_inv_seek(m, v14, l1, int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v63 = F_inv_read(m, v14, v56+int32(4), v53)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v56))) = v55 << (uint(int32(2)) % 32)
							F_pfree(m, v14)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v56
							}
						}
					}
				}
			} else {
				if l2 < int32(0) {
					v32 = v20 - l1
				} else {
					v26 = v20 - l1
					v27 = base.I64_extend_i32_u(l2)
					if v26 < v27 {
						v29 = v26
					} else {
						v29 = v27
					}
					v32 = v29
				}
				if v32 < int64(1073741820) {
					v52 = v32
					v53 = base.I32_wrap_i64(v52)
					v55 = v53 + int32(4)
					v56 = F_palloc(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v59 = F_inv_seek(m, v14, l1, int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v63 = F_inv_read(m, v14, v56+int32(4), v53)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v56))) = v55 << (uint(int32(2)) % 32)
								F_pfree(m, v14)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									return v56
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_lo_get_fragment_internal_1), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_lo_get_fragment_internal_2), int32(779), int32(_a_F_lo_get_fragment_internal_3))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
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
func F_load_critical_index(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_LockRelationOid(m, l1, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		F_LockRelationOid(m, l0, int32(1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = F_RelationBuildDesc(m, l0, int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 == int32(0) {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errcode(m, int32(16779816))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(_a_F_load_critical_index_0), v7)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_load_critical_index_1), int32(_a_F_load_critical_index_2), int32(_a_F_load_critical_index_3))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
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
					v36 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v36
					*(*uint8)(unsafe.Add(mBase, uint32(v16)+25)) = uint8(v36)
					F_UnlockRelationOid(m, l0, v36)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_UnlockRelationOid(m, l1, int32(1))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v47 = F_RelationGetIndexAttOptions(m, v16, int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_load_typcache_tupdesc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
			F_errmsg_internal(m, int32(_a_F_load_typcache_tupdesc_0), v8)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_load_typcache_tupdesc_1), int32(975), int32(_a_F_load_typcache_tupdesc_2))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = F_relation_open(m, v10, int32(1))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
			v33 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v32 + v33
			v36 = int32(_a_F_load_typcache_tupdesc_3)
			v38 = *(*int64)(unsafe.Add(mBase, _c_F_load_typcache_tupdesc[0]))
			v40 = v38 + int64(1)
			*(*int64)(unsafe.Add(mBase, _c_F_load_typcache_tupdesc[0])) = v40
			*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v40
			F_relation_close(m, v28, v33)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_localsub(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var __phi107 int32
	_ = __phi107
	var v112 int64
	_ = v112
	var __phi112 int64
	_ = __phi112
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v226 int64
	_ = v226
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v254 int64
	_ = v254
	var v257 int64
	_ = v257
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int64
	_ = v384
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v538 int64
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v550 int64
	_ = v550
	var v552 int64
	_ = v552
	var v554 int64
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v571 int64
	_ = v571
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v577 int64
	_ = v577
	var v581 int64
	_ = v581
	var v582 int64
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v592 int64
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int64
	_ = v598
	var v601 int64
	_ = v601
	var v603 int64
	_ = v603
	var v604 int64
	_ = v604
	var v611 int32
	_ = v611
	var v614 int64
	_ = v614
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int64
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v708 int32
	_ = v708
	var v711 int64
	_ = v711
	var v714 int32
	_ = v714
	var v715 int64
	_ = v715
	var v717 int64
	_ = v717
	var v726 int64
	_ = v726
	var v728 int64
	_ = v728
	var v731 int64
	_ = v731
	var v732 int64
	_ = v732
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v736 int64
	_ = v736
	var v739 int64
	_ = v739
	var v744 int32
	_ = v744
	var __phi744 int32
	_ = __phi744
	var v749 int64
	_ = v749
	var __phi749 int64
	_ = __phi749
	var v754 int32
	_ = v754
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v770 int64
	_ = v770
	var v778 int32
	_ = v778
	var v780 int64
	_ = v780
	var v786 int32
	_ = v786
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int64
	_ = v858
	var v860 int64
	_ = v860
	var v863 int64
	_ = v863
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int64
	_ = v870
	var v873 int64
	_ = v873
	var v874 int64
	_ = v874
	var v875 int64
	_ = v875
	var v887 int32
	_ = v887
	var v888 int64
	_ = v888
	var v891 int64
	_ = v891
	var v894 int64
	_ = v894
	var v898 int64
	_ = v898
	var v899 int64
	_ = v899
	var v909 int32
	_ = v909
	var v910 int64
	_ = v910
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int64
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v17
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v1179
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_localsub[0]))
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v535 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L5:
	;
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_localsub[0]))
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v25 = F_emscripten_builtin_malloc(m, int32(_a_F_localsub_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[0])) = v25
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v1179 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v32 = F_tzload(m, int32(_a_F_localsub_1), int32(0), v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v32 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v40 = F_tzparse(m, int32(_a_F_localsub_1), v25, int32(1))
	mBase = m.M
	goto L5
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[1])) = v44 + int32(_a_F_localsub_2)
	v1179 = v530
	goto L1
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v55 = v54
	goto L16
L15:
	;
	v55 = v42
	goto L16
L16:
	;
	v61 = v55
	goto L19
L17:
	;
	v98 = int64(86400)
	v99 = base.I64_div_s(v94, v98)
	v102 = v94 - v99*v98
	__phi107 = int32(1970)
	__phi112 = v99
	v107 = __phi107
	v112 = __phi112
	goto L26
L18:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v94 = v91
	v95 = int64(0)
	v97 = int32(0)
	goto L17
L19:
	;
	v71 = v61 - int32(1)
	if v71 < int32(0) {
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
	if v74 != v78 {
		v94 = v74
		v95 = v80
		v97 = int32(0)
		goto L17
	} else {
		goto L23
	}
L21:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v77 = v44 + int32(_a_F_localsub_3) + v71<<(uint(int32(4))%32)
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	if v74 < v78 {
		v61 = v71
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v71 == int32(0) {
		v94 = v74
		v95 = v80
		v97 = base.B2i32(int64(0) < v80)
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v77-int32(8))))
	v94 = v74
	v95 = v80
	v97 = base.B2i32(v89 < v80)
	goto L17
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[2])) = int32(61)
	v530 = int32(0)
	goto L13
L26:
	;
	v117 = base.B2i32(v112 < int64(0))
	if v117 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v220 = base.I32_wrap_i64(v112)
	v221 = base.I64_extend_i32_s(v42)
	v223 = v221 - v95 + v102
	if v223 < int64(0) {
		goto L57
	} else {
		goto L58
	}
L28:
	;
	goto L27
L29:
	;
	if v107&int32(3) != 0 {
		v130 = int32(0)
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui64(int64(1571958030700)) < base.Ui64(v112+int64(785979015533)) {
		goto L25
	} else {
		goto L36
	}
L32:
	;
	v133 = int64(*(*int32)(unsafe.Add(mBase, uint32(v130<<(uint(int32(2))%32))+uint32(_c_F_localsub[3]))))
	if v112 < v133 {
		goto L28
	} else {
		goto L35
	}
L33:
	;
	v125 = base.I32_rem_s(v107, int32(100))
	if v125 != 0 {
		v130 = int32(1)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v127 = base.I32_rem_s(v107, int32(400))
	v130 = base.B2i32(v127 == int32(0))
	goto L32
L35:
	;
	goto L31
L36:
	;
	if v112 < int64(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = int32(-1)
	goto L39
L38:
	;
	v141 = int32(1)
	goto L39
L39:
	;
	v143 = base.I64_div_s(v112, int64(366))
	if base.Ui64(v112+int64(365)) < base.Ui64(int64(731)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v149 = v141
	goto L42
L41:
	;
	v149 = base.I32_wrap_i64(v143)
	goto L42
L42:
	;
	if int32(0) <= v107 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v158 = v149 + v107
	v160 = v158 - int32(1)
	if v160 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v149 <= v107^int32(2147483647) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v149 < int32(-2147483648)-v107 {
		goto L25
	} else {
		goto L48
	}
L47:
	;
	goto L25
L48:
	;
	goto L43
L49:
	;
	v192 = v107 - int32(1)
	if v192 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v164 = int32(0) - v158
	v168 = base.I32_div_u_s(v164, int32(100))
	v171 = base.I32_div_u_s(v164, int32(400))
	v184 = int32(base.Ui32(v164)>>(uint(int32(2))%32)) - v168 + v171 ^ int32(-1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v178 = base.I32_div_u_s(v160, int32(100))
	v181 = base.I32_div_u_s(v160, int32(400))
	v184 = int32(base.Ui32(v160)>>(uint(int32(2))%32)) - v178 + v181
	goto L49
L53:
	;
	__phi107 = v158
	__phi112 = (base.I64_extend_i32_s(v158)-base.I64_extend_i32_s(v107))*int64(-365) + v112 - base.I64_extend_i32_s(v184-v216)
	v107 = __phi107
	v112 = __phi112
	goto L26
L54:
	;
	v196 = int32(0) - v107
	v200 = base.I32_div_u_s(v196, int32(100))
	v203 = base.I32_div_u_s(v196, int32(400))
	v216 = int32(base.Ui32(v196)>>(uint(int32(2))%32)) - v200 + v203 ^ int32(-1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v210 = base.I32_div_u_s(v192, int32(100))
	v213 = base.I32_div_u_s(v192, int32(400))
	v216 = int32(base.Ui32(v192)>>(uint(int32(2))%32)) - v210 + v213
	goto L53
L57:
	;
	v226 = int64(-86400)
	if base.Ui64(v223) <= base.Ui64(v226) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v250 = v220
	v251 = v223
	goto L59
L59:
	;
	if base.Ui64(int64(86400)) <= base.Ui64(v251) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v229 = v226
	goto L62
L61:
	;
	v229 = v223
	goto L62
L62:
	;
	v231 = v95 + v229 - v102
	v233 = base.I64_extend_i32_u(base.B2i32(v231 != v221))
	v236 = int64(86400)
	v237 = base.I64_div_u_s(v231-(v233+v221), v236)
	v238 = v237 + v233
	v250 = base.I32_wrap_i64(v238) ^ int32(-1) + v220
	v251 = v102 + v238*v236 + v221 - v95 + v236
	goto L59
L63:
	;
	v254 = int64(172799)
	if v254 <= v251 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v272 = v250
	v273 = v251
	goto L65
L65:
	;
	if v272 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v257 = v254
	goto L68
L67:
	;
	v257 = v251
	goto L68
L68:
	;
	v261 = int64(86400)
	v262 = base.I64_div_u_s(v251-v257+int64(86399), v261)
	v272 = v250 + base.I32_wrap_i64(v262) + int32(1)
	v273 = v251 + v262*int64(-86400) - v261
	goto L65
L69:
	;
	v277 = v272
	v280 = v107
	goto L72
L70:
	;
	v310 = v272
	v313 = v107
	goto L71
L71:
	;
	v322 = v310
	v325 = v313
	goto L79
L72:
	;
	if v280 == int32(-2147483648) {
		goto L25
	} else {
		goto L74
	}
L73:
	;
	v310 = v307
	v313 = v293
	goto L71
L74:
	;
	v293 = v280 - int32(1)
	if v293&int32(3) != 0 {
		v303 = int32(0)
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303<<(uint(int32(2))%32))+uint32(_c_F_localsub[3])))
	v307 = v306 + v277
	if v307 < int32(0) {
		v277 = v307
		v280 = v293
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v298 = base.I32_rem_s(v293, int32(100))
	if v298 != 0 {
		v303 = int32(1)
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v300 = base.I32_rem_s(v293, int32(400))
	v303 = base.B2i32(v300 == int32(0))
	goto L75
L78:
	;
	goto L73
L79:
	;
	v335 = v325 & int32(3)
	if v335 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[4])) = v325
	if v325 < int32(-2147481748) {
		goto L25
	} else {
		goto L93
	}
L81:
	;
	goto L80
L82:
	;
	if v325 == int32(2147483647) {
		goto L25
	} else {
		goto L92
	}
L83:
	;
	v339 = base.I32_rem_s(v325, int32(100))
	if v339 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	if v322 < int32(365) {
		goto L81
	} else {
		goto L91
	}
L86:
	;
	v343 = base.I32_rem_s(v325, int32(400))
	v345 = base.B2i32(v343 == int32(0))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v345<<(uint(int32(2))%32))+uint32(_c_F_localsub[3])))
	if v322 < v348 {
		goto L81
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v322 < int32(366) {
		goto L81
	} else {
		goto L90
	}
L89:
	;
	v357 = v345
	goto L82
L90:
	;
	v357 = int32(1)
	goto L82
L91:
	;
	v357 = int32(0)
	goto L82
L92:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v357<<(uint(int32(2))%32))+uint32(_c_F_localsub[3])))
	v322 = v322 - v364
	v325 = v325 + int32(1)
	goto L79
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[5])) = v322
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[4])) = v325 - int32(1900)
	v380 = base.I32_rem_s(v325-int32(1970), int32(7))
	v381 = int32(0)
	v384 = base.I64_div_u_s(v273, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _c_F_localsub[6])) = uint32(v384)
	if v325 <= v381 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v417 = int32(7)
	v418 = base.I32_rem_s(v412+(v322+v380)-int32(473), v417)
	if v418 < int32(0) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v390 = int32(0) - v325
	v394 = base.I32_div_u_s(v390, int32(100))
	v397 = base.I32_div_u_s(v390, int32(400))
	v412 = int32(base.Ui32(v390)>>(uint(int32(2))%32)) - v394 + v397 ^ int32(-1)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v402 = v325 - int32(1)
	v406 = base.I32_div_u_s(v402, int32(100))
	v409 = base.I32_div_u_s(v402, int32(400))
	v412 = int32(base.Ui32(v402)>>(uint(int32(2))%32)) - v406 + v409
	goto L94
L98:
	;
	v423 = v418 + v417
	goto L100
L99:
	;
	v423 = v418
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[7])) = v423
	v429 = base.I32_wrap_i64(v273 - v384*int64(3600))
	v430 = int32(_a_F_localsub_4)
	v432 = int32(60)
	v433 = base.I32_div_u_s(v429&v430, v432)
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[8])) = v433
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[9])) = v97 + (v429-v433*v432)&v430
	if v335 != 0 {
		v451 = int32(0)
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v453 = v451 * int32(48)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+uint32(_c_F_localsub[10])))
	if v454 <= v322 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v446 = base.I32_rem_s(v325, int32(100))
	if v446 != 0 {
		v451 = int32(1)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v448 = base.I32_rem_s(v325, int32(400))
	v451 = base.B2i32(v448 == int32(0))
	goto L101
L104:
	;
	v458 = v322
	v460 = v381
	v461 = v454
	goto L107
L105:
	;
	v478 = v322
	v480 = v381
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[11])) = v42
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[12])) = v480
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[13])) = v478 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[14])) = int32(0)
	v530 = int32(_a_F_localsub_5)
	goto L13
L107:
	;
	v470 = v458 - v461
	v472 = v460 + int32(1)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v453+int32(_a_F_localsub_6)+v472<<(uint(int32(2))%32))))
	if v476 <= v470 {
		v458 = v470
		v460 = v472
		v461 = v476
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v478 = v470
	v480 = v472
	goto L106
L109:
	;
	goto L108
L110:
	;
	v611 = v15 + int32(8)
	if v541 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L111:
	;
	v556 = l0 + int32(24)
	if v17 < v554 {
		goto L118
	} else {
		goto L119
	}
L112:
	;
	v538 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < v538 {
		v554 = v538
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v542 != int32(1) {
		goto L110
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v550 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(16)+v541<<(uint(int32(3))%32))))
	if v17 <= v550 {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v552 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v554 = v552
	goto L111
L118:
	;
	v567 = v554 - v17
	goto L120
L119:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v556+v559<<(uint(int32(3))%32)-int32(8))))
	v567 = v17 - v565
	goto L120
L120:
	;
	v571 = base.I64_div_s(v567-int64(1), int64(12622780800))
	v572 = int64(400)
	v573 = v571 * v572
	v575 = v573 + v572
	v577 = v575 * int64(31556952)
	if v17 < v554 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v581 = v577
	goto L123
L122:
	;
	v581 = int64(0) - v577
	goto L123
L123:
	;
	v582 = v581 + v17
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v582
	v584 = int32(0)
	if v582 < v554 {
		v1179 = v584
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v556+v586<<(uint(int32(3))%32)-int32(8))))
	if v592 < v582 {
		v1179 = v584
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v594 = F_localsub(m, l0, v15)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	if v594 == int32(0) {
		v1179 = v584
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v598 = int64(*(*int32)(unsafe.Add(mBase, uint32(v594)+20)))
	v601 = *(*int64)(unsafe.Add(mBase, uint32(v556)))
	if v17 < v601 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v603 = int64(-400) - v573
	goto L130
L129:
	;
	v603 = v575
	goto L130
L130:
	;
	v604 = v598 + v603
	if base.Ui64(int64(4294967295)) < base.Ui64(v604+int64(2147483648)) {
		v1179 = v584
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v594)+20)) = uint32(v604)
	v1179 = v594
	goto L1
L132:
	;
	v678 = l0 + v675<<(uint(int32(4))%32)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v678)+uint32(_c_F_localsub[15])))
	if l0 != 0 {
		goto L149
	} else {
		goto L150
	}
L133:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_localsub[16])))
	v675 = v662
	goto L132
L134:
	;
	v614 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < v614 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v616 = int32(1)
	if v616 < v541 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v622 = v541
	v628 = v616
	goto L139
L137:
	;
	v653 = v616
	goto L138
L138:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v653)+uint32(_c_F_localsub[17]))))
	v675 = v661
	goto L132
L139:
	;
	v634 = int32(1)
	v635 = (v622 + v628) >> (uint(v634) % 32)
	v641 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)+v635<<(uint(int32(3))%32))))
	v642 = base.B2i32(v17 < v641)
	if v17 < v641 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v653 = v643
	goto L138
L141:
	;
	v643 = v628
	goto L143
L142:
	;
	v643 = v635 + v634
	goto L143
L143:
	;
	if v17 < v641 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v644 = v635
	goto L146
L145:
	;
	v644 = v622
	goto L146
L146:
	;
	if v643 < v644 {
		v622 = v644
		v628 = v643
		goto L139
	} else {
		goto L147
	}
L147:
	;
	goto L140
L148:
	;
	if v1167 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L149:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v692 = v691
	goto L151
L150:
	;
	v692 = int32(0)
	goto L151
L151:
	;
	v698 = v692
	goto L154
L152:
	;
	v735 = int64(86400)
	v736 = base.I64_div_s(v731, v735)
	v739 = v731 - v736*v735
	__phi744 = int32(1970)
	__phi749 = v736
	v744 = __phi744
	v749 = __phi749
	goto L161
L153:
	;
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v611)))
	v731 = v728
	v732 = int64(0)
	v734 = int32(0)
	goto L152
L154:
	;
	v708 = v698 - int32(1)
	if v708 < int32(0) {
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v714)+8))
	if v711 != v715 {
		v731 = v711
		v732 = v717
		v734 = int32(0)
		goto L152
	} else {
		goto L158
	}
L156:
	;
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v611)))
	v714 = l0 + int32(_a_F_localsub_3) + v708<<(uint(int32(4))%32)
	v715 = *(*int64)(unsafe.Add(mBase, uint32(v714)))
	if v711 < v715 {
		v698 = v708
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	if v708 == int32(0) {
		v731 = v711
		v732 = v717
		v734 = base.B2i32(int64(0) < v717)
		goto L152
	} else {
		goto L159
	}
L159:
	;
	v726 = *(*int64)(unsafe.Add(mBase, uint32(v714-int32(8))))
	v731 = v711
	v732 = v717
	v734 = base.B2i32(v726 < v717)
	goto L152
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[2])) = int32(61)
	v1167 = int32(0)
	goto L148
L161:
	;
	v754 = base.B2i32(v749 < int64(0))
	if v754 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v857 = base.I32_wrap_i64(v749)
	v858 = base.I64_extend_i32_s(v681)
	v860 = v858 - v732 + v739
	if v860 < int64(0) {
		goto L192
	} else {
		goto L193
	}
L163:
	;
	goto L162
L164:
	;
	if v744&int32(3) != 0 {
		v767 = int32(0)
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	if base.Ui64(int64(1571958030700)) < base.Ui64(v749+int64(785979015533)) {
		goto L160
	} else {
		goto L171
	}
L167:
	;
	v770 = int64(*(*int32)(unsafe.Add(mBase, uint32(v767<<(uint(int32(2))%32))+uint32(_c_F_localsub[3]))))
	if v749 < v770 {
		goto L163
	} else {
		goto L170
	}
L168:
	;
	v762 = base.I32_rem_s(v744, int32(100))
	if v762 != 0 {
		v767 = int32(1)
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v764 = base.I32_rem_s(v744, int32(400))
	v767 = base.B2i32(v764 == int32(0))
	goto L167
L170:
	;
	goto L166
L171:
	;
	if v749 < int64(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v778 = int32(-1)
	goto L174
L173:
	;
	v778 = int32(1)
	goto L174
L174:
	;
	v780 = base.I64_div_s(v749, int64(366))
	if base.Ui64(v749+int64(365)) < base.Ui64(int64(731)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v786 = v778
	goto L177
L176:
	;
	v786 = base.I32_wrap_i64(v780)
	goto L177
L177:
	;
	if int32(0) <= v744 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v795 = v786 + v744
	v797 = v795 - int32(1)
	if v797 < int32(0) {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	if v786 <= v744^int32(2147483647) {
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v786 < int32(-2147483648)-v744 {
		goto L160
	} else {
		goto L183
	}
L182:
	;
	goto L160
L183:
	;
	goto L178
L184:
	;
	v829 = v744 - int32(1)
	if v829 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v801 = int32(0) - v795
	v805 = base.I32_div_u_s(v801, int32(100))
	v808 = base.I32_div_u_s(v801, int32(400))
	v821 = int32(base.Ui32(v801)>>(uint(int32(2))%32)) - v805 + v808 ^ int32(-1)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v815 = base.I32_div_u_s(v797, int32(100))
	v818 = base.I32_div_u_s(v797, int32(400))
	v821 = int32(base.Ui32(v797)>>(uint(int32(2))%32)) - v815 + v818
	goto L184
L188:
	;
	__phi744 = v795
	__phi749 = (base.I64_extend_i32_s(v795)-base.I64_extend_i32_s(v744))*int64(-365) + v749 - base.I64_extend_i32_s(v821-v853)
	v744 = __phi744
	v749 = __phi749
	goto L161
L189:
	;
	v833 = int32(0) - v744
	v837 = base.I32_div_u_s(v833, int32(100))
	v840 = base.I32_div_u_s(v833, int32(400))
	v853 = int32(base.Ui32(v833)>>(uint(int32(2))%32)) - v837 + v840 ^ int32(-1)
	goto L188
L190:
	;
	goto L191
L191:
	;
	v847 = base.I32_div_u_s(v829, int32(100))
	v850 = base.I32_div_u_s(v829, int32(400))
	v853 = int32(base.Ui32(v829)>>(uint(int32(2))%32)) - v847 + v850
	goto L188
L192:
	;
	v863 = int64(-86400)
	if base.Ui64(v860) <= base.Ui64(v863) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	v887 = v857
	v888 = v860
	goto L194
L194:
	;
	if base.Ui64(int64(86400)) <= base.Ui64(v888) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v866 = v863
	goto L197
L196:
	;
	v866 = v860
	goto L197
L197:
	;
	v868 = v732 + v866 - v739
	v870 = base.I64_extend_i32_u(base.B2i32(v868 != v858))
	v873 = int64(86400)
	v874 = base.I64_div_u_s(v868-(v870+v858), v873)
	v875 = v874 + v870
	v887 = base.I32_wrap_i64(v875) ^ int32(-1) + v857
	v888 = v739 + v875*v873 + v858 - v732 + v873
	goto L194
L198:
	;
	v891 = int64(172799)
	if v891 <= v888 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v909 = v887
	v910 = v888
	goto L200
L200:
	;
	if v909 < int32(0) {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v894 = v891
	goto L203
L202:
	;
	v894 = v888
	goto L203
L203:
	;
	v898 = int64(86400)
	v899 = base.I64_div_u_s(v888-v894+int64(86399), v898)
	v909 = v887 + base.I32_wrap_i64(v899) + int32(1)
	v910 = v888 + v899*int64(-86400) - v898
	goto L200
L204:
	;
	v914 = v909
	v917 = v744
	goto L207
L205:
	;
	v947 = v909
	v950 = v744
	goto L206
L206:
	;
	v959 = v947
	v962 = v950
	goto L214
L207:
	;
	if v917 == int32(-2147483648) {
		goto L160
	} else {
		goto L209
	}
L208:
	;
	v947 = v944
	v950 = v930
	goto L206
L209:
	;
	v930 = v917 - int32(1)
	if v930&int32(3) != 0 {
		v940 = int32(0)
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v940<<(uint(int32(2))%32))+uint32(_c_F_localsub[3])))
	v944 = v943 + v914
	if v944 < int32(0) {
		v914 = v944
		v917 = v930
		goto L207
	} else {
		goto L213
	}
L211:
	;
	v935 = base.I32_rem_s(v930, int32(100))
	if v935 != 0 {
		v940 = int32(1)
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v937 = base.I32_rem_s(v930, int32(400))
	v940 = base.B2i32(v937 == int32(0))
	goto L210
L213:
	;
	goto L208
L214:
	;
	v972 = v962 & int32(3)
	if v972 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[4])) = v962
	if v962 < int32(-2147481748) {
		goto L160
	} else {
		goto L228
	}
L216:
	;
	goto L215
L217:
	;
	if v962 == int32(2147483647) {
		goto L160
	} else {
		goto L227
	}
L218:
	;
	v976 = base.I32_rem_s(v962, int32(100))
	if v976 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	if v959 < int32(365) {
		goto L216
	} else {
		goto L226
	}
L221:
	;
	v980 = base.I32_rem_s(v962, int32(400))
	v982 = base.B2i32(v980 == int32(0))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v982<<(uint(int32(2))%32))+uint32(_c_F_localsub[3])))
	if v959 < v985 {
		goto L216
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v959 < int32(366) {
		goto L216
	} else {
		goto L225
	}
L224:
	;
	v994 = v982
	goto L217
L225:
	;
	v994 = int32(1)
	goto L217
L226:
	;
	v994 = int32(0)
	goto L217
L227:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v994<<(uint(int32(2))%32))+uint32(_c_F_localsub[3])))
	v959 = v959 - v1001
	v962 = v962 + int32(1)
	goto L214
L228:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[5])) = v959
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[4])) = v962 - int32(1900)
	v1017 = base.I32_rem_s(v962-int32(1970), int32(7))
	v1018 = int32(0)
	v1021 = base.I64_div_u_s(v910, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _c_F_localsub[6])) = uint32(v1021)
	if v962 <= v1018 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1054 = int32(7)
	v1055 = base.I32_rem_s(v1049+(v959+v1017)-int32(473), v1054)
	if v1055 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v1027 = int32(0) - v962
	v1031 = base.I32_div_u_s(v1027, int32(100))
	v1034 = base.I32_div_u_s(v1027, int32(400))
	v1049 = int32(base.Ui32(v1027)>>(uint(int32(2))%32)) - v1031 + v1034 ^ int32(-1)
	goto L229
L231:
	;
	goto L232
L232:
	;
	v1039 = v962 - int32(1)
	v1043 = base.I32_div_u_s(v1039, int32(100))
	v1046 = base.I32_div_u_s(v1039, int32(400))
	v1049 = int32(base.Ui32(v1039)>>(uint(int32(2))%32)) - v1043 + v1046
	goto L229
L233:
	;
	v1060 = v1055 + v1054
	goto L235
L234:
	;
	v1060 = v1055
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[7])) = v1060
	v1066 = base.I32_wrap_i64(v910 - v1021*int64(3600))
	v1067 = int32(_a_F_localsub_4)
	v1069 = int32(60)
	v1070 = base.I32_div_u_s(v1066&v1067, v1069)
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[8])) = v1070
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[9])) = v734 + (v1066-v1070*v1069)&v1067
	if v972 != 0 {
		v1088 = int32(0)
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1090 = v1088 * int32(48)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+uint32(_c_F_localsub[10])))
	if v1091 <= v959 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	v1083 = base.I32_rem_s(v962, int32(100))
	if v1083 != 0 {
		v1088 = int32(1)
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1085 = base.I32_rem_s(v962, int32(400))
	v1088 = base.B2i32(v1085 == int32(0))
	goto L236
L239:
	;
	v1095 = v959
	v1097 = v1018
	v1098 = v1091
	goto L242
L240:
	;
	v1115 = v959
	v1117 = v1018
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[11])) = v681
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[12])) = v1117
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[13])) = v1115 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_localsub[14])) = int32(0)
	v1167 = int32(_a_F_localsub_5)
	goto L148
L242:
	;
	v1107 = v1095 - v1098
	v1109 = v1097 + int32(1)
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1090+int32(_a_F_localsub_6)+v1109<<(uint(int32(2))%32))))
	if v1113 <= v1107 {
		v1095 = v1107
		v1097 = v1109
		v1098 = v1113
		goto L242
	} else {
		goto L244
	}
L243:
	;
	v1115 = v1107
	v1117 = v1109
	goto L241
L244:
	;
	goto L243
L245:
	;
	v1179 = int32(0)
	goto L1
L246:
	;
	goto L247
L247:
	;
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+uint32(_c_F_localsub[18]))))
	*(*int32)(unsafe.Add(mBase, uint32(v1167)+32)) = v1171
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v678)+uint32(_c_F_localsub[19])))
	*(*int32)(unsafe.Add(mBase, uint32(v1167)+40)) = l0 + v1173 + int32(_a_F_localsub_2)
	v1179 = v1167
	goto L1
}
func F_locate_windowfunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(-1)
	v13 = F_query_or_expression_tree_walker_impl(m, l0, int32(1046), v5+int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		m.G0 = v5 + int32(16)
		return v17
	}
}
func F_logfile_rotate_dest(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v8 = int32(1)
	if l3 == v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0|l1&l3 == int32(0) {
		v129 = v8
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_rotate_dest[0]))
	if v12&l3 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = F_fclose(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	F_pfree(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	return int32(1)
L12:
	;
	goto L11
L13:
	;
	return v129
L14:
	;
	if l3 == int32(8) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v37 = int32(_a_F_logfile_rotate_dest_0)
	goto L17
L16:
	;
	v37 = int32(0)
	goto L17
L17:
	;
	if l3 == int32(16) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v40 = int32(_a_F_logfile_rotate_dest_1)
	goto L20
L19:
	;
	v40 = v37
	goto L20
L20:
	;
	v41 = F_logfile_getname(m, l2, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v87 = F_logfile_open(m, v41, v84, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L35
	}
L23:
	;
	v84 = int32(_a_F_logfile_rotate_dest_2)
	goto L22
L24:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logfile_rotate_dest[1])))
	if v46&int32(1) == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v51 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if base.B2i32(v57 == int32(0))|base.B2i32(v57 != v60) != 0 {
		v78 = v57
		v79 = v60
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v78-v79 != 0 {
		v84 = int32(_a_F_logfile_rotate_dest_3)
		goto L22
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v63 = v41
	v64 = v51
	goto L30
L30:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v68 == int32(0) {
		v78 = v68
		v79 = v67
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v78 = v68
	v79 = v67
	goto L28
L32:
	;
	v71 = int32(1)
	if v68 == v67 {
		v63 = v63 + v71
		v64 = v64 + v71
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L23
L35:
	;
	if v87 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_rotate_dest[2]))
	switch v92 - int32(33) {
	case 0, 8:
		goto L39
	default:
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v118 != 0 {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v111 = int32(0)
	if v41 == v111 {
		v129 = v111
		goto L13
	} else {
		goto L47
	}
L40:
	;
	v97 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if v97 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_errmsg(m, int32(_a_F_logfile_rotate_dest_4), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_logfile_rotate_dest[3])) = uint8(v109)
	goto L39
L45:
	;
	F_errfinish(m, int32(_a_F_logfile_rotate_dest_5), int32(1337), int32(_a_F_logfile_rotate_dest_6))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_pfree(m, v41)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	return int32(0)
L49:
	;
	v119 = F_fclose(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v87
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v122 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	F_pfree(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v41
	v129 = v8
	goto L13
L56:
	;
	goto L55
}
func F_lookup_rowtype_tupdesc_domain(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != int32(2249) {
		v11 = F_lookup_type_cache(m, l0, int32(_a_F_lookup_rowtype_tupdesc_domain_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
			if v15 == int32(100) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+300))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+304))
				v20 = F_lookup_rowtype_tupdesc_internal(m, v18, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 == int32(0) {
						v60 = int32(0)
						m.G0 = v6 + int32(16)
						return v60
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
						if int32(0) <= v25 {
							v57 = v20
							F_IncrTupleDescRefCount(m, v57)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = v57
								m.G0 = v6 + int32(16)
								return v60
							}
						} else {
							v60 = v20
							m.G0 = v6 + int32(16)
							return v60
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
				if v28 != 0 {
					v50 = v28
					if v50 == int32(0) {
						v60 = int32(0)
						m.G0 = v6 + int32(16)
						return v60
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
						if v54 < int32(0) {
							v60 = v50
							m.G0 = v6 + int32(16)
							return v60
						} else {
							v57 = v50
							F_IncrTupleDescRefCount(m, v57)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = v57
								m.G0 = v6 + int32(16)
								return v60
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = F_format_type_be(m, l0)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v36
								F_errmsg(m, int32(_a_F_lookup_rowtype_tupdesc_domain_1), v6)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_lookup_rowtype_tupdesc_domain_2), int32(2001), int32(_a_F_lookup_rowtype_tupdesc_domain_3))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v48 = F_lookup_rowtype_tupdesc_internal(m, int32(2249), l1)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v50 = v48
			if v50 == int32(0) {
				v60 = int32(0)
				m.G0 = v6 + int32(16)
				return v60
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
				if v54 < int32(0) {
					v60 = v50
					m.G0 = v6 + int32(16)
					return v60
				} else {
					v57 = v50
					F_IncrTupleDescRefCount(m, v57)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = v57
						m.G0 = v6 + int32(16)
						return v60
					}
				}
			}
		}
	}
}
func F_lose_s(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	if l0 != 0 {
		F_pfree(m, l0-int32(8))
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_lower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v17 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = F_str_tolower(m, v18, v45, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = F_cstring_to_text(m, v47)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v47)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v49
				}
			}
		}
	}
}
func F_lquery_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_deparse_lquery(m, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_pq_begintypsend(m, v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_enlargeStringInfo(m, v6, int32(1))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v20+v21))) = uint8(v23)
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v20 + v23
					v28 = F_strlen(m, v13)
					mBase = m.M
					F_pq_sendtext(m, v6, v13, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v13)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = v35 << (uint(int32(2)) % 32)
							m.G0 = v6 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_ltq_rregex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltq_rregex_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}

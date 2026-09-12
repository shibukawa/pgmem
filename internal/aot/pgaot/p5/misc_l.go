package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(4438508)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	*(*int32)(unsafe.Add(mBase, _consts[415])) = v15 + int32(1)
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[719]))
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
	v148 = int32(4438508)
	v150 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	*(*int32)(unsafe.Add(mBase, _consts[415])) = v150 - int32(1)
	m.G0 = v11 + int32(16)
	return
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v42 = v34 + v35&int32(15)<<(uint(int32(7))%32) + int32(23296)
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
	v47 = *(*int32)(unsafe.Add(mBase, _consts[296]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[719])) = int32(0)
	F_LWLockRelease(m, v42)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[718]))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[719]))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+32)) = v58 + int64(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v53 < v63 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L8
L15:
	;
	v67 = v53
	goto L18
L16:
	;
	v95 = int32(0)
	goto L17
L17:
	;
	v98 = v62 + v95<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v55
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+40)) = v102 + int32(1)
	if v55 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v76 = v62 + v67<<(uint(int32(4))%32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v55 == v77 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v95 = v63
	goto L17
L20:
	;
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v79 + int64(1)
	goto L14
L21:
	;
	goto L22
L22:
	;
	v84 = v67 + int32(1)
	if v84 != v63 {
		v67 = v84
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+18)))
	if base.Ui32(v107) <= base.Ui32(int32(15)) {
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
	if v107 != int32(15) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v55+v107<<(uint(int32(2))%32))+292)) = v57
	goto L33
L32:
	;
	goto L33
L33:
	;
	v117 = v107 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+18)) = uint8(v117)
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
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[779]))
	v20 = F_get_hash_value(m, v19, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		v30 = v23 + v20&int32(15)<<(uint(int32(7))%32) + int32(23296)
		v32 = F_LWLockAcquire(m, v30, int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _consts[779]))
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
					v43 = *(*int32)(unsafe.Add(mBase, _consts[699]))
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
							v191 = m.ExcPending
							if v191 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(103992), int32(0))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return
								} else {
									F_errfinish(m, int32(476003), int32(3296), int32(345473))
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
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
											F_errmsg_internal(m, int32(182795), v16)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errfinish(m, int32(476003), int32(3307), int32(345473))
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
												if v132 != 0 {
													m.G0 = v16 + int32(16)
													return
												} else {
													if l3 < int32(5) {
														m.G0 = v16 + int32(16)
														return
													} else {
														v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
														if v135 == int32(0) {
															m.G0 = v16 + int32(16)
															return
														} else {
															v139 = *(*int32)(unsafe.Add(mBase, _consts[784]))
															v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
															*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(1)
															if v140 != 0 {
																v146 = *(*int32)(unsafe.Add(mBase, _consts[784]))
																F_s_lock(m, v146, int32(476003), int32(3330), int32(345473))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return
																} else {
																	v153 = *(*int32)(unsafe.Add(mBase, _consts[784]))
																	v158 = v153 + v20&int32(1023)<<(uint(int32(2))%32) + int32(4)
																	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
																	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v159 - int32(1)
																	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(0)
																	m.G0 = v16 + int32(16)
																	return
																}
															} else {
																v153 = *(*int32)(unsafe.Add(mBase, _consts[784]))
																v158 = v153 + v20&int32(1023)<<(uint(int32(2))%32) + int32(4)
																v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
																*(*int32)(unsafe.Add(mBase, uint32(v158))) = v159 - int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(0)
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
					}
				} else {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(104033), int32(0))
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return
						} else {
							F_errfinish(m, int32(476003), int32(3280), int32(345473))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(224539)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[248])))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v14 == int32(0) {
		v33 = v13
		v34 = v14
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L31
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v70
L3:
	;
	if v34-v33 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	goto L3
L5:
	;
	if v13 != v14 {
		v33 = v13
		v34 = v14
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v18 = l0
	v19 = v10
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v22
		v34 = v23
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v33 = v22
	v34 = v23
	goto L4
L9:
	;
	v26 = int32(1)
	if v22 == v23 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	if v39 != 0 {
		v70 = v39
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v42 = int32(0)
	v45 = F_GetSysCacheOid(m, int32(37), l0, v42, v42, v42)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	return int32(0)
L16:
	;
	if l1|v45 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if l1 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = int32(0)
	if v45 == v52 {
		v70 = v52
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v60 = F_object_aclcheck(m, int32(2615), v45, v58, int64(256))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_aclcheck_error(m, v60, int32(36), l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v68 = F_RunNamespaceSearchHook(m, v45, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v70 = v45
	goto L2
L30:
	;
	goto L29
L31:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(69579), v8)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(477805), int32(3547), int32(415679))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[295]))
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
						F_errmsg(m, int32(256790), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(477791), int32(911), int32(294782))
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
				v24 = *(*int32)(unsafe.Add(mBase, _consts[296]))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
				v27 = *(*int32)(unsafe.Add(mBase, _consts[295]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				if v25 != v28 {
					v30 = int32(4443908)
					v31 = *(*int32)(unsafe.Add(mBase, _consts[10]))
					v34 = *(*int32)(unsafe.Add(mBase, _consts[297]))
					*(*int32)(unsafe.Add(mBase, _consts[10])) = v34
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					F_LockRelationOid(m, v36, int32(3))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v25
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v46 = F_sequence_open(m, v44, int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, _consts[295]))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							v52 = *(*int32)(unsafe.Add(mBase, _consts[239]))
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
											F_errmsg(m, int32(186143), v8)
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(477791), int32(923), int32(294782))
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
									v57 = *(*int32)(unsafe.Add(mBase, _consts[295]))
									v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+16))
									F_sequence_close(m, v46, int32(0))
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
						v49 = *(*int32)(unsafe.Add(mBase, _consts[295]))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						v52 = *(*int32)(unsafe.Add(mBase, _consts[239]))
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
										F_errmsg(m, int32(186143), v8)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(477791), int32(923), int32(294782))
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
								v57 = *(*int32)(unsafe.Add(mBase, _consts[295]))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+16))
								F_sequence_close(m, v46, int32(0))
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
				F_errmsg(m, int32(256790), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(477791), int32(905), int32(294782))
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
	mBase := m.M
	_ = mBase
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
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(8), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic(m, v6, v5, v10, int32(129), int32(8), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func F_latin2mic_with_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v62 int32
	_ = v62
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
	return v62 - l0
L2:
	;
	v58 = l1
	v62 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = l1
	v17 = l2
	v20 = l0
	goto L5
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l6 != 0 {
		v58 = v16
		v62 = v20
		goto L1
	} else {
		goto L20
	}
L7:
	;
	if l6 != 0 {
		v58 = v16
		v62 = v20
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v32 = base.I32_extend8_s(v25)
	if int32(0) <= v32 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_report_invalid_encoding(m, l4, v20, v17)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
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
	if v48 < v17 {
		v16 = v47
		v17 = v17 - v48
		v20 = v49
		goto L5
	} else {
		goto L19
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v32)
	v47 = v16 + int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+(l5-int32(128))))))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)) = uint8(v39)
	v47 = v16 + int32(2)
	goto L14
L19:
	;
	v58 = v47
	v62 = v49
	goto L1
L20:
	;
	F_report_untranslatable_char(m, l4, int32(7), v20, v17)
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
	mBase := m.M
	_ = mBase
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
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(10), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic(m, v6, v5, v10, int32(131), int32(10), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func F_lazy_check_wraparound_failsafe(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v65 int32
	_ = v65
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int64
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[45])))
	if v16 != 0 {
		v322 = int32(1)
		m.G0 = v12 + int32(48)
		return v322
	} else {
		v19 = l0 + int32(28)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[46]))
		v22 = base.F64_convert_i32_s(v21)
		v24 = *(*int32)(unsafe.Add(mBase, _consts[47]))
		v27 = base.F64_mul(base.F64_convert_i32_s(v24), float64(1.05))
		if base.F64_gt(v22, v27) != 0 {
			v29 = v22
		} else {
			v29 = v27
		}
		if base.F64_lt(base.F64_abs(v29), float64(2.147483648e+09)) != 0 {
			v33 = base.I32_trunc_f64_s(v29)
			v35 = v33
		} else {
			v35 = int32(-2147483648)
		}
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v39 = F_ReadNextFullTransactionId(m)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v44 = base.I32_wrap_i64(v39) - v35
			if base.Ui32(v44) <= base.Ui32(int32(3)) {
				v47 = int32(3)
			} else {
				v47 = v44
			}
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v47))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v37)) == int32(0) {
				v59 = base.B2i32(base.Ui32(v37) < base.Ui32(v47))
			} else {
				v59 = int32(base.Ui32(v37-v47) >> (uint(int32(31)) % 32))
			}
			if v59 != 0 {
				v90 = int32(1)
				if v90 == int32(0) {
					v322 = int32(0)
					m.G0 = v12 + int32(48)
					return v322
				} else {
					v93 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[45])) = uint8(v93)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(38654705672)
					v99 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v99
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v99
					v103 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v103)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+23)) = uint16(v103)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v103
					v121 = *(*int32)(unsafe.Add(mBase, _consts[48]))
					if v121 == v103 {
					} else {
						v127 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
						if v127&int32(1) == int32(0) {
						} else {
							v132 = int32(4438516)
							v134 = *(*int32)(unsafe.Add(mBase, _consts[13]))
							v135 = int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[13])) = v134 + v135
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
							*(*int32)(unsafe.Add(mBase, uint32(v121))) = v138 + v135
							v233 = int32(0)
							v236 = v103
							for {
								v245 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(40)+v236<<(uint(int32(2))%32))))
								v246 = int32(3)
								v252 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(16)+v236<<(uint(v246)%32))))
								*(*int64)(unsafe.Add(mBase, uint32(v121+int32(232)+v245<<(uint(v246)%32)))) = v252
								v254 = int32(1)
								v257 = v233 + v254
								if v257 != int32(2) {
									v233 = v257
									v236 = v236 + v254
									continue
								} else {
									break
								}
								break
							}
							v268 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
							v269 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v121))) = v268 + v269
							v272 = int32(4438516)
							v274 = *(*int32)(unsafe.Add(mBase, _consts[13]))
							*(*int32)(unsafe.Add(mBase, _consts[13])) = v274 - v269
						}
					}
					v289 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v290 = m.ExcPending
					if v290 != 0 {
						return int32(0)
					} else {
						if v289 != 0 {
							v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							v292 = *(*int64)(unsafe.Add(mBase, uint32(l0)+68))
							v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v293
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v292
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v291
							F_errmsg(m, int32(141160), v12)
							mBase = m.M
							v299 = m.ExcPending
							if v299 != 0 {
								return int32(0)
							} else {
								F_errdetail(m, int32(540147), int32(0))
								mBase = m.M
								v303 = m.ExcPending
								if v303 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(564130), int32(0))
									mBase = m.M
									v307 = m.ExcPending
									if v307 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(470959), int32(2987), int32(391612))
										mBase = m.M
										v312 = m.ExcPending
										if v312 != 0 {
											return int32(0)
										} else {
											v316 = int32(0)
											*(*int32)(unsafe.Add(mBase, _consts[50])) = v316
											*(*uint8)(unsafe.Add(mBase, _consts[51])) = uint8(v316)
											v322 = v93
											m.G0 = v12 + int32(48)
											return v322
										}
									}
								}
							}
						} else {
							v316 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[50])) = v316
							*(*uint8)(unsafe.Add(mBase, _consts[51])) = uint8(v316)
							v322 = v93
							m.G0 = v12 + int32(48)
							return v322
						}
					}
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, _consts[52]))
				v63 = base.F64_convert_i32_s(v62)
				v65 = *(*int32)(unsafe.Add(mBase, _consts[53]))
				v68 = base.F64_mul(base.F64_convert_i32_s(v65), float64(1.05))
				if base.F64_gt(v63, v68) != 0 {
					v70 = v63
				} else {
					v70 = v68
				}
				if base.F64_lt(base.F64_abs(v70), float64(2.147483648e+09)) != 0 {
					v74 = base.I32_trunc_f64_s(v70)
					v76 = v74
				} else {
					v76 = int32(-2147483648)
				}
				v78 = F_ReadNextMultiXactId(m)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					if v76 == v78 {
						v82 = int32(1)
					} else {
						v82 = v78 - v76
					}
					v90 = int32(base.Ui32(v36-v82) >> (uint(int32(31)) % 32))
					if v90 == int32(0) {
						v322 = int32(0)
						m.G0 = v12 + int32(48)
						return v322
					} else {
						v93 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[45])) = uint8(v93)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(38654705672)
						v99 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v99
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v99
						v103 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v103)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+23)) = uint16(v103)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v103
						v121 = *(*int32)(unsafe.Add(mBase, _consts[48]))
						if v121 == v103 {
						} else {
							v127 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
							if v127&int32(1) == int32(0) {
							} else {
								v132 = int32(4438516)
								v134 = *(*int32)(unsafe.Add(mBase, _consts[13]))
								v135 = int32(1)
								*(*int32)(unsafe.Add(mBase, _consts[13])) = v134 + v135
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
								*(*int32)(unsafe.Add(mBase, uint32(v121))) = v138 + v135
								v233 = int32(0)
								v236 = v103
								for {
									v245 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(40)+v236<<(uint(int32(2))%32))))
									v246 = int32(3)
									v252 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(16)+v236<<(uint(v246)%32))))
									*(*int64)(unsafe.Add(mBase, uint32(v121+int32(232)+v245<<(uint(v246)%32)))) = v252
									v254 = int32(1)
									v257 = v233 + v254
									if v257 != int32(2) {
										v233 = v257
										v236 = v236 + v254
										continue
									} else {
										break
									}
									break
								}
								v268 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
								v269 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v121))) = v268 + v269
								v272 = int32(4438516)
								v274 = *(*int32)(unsafe.Add(mBase, _consts[13]))
								*(*int32)(unsafe.Add(mBase, _consts[13])) = v274 - v269
							}
						}
						v289 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v290 = m.ExcPending
						if v290 != 0 {
							return int32(0)
						} else {
							if v289 != 0 {
								v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
								v292 = *(*int64)(unsafe.Add(mBase, uint32(l0)+68))
								v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v293
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = v292
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v291
								F_errmsg(m, int32(141160), v12)
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return int32(0)
								} else {
									F_errdetail(m, int32(540147), int32(0))
									mBase = m.M
									v303 = m.ExcPending
									if v303 != 0 {
										return int32(0)
									} else {
										F_errhint(m, int32(564130), int32(0))
										mBase = m.M
										v307 = m.ExcPending
										if v307 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(470959), int32(2987), int32(391612))
											mBase = m.M
											v312 = m.ExcPending
											if v312 != 0 {
												return int32(0)
											} else {
												v316 = int32(0)
												*(*int32)(unsafe.Add(mBase, _consts[50])) = v316
												*(*uint8)(unsafe.Add(mBase, _consts[51])) = uint8(v316)
												v322 = v93
												m.G0 = v12 + int32(48)
												return v322
											}
										}
									}
								}
							} else {
								v316 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[50])) = v316
								*(*uint8)(unsafe.Add(mBase, _consts[51])) = uint8(v316)
								v322 = v93
								m.G0 = v12 + int32(48)
								return v322
							}
						}
					}
				}
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = v11
	v18 = v5
	v20 = v5
	goto L3
L1:
	;
	if base.Ui32(v54) < base.Ui32(l2) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
	v54 = v18
	v56 = v20
	goto L1
L3:
	;
	if base.Ui32(v12) <= base.Ui32(v17) {
		v54 = v18
		v56 = v20
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v54 = l3
	v56 = v47
	goto L1
L5:
	;
	v25 = v17 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v29 = v27 - int32(48)
	if base.Ui32(int32(54)) < base.Ui32(v29) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(35465847073801215))>>(uint(base.I64_extend_i32_u(v29))%64)))&int32(1) == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_consts[509])))
	if base.Ui32(l1) <= base.Ui32(v44) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v47 = v44 + l1*v20
	v49 = v18 + int32(1)
	if v49 != l3 {
		v17 = v25
		v18 = v49
		v20 = v47
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v61 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	return v56
L13:
	;
	v63 = v61
	goto L15
L14:
	;
	v63 = int32(5)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v63
	goto L12
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v332 float64
	_ = v332
	var v335 int32
	_ = v335
	var v348 float64
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v366 int32
	_ = v366
	var v382 float64
	_ = v382
	var v383 float64
	_ = v383
	var v386 float64
	_ = v386
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	v6 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_consts[1059])))
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L12
	} else {
		goto L117
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L12
	} else {
		goto L113
	}
L4:
	;
	v150 = F_palloc(m, v144+int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L53
	}
L5:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v97 = F_pg_detoast_datum_packed(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L34
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = F_text_to_cstring(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L15
	}
L7:
	;
	if v16 == int32(17) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v16 == int32(17) {
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
	v29 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v35 = v29
	goto L6
L14:
	;
	v35 = v6
	goto L6
L15:
	;
	if v37&int32(3) == int32(0) {
		v62 = v37
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v144 = v95
	v145 = v35
	v146 = v37
	v147 = v6
	goto L4
L17:
	;
	v95 = v87 - v37
	goto L16
L18:
	;
	v66 = v62
	goto L27
L19:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v46 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v95 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v51 = v37
	goto L23
L23:
	;
	v55 = v51 + int32(1)
	if v55&int32(3) == int32(0) {
		v62 = v55
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v87 = v55
	goto L17
L25:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v60 != 0 {
		v51 = v55
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v75 = int32(-2139062144)
	if (int32(16843008)-v72|v72)&v75 == v75 {
		v66 = v66 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v81 = v66
	goto L30
L29:
	;
	goto L28
L30:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 != 0 {
		v81 = v81 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v87 = v81
	goto L17
L32:
	;
	goto L31
L33:
	;
	v131 = F_palloc(m, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L45
	}
L34:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v99 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v102 = int32(4)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v104&int32(254) == int32(2) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v117 = int32(1)
	if v99&v117 != 0 {
		v129 = int32(base.Ui32(v99)>>(uint(v117)%32)) - v117
		goto L33
	} else {
		goto L44
	}
L38:
	;
	v113 = v102
	goto L40
L39:
	;
	v113 = base.B2i32(v104 == int32(18)) << (uint(v102) % 32)
	goto L40
L40:
	;
	if v104 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v116 = v102
	goto L43
L42:
	;
	v116 = v113
	goto L43
L43:
	;
	v129 = v116
	goto L33
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v129 = int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(4)
	goto L33
L45:
	;
	v133 = int32(1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v135&v133 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v138 = v133
	goto L48
L47:
	;
	v138 = int32(4)
	goto L48
L48:
	;
	if v129 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v144 = v129
	v145 = v6
	v146 = v131
	v147 = int32(1)
	goto L4
L50:
	;
	v140 = F__emscripten_memcpy_bulkmem(m, v131, v97+v138, v129)
	mBase = m.M
	goto L52
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v152 = int32(0)
	if v144 <= v152 {
		v241 = v152
		v243 = v152
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v150))) = uint8(v257)
	if v147 == v257 {
		goto L74
	} else {
		goto L75
	}
L55:
	;
	v158 = v152
	v160 = v152
	goto L56
L56:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158+v146))))
	switch v174 - int32(92) {
	case 0:
		goto L59
	case 1, 2:
		v182 = v158
		goto L58
	case 3:
		v241 = v158
		v243 = v160
		goto L54
	default:
		goto L60
	}
L57:
	;
	v241 = v239
	v243 = v237
	goto L54
L58:
	;
	if l1 != 0 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v180 = v158 + int32(1)
	if v144 <= v180 {
		v241 = v180
		v243 = v160
		goto L54
	} else {
		goto L62
	}
L60:
	;
	if v174 != int32(37) {
		v182 = v158
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v241 = v158
	v243 = v160
	goto L54
L62:
	;
	v182 = v180
	goto L58
L63:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v146))))
	v185 = base.I32_extend8_s(v184)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+3)))
	if v194 == int32(1) {
		v229 = base.B2i32(base.Ui32((v185&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)))
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v150))) = uint8(v234)
	v236 = int32(1)
	v237 = v160 + v236
	v239 = v182 + v236
	if v239 < v144 {
		v158 = v239
		v160 = v237
		goto L56
	} else {
		goto L72
	}
L66:
	;
	if v229 != 0 {
		v241 = v182
		v243 = v160
		goto L54
	} else {
		goto L71
	}
L67:
	;
	v197 = int32(0)
	if base.B2i32(base.B2i32(v24 < int32(2)) == v197)&base.B2i32(v185 < v197) != 0 {
		v241 = v182
		v243 = v160
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v204 = int32(255)
	v206 = int32(26)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v219 != int32(99) {
		v229 = base.B2i32(base.Ui32((v185-int32(65))&v204) < base.Ui32(v206)) | base.B2i32(base.Ui32((v185-int32(97))&v204) < base.Ui32(v206)) | base.B2i32(base.I32_extend8_s(v185) < int32(0))
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L70
L70:
	;
	v229 = base.B2i32(base.Ui32(v184|int32(32)-int32(97)) < base.Ui32(int32(26)))
	goto L66
L71:
	;
	goto L65
L72:
	;
	goto L57
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v284
	if l4 != 0 {
		goto L84
	} else {
		goto L85
	}
L74:
	;
	v261 = F_string_to_const(m, v150, v16)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L12
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v264 = v243 + int32(4)
	v265 = F_palloc(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L78
	}
L77:
	;
	v284 = v261
	goto L73
L78:
	;
	if v243 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v264 << (uint(int32(2)) % 32)
	v275 = int32(-1)
	v276 = int32(0)
	v280 = F_makeConst(m, int32(17), v275, v276, v275, v265, v276, v276)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L12
	} else {
		goto L83
	}
L80:
	;
	v269 = F__emscripten_memcpy_bulkmem(m, v265+int32(4), v150, v243)
	mBase = m.M
	goto L82
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v284 = v280
	goto L73
L84:
	;
	v286 = v241 + v146
	v287 = int32(0)
	v288 = v144 - v241
	if v288 <= v287 {
		v318 = v287
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L86
L86:
	;
	F_pfree(m, v146)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L12
	} else {
		goto L108
	}
L87:
	;
	v383 = float64(1)
	if base.F64_gt(v382, v383) != 0 {
		goto L105
	} else {
		goto L106
	}
L88:
	;
	v332 = float64(1)
	if v288 <= v318 {
		v382 = v332
		goto L87
	} else {
		goto L94
	}
L89:
	;
	v292 = v287
	goto L90
L90:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v286))))
	if base.B2i32(v307 != int32(95))&base.B2i32(v307 != int32(37)) != 0 {
		v318 = v292
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v382 = float64(1)
	goto L87
L92:
	;
	v314 = v292 + int32(1)
	if v314 != v288 {
		v292 = v314
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v335 = v318
	v348 = v332
	goto L95
L95:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335+v286))))
	switch v350 - int32(92) {
	case 0:
		goto L99
	case 1, 2:
		v360 = v335
		goto L98
	case 3:
		goto L100
	default:
		goto L101
	}
L96:
	;
	v382 = v364
	goto L87
L97:
	;
	v364 = base.F64_mul(v348, v363)
	v366 = v362 + int32(1)
	if v366 < v288 {
		v335 = v366
		v348 = v364
		goto L95
	} else {
		goto L104
	}
L98:
	;
	v362 = v360
	v363 = float64(0.2)
	goto L97
L99:
	;
	v358 = v335 + int32(1)
	if v288 <= v358 {
		v382 = v348
		goto L87
	} else {
		goto L103
	}
L100:
	;
	v362 = v335
	v363 = float64(0.9)
	goto L97
L101:
	;
	if v350 != int32(37) {
		v360 = v335
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v362 = v335
	v363 = float64(5)
	goto L97
L103:
	;
	v360 = v358
	goto L98
L104:
	;
	goto L96
L105:
	;
	v386 = v383
	goto L107
L106:
	;
	v386 = v382
	goto L107
L107:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v386
	goto L86
L108:
	;
	F_pfree(m, v150)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	if v241 == v144 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v411 = int32(2)
	goto L112
L111:
	;
	v411 = base.B2i32(int32(0) < v243)
	goto L112
L112:
	;
	return v411
L113:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(485038), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(471451), int32(1009), int32(25619))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(518689), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	F_errhint(m, int32(536375), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(471451), int32(1020), int32(25619))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
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
	*(*uint8)(unsafe.Add(mBase, _consts[371])) = uint8(v9)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v14 = F_inv_open(m, l0, int32(262144), v13)
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
							F_errmsg(m, int32(382505), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(473498), int32(779), int32(296713))
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
							F_errmsg_internal(m, int32(37899), v7)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(477505), int32(4408), int32(26137))
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
			F_errmsg_internal(m, int32(48078), v8)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errfinish(m, int32(477482), int32(975), int32(467405))
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
			v36 = int32(4063840)
			v38 = *(*int64)(unsafe.Add(mBase, _consts[1147]))
			v40 = v38 + int64(1)
			*(*int64)(unsafe.Add(mBase, _consts[1147])) = v40
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
	var v3 int32
	_ = v3
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
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
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
	var v135 int64
	_ = v135
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int64
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int64
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v570 int64
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v581 int64
	_ = v581
	var v583 int64
	_ = v583
	var v587 int64
	_ = v587
	var v588 int64
	_ = v588
	var v589 int64
	_ = v589
	var v591 int64
	_ = v591
	var v593 int64
	_ = v593
	var v597 int64
	_ = v597
	var v598 int64
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v608 int64
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int64
	_ = v614
	var v617 int64
	_ = v617
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int64
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int64
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v729 int32
	_ = v729
	var v732 int64
	_ = v732
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	var v738 int64
	_ = v738
	var v747 int64
	_ = v747
	var v749 int64
	_ = v749
	var v753 int64
	_ = v753
	var v754 int64
	_ = v754
	var v755 int32
	_ = v755
	var v756 int64
	_ = v756
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v765 int32
	_ = v765
	var __phi765 int32
	_ = __phi765
	var v770 int64
	_ = v770
	var __phi770 int64
	_ = __phi770
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v793 int64
	_ = v793
	var v801 int32
	_ = v801
	var v803 int64
	_ = v803
	var v809 int32
	_ = v809
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int64
	_ = v881
	var v883 int64
	_ = v883
	var v886 int64
	_ = v886
	var v889 int64
	_ = v889
	var v891 int64
	_ = v891
	var v893 int64
	_ = v893
	var v896 int64
	_ = v896
	var v897 int64
	_ = v897
	var v898 int64
	_ = v898
	var v910 int32
	_ = v910
	var v911 int64
	_ = v911
	var v915 int64
	_ = v915
	var v918 int64
	_ = v918
	var v921 int64
	_ = v921
	var v922 int64
	_ = v922
	var v932 int32
	_ = v932
	var v933 int64
	_ = v933
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1024 int32
	_ = v1024
	var v1033 int32
	_ = v1033
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int64
	_ = v1078
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v17
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v1210
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1308]))
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
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v547 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L5:
	;
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1308]))
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v25 = F_emscripten_builtin_malloc(m, int32(23440))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1308])) = v25
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v1210 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v32 = F_tzload(m, int32(1790976), int32(0), v25)
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
	v40 = F_tzparse(m, int32(1790976), v25, int32(1))
	mBase = m.M
	goto L5
L13:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _consts[1308]))
	*(*int32)(unsafe.Add(mBase, _consts[1309])) = v543 + int32(22120)
	v1210 = v540
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
	v99 = base.I64_div_s(v95, v98)
	v102 = v95 - v99*v98
	__phi107 = int32(1970)
	__phi112 = v99
	v107 = __phi107
	v112 = __phi112
	goto L26
L18:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v95 = v91
	v96 = int64(0)
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
	if v78 != v74 {
		v95 = v74
		v96 = v80
		v97 = int32(0)
		goto L17
	} else {
		goto L23
	}
L21:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v77 = v44 + int32(22632) + v71<<(uint(int32(4))%32)
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
		v95 = v74
		v96 = v80
		v97 = base.B2i32(int64(0) < v80)
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v77-int32(8))))
	v95 = v74
	v96 = v80
	v97 = base.B2i32(v89 < v80)
	goto L17
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(61)
	v540 = int32(0)
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
	v222 = base.I32_wrap_i64(v112)
	v223 = base.I64_extend_i32_s(v42)
	v225 = v223 - v96 + v102
	if v225 < int64(0) {
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
	v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v130<<(uint(int32(2))%32))+uint32(_consts[1310]))))
	if v112 < v135 {
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
	v143 = int32(-1)
	goto L39
L38:
	;
	v143 = int32(1)
	goto L39
L39:
	;
	v145 = base.I64_div_s(v112, int64(366))
	if base.Ui64(v112+int64(365)) < base.Ui64(int64(731)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v151 = v143
	goto L42
L41:
	;
	v151 = base.I32_wrap_i64(v145)
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
	v160 = v151 + v107
	v162 = v160 - int32(1)
	if v162 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v151 <= v107^int32(2147483647) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v151 < int32(-2147483648)-v107 {
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
	v194 = v107 - int32(1)
	if v194 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v166 = int32(0) - v160
	v170 = base.I32_div_u_s(v166, int32(100))
	v173 = base.I32_div_u_s(v166, int32(400))
	v186 = int32(base.Ui32(v166)>>(uint(int32(2))%32)) - v170 + v173 ^ int32(-1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v180 = base.I32_div_u_s(v162, int32(100))
	v183 = base.I32_div_u_s(v162, int32(400))
	v186 = int32(base.Ui32(v162)>>(uint(int32(2))%32)) - v180 + v183
	goto L49
L53:
	;
	__phi107 = v160
	__phi112 = (base.I64_extend_i32_s(v160)-base.I64_extend_i32_s(v107))*int64(-365) + v112 - base.I64_extend_i32_s(v186-v218)
	v107 = __phi107
	v112 = __phi112
	goto L26
L54:
	;
	v198 = int32(0) - v107
	v202 = base.I32_div_u_s(v198, int32(100))
	v205 = base.I32_div_u_s(v198, int32(400))
	v218 = int32(base.Ui32(v198)>>(uint(int32(2))%32)) - v202 + v205 ^ int32(-1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v212 = base.I32_div_u_s(v194, int32(100))
	v215 = base.I32_div_u_s(v194, int32(400))
	v218 = int32(base.Ui32(v194)>>(uint(int32(2))%32)) - v212 + v215
	goto L53
L57:
	;
	v228 = int64(-86400)
	if base.Ui64(v225) <= base.Ui64(v228) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v252 = v222
	v253 = v225
	goto L59
L59:
	;
	if int64(86400) <= v253 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v231 = v228
	goto L62
L61:
	;
	v231 = v225
	goto L62
L62:
	;
	v233 = v96 + v231 - v102
	v235 = base.I64_extend_i32_u(base.B2i32(v233 != v223))
	v238 = int64(86400)
	v239 = base.I64_div_u_s(v233-(v235+v223), v238)
	v240 = v239 + v235
	v252 = base.I32_wrap_i64(v240) ^ int32(-1) + v222
	v253 = v102 + v240*v238 + v223 - v96 + v238
	goto L59
L63:
	;
	v257 = v253 - int64(172799)
	if base.Ui64(v257) <= base.Ui64(v253) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v274 = v252
	v275 = v253
	goto L65
L65:
	;
	if v274 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v260 = v257
	goto L68
L67:
	;
	v260 = int64(0)
	goto L68
L68:
	;
	v263 = int64(86400)
	v264 = base.I64_div_u_s(v260+int64(86399), v263)
	v274 = v252 + base.I32_wrap_i64(v264) + int32(1)
	v275 = v253 + v264*int64(-86400) - v263
	goto L65
L69:
	;
	v279 = v274
	v282 = v107
	goto L72
L70:
	;
	v314 = v274
	v317 = v107
	goto L71
L71:
	;
	v326 = v314
	v329 = v317
	goto L79
L72:
	;
	if v282 == int32(-2147483648) {
		goto L25
	} else {
		goto L74
	}
L73:
	;
	v314 = v311
	v317 = v295
	goto L71
L74:
	;
	v295 = v282 - int32(1)
	if v295&int32(3) != 0 {
		v305 = int32(0)
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v305<<(uint(int32(2))%32))+uint32(_consts[1310])))
	v311 = v310 + v279
	if v311 < int32(0) {
		v279 = v311
		v282 = v295
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v300 = base.I32_rem_s(v295, int32(100))
	if v300 != 0 {
		v305 = int32(1)
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v302 = base.I32_rem_s(v295, int32(400))
	v305 = base.B2i32(v302 == int32(0))
	goto L75
L78:
	;
	goto L73
L79:
	;
	v339 = v329 & int32(3)
	if v339 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1311])) = v329
	if v329 < int32(-2147481748) {
		goto L25
	} else {
		goto L93
	}
L81:
	;
	goto L80
L82:
	;
	if v329 == int32(2147483647) {
		goto L25
	} else {
		goto L92
	}
L83:
	;
	v343 = base.I32_rem_s(v329, int32(100))
	if v343 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	if v326 < int32(365) {
		goto L81
	} else {
		goto L91
	}
L86:
	;
	v347 = base.I32_rem_s(v329, int32(400))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v347 == int32(0))<<(uint(int32(2))%32))+uint32(_consts[1310])))
	if v326 < v354 {
		goto L81
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v326 < int32(366) {
		goto L81
	} else {
		goto L90
	}
L89:
	;
	v357 = base.I32_rem_s(v329, int32(400))
	v366 = base.B2i32(v357 == int32(0))
	goto L82
L90:
	;
	v366 = int32(1)
	goto L82
L91:
	;
	v366 = int32(0)
	goto L82
L92:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v366<<(uint(int32(2))%32))+uint32(_consts[1310])))
	v326 = v326 - v375
	v329 = v329 + int32(1)
	goto L79
L93:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1312])) = v326
	*(*int32)(unsafe.Add(mBase, _consts[1311])) = v329 - int32(1900)
	v390 = base.I32_rem_s(v329-int32(1970), int32(7))
	if v329 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v417 = int32(0)
	v420 = base.I64_div_u_s(v275, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _consts[1313])) = uint32(v420)
	v427 = int32(7)
	v428 = base.I32_rem_s(v326+v390+v416-int32(473), v427)
	if v428 < v417 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v394 = int32(0) - v329
	v398 = base.I32_div_u_s(v394, int32(100))
	v401 = base.I32_div_u_s(v394, int32(400))
	v416 = int32(base.Ui32(v394)>>(uint(int32(2))%32)) - v398 + v401 ^ int32(-1)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v406 = v329 - int32(1)
	v410 = base.I32_div_u_s(v406, int32(100))
	v413 = base.I32_div_u_s(v406, int32(400))
	v416 = int32(base.Ui32(v406)>>(uint(int32(2))%32)) - v410 + v413
	goto L94
L98:
	;
	v433 = v428 + v427
	goto L100
L99:
	;
	v433 = v428
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1314])) = v433
	v439 = base.I32_wrap_i64(v275 - v420*int64(3600))
	v440 = int32(65535)
	v442 = int32(60)
	v443 = base.I32_div_u_s(v439&v440, v442)
	*(*int32)(unsafe.Add(mBase, _consts[1315])) = v443
	*(*int32)(unsafe.Add(mBase, _consts[1316])) = v97 + (v439-v443*v442)&v440
	if v339 != 0 {
		v461 = int32(0)
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v463 = v461 * int32(48)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_consts[1304])))
	if v466 <= v326 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v456 = base.I32_rem_s(v329, int32(100))
	if v456 != 0 {
		v461 = int32(1)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v458 = base.I32_rem_s(v329, int32(400))
	v461 = base.B2i32(v458 == int32(0))
	goto L101
L104:
	;
	v468 = v326
	v471 = v466
	v472 = v417
	goto L107
L105:
	;
	v488 = v326
	v492 = v417
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1317])) = v42
	*(*int32)(unsafe.Add(mBase, _consts[1318])) = v492
	*(*int32)(unsafe.Add(mBase, _consts[1319])) = v488 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1320])) = int32(0)
	v540 = int32(4444200)
	goto L13
L107:
	;
	v480 = v468 - v471
	v482 = v472 + int32(1)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v463+int32(1790880)+v482<<(uint(int32(2))%32))))
	if v486 <= v480 {
		v468 = v480
		v471 = v486
		v472 = v482
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v488 = v480
	v492 = v482
	goto L106
L109:
	;
	goto L108
L110:
	;
	v629 = v15 + int32(8)
	if v626 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L111:
	;
	v572 = l0 + int32(24)
	if v17 < v570 {
		goto L120
	} else {
		goto L121
	}
L112:
	;
	v550 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < v550 {
		v570 = v550
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v553 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v626 = v556
	goto L110
L117:
	;
	goto L118
L118:
	;
	v558 = l0 + int32(24)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v558+v559<<(uint(int32(3))%32)-int32(8))))
	if v17 <= v565 {
		v626 = v559
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v558)))
	v570 = v567
	goto L111
L120:
	;
	v583 = v570 - v17
	goto L122
L121:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v575<<(uint(int32(3))%32)+v572-int32(8))))
	v583 = v17 - v581
	goto L122
L122:
	;
	v587 = base.I64_div_s(v583-int64(1), int64(12622780800))
	v588 = int64(400)
	v589 = v587 * v588
	v591 = v589 + v588
	v593 = v591 * int64(31556952)
	if v17 < v570 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v597 = v593
	goto L125
L124:
	;
	v597 = int64(0) - v593
	goto L125
L125:
	;
	v598 = v597 + v17
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v598
	v600 = int32(0)
	if v598 < v570 {
		v1210 = v600
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v602<<(uint(int32(3))%32)+v572-int32(8))))
	if v608 < v598 {
		v1210 = v600
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v610 = F_localsub(m, l0, v15)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	if v610 == int32(0) {
		v1210 = v600
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v614 = int64(*(*int32)(unsafe.Add(mBase, uint32(v610)+20)))
	v617 = *(*int64)(unsafe.Add(mBase, uint32(v572)))
	if v17 < v617 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v619 = int64(-400) - v589
	goto L132
L131:
	;
	v619 = v591
	goto L132
L132:
	;
	v620 = v614 + v619
	if base.Ui64(int64(4294967295)) < base.Ui64(v620+int64(2147483648)) {
		v1210 = v600
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v610)+20)) = uint32(v620)
	v1210 = v610
	goto L1
L134:
	;
	v699 = l0 + v696<<(uint(int32(4))%32)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v699)+uint32(_consts[1321])))
	if l0 != 0 {
		goto L151
	} else {
		goto L152
	}
L135:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1322])))
	v696 = v683
	goto L134
L136:
	;
	v632 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v17 < v632 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v634 = int32(1)
	if v634 < v626 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v640 = v626
	v641 = v634
	goto L141
L139:
	;
	v678 = v3
	goto L140
L140:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678+l0)+uint32(_consts[1323]))))
	v696 = v682
	goto L134
L141:
	;
	v652 = int32(1)
	v653 = (v640 + v641) >> (uint(v652) % 32)
	v659 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)+v653<<(uint(int32(3))%32))))
	v660 = base.B2i32(v17 < v659)
	if v17 < v659 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v678 = v661 - int32(1)
	goto L140
L143:
	;
	v661 = v641
	goto L145
L144:
	;
	v661 = v653 + v652
	goto L145
L145:
	;
	if v17 < v659 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v662 = v653
	goto L148
L147:
	;
	v662 = v640
	goto L148
L148:
	;
	if v661 < v662 {
		v640 = v662
		v641 = v661
		goto L141
	} else {
		goto L149
	}
L149:
	;
	goto L142
L150:
	;
	if v1198 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L151:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v713 = v712
	goto L153
L152:
	;
	v713 = int32(0)
	goto L153
L153:
	;
	v719 = v713
	goto L156
L154:
	;
	v756 = int64(86400)
	v757 = base.I64_div_s(v753, v756)
	v760 = v753 - v757*v756
	__phi765 = int32(1970)
	__phi770 = v757
	v765 = __phi765
	v770 = __phi770
	goto L163
L155:
	;
	v749 = *(*int64)(unsafe.Add(mBase, uint32(v629)))
	v753 = v749
	v754 = int64(0)
	v755 = int32(0)
	goto L154
L156:
	;
	v729 = v719 - int32(1)
	if v729 < int32(0) {
		goto L155
	} else {
		goto L158
	}
L157:
	;
	v738 = *(*int64)(unsafe.Add(mBase, uint32(v735)+8))
	if v736 != v732 {
		v753 = v732
		v754 = v738
		v755 = int32(0)
		goto L154
	} else {
		goto L160
	}
L158:
	;
	v732 = *(*int64)(unsafe.Add(mBase, uint32(v629)))
	v735 = l0 + int32(22632) + v729<<(uint(int32(4))%32)
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v735)))
	if v732 < v736 {
		v719 = v729
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	if v729 == int32(0) {
		v753 = v732
		v754 = v738
		v755 = base.B2i32(int64(0) < v738)
		goto L154
	} else {
		goto L161
	}
L161:
	;
	v747 = *(*int64)(unsafe.Add(mBase, uint32(v735-int32(8))))
	v753 = v732
	v754 = v738
	v755 = base.B2i32(v747 < v738)
	goto L154
L162:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(61)
	v1198 = int32(0)
	goto L150
L163:
	;
	v775 = base.B2i32(v770 < int64(0))
	if v775 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v880 = base.I32_wrap_i64(v770)
	v881 = base.I64_extend_i32_s(v702)
	v883 = v881 - v754 + v760
	if v883 < int64(0) {
		goto L194
	} else {
		goto L195
	}
L165:
	;
	goto L164
L166:
	;
	if v765&int32(3) != 0 {
		v788 = int32(0)
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L168
L168:
	;
	if base.Ui64(int64(1571958030700)) < base.Ui64(v770+int64(785979015533)) {
		goto L162
	} else {
		goto L173
	}
L169:
	;
	v793 = int64(*(*int32)(unsafe.Add(mBase, uint32(v788<<(uint(int32(2))%32))+uint32(_consts[1310]))))
	if v770 < v793 {
		goto L165
	} else {
		goto L172
	}
L170:
	;
	v783 = base.I32_rem_s(v765, int32(100))
	if v783 != 0 {
		v788 = int32(1)
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v785 = base.I32_rem_s(v765, int32(400))
	v788 = base.B2i32(v785 == int32(0))
	goto L169
L172:
	;
	goto L168
L173:
	;
	if v770 < int64(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v801 = int32(-1)
	goto L176
L175:
	;
	v801 = int32(1)
	goto L176
L176:
	;
	v803 = base.I64_div_s(v770, int64(366))
	if base.Ui64(v770+int64(365)) < base.Ui64(int64(731)) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v809 = v801
	goto L179
L178:
	;
	v809 = base.I32_wrap_i64(v803)
	goto L179
L179:
	;
	if int32(0) <= v765 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v818 = v809 + v765
	v820 = v818 - int32(1)
	if v820 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L181:
	;
	if v809 <= v765^int32(2147483647) {
		goto L180
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	if v809 < int32(-2147483648)-v765 {
		goto L162
	} else {
		goto L185
	}
L184:
	;
	goto L162
L185:
	;
	goto L180
L186:
	;
	v852 = v765 - int32(1)
	if v852 < int32(0) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v824 = int32(0) - v818
	v828 = base.I32_div_u_s(v824, int32(100))
	v831 = base.I32_div_u_s(v824, int32(400))
	v844 = int32(base.Ui32(v824)>>(uint(int32(2))%32)) - v828 + v831 ^ int32(-1)
	goto L186
L188:
	;
	goto L189
L189:
	;
	v838 = base.I32_div_u_s(v820, int32(100))
	v841 = base.I32_div_u_s(v820, int32(400))
	v844 = int32(base.Ui32(v820)>>(uint(int32(2))%32)) - v838 + v841
	goto L186
L190:
	;
	__phi765 = v818
	__phi770 = (base.I64_extend_i32_s(v818)-base.I64_extend_i32_s(v765))*int64(-365) + v770 - base.I64_extend_i32_s(v844-v876)
	v765 = __phi765
	v770 = __phi770
	goto L163
L191:
	;
	v856 = int32(0) - v765
	v860 = base.I32_div_u_s(v856, int32(100))
	v863 = base.I32_div_u_s(v856, int32(400))
	v876 = int32(base.Ui32(v856)>>(uint(int32(2))%32)) - v860 + v863 ^ int32(-1)
	goto L190
L192:
	;
	goto L193
L193:
	;
	v870 = base.I32_div_u_s(v852, int32(100))
	v873 = base.I32_div_u_s(v852, int32(400))
	v876 = int32(base.Ui32(v852)>>(uint(int32(2))%32)) - v870 + v873
	goto L190
L194:
	;
	v886 = int64(-86400)
	if base.Ui64(v883) <= base.Ui64(v886) {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	v910 = v880
	v911 = v883
	goto L196
L196:
	;
	if int64(86400) <= v911 {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	v889 = v886
	goto L199
L198:
	;
	v889 = v883
	goto L199
L199:
	;
	v891 = v754 + v889 - v760
	v893 = base.I64_extend_i32_u(base.B2i32(v891 != v881))
	v896 = int64(86400)
	v897 = base.I64_div_u_s(v891-(v893+v881), v896)
	v898 = v897 + v893
	v910 = base.I32_wrap_i64(v898) ^ int32(-1) + v880
	v911 = v760 + v898*v896 + v881 - v754 + v896
	goto L196
L200:
	;
	v915 = v911 - int64(172799)
	if base.Ui64(v915) <= base.Ui64(v911) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v932 = v910
	v933 = v911
	goto L202
L202:
	;
	if v932 < int32(0) {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v918 = v915
	goto L205
L204:
	;
	v918 = int64(0)
	goto L205
L205:
	;
	v921 = int64(86400)
	v922 = base.I64_div_u_s(v918+int64(86399), v921)
	v932 = v910 + base.I32_wrap_i64(v922) + int32(1)
	v933 = v911 + v922*int64(-86400) - v921
	goto L202
L206:
	;
	v937 = v932
	v940 = v765
	goto L209
L207:
	;
	v972 = v932
	v975 = v765
	goto L208
L208:
	;
	v984 = v972
	v987 = v975
	goto L216
L209:
	;
	if v940 == int32(-2147483648) {
		goto L162
	} else {
		goto L211
	}
L210:
	;
	v972 = v969
	v975 = v953
	goto L208
L211:
	;
	v953 = v940 - int32(1)
	if v953&int32(3) != 0 {
		v963 = int32(0)
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v963<<(uint(int32(2))%32))+uint32(_consts[1310])))
	v969 = v968 + v937
	if v969 < int32(0) {
		v937 = v969
		v940 = v953
		goto L209
	} else {
		goto L215
	}
L213:
	;
	v958 = base.I32_rem_s(v953, int32(100))
	if v958 != 0 {
		v963 = int32(1)
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v960 = base.I32_rem_s(v953, int32(400))
	v963 = base.B2i32(v960 == int32(0))
	goto L212
L215:
	;
	goto L210
L216:
	;
	v997 = v987 & int32(3)
	if v997 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1311])) = v987
	if v987 < int32(-2147481748) {
		goto L162
	} else {
		goto L230
	}
L218:
	;
	goto L217
L219:
	;
	if v987 == int32(2147483647) {
		goto L162
	} else {
		goto L229
	}
L220:
	;
	v1001 = base.I32_rem_s(v987, int32(100))
	if v1001 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	if v984 < int32(365) {
		goto L218
	} else {
		goto L228
	}
L223:
	;
	v1005 = base.I32_rem_s(v987, int32(400))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v1005 == int32(0))<<(uint(int32(2))%32))+uint32(_consts[1310])))
	if v984 < v1012 {
		goto L218
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	if v984 < int32(366) {
		goto L218
	} else {
		goto L227
	}
L226:
	;
	v1015 = base.I32_rem_s(v987, int32(400))
	v1024 = base.B2i32(v1015 == int32(0))
	goto L219
L227:
	;
	v1024 = int32(1)
	goto L219
L228:
	;
	v1024 = int32(0)
	goto L219
L229:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1024<<(uint(int32(2))%32))+uint32(_consts[1310])))
	v984 = v984 - v1033
	v987 = v987 + int32(1)
	goto L216
L230:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1312])) = v984
	*(*int32)(unsafe.Add(mBase, _consts[1311])) = v987 - int32(1900)
	v1048 = base.I32_rem_s(v987-int32(1970), int32(7))
	if v987 <= int32(0) {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1075 = int32(0)
	v1078 = base.I64_div_u_s(v933, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _consts[1313])) = uint32(v1078)
	v1085 = int32(7)
	v1086 = base.I32_rem_s(v984+v1048+v1074-int32(473), v1085)
	if v1086 < v1075 {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	v1052 = int32(0) - v987
	v1056 = base.I32_div_u_s(v1052, int32(100))
	v1059 = base.I32_div_u_s(v1052, int32(400))
	v1074 = int32(base.Ui32(v1052)>>(uint(int32(2))%32)) - v1056 + v1059 ^ int32(-1)
	goto L231
L233:
	;
	goto L234
L234:
	;
	v1064 = v987 - int32(1)
	v1068 = base.I32_div_u_s(v1064, int32(100))
	v1071 = base.I32_div_u_s(v1064, int32(400))
	v1074 = int32(base.Ui32(v1064)>>(uint(int32(2))%32)) - v1068 + v1071
	goto L231
L235:
	;
	v1091 = v1086 + v1085
	goto L237
L236:
	;
	v1091 = v1086
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1314])) = v1091
	v1097 = base.I32_wrap_i64(v933 - v1078*int64(3600))
	v1098 = int32(65535)
	v1100 = int32(60)
	v1101 = base.I32_div_u_s(v1097&v1098, v1100)
	*(*int32)(unsafe.Add(mBase, _consts[1315])) = v1101
	*(*int32)(unsafe.Add(mBase, _consts[1316])) = v755 + (v1097-v1101*v1100)&v1098
	if v997 != 0 {
		v1119 = int32(0)
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1121 = v1119 * int32(48)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+uint32(_consts[1304])))
	if v1124 <= v984 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v1114 = base.I32_rem_s(v987, int32(100))
	if v1114 != 0 {
		v1119 = int32(1)
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v1116 = base.I32_rem_s(v987, int32(400))
	v1119 = base.B2i32(v1116 == int32(0))
	goto L238
L241:
	;
	v1126 = v984
	v1129 = v1124
	v1130 = v1075
	goto L244
L242:
	;
	v1146 = v984
	v1150 = v1075
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1317])) = v702
	*(*int32)(unsafe.Add(mBase, _consts[1318])) = v1150
	*(*int32)(unsafe.Add(mBase, _consts[1319])) = v1146 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1320])) = int32(0)
	v1198 = int32(4444200)
	goto L150
L244:
	;
	v1138 = v1126 - v1129
	v1140 = v1130 + int32(1)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1121+int32(1790880)+v1140<<(uint(int32(2))%32))))
	if v1144 <= v1138 {
		v1126 = v1138
		v1129 = v1144
		v1130 = v1140
		goto L244
	} else {
		goto L246
	}
L245:
	;
	v1146 = v1138
	v1150 = v1140
	goto L243
L246:
	;
	goto L245
L247:
	;
	v1210 = int32(0)
	goto L1
L248:
	;
	goto L249
L249:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+uint32(_consts[1324]))))
	*(*int32)(unsafe.Add(mBase, uint32(v1198)+32)) = v1202
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v699)+uint32(_consts[1325])))
	*(*int32)(unsafe.Add(mBase, uint32(v1198)+40)) = l0 + v1204 + int32(22120)
	v1210 = v1198
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
	v13 = F_query_or_expression_tree_walker_impl(m, l0, int32(1045), v5+int32(12), int32(0))
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[462]))
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
	v37 = int32(30435)
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
	v40 = int32(233875)
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
	v86 = F_logfile_open(m, v41, v83, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L36
	}
L23:
	;
	v83 = int32(485315)
	goto L22
L24:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[493])))
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
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v58 == int32(0) {
		v77 = v57
		v78 = v58
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v78-v77 != 0 {
		v83 = int32(30419)
		goto L22
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	if v57 != v58 {
		v77 = v57
		v78 = v58
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v62 = v41
	v63 = v51
	goto L31
L31:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v66
		v78 = v67
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v77 = v66
	v78 = v67
	goto L28
L33:
	;
	v70 = int32(1)
	if v66 == v67 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L23
L36:
	;
	if v86 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	switch v91 - int32(33) {
	case 0, 8:
		goto L40
	default:
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v117 != 0 {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v110 = int32(0)
	if v41 == v110 {
		v129 = v110
		goto L13
	} else {
		goto L48
	}
L41:
	;
	v96 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	if v96 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_errmsg(m, int32(637287), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[494])) = uint8(v108)
	goto L40
L46:
	;
	F_errfinish(m, int32(473982), int32(1337), int32(74444))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_pfree(m, v41)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	return int32(0)
L50:
	;
	v118 = F_fclose(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v86
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v121 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	F_pfree(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v41
	v129 = int32(1)
	goto L13
L57:
	;
	goto L56
}
func F_lookup_rowtype_tupdesc_domain(m *base.Module, l0 int32) int32 {
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != int32(2249) {
		v11 = F_lookup_type_cache(m, l0, int32(4352))
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
						v61 = int32(0)
						m.G0 = v6 + int32(16)
						return v61
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
						if int32(0) <= v25 {
							v58 = v20
							F_IncrTupleDescRefCount(m, v58)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = v58
								m.G0 = v6 + int32(16)
								return v61
							}
						} else {
							v61 = v20
							m.G0 = v6 + int32(16)
							return v61
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
				if v28 != 0 {
					v51 = v28
					if v51 == int32(0) {
						v61 = int32(0)
						m.G0 = v6 + int32(16)
						return v61
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
						if v55 < int32(0) {
							v61 = v51
							m.G0 = v6 + int32(16)
							return v61
						} else {
							v58 = v51
							F_IncrTupleDescRefCount(m, v58)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = v58
								m.G0 = v6 + int32(16)
								return v61
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
								F_errmsg(m, int32(333559), v6)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(477482), int32(2001), int32(265105))
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
		v49 = F_lookup_rowtype_tupdesc_internal(m, int32(2249), int32(-1))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			v51 = v49
			if v51 == int32(0) {
				v61 = int32(0)
				m.G0 = v6 + int32(16)
				return v61
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
				if v55 < int32(0) {
					v61 = v51
					m.G0 = v6 + int32(16)
					return v61
				} else {
					v58 = v51
					F_IncrTupleDescRefCount(m, v58)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = v58
						m.G0 = v6 + int32(16)
						return v61
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v48 = F_str_tolower(m, v18, v46, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v50 = F_cstring_to_text(m, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v48)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					return v50
				}
			}
		}
	}
}

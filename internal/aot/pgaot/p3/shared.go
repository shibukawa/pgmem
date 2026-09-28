package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AttachSharedMemoryStructs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	v9 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[0]))
	if v12&(v12-v9) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[1])) = int32(3)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[2]))
	v37 = F_LWLockAcquire(m, v33+int32(16), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v19 = v9 << (uint(int32(32)-base.I32_clz(v12)) % 32)
	goto L4
L3:
	;
	v19 = v12
	goto L4
L4:
	;
	if base.Ui32(v19) <= base.Ui32(int32(31)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = int32(31)
	goto L7
L6:
	;
	v22 = v19
	goto L7
L7:
	;
	if base.Ui32(int32(_a_F_AttachSharedMemoryStructs_0)) <= base.Ui32(v19) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = int32(1024)
	goto L10
L9:
	;
	v27 = int32(base.Ui32(v22) >> (uint(int32(4)) % 32))
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[3])) = v27
	goto L1
L11:
	;
	return
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[4]))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if int32(0) < v41 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v74 = int32(0)
	goto L15
L15:
	;
	F_list_free_deep(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L24
	}
L16:
	;
	v44 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[4]))
	v74 = v68
	goto L15
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v44<<(uint(int32(2))%32))))
	v54 = F_AttachShmemIndexEntry(m, v52, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	F_pfree(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v60 = v44 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v60 < v61 {
		v44 = v60
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[4])) = v78
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[5]))
	if v81 == v78 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[2]))
	F_LWLockRelease(m, v110+int32(16))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L35
	}
L26:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v85 <= v84 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v88 = v84
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v88<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	if v97 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L25
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	m.T0[v97].(func(*base.Module, int32))(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v102 = v88 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v102 < v103 {
		v88 = v102
		goto L28
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[1])) = int32(5)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_AttachSharedMemoryStructs[6]))
	if v119 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	m.T0[v119].(func(*base.Module))(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L11
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	return
L39:
	;
	goto L38
}
func F_InitSharedLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v6 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v6)
	return
}
func F_MarkSharedBufferDirtyHint(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v82 int64
	_ = v82
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v151 int64
	_ = v151
	var v155 int64
	_ = v155
	var v165 int64
	_ = v165
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l2&int64(8388608) != int64(0) {
		m.G0 = v10 + int32(32)
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[0]))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[1])))
		if v19 == int32(0) {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[2]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+268))
			if l2&int64(2147483648) == int64(0) {
				v144 = F_LockBufHdr(m, l1)
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return
				} else {
					v151 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v144, v144&int64(-12582913)|int64(8388608))
					if v144 == v151 {
					} else {
						v155 = v151
						for {
							v165 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v155, v155&int64(-12582913)|int64(8388608))
							if v155 != v165 {
								v155 = v165
								continue
							} else {
								break
							}
							break
						}
					}
					v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
					v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
					*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
					v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
					if v181 != int32(1) {
					} else {
						v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
						v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
						v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
						*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
					}
					m.G0 = v10 + int32(32)
					return
				}
			} else {
				if v24 != int32(0) {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[7])))
					if v37 == int32(1) {
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[8]))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+308))
						v45 = base.B2i32(v43 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[7])) = uint8(v45)
						v47 = v45
					} else {
						v47 = int32(0)
					}
					if v47 != 0 {
						m.G0 = v10 + int32(32)
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v50
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v52
						v54 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v52
						v59 = F_RelFileLocatorSkippingWAL(m, v10+int32(8))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							if v59 != 0 {
								m.G0 = v10 + int32(32)
								return
							} else {
								v61 = F_LockBufHdr(m, l1)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v68 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v61, v61&int64(-12582913)|int64(8388608))
									if v61 != v68 {
										v72 = v68
										for {
											v82 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v72, v72&int64(-12582913)|int64(8388608))
											if v72 != v82 {
												v72 = v82
												continue
											} else {
												break
											}
											break
										}
									} else {
									}
									v91 = F_GetRedoRecPtr(m)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										if l0 < int32(0) {
											v96 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[9]))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+(l0^int32(-1))<<(uint(int32(2))%32))))
											v110 = v102
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[0]))
											v110 = v104 + l0<<(uint(int32(13))%32) + int32(-8192)
										}
										v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
										if base.Ui64(base.I64_rotl(v111, int64(32))) <= base.Ui64(v91) {
											F_XLogBeginInsert(m)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												v117 = int32(0)
												if l3 != 0 {
													v120 = int32(8)
												} else {
													v120 = v117
												}
												F_XLogRegisterBuffer(m, v117, l0, v120)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return
												} else {
													v125 = F_XLogInsert(m, int32(0), int32(160))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														v128 = v125
														if v128 == int64(0) {
															v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
															v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
															*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
															v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
															if v181 != int32(1) {
															} else {
																v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
																v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
																v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
																*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
															}
															m.G0 = v10 + int32(32)
															return
														} else {
															v131 = F_LockBufHdr(m, l1)
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v17+l0<<(uint(int32(13))%32))+uint32(_c_F_MarkSharedBufferDirtyHint[10]))) = base.I64_rotl(v128, int64(32))
																v143 = base.AtomicRmwSub64(m, l1, int32(24), int64(4194304))
																v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
																v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
																*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
																v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
																if v181 != int32(1) {
																} else {
																	v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
																	v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
																	v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
																	*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
																}
																m.G0 = v10 + int32(32)
																return
															}
														}
													}
												}
											}
										} else {
											v128 = int64(0)
											if v128 == int64(0) {
												v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
												v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
												*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
												v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
												if v181 != int32(1) {
												} else {
													v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
													v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
													*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
												}
												m.G0 = v10 + int32(32)
												return
											} else {
												v131 = F_LockBufHdr(m, l1)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v17+l0<<(uint(int32(13))%32))+uint32(_c_F_MarkSharedBufferDirtyHint[10]))) = base.I64_rotl(v128, int64(32))
													v143 = base.AtomicRmwSub64(m, l1, int32(24), int64(4194304))
													v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
													v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
													*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
													v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
													if v181 != int32(1) {
													} else {
														v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
														v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
														v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
														*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
													}
													m.G0 = v10 + int32(32)
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
					v144 = F_LockBufHdr(m, l1)
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return
					} else {
						v151 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v144, v144&int64(-12582913)|int64(8388608))
						if v144 == v151 {
						} else {
							v155 = v151
							for {
								v165 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v155, v155&int64(-12582913)|int64(8388608))
								if v155 != v165 {
									v155 = v165
									continue
								} else {
									break
								}
								break
							}
						}
						v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
						v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
						*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
						v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
						if v181 != int32(1) {
						} else {
							v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
							v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
							v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
							*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
						}
						m.G0 = v10 + int32(32)
						return
					}
				}
			}
		} else {
			if l2&int64(2147483648) == int64(0) {
				v144 = F_LockBufHdr(m, l1)
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return
				} else {
					v151 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v144, v144&int64(-12582913)|int64(8388608))
					if v144 == v151 {
					} else {
						v155 = v151
						for {
							v165 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v155, v155&int64(-12582913)|int64(8388608))
							if v155 != v165 {
								v155 = v165
								continue
							} else {
								break
							}
							break
						}
					}
					v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
					v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
					*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
					v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
					if v181 != int32(1) {
					} else {
						v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
						v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
						v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
						*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
					}
					m.G0 = v10 + int32(32)
					return
				}
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[7])))
				if v37 == int32(1) {
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[8]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+308))
					v45 = base.B2i32(v43 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[7])) = uint8(v45)
					v47 = v45
				} else {
					v47 = int32(0)
				}
				if v47 != 0 {
					m.G0 = v10 + int32(32)
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v50
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v52
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v10)+20))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v52
					v59 = F_RelFileLocatorSkippingWAL(m, v10+int32(8))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						if v59 != 0 {
							m.G0 = v10 + int32(32)
							return
						} else {
							v61 = F_LockBufHdr(m, l1)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v68 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v61, v61&int64(-12582913)|int64(8388608))
								if v61 != v68 {
									v72 = v68
									for {
										v82 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v72, v72&int64(-12582913)|int64(8388608))
										if v72 != v82 {
											v72 = v82
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v91 = F_GetRedoRecPtr(m)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									if l0 < int32(0) {
										v96 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[9]))
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+(l0^int32(-1))<<(uint(int32(2))%32))))
										v110 = v102
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[0]))
										v110 = v104 + l0<<(uint(int32(13))%32) + int32(-8192)
									}
									v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
									if base.Ui64(base.I64_rotl(v111, int64(32))) <= base.Ui64(v91) {
										F_XLogBeginInsert(m)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return
										} else {
											v117 = int32(0)
											if l3 != 0 {
												v120 = int32(8)
											} else {
												v120 = v117
											}
											F_XLogRegisterBuffer(m, v117, l0, v120)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												v125 = F_XLogInsert(m, int32(0), int32(160))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													v128 = v125
													if v128 == int64(0) {
														v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
														v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
														*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
														v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
														if v181 != int32(1) {
														} else {
															v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
															v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
															v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
															*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
														}
														m.G0 = v10 + int32(32)
														return
													} else {
														v131 = F_LockBufHdr(m, l1)
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v17+l0<<(uint(int32(13))%32))+uint32(_c_F_MarkSharedBufferDirtyHint[10]))) = base.I64_rotl(v128, int64(32))
															v143 = base.AtomicRmwSub64(m, l1, int32(24), int64(4194304))
															v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
															v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
															*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
															v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
															if v181 != int32(1) {
															} else {
																v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
																v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
																v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
																*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
															}
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										v128 = int64(0)
										if v128 == int64(0) {
											v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
											v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
											*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
											v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
											if v181 != int32(1) {
											} else {
												v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
												v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
												v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
											}
											m.G0 = v10 + int32(32)
											return
										} else {
											v131 = F_LockBufHdr(m, l1)
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v17+l0<<(uint(int32(13))%32))+uint32(_c_F_MarkSharedBufferDirtyHint[10]))) = base.I64_rotl(v128, int64(32))
												v143 = base.AtomicRmwSub64(m, l1, int32(24), int64(4194304))
												v174 = int32(_a_F_MarkSharedBufferDirtyHint_0)
												v176 = *(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3]))
												*(*int64)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[3])) = v176 + int64(1)
												v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[4])))
												if v181 != int32(1) {
												} else {
													v184 = int32(_a_F_MarkSharedBufferDirtyHint_1)
													v186 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5]))
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[6]))
													*(*int32)(unsafe.Add(mBase, _c_F_MarkSharedBufferDirtyHint[5])) = v186 + v188
												}
												m.G0 = v10 + int32(32)
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
}
func F_SharedBufferBeginSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v40 int64
	_ = v40
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v62 int64
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_SharedBufferBeginSetHintBits[0]))
	if v8 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L24
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L21
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v12 = v8 << (uint(int32(4)) % 32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_SharedBufferBeginSetHintBits[1])))
	if v15 == l0 {
		v25 = v12 + int32(_a_F_SharedBufferBeginSetHintBits_0)
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v19 = F_GetPrivateRefCountEntrySlow(m, l0, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	return int32(0)
L9:
	;
	if v19 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v25 = v19
	goto L3
L11:
	;
	v29 = int64(0)
	v32 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v29, v29)
	if v26&int32(-2) != int32(2) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = v32
	goto L15
L13:
	;
	v62 = v32
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v62
	return int32(1)
L15:
	;
	if v40&int64(13510798882111488) != int64(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = int32(2)
	v62 = v50
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v50 = v40 + int64(4503582447501312)
	v52 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v40, v50)
	if v40 != v52 {
		v40 = v52
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	F_errmsg_internal(m, int32(_a_F_SharedBufferBeginSetHintBits_1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_SharedBufferBeginSetHintBits_2), int32(_a_F_SharedBufferBeginSetHintBits_3), int32(_a_F_SharedBufferBeginSetHintBits_4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errmsg_internal(m, int32(_a_F_SharedBufferBeginSetHintBits_5), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_SharedBufferBeginSetHintBits_2), int32(_a_F_SharedBufferBeginSetHintBits_6), int32(_a_F_SharedBufferBeginSetHintBits_4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SharedInvalBackendInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v1 = l0
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[0]))
	if int32(0) <= v13 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[1]))
		if v17+int32(38) <= v13 {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return
			} else {
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v93
				v96 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v96 + int32(38)
				F_errmsg_internal(m, int32(_a_F_SharedInvalBackendInit_0), v10+int32(16))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_SharedInvalBackendInit_1), int32(284), int32(_a_F_SharedInvalBackendInit_2))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[2]))
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[3]))
			v28 = F_LWLockAcquire(m, v24+int32(768), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v32 = v22 + v13<<(uint(int32(4))%32)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[4])))
				if v33 != 0 {
					v111 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[3]))
					F_LWLockRelease(m, v111+int32(768))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v33
							v122 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v122
							F_errmsg_internal(m, int32(_a_F_SharedInvalBackendInit_3), v10)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_SharedInvalBackendInit_1), int32(299), int32(_a_F_SharedInvalBackendInit_2))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
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
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[0]))
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[2]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_SharedInvalBackendInit[5])))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_SharedInvalBackendInit[5]))) = v38 + int32(1)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_SharedInvalBackendInit[6])))
					*(*int32)(unsafe.Add(mBase, uint32(v42+v38<<(uint(int32(2))%32)))) = v35
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[7])))
					*(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[8])) = v50
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[9]))
					*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[4]))) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[10]))) = uint8(v1)
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[11]))) = uint8(v57)
					*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[12]))) = uint16(v57)
					*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_SharedInvalBackendInit[13]))) = v55
					v63 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalBackendInit[3]))
					F_LWLockRelease(m, v63+int32(768))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_on_shmem_exit(m, int32(1209), base.I64_extend_i32_u(v22))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_SharedInvalBackendInit_4), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_SharedInvalBackendInit_1), int32(281), int32(_a_F_SharedInvalBackendInit_2))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
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
func F_SharedInvalShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v2 = int32(0)
	v3 = int32(_a_F_SharedInvalShmemInit_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(8796093022208)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v8)+12)), uint32(v2))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[1]))
	if v2 < v15+int32(38) {
		v21 = v2
		for {
			v23 = v21 << (uint(int32(4)) % 32)
			v24 = int32(_a_F_SharedInvalShmemInit_0)
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23+v25)+uint32(_c_F_SharedInvalShmemInit[2]))) = v27
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v30+v23)+uint32(_c_F_SharedInvalShmemInit[3]))) = v27
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v37+v23)+uint32(_c_F_SharedInvalShmemInit[4]))) = uint8(v27)
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v44+v23)+uint32(_c_F_SharedInvalShmemInit[5]))) = uint8(v27)
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v51+v23)+uint32(_c_F_SharedInvalShmemInit[6]))) = uint8(v27)
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v58+v23)+uint32(_c_F_SharedInvalShmemInit[7]))) = v27
			v65 = v21 + int32(1)
			v67 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[1]))
			if v65 < v67+int32(38) {
				v21 = v65
				continue
			} else {
				break
			}
			break
		}
		v72 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemInit[0]))
		v73 = v72
		v74 = v65
	} else {
		v73 = v8
		v74 = v2
	}
	*(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_SharedInvalShmemInit[8]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_SharedInvalShmemInit[9]))) = v73 + v74<<(uint(int32(4))%32) + int32(_a_F_SharedInvalShmemInit_1)
	return
}
func F_StartSharedBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v40 int64
	_ = v40
	var v52 int64
	_ = v52
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v120 int64
	_ = v120
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v172 int64
	_ = v172
	var v178 int64
	_ = v178
	var v187 int64
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_StartSharedBufferIO[0]))
	F_ResourceOwnerEnlarge(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = l0 + int32(36)
	goto L4
L3:
	;
	m.G0 = v11 + int32(32)
	return v215
L4:
	;
	v29 = int64(4194304)
	v31 = base.AtomicRmwOr64(m, l0, int32(24), v29)
	if v31&v29 != int64(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if l1 != 0 {
		goto L42
	} else {
		goto L43
	}
L6:
	;
	v40 = v31
	goto L9
L7:
	;
	v120 = v31
	goto L8
L8:
	;
	if v120&int64(67108864) != int64(0) {
		goto L30
	} else {
		goto L31
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = int32(_a_F_StartSharedBufferIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(_a_F_StartSharedBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(_a_F_StartSharedBufferIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(0)
	v52 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
	if v40&int64(4194304) != v52 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v120 = v111
	goto L8
L11:
	;
	goto L14
L12:
	;
	goto L13
L13:
	;
	v89 = int32(_a_F_StartSharedBufferIO_3)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_StartSharedBufferIO[1]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(8))+8))
	if v92 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	F_perform_spin_delay(m, v11+int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v70 = int64(0)
	v73 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v70, v70)
	if v73&int64(4194304) != v70 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v109 = int64(4194304)
	v111 = base.AtomicRmwOr64(m, l0, int32(24), v109)
	if v111&v109 != int64(0) {
		v40 = v111
		goto L9
	} else {
		goto L29
	}
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartSharedBufferIO[1])) = v107
	goto L19
L21:
	;
	if int32(999) < v90 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v90 < int32(11) {
		goto L19
	} else {
		goto L28
	}
L24:
	;
	v97 = int32(900)
	if v97 <= v90 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v100 = v97
	goto L27
L26:
	;
	v100 = v90
	goto L27
L27:
	;
	v107 = v100 + int32(100)
	goto L20
L28:
	;
	v107 = v90 - int32(1)
	goto L20
L29:
	;
	goto L10
L30:
	;
	if l3 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	goto L5
L33:
	;
	v145 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	if l2 == int32(0) {
		v215 = int32(1)
		goto L3
	} else {
		goto L37
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	goto L35
L35:
	;
	if base.B2i32(v130 != int32(-1)) == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v137
	v141 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	v215 = int32(1)
	goto L3
L37:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L4
L40:
	;
	v166 = v120 | int64(4194304)
	v172 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v166, v120&int64(-71303169)|int64(67108864))
	if v172 != v166 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v163 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	v215 = int32(0)
	goto L3
L42:
	;
	if v120&int64(16777216) != int64(0) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v120&int64(8388608) != int64(0) {
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L40
L46:
	;
	goto L41
L47:
	;
	v178 = v172
	goto L50
L48:
	;
	goto L49
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_StartSharedBufferIO[0]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerRemember(m, v198, base.I64_extend_i32_s(v199+int32(1)), int32(_a_F_StartSharedBufferIO_4))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v187 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v178, v178&int64(-71303169)|int64(67108864))
	if v178 != v187 {
		v178 = v187
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	goto L51
L53:
	;
	v215 = int32(2)
	goto L3
}
func F_UnlockSharedObjectForSession(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(1262)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v2
	v18 = F_LockRelease(m, v5, int32(8), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_process_shared_preload_libraries(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_shared_preload_libraries[0])) = uint8(v2)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_process_shared_preload_libraries[1]))
	F_load_libraries(m, v5, int32(_a_F_process_shared_preload_libraries_0), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_process_shared_preload_libraries[2])) = uint8(v11)
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_process_shared_preload_libraries[0])) = uint8(v14)
		return
	}
}
func F_shared_buffer_readv_complete_local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = int32(0)
	if base.B2i32(v5&int32(448) == int32(64))|base.B2i32(v5&int32(33292288) == v12) == v12 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(104))+4))
		F_pgstat_report_checksum_failures_in_db(m, v19, int32(base.Ui32(v5)>>(uint(int32(18))%32))&int32(127))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v26
			return
		}
	} else {
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v26
		return
	}
}
func F_shared_record_typmod_registry_detach(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_shared_record_typmod_registry_detach[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v6 != 0 {
		F_pfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_shared_record_typmod_registry_detach[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
			v13 = v10
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			if v14 != 0 {
				F_pfree(m, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, _c_F_shared_record_typmod_registry_detach[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
					v21 = v18
					*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
					return
				}
			} else {
				v21 = v13
				*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
				return
			}
		}
	} else {
		v13 = v5
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		if v14 != 0 {
			F_pfree(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_shared_record_typmod_registry_detach[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
				v21 = v18
				*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
				return
			}
		} else {
			v21 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
			return
		}
	}
}
func F_shared_ts_extend_down(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
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
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v3 = l2
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = F_dsa_allocate_extended(m, v8, int32(24), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_dsa_get_address(m, v15, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
	v22 = l3 - int32(8)
	if int32(0) < v22 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = v16
	v32 = base.I64_extend_i32_u(v22)
	goto L7
L5:
	;
	v56 = v16
	goto L6
L6:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+2)) = uint8(v59)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+3)) = uint8(v3)
	return v56 + int32(8)
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = F_dsa_allocate_extended(m, v33, int32(24), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v56 = v39
	goto L6
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = F_dsa_get_address(m, v38, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v36
	v44 = int64(base.Ui64(v3) >> (uint(v32) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+3)) = uint8(v44)
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+2)) = uint8(v46)
	v48 = int64(8)
	if base.Ui64(v48) < base.Ui64(v32) {
		v30 = v39
		v32 = v32 - v48
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}

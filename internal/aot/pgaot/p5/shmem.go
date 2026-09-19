package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ShmemAlloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[0]))
	v15 = base.AtomicRmwXchg32(m, v12, int32(0), int32(1))
	if v15 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[0]))
		F_s_lock(m, v17, int32(_a_F_ShmemAlloc_0), int32(208), int32(_a_F_ShmemAlloc_1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[1]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
			v32 = v27 + (l0+int32(127))&int32(-128)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			if base.Ui32(v33) < base.Ui32(v32) {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[0]))
				v37 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36))), uint32(v37))
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_ShmemAlloc_2))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(_a_F_ShmemAlloc_3), v9)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ShmemAlloc_0), int32(162), int32(_a_F_ShmemAlloc_4))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v32
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[0]))
				v45 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v44))), uint32(v45))
				if v41 != 0 {
					m.G0 = v9 + int32(16)
					return v41 + v27
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(_a_F_ShmemAlloc_2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(_a_F_ShmemAlloc_3), v9)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ShmemAlloc_0), int32(162), int32(_a_F_ShmemAlloc_4))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[1]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
		v32 = v27 + (l0+int32(127))&int32(-128)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		if base.Ui32(v33) < base.Ui32(v32) {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[0]))
			v37 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36))), uint32(v37))
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_ShmemAlloc_2))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					F_errmsg(m, int32(_a_F_ShmemAlloc_3), v9)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ShmemAlloc_0), int32(162), int32(_a_F_ShmemAlloc_4))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
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
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v32
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAlloc[0]))
			v45 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v44))), uint32(v45))
			if v41 != 0 {
				m.G0 = v9 + int32(16)
				return v41 + v27
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_ShmemAlloc_2))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(_a_F_ShmemAlloc_3), v9)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ShmemAlloc_0), int32(162), int32(_a_F_ShmemAlloc_4))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
func F_ShmemInitStruct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
	v19 = F_LWLockAcquire(m, v15+int32(128), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[1]))
		if v24 == int32(0) {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[2]))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShmemInitStruct[3])))
			if v30 == int32(1) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
				v94 = v33
				v102 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
				F_LWLockRelease(m, v102+int32(128))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					m.G0 = v12 + int32(48)
					return v94
				}
			} else {
				v36 = F_ShmemAlloc(m, l1)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v36
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v39)
					v94 = v36
					v102 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
					F_LWLockRelease(m, v102+int32(128))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(48)
						return v94
					}
				}
			}
		} else {
			v42 = F_hash_search(m, v24, l0, int32(3), l2)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 == int32(0) {
					v112 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
					F_LWLockRelease(m, v112+int32(128))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_ShmemInitStruct_0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
								F_errmsg(m, int32(_a_F_ShmemInitStruct_1), v12)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ShmemInitStruct_2), int32(437), int32(_a_F_ShmemInitStruct_3))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
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
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
					if v46 == int32(1) {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
						if v49 != l1 {
							v134 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
							F_LWLockRelease(m, v134+int32(128))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v143
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
									F_errmsg(m, int32(_a_F_ShmemInitStruct_4), v12+int32(16))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ShmemInitStruct_2), int32(453), int32(_a_F_ShmemInitStruct_3))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
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
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
							v94 = v51
							v102 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
							F_LWLockRelease(m, v102+int32(128))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(48)
								return v94
							}
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[4]))
						v56 = base.AtomicRmwXchg32(m, v53, int32(0), int32(1))
						if v56 != 0 {
							v58 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[4]))
							F_s_lock(m, v58, int32(_a_F_ShmemInitStruct_2), int32(208), int32(_a_F_ShmemInitStruct_5))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v67 = (l1 + int32(127)) & int32(-128)
								v69 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[2]))
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
								v71 = v67 + v70
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
								if base.Ui32(v72) < base.Ui32(v71) {
									v75 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[4]))
									v76 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v75))), uint32(v76))
									v159 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[1]))
									v162 = F_hash_search(m, v159, l0, int32(2), int32(0))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
										F_LWLockRelease(m, v165+int32(128))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(_a_F_ShmemInitStruct_0))
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
													F_errmsg(m, int32(_a_F_ShmemInitStruct_6), v12+int32(32))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_ShmemInitStruct_2), int32(472), int32(_a_F_ShmemInitStruct_3))
														mBase = m.M
														v188 = m.ExcPending
														if v188 != 0 {
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
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v71
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[4]))
									v84 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83))), uint32(v84))
									if v80 == v84 {
										v159 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[1]))
										v162 = F_hash_search(m, v159, l0, int32(2), int32(0))
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return int32(0)
										} else {
											v165 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
											F_LWLockRelease(m, v165+int32(128))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(_a_F_ShmemInitStruct_0))
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
														F_errmsg(m, int32(_a_F_ShmemInitStruct_6), v12+int32(32))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_ShmemInitStruct_2), int32(472), int32(_a_F_ShmemInitStruct_3))
															mBase = m.M
															v188 = m.ExcPending
															if v188 != 0 {
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
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v67
										*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = l1
										v91 = v80 + v70
										*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v91
										v94 = v91
										v102 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
										F_LWLockRelease(m, v102+int32(128))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(48)
											return v94
										}
									}
								}
							}
						} else {
							v67 = (l1 + int32(127)) & int32(-128)
							v69 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[2]))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
							v71 = v67 + v70
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
							if base.Ui32(v72) < base.Ui32(v71) {
								v75 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[4]))
								v76 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v75))), uint32(v76))
								v159 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[1]))
								v162 = F_hash_search(m, v159, l0, int32(2), int32(0))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									v165 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
									F_LWLockRelease(m, v165+int32(128))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(_a_F_ShmemInitStruct_0))
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
												F_errmsg(m, int32(_a_F_ShmemInitStruct_6), v12+int32(32))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_ShmemInitStruct_2), int32(472), int32(_a_F_ShmemInitStruct_3))
													mBase = m.M
													v188 = m.ExcPending
													if v188 != 0 {
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
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v71
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[4]))
								v84 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83))), uint32(v84))
								if v80 == v84 {
									v159 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[1]))
									v162 = F_hash_search(m, v159, l0, int32(2), int32(0))
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
										F_LWLockRelease(m, v165+int32(128))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(_a_F_ShmemInitStruct_0))
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
													F_errmsg(m, int32(_a_F_ShmemInitStruct_6), v12+int32(32))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_ShmemInitStruct_2), int32(472), int32(_a_F_ShmemInitStruct_3))
														mBase = m.M
														v188 = m.ExcPending
														if v188 != 0 {
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
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v67
									*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = l1
									v91 = v80 + v70
									*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v91
									v94 = v91
									v102 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
									F_LWLockRelease(m, v102+int32(128))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(48)
										return v94
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
func F_process_shmem_requests(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v3)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_process_shmem_requests[1]))
	if v6 != 0 {
		m.T0[v6].(func(*base.Module))(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v10)
			return
		}
	} else {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v10)
		return
	}
}

package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_InitShmemAllocator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
	if v14 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[1])) = int32(2)
	} else {
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = (v20 + int32(159)) & int32(-128)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v24) <= base.Ui32(v25) {
		v28 = l0 + v20
		*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2])) = v28
		if v14 == int32(0) {
			v32 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v28)+4)), uint32(v32))
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = v24
			F_LWLockInitialize(m, v28+int32(16), int32(99))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v42 = v41
				*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[3])) = l0
				*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[4])) = l0
				*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[5])) = l0 + v42
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[6]))
				if v51 != 0 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
					v57 = base.I64_extend_i32_s(v52 + int32(128))
				} else {
					v57 = int64(128)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(274877906992)
				v60 = int32(1)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
				if v62 == v60 {
					v66 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
					v136 = v60
					v137 = v66
					v138 = v67
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
					v143 = m.G0
					v145 = v143 - int32(16)
					m.G0 = v145
					v148 = v9 + int32(-48)
					v149 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v149
					*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = int32(1208)
					if v136&int32(1) == v149 {
						*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v138
						*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v138 + v142
						*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v145 + int32(8)
						v167 = int32(_a_F_InitShmemAllocator_0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v138
						v167 = int32(_a_F_InitShmemAllocator_1)
					}
					v168 = F_hash_create(m, int32(_a_F_InitShmemAllocator_2), v57, v148, v167)
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return
					} else {
						m.G0 = v145 + int32(16)
						*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[7])) = v168
						v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
						if v176 == int32(0) {
							v183 = F_hash_search(m, v168, int32(_a_F_InitShmemAllocator_2), int32(1), v9+int32(-49))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								v186 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
								v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v187
								*(*int32)(unsafe.Add(mBase, uint32(v183)+52)) = v187
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
								v191 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v183)+60)) = uint8(v191)
								*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v190
								m.G0 = v11 - int32(-64)
								return
							}
						} else {
							m.G0 = v11 - int32(-64)
							return
						}
					}
				} else {
					v69 = F_hash_estimate_size(m, v57, int32(64))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = int32(_a_F_InitShmemAllocator_3)
						v72 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v69
						v74 = m.G0
						v76 = v74 - int32(16)
						m.G0 = v76
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
						v82 = base.AtomicRmwXchg32(m, v79, int32(4), int32(1))
						if v82 != 0 {
							F_s_lock(m, v79+int32(4), int32(_a_F_InitShmemAllocator_4))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								v94 = (v90 + int32(127)) & int32(-128)
								v95 = v94 + v69
								if base.Ui32(v94) <= base.Ui32(v95) {
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[4]))
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
									if base.Ui32(v95) <= base.Ui32(v99) {
										v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[3]))
										*(*int32)(unsafe.Add(mBase, uint32(v89))) = v95
										v107 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v107))
										if v105 != 0 {
											m.G0 = v76 + int32(16)
											v131 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
											v132 = v94 + v105
											*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v132
											v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
											v136 = v135
											v137 = v131
											v138 = v132
											v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
											v143 = m.G0
											v145 = v143 - int32(16)
											m.G0 = v145
											v148 = v9 + int32(-48)
											v149 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v149
											*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = int32(1208)
											if v136&int32(1) == v149 {
												*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v138
												*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v138 + v142
												*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v145 + int32(8)
												v167 = int32(_a_F_InitShmemAllocator_0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v138
												v167 = int32(_a_F_InitShmemAllocator_1)
											}
											v168 = F_hash_create(m, int32(_a_F_InitShmemAllocator_2), v57, v148, v167)
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return
											} else {
												m.G0 = v145 + int32(16)
												*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[7])) = v168
												v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
												if v176 == int32(0) {
													v183 = F_hash_search(m, v168, int32(_a_F_InitShmemAllocator_2), int32(1), v9+int32(-49))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return
													} else {
														v186 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
														v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v187
														*(*int32)(unsafe.Add(mBase, uint32(v183)+52)) = v187
														v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
														v191 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v183)+60)) = uint8(v191)
														*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v190
														m.G0 = v11 - int32(-64)
														return
													}
												} else {
													m.G0 = v11 - int32(-64)
													return
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return
											} else {
												F_errcode(m, int32(_a_F_InitShmemAllocator_5))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
													F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
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
										v101 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											F_errcode(m, int32(_a_F_InitShmemAllocator_5))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
												F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
									v101 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										F_errcode(m, int32(_a_F_InitShmemAllocator_5))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
											F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							v94 = (v90 + int32(127)) & int32(-128)
							v95 = v94 + v69
							if base.Ui32(v94) <= base.Ui32(v95) {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[4]))
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
								if base.Ui32(v95) <= base.Ui32(v99) {
									v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[3]))
									*(*int32)(unsafe.Add(mBase, uint32(v89))) = v95
									v107 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v107))
									if v105 != 0 {
										m.G0 = v76 + int32(16)
										v131 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
										v132 = v94 + v105
										*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v132
										v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
										v136 = v135
										v137 = v131
										v138 = v132
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
										v143 = m.G0
										v145 = v143 - int32(16)
										m.G0 = v145
										v148 = v9 + int32(-48)
										v149 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v149
										*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = int32(1208)
										if v136&int32(1) == v149 {
											*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v138
											*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v138 + v142
											*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v145 + int32(8)
											v167 = int32(_a_F_InitShmemAllocator_0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v138
											v167 = int32(_a_F_InitShmemAllocator_1)
										}
										v168 = F_hash_create(m, int32(_a_F_InitShmemAllocator_2), v57, v148, v167)
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return
										} else {
											m.G0 = v145 + int32(16)
											*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[7])) = v168
											v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
											if v176 == int32(0) {
												v183 = F_hash_search(m, v168, int32(_a_F_InitShmemAllocator_2), int32(1), v9+int32(-49))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													v186 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
													v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v187
													*(*int32)(unsafe.Add(mBase, uint32(v183)+52)) = v187
													v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
													v191 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v183)+60)) = uint8(v191)
													*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v190
													m.G0 = v11 - int32(-64)
													return
												}
											} else {
												m.G0 = v11 - int32(-64)
												return
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											F_errcode(m, int32(_a_F_InitShmemAllocator_5))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
												F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
									v101 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										F_errcode(m, int32(_a_F_InitShmemAllocator_5))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
											F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
								v101 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_InitShmemAllocator_5))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
										F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
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
			}
		} else {
			v42 = v25
			*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[3])) = l0
			*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[4])) = l0
			*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[5])) = l0 + v42
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[6]))
			if v51 != 0 {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
				v57 = base.I64_extend_i32_s(v52 + int32(128))
			} else {
				v57 = int64(128)
			}
			*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(274877906992)
			v60 = int32(1)
			v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
			if v62 == v60 {
				v66 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
				v136 = v60
				v137 = v66
				v138 = v67
				v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
				v143 = m.G0
				v145 = v143 - int32(16)
				m.G0 = v145
				v148 = v9 + int32(-48)
				v149 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v149
				*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = int32(1208)
				if v136&int32(1) == v149 {
					*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v138
					*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v138 + v142
					*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v145 + int32(8)
					v167 = int32(_a_F_InitShmemAllocator_0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v138
					v167 = int32(_a_F_InitShmemAllocator_1)
				}
				v168 = F_hash_create(m, int32(_a_F_InitShmemAllocator_2), v57, v148, v167)
				mBase = m.M
				v169 = m.ExcPending
				if v169 != 0 {
					return
				} else {
					m.G0 = v145 + int32(16)
					*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[7])) = v168
					v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
					if v176 == int32(0) {
						v183 = F_hash_search(m, v168, int32(_a_F_InitShmemAllocator_2), int32(1), v9+int32(-49))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return
						} else {
							v186 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
							v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v187
							*(*int32)(unsafe.Add(mBase, uint32(v183)+52)) = v187
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
							v191 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v183)+60)) = uint8(v191)
							*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v190
							m.G0 = v11 - int32(-64)
							return
						}
					} else {
						m.G0 = v11 - int32(-64)
						return
					}
				}
			} else {
				v69 = F_hash_estimate_size(m, v57, int32(64))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					v71 = int32(_a_F_InitShmemAllocator_3)
					v72 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v69
					v74 = m.G0
					v76 = v74 - int32(16)
					m.G0 = v76
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
					v82 = base.AtomicRmwXchg32(m, v79, int32(4), int32(1))
					if v82 != 0 {
						F_s_lock(m, v79+int32(4), int32(_a_F_InitShmemAllocator_4))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							v94 = (v90 + int32(127)) & int32(-128)
							v95 = v94 + v69
							if base.Ui32(v94) <= base.Ui32(v95) {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[4]))
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
								if base.Ui32(v95) <= base.Ui32(v99) {
									v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[3]))
									*(*int32)(unsafe.Add(mBase, uint32(v89))) = v95
									v107 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v107))
									if v105 != 0 {
										m.G0 = v76 + int32(16)
										v131 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
										v132 = v94 + v105
										*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v132
										v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
										v136 = v135
										v137 = v131
										v138 = v132
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
										v143 = m.G0
										v145 = v143 - int32(16)
										m.G0 = v145
										v148 = v9 + int32(-48)
										v149 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v149
										*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = int32(1208)
										if v136&int32(1) == v149 {
											*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v138
											*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v138 + v142
											*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v145 + int32(8)
											v167 = int32(_a_F_InitShmemAllocator_0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v138
											v167 = int32(_a_F_InitShmemAllocator_1)
										}
										v168 = F_hash_create(m, int32(_a_F_InitShmemAllocator_2), v57, v148, v167)
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return
										} else {
											m.G0 = v145 + int32(16)
											*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[7])) = v168
											v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
											if v176 == int32(0) {
												v183 = F_hash_search(m, v168, int32(_a_F_InitShmemAllocator_2), int32(1), v9+int32(-49))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													v186 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
													v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v187
													*(*int32)(unsafe.Add(mBase, uint32(v183)+52)) = v187
													v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
													v191 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v183)+60)) = uint8(v191)
													*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v190
													m.G0 = v11 - int32(-64)
													return
												}
											} else {
												m.G0 = v11 - int32(-64)
												return
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											F_errcode(m, int32(_a_F_InitShmemAllocator_5))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
												F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
									v101 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										F_errcode(m, int32(_a_F_InitShmemAllocator_5))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
											F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
								v101 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_InitShmemAllocator_5))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
										F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
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
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
						v94 = (v90 + int32(127)) & int32(-128)
						v95 = v94 + v69
						if base.Ui32(v94) <= base.Ui32(v95) {
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[4]))
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
							if base.Ui32(v95) <= base.Ui32(v99) {
								v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v89))) = v95
								v107 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v107))
								if v105 != 0 {
									m.G0 = v76 + int32(16)
									v131 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
									v132 = v94 + v105
									*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v132
									v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
									v136 = v135
									v137 = v131
									v138 = v132
									v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
									v143 = m.G0
									v145 = v143 - int32(16)
									m.G0 = v145
									v148 = v9 + int32(-48)
									v149 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v149
									*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = int32(1208)
									if v136&int32(1) == v149 {
										*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v138
										*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v138 + v142
										*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v145 + int32(8)
										v167 = int32(_a_F_InitShmemAllocator_0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v138
										v167 = int32(_a_F_InitShmemAllocator_1)
									}
									v168 = F_hash_create(m, int32(_a_F_InitShmemAllocator_2), v57, v148, v167)
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return
									} else {
										m.G0 = v145 + int32(16)
										*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[7])) = v168
										v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitShmemAllocator[0])))
										if v176 == int32(0) {
											v183 = F_hash_search(m, v168, int32(_a_F_InitShmemAllocator_2), int32(1), v9+int32(-49))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return
											} else {
												v186 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemAllocator[2]))
												v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v187
												*(*int32)(unsafe.Add(mBase, uint32(v183)+52)) = v187
												v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
												v191 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v183)+60)) = uint8(v191)
												*(*int32)(unsafe.Add(mBase, uint32(v183)+48)) = v190
												m.G0 = v11 - int32(-64)
												return
											}
										} else {
											m.G0 = v11 - int32(-64)
											return
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										F_errcode(m, int32(_a_F_InitShmemAllocator_5))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
											F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
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
								v101 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_InitShmemAllocator_5))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
										F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
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
							v101 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+4)), uint32(v101))
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_InitShmemAllocator_5))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v76))) = v69
									F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v76)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(821), int32(_a_F_InitShmemAllocator_8))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
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
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v203 = m.ExcPending
		if v203 != 0 {
			return
		} else {
			F_errcode(m, int32(_a_F_InitShmemAllocator_5))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v24
				F_errmsg(m, int32(_a_F_InitShmemAllocator_6), v11)
				mBase = m.M
				v210 = m.ExcPending
				if v210 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InitShmemAllocator_7), int32(723), int32(_a_F_InitShmemAllocator_9))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
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
func F_RegisterShmemCallbacks(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[0]))
	switch v13 {
	case 0:
		goto L2
	default:
		goto L3
	case 5:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L104
	} else {
		goto L110
	}
L2:
	;
	v492 = int32(_a_F_RegisterShmemCallbacks_0)
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[1]))
	v495 = F_lappend(m, v494, l0)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L104
	} else {
		goto L109
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L104
	} else {
		goto L106
	}
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v14&int32(1) == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = m.G0
	v21 = v19 - int32(192)
	m.G0 = v21
	v25 = int32(-1)
	v30 = v2
	v31 = v2
	v32 = v2
	goto L7
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[2])) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[3])) = v51
	m.G0 = v21 + int32(192)
	return
L7:
	;
	if v25 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[2]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[3]))
	v43 = v21 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v21 + int32(28)
	goto L12
L10:
	;
	v49 = v30
	v50 = v31
	v51 = v32
	goto L11
L11:
	;
	goto L14
L12:
	;
	v49 = int32(0)
	v50 = v39
	v51 = v41
	goto L11
L13:
	;
	goto L8
L14:
	;
	if v49 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	v429 = int32(m.ExcTag)
	v430 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v429 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[0])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[3])) = v21 + int32(32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[3])) = v51
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[2])) = v50
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4]))
	if v353 != 0 {
		goto L82
	} else {
		goto L83
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v61].(func(*base.Module, int32))(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[0])) = int32(4)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4]))
	if v69 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[2])) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[3])) = v51
	F_list_free_deep(m, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[5]))
	v90 = F_LWLockAcquire(m, v86+int32(16), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[0])) = int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4])) = int32(0)
	goto L6
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4]))
	if v93 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4]))
	if v289 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L30:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v264].(func(*base.Module, int32))(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L16
	} else {
		goto L74
	}
L31:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v260 == int32(0) {
		goto L29
	} else {
		goto L73
	}
L32:
	;
	v96 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v99 <= v96 {
		v193 = v96
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4]))
	if v203 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L34:
	;
	v103 = v96
	v104 = v96
	v111 = v96
	goto L36
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L16
	} else {
		goto L55
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[6]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v103<<(uint(int32(2))%32))))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v122 = int32(0)
	v124 = F_hash_search(m, v114, v121, v122, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L39
	}
L37:
	;
	if v150&v151 == int32(0) {
		v193 = v150
		goto L33
	} else {
		goto L51
	}
L38:
	;
	v153 = v103 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v153 < v154 {
		v103 = v153
		v104 = v150
		v111 = v151
		goto L36
	} else {
		goto L50
	}
L39:
	;
	if v124 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+60)))
	if v126 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v146 == int32(-1) {
		goto L35
	} else {
		goto L49
	}
L43:
	;
	v150 = int32(1)
	v151 = v111
	goto L38
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v133
	F_errmsg(m, int32(_a_F_RegisterShmemCallbacks_1), v21+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_RegisterShmemCallbacks_2), int32(1034), int32(_a_F_RegisterShmemCallbacks_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	goto L13
L49:
	;
	v150 = v104
	v151 = int32(1)
	goto L38
L50:
	;
	goto L37
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	F_errmsg_internal(m, int32(_a_F_RegisterShmemCallbacks_4), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_RegisterShmemCallbacks_2), int32(1048), int32(_a_F_RegisterShmemCallbacks_3))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	goto L13
L55:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v177
	F_errmsg(m, int32(_a_F_RegisterShmemCallbacks_5), v21)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v184 = F_errdetail(m, int32(_a_F_RegisterShmemCallbacks_6), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_RegisterShmemCallbacks_2), int32(1043), int32(_a_F_RegisterShmemCallbacks_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	goto L13
L59:
	;
	if v193 == int32(0) {
		goto L31
	} else {
		goto L71
	}
L60:
	;
	v206 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if v207 <= v206 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v211 = v206
	goto L62
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221+v211<<(uint(int32(2))%32))))
	if v193 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L59
L64:
	;
	v232 = v211 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if v232 < v233 {
		v211 = v232
		goto L62
	} else {
		goto L70
	}
L65:
	;
	v227 = F_AttachShmemIndexEntry(m, v225, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L16
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_InitShmemIndexEntry(m, v225)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L16
	} else {
		goto L69
	}
L68:
	;
	goto L64
L69:
	;
	goto L64
L70:
	;
	goto L63
L71:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v248 != 0 {
		v264 = v248
		goto L30
	} else {
		goto L72
	}
L72:
	;
	goto L29
L73:
	;
	v264 = v260
	goto L30
L74:
	;
	goto L29
L75:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[5]))
	F_LWLockRelease(m, v331+int32(16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L16
	} else {
		goto L81
	}
L76:
	;
	v292 = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v293 <= v292 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v297 = v292
	goto L78
L78:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307+v297<<(uint(int32(2))%32))))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+8))
	v313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v312)+60)) = uint8(v313)
	v316 = v297 + v313
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v316 < v317 {
		v297 = v316
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L75
L80:
	;
	goto L79
L81:
	;
	goto L19
L82:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if int32(0) < v354 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v405 = int32(0)
	goto L84
L84:
	;
	F_list_free_deep(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L16
	} else {
		goto L92
	}
L85:
	;
	v358 = int32(0)
	goto L88
L86:
	;
	goto L87
L87:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4]))
	v405 = v392
	goto L84
L88:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v368+v358<<(uint(int32(2))%32))))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	F_pfree(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L16
	} else {
		goto L90
	}
L89:
	;
	goto L87
L90:
	;
	v377 = v358 + int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v377 < v378 {
		v358 = v377
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[0])) = int32(5)
	v412 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[4])) = v412
	if v49 == v412 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	goto L15
L95:
	;
	v434 = int32(v430)
	m.G0 = v21
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if v21+int32(28) == v440 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	m.ExcPending = 1
	goto L104
L97:
	;
	if v444 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v444 = v442
	goto L100
L99:
	;
	v444 = int32(0)
	goto L100
L100:
	;
	goto L97
L101:
	;
	F___wasm_longjmp(m, v437, v436)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v25 = v444
	v30 = v436
	v31 = v50
	v32 = v51
	goto L7
L104:
	;
	return
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errmsg_internal(m, int32(_a_F_RegisterShmemCallbacks_7), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_RegisterShmemCallbacks_2), int32(952), int32(_a_F_RegisterShmemCallbacks_8))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterShmemCallbacks[1])) = v495
	return
L110:
	;
	F_errmsg_internal(m, int32(_a_F_RegisterShmemCallbacks_7), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_RegisterShmemCallbacks_2), int32(940), int32(_a_F_RegisterShmemCallbacks_8))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L104
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}

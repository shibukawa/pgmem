package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ReadDirExtended(m, l0, l1, int32(21))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ReadNextFullTransactionId(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextFullTransactionId[0]))
	v7 = F_LWLockAcquire(m, v3+int32(384), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextFullTransactionId[1]))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextFullTransactionId[0]))
		F_LWLockRelease(m, v15+int32(384))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			return v13
		}
	}
}
func F_RegisterCatcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v1 = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[1]))
	if v13 <= v11 {
		if v10 == int32(0) {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[2]))
			v21 = F_MemoryContextAlloc(m, v19, int32(512))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v29 = int32(32)
				v30 = v21
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[1])) = v29
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[0])) = v30
				v35 = v30
				v39 = v35 + v11<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43 + int32(1)
				return
			}
		} else {
			v27 = F_repalloc(m, v10, v13<<(uint(int32(5))%32))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = v13 << (uint(int32(1)) % 32)
				v30 = v27
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[1])) = v29
				*(*int32)(unsafe.Add(mBase, _c_F_RegisterCatcacheInvalidation[0])) = v30
				v35 = v30
				v39 = v35 + v11<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43 + int32(1)
				return
			}
		}
	} else {
		v35 = v10
		v39 = v35 + v11<<(uint(int32(4))%32)
		*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v1)
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v43 + int32(1)
		return
	}
}
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
	v4 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v3, v4, l0, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
		F_ResourceOwnerReleaseInternal(m, v9, int32(2), l0, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
			F_ResourceOwnerReleaseInternal(m, v15, int32(3), l0, int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResources[0]))
				v22 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)) = uint16(v22)
				return
			}
		}
	}
}
func F_RepairGraphElement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
	v24 = F_add_size(m, v22, int32(2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v26 = F_mul_size(m, v24, v19)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = F_mul_size(m, int32(6), v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = F_add_size(m, int32(4), v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					if l2 == int32(0) {
						v40 = int32(0)
						F_HnswInitNeighbors(m, v40, l1, v19, v40)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v44)
							v48 = l0 + int32(24)
							F_HnswFindElementNeighbors(m, v44, l1, l2, v18, v48, v19, v17, int32(1))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v52 = int32(0)
								base.MemoryFill(m, v15, v52, int32(_a_F_RepairGraphElement_0))
								v67 = int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v67)
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
								v79 = v71
								v83 = v52
								for {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v79<<(uint(int32(2))%32))))
									if base.B2i32(v19 <= v52) == int32(0) {
										v112 = int32(0)
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
										v121 = v112
										v126 = v83
										for {
											v134 = v15 + int32(4) + v126*int32(6)
											if v121 < v115 {
												v141 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(8)+v121*int32(12))))
												v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+80)))
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
												v150 = int32(base.Ui32(v148) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v150)
												v156 = v147
												v157 = v148
											} else {
												v152 = int32(_a_F_RepairGraphElement_1)
												*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v152)
												v156 = int32(0)
												v157 = v152
											}
											v158 = int32(1)
											v159 = v126 + v158
											*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v156)
											*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)) = uint16(v157)
											v163 = v121 + v158
											if v163 != v19<<(uint(base.B2i32(v79 == v112))%32) {
												v121 = v163
												v126 = v159
												continue
											} else {
												break
											}
											break
										}
										v174 = v159
									} else {
										v174 = v83
									}
									if int32(0) < v79 {
										v79 = v79 - int32(1)
										v83 = v174
										continue
									} else {
										break
									}
									break
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v174)
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v185)
								v187 = int32(0)
								v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
								v190 = F_ReadBufferExtended(m, v18, v187, v188, v187, v16)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return
								} else {
									F_LockBufferInternal(m, v190, int32(3))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return
									} else {
										v195 = F_GenericXLogStart(m, v18)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return
										} else {
											v198 = F_GenericXLogRegisterBuffer(m, v195, v190, int32(0))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return
											} else {
												v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
												v205 = F_PageIndexTupleOverwrite(m, v198, v200, v15, (v30+int32(7))&int32(-8))
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
													return
												} else {
													if v205 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return
														} else {
															v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227 + int32(4)
															F_errmsg_internal(m, int32(_a_F_RepairGraphElement_2), v13)
															mBase = m.M
															v233 = m.ExcPending
															if v233 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_RepairGraphElement_3), int32(266), int32(_a_F_RepairGraphElement_4))
																mBase = m.M
																v238 = m.ExcPending
																if v238 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														F_GenericXLogFinish(m, v195)
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return
														} else {
															F_UnlockReleaseBuffer(m, v190)
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																F_HnswUpdateNeighborsOnDisk(m, v18, v48, l1, v19, int32(1), int32(0))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(16)
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
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
						if v34 != v35 {
							v40 = int32(0)
							F_HnswInitNeighbors(m, v40, l1, v19, v40)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								v44 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v44)
								v48 = l0 + int32(24)
								F_HnswFindElementNeighbors(m, v44, l1, l2, v18, v48, v19, v17, int32(1))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									v52 = int32(0)
									base.MemoryFill(m, v15, v52, int32(_a_F_RepairGraphElement_0))
									v67 = int32(2)
									*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v67)
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
									v79 = v71
									v83 = v52
									for {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v79<<(uint(int32(2))%32))))
										if base.B2i32(v19 <= v52) == int32(0) {
											v112 = int32(0)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
											v121 = v112
											v126 = v83
											for {
												v134 = v15 + int32(4) + v126*int32(6)
												if v121 < v115 {
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(8)+v121*int32(12))))
													v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+80)))
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
													v150 = int32(base.Ui32(v148) >> (uint(int32(16)) % 32))
													*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v150)
													v156 = v147
													v157 = v148
												} else {
													v152 = int32(_a_F_RepairGraphElement_1)
													*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v152)
													v156 = int32(0)
													v157 = v152
												}
												v158 = int32(1)
												v159 = v126 + v158
												*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v156)
												*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)) = uint16(v157)
												v163 = v121 + v158
												if v163 != v19<<(uint(base.B2i32(v79 == v112))%32) {
													v121 = v163
													v126 = v159
													continue
												} else {
													break
												}
												break
											}
											v174 = v159
										} else {
											v174 = v83
										}
										if int32(0) < v79 {
											v79 = v79 - int32(1)
											v83 = v174
											continue
										} else {
											break
										}
										break
									}
									*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v174)
									v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v185)
									v187 = int32(0)
									v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
									v190 = F_ReadBufferExtended(m, v18, v187, v188, v187, v16)
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return
									} else {
										F_LockBufferInternal(m, v190, int32(3))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return
										} else {
											v195 = F_GenericXLogStart(m, v18)
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return
											} else {
												v198 = F_GenericXLogRegisterBuffer(m, v195, v190, int32(0))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return
												} else {
													v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
													v205 = F_PageIndexTupleOverwrite(m, v198, v200, v15, (v30+int32(7))&int32(-8))
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
														return
													} else {
														if v205 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return
															} else {
																v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227 + int32(4)
																F_errmsg_internal(m, int32(_a_F_RepairGraphElement_2), v13)
																mBase = m.M
																v233 = m.ExcPending
																if v233 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_RepairGraphElement_3), int32(266), int32(_a_F_RepairGraphElement_4))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															F_GenericXLogFinish(m, v195)
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return
															} else {
																F_UnlockReleaseBuffer(m, v190)
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	F_HnswUpdateNeighborsOnDisk(m, v18, v48, l1, v19, int32(1), int32(0))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(16)
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
							v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+80)))
							v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
							if v37 == v38 {
								m.G0 = v13 + int32(16)
								return
							} else {
								v40 = int32(0)
								F_HnswInitNeighbors(m, v40, l1, v19, v40)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									v44 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v44)
									v48 = l0 + int32(24)
									F_HnswFindElementNeighbors(m, v44, l1, l2, v18, v48, v19, v17, int32(1))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v52 = int32(0)
										base.MemoryFill(m, v15, v52, int32(_a_F_RepairGraphElement_0))
										v67 = int32(2)
										*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v67)
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
										v79 = v71
										v83 = v52
										for {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v79<<(uint(int32(2))%32))))
											if base.B2i32(v19 <= v52) == int32(0) {
												v112 = int32(0)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
												v121 = v112
												v126 = v83
												for {
													v134 = v15 + int32(4) + v126*int32(6)
													if v121 < v115 {
														v141 = *(*int32)(unsafe.Add(mBase, uint32(v102+int32(8)+v121*int32(12))))
														v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+80)))
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
														v150 = int32(base.Ui32(v148) >> (uint(int32(16)) % 32))
														*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v150)
														v156 = v147
														v157 = v148
													} else {
														v152 = int32(_a_F_RepairGraphElement_1)
														*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v152)
														v156 = int32(0)
														v157 = v152
													}
													v158 = int32(1)
													v159 = v126 + v158
													*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v156)
													*(*uint16)(unsafe.Add(mBase, uint32(v134)+2)) = uint16(v157)
													v163 = v121 + v158
													if v163 != v19<<(uint(base.B2i32(v79 == v112))%32) {
														v121 = v163
														v126 = v159
														continue
													} else {
														break
													}
													break
												}
												v174 = v159
											} else {
												v174 = v83
											}
											if int32(0) < v79 {
												v79 = v79 - int32(1)
												v83 = v174
												continue
											} else {
												break
											}
											break
										}
										*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)) = uint16(v174)
										v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+67)))
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v185)
										v187 = int32(0)
										v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
										v190 = F_ReadBufferExtended(m, v18, v187, v188, v187, v16)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return
										} else {
											F_LockBufferInternal(m, v190, int32(3))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return
											} else {
												v195 = F_GenericXLogStart(m, v18)
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return
												} else {
													v198 = F_GenericXLogRegisterBuffer(m, v195, v190, int32(0))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return
													} else {
														v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+82)))
														v205 = F_PageIndexTupleOverwrite(m, v198, v200, v15, (v30+int32(7))&int32(-8))
														mBase = m.M
														v206 = m.ExcPending
														if v206 != 0 {
															return
														} else {
															if v205 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
																	return
																} else {
																	v227 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v227 + int32(4)
																	F_errmsg_internal(m, int32(_a_F_RepairGraphElement_2), v13)
																	mBase = m.M
																	v233 = m.ExcPending
																	if v233 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_RepairGraphElement_3), int32(266), int32(_a_F_RepairGraphElement_4))
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																F_GenericXLogFinish(m, v195)
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
																	return
																} else {
																	F_UnlockReleaseBuffer(m, v190)
																	mBase = m.M
																	v212 = m.ExcPending
																	if v212 != 0 {
																		return
																	} else {
																		F_HnswUpdateNeighborsOnDisk(m, v18, v48, l1, v19, int32(1), int32(0))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(16)
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
						}
					}
				}
			}
		}
	}
}
func F_ReplaceVarsFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
	v23 = F_replace_rte_variables(m, l0, l1, int32(0), int32(1136), v12+int32(12), l7)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		m.G0 = v12 + int32(32)
		return v23
	}
}
func F_ReplaceVarsFromTargetList_callback(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v14 = F_ReplaceVarFromTargetList(m, l0, v9, v10, v11, v12, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v18
			v26 = F_query_or_expression_tree_walker_impl(m, v14, int32(1128), v6+int32(8), int32(16))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v14
			}
		} else {
			m.G0 = v6 + int32(16)
			return v14
		}
	}
}
func F_RequestCheckpoint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RequestCheckpoint[0])))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = F_CreateCheckPoint(m, l0|int32(4))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	v21 = base.AtomicRmwXchg32(m, v18, int32(4), int32(1))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return
L6:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	F_s_lock(m, v18+int32(4), int32(_a_F_RequestCheckpoint_0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v27 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = l0 | v30 | int32(64)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+4)), uint32(v27))
	v41 = l0 & int32(32)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[2]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	if v44 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	if v41 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L13:
	;
	v142 = F_errstart(m, v136, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L44
	}
L14:
	;
	if v41 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v72 = v43
	v75 = v44
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v83 = v78 + v75*int32(768) + int32(316)
	v84 = int32(0)
	v87 = base.AtomicRmwOr32(m, v84, int32(_a_F_RequestCheckpoint_1), v84)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v88 != 0 {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v136 = int32(15)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v51 = v27
	goto L20
L20:
	;
	if v51 == int32(600) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v72 = v68
	v75 = v69
	goto L16
L22:
	;
	v136 = int32(21)
	goto L13
L23:
	;
	goto L24
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[3]))
	if v60 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_pg_usleep(m, int32(_a_F_RequestCheckpoint_2))
	mBase = m.M
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[2]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
	if v69 == int32(-1) {
		v51 = v51 + int32(1)
		goto L20
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	goto L21
L30:
	;
	goto L12
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(1)
	v91 = int32(0)
	v94 = base.AtomicRmwOr32(m, v91, int32(_a_F_RequestCheckpoint_1), v91)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v95 == v91 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v98 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[4]))
	if v102 == v98 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = m.G0
	v106 = v104 - int32(16)
	m.G0 = v106
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[5]))
	if v109 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v132 = F_pgmem_kill(m, v98, int32(23))
	mBase = m.M
	goto L31
L38:
	;
	m.G0 = v106 + int32(16)
	goto L30
L39:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+15)) = uint8(v112)
	goto L40
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[6]))
	v120 = F_write(m, v116, v106+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v120 {
		goto L38
	} else {
		goto L42
	}
L41:
	;
	goto L38
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[7]))
	if v124 == int32(27) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v142 == int32(0) {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	F_errmsg_internal(m, int32(_a_F_RequestCheckpoint_3), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_RequestCheckpoint_4), int32(1132), int32(_a_F_RequestCheckpoint_5))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L12
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	F_ConditionVariablePrepareToSleep(m, v164+int32(24))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	goto L50
L50:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	v179 = base.AtomicRmwXchg32(m, v176, int32(4), int32(1))
	if v179 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L60
	}
L52:
	;
	F_s_lock(m, v176+int32(4), int32(_a_F_RequestCheckpoint_0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v188 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v186)+4)), uint32(v188))
	if v187 == v35 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_ConditionVariableSleep(m, v186+int32(24), int32(134217740))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L51
L59:
	;
	goto L50
L60:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	F_ConditionVariablePrepareToSleep(m, v200+int32(36))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	goto L62
L62:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	v215 = base.AtomicRmwXchg32(m, v212, int32(4), int32(1))
	if v215 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L72
	}
L64:
	;
	F_s_lock(m, v212+int32(4), int32(_a_F_RequestCheckpoint_0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_RequestCheckpoint[1]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v225 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v222)+4)), uint32(v225))
	if v224-v187 < v225 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	F_ConditionVariableSleep(m, v222+int32(36), int32(134217739))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L63
L71:
	;
	goto L62
L72:
	;
	if v223 == v36 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_RequestCheckpoint_6), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	F_errhint(m, int32(_a_F_RequestCheckpoint_7), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_RequestCheckpoint_4), int32(1196), int32(_a_F_RequestCheckpoint_5))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReserveExternalFD(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[1]))
	if v6 <= int32(0) {
		v36 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[0])) = v36 + int32(1)
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[2]))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[3]))
	if v12+v6+v4 < v10 {
		v36 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[4]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	F_LruDelete(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v36 = v24
	goto L1
L6:
	;
	return
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[0]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[1]))
	if v26 <= int32(0) {
		v36 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[2]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ReserveExternalFD[3]))
	if v30 <= v32+v26+v24 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_r_C_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = int32(2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5-v12 < v7 {
		v22 = v2
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = F_memcmp(m, v15+v5-v7, int32(_a_F_r_C_2_0), v7)
		mBase = m.M
		if v18 != 0 {
			v22 = v2
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v5 - v7
			v22 = int32(1)
		}
	}
	if v22 != 0 {
		v162 = v2
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v24 = v6 - v5
		v25 = v23 - v24
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v25 <= v40 {
			v149 = int32(-1)
			v156 = v149
		} else {
			v57 = int32(1)
			v58 = v25 - v57
			v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41+v58))))
			v62 = v60 & int32(255)
			if base.B2i32(v58 == v40)|base.B2i32(int32(0) <= v60) != 0 {
				v120 = v62
				v124 = v57
			} else {
				v69 = v62 & int32(63)
				v71 = v25 - int32(2)
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v71))))
				v75 = v73 << (uint(int32(6)) % 32)
				if base.B2i32(v71 != v40)&base.B2i32(base.Ui32(v73) < base.Ui32(int32(192))) == int32(0) {
					v120 = v75&int32(1984) | v69
					v124 = int32(2)
				} else {
					v88 = v75&int32(4032) | v69
					v90 = v25 - int32(3)
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v90))))
					if base.B2i32(v90 != v40)&base.B2i32(base.Ui32(v92) < base.Ui32(int32(224))) == int32(0) {
						v120 = v92<<(uint(int32(12))%32)&int32(_a_F_r_C_2_1) | v88
						v124 = int32(3)
					} else {
						v110 = int32(4)
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v41-v110))))
						v120 = v92<<(uint(int32(12))%32)&int32(_a_F_r_C_2_2) | v112&int32(7)<<(uint(int32(18))%32) | v88
						v124 = v110
					}
				}
			}
			if int32(252) < v120 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - v124
				v149 = int32(0)
				v156 = v149
			} else {
				v126 = v120 - int32(97)
				if v126 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - v124
					v149 = int32(0)
					v156 = v149
				} else {
					v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v126)>>(uint(int32(3))%32)))+uint32(_c_F_r_C_2[0]))))
					if int32(base.Ui32(v132)>>(uint(v126&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - v124
						v149 = int32(0)
						v156 = v149
					} else {
						v156 = v124
					}
				}
			}
		}
		if v156 != 0 {
			v162 = v2
		} else {
			v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v157 - v24
			v162 = int32(1)
		}
	}
	return v162
}
func F_r_Step_5b(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v9 {
		v43 = v2
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v7-int32(1)))))
		if v15 != int32(108) {
			v43 = v2
		} else {
			v19 = v7 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if base.B2i32(v19 <= v9)|base.B2i32(v7 <= v23) != 0 {
				v43 = v2
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v11-int32(1)))))
				if v29 != int32(108) {
					v43 = v2
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - int32(2)
					v36 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v36 {
						v39 = int32(1)
					} else {
						v39 = v36
					}
					v43 = v39
				}
			}
		}
	}
	return v43
}
func F_r_V_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = v4 - v5
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 <= v16 {
		v59 = int32(-1)
	} else {
		v28 = int32(1)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v5-v28))))
		if int32(252) < v33 {
			v55 = v28
		} else {
			v35 = v33 - int32(97)
			if v35 < int32(0) {
				v55 = v28
			} else {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_c_F_r_V_1[0]))))
				if int32(base.Ui32(v41)>>(uint(v35&int32(7))%32))&int32(1) == int32(0) {
					v55 = v28
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v5 - int32(1)
					v55 = int32(0)
				}
			}
		}
		v59 = v55
	}
	if v59 != 0 {
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v61 = v60 - v6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
		v63 = int32(2)
		v65 = int32(0)
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v61-v68 < v63 {
			v78 = v65
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v74 = F_memcmp(m, v71+v61-v63, int32(_a_F_r_V_1_0), v63)
			mBase = m.M
			if v74 != 0 {
				v78 = v65
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 - v63
				v78 = int32(1)
			}
		}
		if v78 == int32(0) {
			v85 = int32(0)
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 - v6
			v85 = int32(1)
		}
	} else {
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 - v6
		v85 = int32(1)
	}
	return v85
}
func F_r_e_ending_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v2)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v9 {
		v170 = v2
		return v170
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v7-int32(1)))))
		if v15 != int32(101) {
			v170 = v2
			return v170
		} else {
			v19 = v7 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v7 <= v22 {
				v170 = v2
				return v170
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v37 <= v38 {
					v147 = int32(-1)
					v154 = v147
				} else {
					v55 = int32(1)
					v56 = v37 - v55
					v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39+v56))))
					v60 = v58 & int32(255)
					if base.B2i32(v56 == v38)|base.B2i32(int32(0) <= v58) != 0 {
						v118 = v60
						v122 = v55
					} else {
						v67 = v60 & int32(63)
						v69 = v37 - int32(2)
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v69))))
						v73 = v71 << (uint(int32(6)) % 32)
						if base.B2i32(v69 != v38)&base.B2i32(base.Ui32(v71) < base.Ui32(int32(192))) == int32(0) {
							v118 = v73&int32(1984) | v67
							v122 = int32(2)
						} else {
							v86 = v73&int32(4032) | v67
							v88 = v37 - int32(3)
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v88))))
							if base.B2i32(v88 != v38)&base.B2i32(base.Ui32(v90) < base.Ui32(int32(224))) == int32(0) {
								v118 = v90<<(uint(int32(12))%32)&int32(_a_F_r_e_ending_2_0) | v86
								v122 = int32(3)
							} else {
								v108 = int32(4)
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v39-v108))))
								v118 = v90<<(uint(int32(12))%32)&int32(_a_F_r_e_ending_2_1) | v110&int32(7)<<(uint(int32(18))%32) | v86
								v122 = v108
							}
						}
					}
					if int32(232) < v118 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37 - v122
						v147 = int32(0)
						v154 = v147
					} else {
						v124 = v118 - int32(97)
						if v124 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37 - v122
							v147 = int32(0)
							v154 = v147
						} else {
							v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v124)>>(uint(int32(3))%32)))+uint32(_c_F_r_e_ending_2[0]))))
							if int32(base.Ui32(v130)>>(uint(v124&int32(7))%32))&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37 - v122
								v147 = int32(0)
								v154 = v147
							} else {
								v154 = v122
							}
						}
					}
				}
				if v154 != 0 {
					v170 = v2
					return v170
				} else {
					v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155 + (v19 - v24)
					v159 = F_slice_del(m, l0)
					mBase = m.M
					if v159 < int32(0) {
						v170 = v159
						return v170
					} else {
						v162 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v162)
						v164 = F_r_undouble_3(m, l0)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v170 = v164
							return v170
						}
					}
				}
			}
		}
	}
}
func F_r_i_plural(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v6 <= v5 {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v5
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
		if v6 < v5 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v5-int32(1)))))
			if base.Ui32((v16-int32(105))&int32(255)) < base.Ui32(int32(2)) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
				v27 = int32(1)
				v28 = v5 - v27
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
				v32 = F_slice_del(m, l0)
				mBase = m.M
				if int32(0) <= v32 {
					v38 = v27
				} else {
					v38 = v32 >> (uint(int32(31)) % 32) & v32
				}
				v43 = v38
				return v43
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
				return int32(0)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
			return int32(0)
		}
	} else {
		v43 = int32(0)
		return v43
	}
}
func F_r_remove_suffix_1(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14366(m, l0, int32(_a_F_r_remove_suffix_1_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_r_stem_suffix_chain_before_ki(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = int32(2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6-v13 < v8 {
		v23 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v552
L2:
	;
	if v23 <= int32(0) {
		v552 = v23
		goto L1
	} else {
		goto L6
	}
L3:
	;
	goto L2
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = F_memcmp(m, v16+v6-v8, int32(_a_F_r_stem_suffix_chain_before_ki_0), v8)
	mBase = m.M
	if v19 != 0 {
		v23 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v8
	v23 = int32(1)
	goto L3
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v552 = int32(1)
	goto L1
L8:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v141 = v27 - v26
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v140 - v141
	v144 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v144 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v31 - int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v33 <= v34 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v33))))
	switch v38 - int32(97) {
	case 0, 4:
		goto L11
	default:
		goto L8
	}
L11:
	;
	v44 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_6), int32(4), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v44 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v50
	v52 = F_slice_del(m, l0)
	mBase = m.M
	if v52 < int32(0) {
		v552 = v52
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v58 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v98 = v55 - v57
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98 + v99
	v102 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L28
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v62-int32(2) <= v61 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v62-int32(1)))))
	if v70 != int32(114) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v76 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_7), int32(2), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v76 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v80
	v82 = F_slice_del(m, l0)
	mBase = m.M
	if v82 < int32(0) {
		v552 = v82
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	if v87 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91 + (v85 - v86)
	goto L7
L25:
	;
	goto L26
L26:
	;
	if int32(0) <= v87 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v552 = v87
	goto L1
L28:
	;
	if v102 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106 + v98
	goto L7
L30:
	;
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v109
	v111 = F_slice_del(m, l0)
	mBase = m.M
	if v111 < int32(0) {
		v552 = v111
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v117 = v116 - v114
	v118 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	if v118 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v122 - v117
	goto L7
L35:
	;
	goto L36
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v125
	v127 = F_slice_del(m, l0)
	mBase = m.M
	if v127 < int32(0) {
		v552 = v127
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v130 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	if v130 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134 - v117
	goto L7
L40:
	;
	goto L41
L41:
	;
	if int32(0) <= v130 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v552 = v130
	goto L1
L43:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v542
	v544 = F_slice_del(m, l0)
	mBase = m.M
	if v544 < int32(0) {
		v552 = v544
		goto L1
	} else {
		goto L154
	}
L44:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v537
	v539 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v539 {
		goto L7
	} else {
		goto L153
	}
L45:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v351 - v141
	v354 = int32(0)
	v355 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v355 == v354 {
		v552 = v354
		goto L1
	} else {
		goto L104
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v149 = v147 - int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v149 <= v150 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v149))))
	if v154 != int32(110) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v160 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_5), int32(4), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	if v160 == int32(0) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v165 = Fn14364(m, l0, int32(110))
	mBase = m.M
	goto L51
L51:
	;
	if v165 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v168
	v170 = F_slice_del(m, l0)
	mBase = m.M
	if v170 < int32(0) {
		v552 = v170
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v173-int32(3) <= v176 {
		v196 = v175
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v197 = v175 - v173
	v198 = v196 - v197
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v198
	v201 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L63
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v173-int32(1)))))
	if v184 != int32(177) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v184 != int32(105) {
		v196 = v175
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v192 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_4), int32(2), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v192 != 0 {
		goto L44
	} else {
		goto L61
	}
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v196 = v194
	goto L54
L62:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v337 - v197
	v340 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L12
	} else {
		goto L99
	}
L63:
	;
	if v201 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v205 - v197
	v208 = int32(0)
	v214 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v214 == v208 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v307
	v309 = F_slice_del(m, l0)
	mBase = m.M
	if v309 < int32(0) {
		v552 = v309
		goto L1
	} else {
		goto L88
	}
L67:
	;
	if v304 == int32(0) {
		goto L62
	} else {
		goto L87
	}
L68:
	;
	v304 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v222 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_2), int32(105), int32(305), int32(0))
	mBase = m.M
	if v222 != 0 {
		v296 = v208
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v304 = v296
	goto L67
L72:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v225 <= v226 {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v285
	v296 = v290
	goto L71
L74:
	;
	v285 = v283
	v290 = int32(1)
	goto L73
L75:
	;
	v283 = v235 - v223 + v242
	goto L74
L76:
	;
	v250 = v225 - v223
	v251 = v247 + v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v251
	if v251 <= v248 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v247 = v223
	v248 = v226
	v249 = v224
	goto L76
L78:
	;
	goto L79
L79:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v225-int32(1)))))
	if v231 != int32(115) {
		v247 = v223
		v248 = v226
		v249 = v224
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v235 = v225 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235
	v240 = int32(0)
	v241 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), v240)
	mBase = m.M
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v241 == v240 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v247 = v242
	v248 = v246
	v249 = v245
	goto L76
L82:
	;
	v263 = int32(0)
	v265 = F_skip_b_utf8(m, v249, v251, v248, int32(1))
	mBase = m.M
	if v265 < v263 {
		v296 = v263
		goto L71
	} else {
		goto L85
	}
L83:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251+v249-int32(1)))))
	if v257 != int32(115) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v285 = v251 - int32(1)
	v290 = int32(0)
	goto L73
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v265
	v273 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), int32(0))
	mBase = m.M
	if v273 != 0 {
		v296 = v263
		goto L71
	} else {
		goto L86
	}
L86:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = v274 + v250
	goto L74
L87:
	;
	goto L66
L88:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v315 = v314 - v312
	v316 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L89
	}
L89:
	;
	if v316 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v320 - v315
	goto L7
L91:
	;
	goto L92
L92:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v323
	v325 = F_slice_del(m, l0)
	mBase = m.M
	if v325 < int32(0) {
		v552 = v325
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v328 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L12
	} else {
		goto L94
	}
L94:
	;
	if v328 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v332 - v315
	goto L7
L96:
	;
	goto L97
L97:
	;
	if int32(0) <= v328 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	v552 = v328
	goto L1
L99:
	;
	if v340 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344 + (v173 - v175)
	goto L7
L101:
	;
	goto L102
L102:
	;
	if int32(0) <= v340 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v552 = v340
	goto L1
L104:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v359-int32(2) <= v358 {
		v552 = v354
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v359-int32(1)))))
	switch v367 - int32(97) {
	case 0, 4:
		goto L106
	default:
		v552 = v354
		goto L1
	}
L106:
	;
	v373 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_1), int32(2), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	if v373 == int32(0) {
		v552 = v354
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v379-int32(3) <= v378 {
		v399 = v377
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v400 = v377 - v379
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v400
	v403 = int32(0)
	v409 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v409 == v403 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+v379-int32(1)))))
	if v387 != int32(177) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v387 != int32(105) {
		v399 = v377
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v395 = F_find_among_b(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_4), int32(2), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L12
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	if v395 != 0 {
		goto L43
	} else {
		goto L116
	}
L116:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v399 = v397
	goto L109
L117:
	;
	if v499 != 0 {
		goto L137
	} else {
		goto L138
	}
L118:
	;
	v499 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v417 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_2), int32(105), int32(305), int32(0))
	mBase = m.M
	if v417 != 0 {
		v491 = v403
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v499 = v491
	goto L117
L122:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v420 <= v421 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480
	v491 = v485
	goto L121
L124:
	;
	v480 = v478
	v485 = int32(1)
	goto L123
L125:
	;
	v478 = v430 - v418 + v437
	goto L124
L126:
	;
	v445 = v420 - v418
	v446 = v442 + v445
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v446
	if v446 <= v443 {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v442 = v418
	v443 = v421
	v444 = v419
	goto L126
L128:
	;
	goto L129
L129:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419+v420-int32(1)))))
	if v426 != int32(115) {
		v442 = v418
		v443 = v421
		v444 = v419
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v430 = v420 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v430
	v435 = int32(0)
	v436 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), v435)
	mBase = m.M
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v436 == v435 {
		goto L125
	} else {
		goto L131
	}
L131:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v442 = v437
	v443 = v441
	v444 = v440
	goto L126
L132:
	;
	v458 = int32(0)
	v460 = F_skip_b_utf8(m, v444, v446, v443, int32(1))
	mBase = m.M
	if v460 < v458 {
		v491 = v458
		goto L121
	} else {
		goto L135
	}
L133:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446+v444-int32(1)))))
	if v452 != int32(115) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v480 = v446 - int32(1)
	v485 = int32(0)
	goto L123
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460
	v468 = F_in_grouping_b_U(m, l0, int32(_a_F_r_stem_suffix_chain_before_ki_3), int32(97), int32(305), int32(0))
	mBase = m.M
	if v468 != 0 {
		v491 = v458
		goto L121
	} else {
		goto L136
	}
L136:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v478 = v469 + v445
	goto L124
L137:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v500
	v502 = F_slice_del(m, l0)
	mBase = m.M
	if v502 < int32(0) {
		v552 = v502
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v530 - v400
	v533 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L12
	} else {
		goto L151
	}
L140:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v508 = v507 - v505
	v509 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	if v509 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v513 - v508
	goto L7
L143:
	;
	goto L144
L144:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v516
	v518 = F_slice_del(m, l0)
	mBase = m.M
	if v518 < int32(0) {
		v552 = v518
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v521 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	if v521 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v525 - v508
	goto L7
L148:
	;
	goto L149
L149:
	;
	if int32(0) <= v521 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	v552 = v521
	goto L1
L151:
	;
	if v533 <= int32(0) {
		v552 = v533
		goto L1
	} else {
		goto L152
	}
L152:
	;
	goto L7
L153:
	;
	v552 = v539
	goto L1
L154:
	;
	goto L7
}
func F_r_undouble_3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = v5 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v108 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v108
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
	if base.B2i32(v12&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v12)%32)&int32(_a_F_r_undouble_3_0) == int32(0)) != 0 {
		v108 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = F_find_among_b(m, l0, int32(_a_F_r_undouble_3_1), int32(3), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v28 == int32(0) {
		v108 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = v34 + (v5 - v24)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L9
L7:
	;
	if v92 < int32(0) {
		v108 = v2
		goto L1
	} else {
		goto L26
	}
L9:
	;
	goto L10
L10:
	;
	goto L11
L11:
	;
	v47 = v36
	v49 = int32(1)
	goto L14
L13:
	;
	v92 = v74
	goto L7
L14:
	;
	if v47 <= v40 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v92 = int32(-1)
	goto L7
L17:
	;
	goto L18
L18:
	;
	v54 = v47 - int32(1)
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39+v54))))
	if base.B2i32(int32(0) <= v56)|base.B2i32(v54 <= v40) != 0 {
		v74 = v54
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = int32(1)
	if v78 < v49 {
		v47 = v74
		v49 = v49 - v78
		goto L14
	} else {
		goto L25
	}
L20:
	;
	v62 = v54
	goto L21
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v62))))
	if base.Ui32(int32(191)) < base.Ui32(v67) {
		v74 = v62
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v74 = v40
	goto L19
L23:
	;
	v71 = v62 - int32(1)
	if v40 < v71 {
		v62 = v71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L15
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
	v98 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v98 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v104 = int32(1)
	goto L29
L28:
	;
	v104 = v98 >> (uint(int32(31)) % 32) & v98
	goto L29
L29:
	;
	v108 = v104
	goto L1
}
func F_rangeTableEntry_used(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1129), v6+int32(8), v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_rboolop(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_rboolop_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_read_gucstate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v6) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v6
	goto L5
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L13
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
	return v6
L5:
	;
	v14 = v10 + int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v14 != l1 {
		v10 = v14
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
	F_errmsg_internal(m, int32(_a_F_read_gucstate_0), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_read_gucstate_1), int32(_a_F_read_gucstate_2), int32(_a_F_read_gucstate_3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_read_gucstate_4), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_read_gucstate_1), int32(_a_F_read_gucstate_5), int32(_a_F_read_gucstate_3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_readdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 <= v4 {
		v7 = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = m.Env.X__syscall_getdents64(m, v8, l0+int32(24), int32(2048))
		mBase = m.M
		if v12 <= v7 {
			if base.B2i32(v12 == int32(0))|base.B2i32(v12 == int32(-44)) != 0 {
				v38 = v7
				return v38
			} else {
				v21 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_readdir[0])) = v21 - v12
				return v21
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v12
			v28 = v7
			v29 = l0 + v28
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+40)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v28 + v30
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v33
			v38 = v29 + int32(24)
			return v38
		}
	} else {
		v28 = v4
		v29 = l0 + v28
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+40)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v28 + v30
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v33
		v38 = v29 + int32(24)
		return v38
	}
}
func F_record_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	v4 = F_record_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if int32(0) < v4 {
			v10 = int32(24)
		} else {
			v10 = int32(40)
		}
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_record_le(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v2 <= int32(0)))
	}
}
func F_record_ne(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2 ^ int64(1)
	}
}
func F_reduce_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = base.I32_div_s(l3, int32(2))
	v21 = l1 - int32(1)
	if v19 <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l5
	v26 = F_palloc_mul(m, int32(8), l3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v218 = l1
	goto L3
L3:
	;
	m.G0 = v16 + int32(16)
	return v218
L4:
	;
	return int32(0)
L5:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v30
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0+v21*int32(24))+8))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v35
	if l3 <= int32(3) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v174 = int32(0)
	goto L23
L7:
	;
	v40 = int32(8)
	F_qsort_arg(m, v26, int32(2), v40, int32(22), v16+v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v47 = int32(0)
	v48 = int32(2)
	if int32(6) <= l3 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v169 = int32(1)
	goto L6
L11:
	;
	v148 = int32(8)
	F_qsort_arg(m, v26, v140, v148, int32(22), v16+v148)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L22
	}
L12:
	;
	v51 = int32(2)
	if v19 <= v51 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v110 = v47
	v111 = v48
	goto L14
L14:
	;
	v121 = v26 + v111<<(uint(int32(3))%32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2+v110<<(uint(int32(4))%32))))
	v128 = l0 + v125*int32(24)
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v128)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v131
	v140 = v111 + int32(2)
	goto L11
L15:
	;
	v54 = v51
	goto L17
L16:
	;
	v54 = v19
	goto L17
L17:
	;
	v55 = int32(1)
	v56 = v54 - v55
	v65 = int32(0)
	v66 = v47
	v67 = v48
	goto L18
L18:
	;
	v77 = v26 + v67<<(uint(int32(3))%32)
	v78 = int32(4)
	v80 = l2 + v66<<(uint(v78)%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = int32(24)
	v84 = l0 + v81*v82
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v84)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v92 = l0 + v89*v82
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v92)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+24)) = v95
	v97 = int32(2)
	v98 = v66 + v97
	v100 = v67 + v78
	v102 = v65 + v97
	if v102 != v56&int32(-2) {
		v65 = v102
		v66 = v98
		v67 = v100
		goto L18
	} else {
		goto L20
	}
L19:
	;
	if v56&v55 == int32(0) {
		v140 = v100
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v110 = v98
	v111 = v100
	goto L14
L22:
	;
	v169 = int32(base.Ui32(v140) >> (uint(int32(1)) % 32))
	goto L6
L23:
	;
	v186 = l0 + v174*int32(24)
	v189 = v26 + v174<<(uint(int32(4))%32)
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	*(*int64)(unsafe.Add(mBase, uint32(v186))) = v190
	v193 = v189 + int32(8)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	v199 = F_FunctionCall2Coll(m, v196, v197, v198, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	v218 = v169
	goto L3
L25:
	;
	if v199 == int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	v207 = F_FunctionCall2Coll(m, v203, v204, v205, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	v210 = int64(1)
	goto L28
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+16)) = uint8(base.B2i32(v210 == int64(0)))
	v215 = v174 + int32(1)
	if v215 != v169 {
		v174 = v215
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v210 = v207
	goto L28
L30:
	;
	goto L24
}
func F_reduce_outer_joins_pass1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = F_palloc(m, int32(16))
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
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v17 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v17)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v17
	if l0 == v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(48)
	return v11
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v23 - int32(63) {
	case 0:
		goto L5
	case 1:
		goto L7
	case 2:
		goto L8
	default:
		goto L6
	}
L5:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = F_bms_make_singleton(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L40
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L7:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v67 = F_reduce_outer_joins_pass1(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v36 = int32(0)
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v36<<(uint(int32(2))%32))))
	v42 = F_reduce_outer_joins_pass1(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L3
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v46 = F_bms_add_members(m, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	v51 = v49 | v50
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v55 = F_bms_add_members(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v59 = F_lappend(m, v58, v42)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v59
	v63 = v36 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v63 < v64 {
		v36 = v63
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = F_reduce_outer_joins_pass1(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v74 = F_bms_union(m, v72, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
	v85 = F_list_make2_impl(m, v8+int32(36), v8+int32(32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v85
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v88 {
	case 0, 4:
		goto L26
	case 1, 5:
		goto L25
	case 2:
		goto L23
	case 3:
		goto L24
	default:
		goto L22
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L34
	}
L23:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v119 = F_bms_union(m, v117, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L33
	}
L24:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v108)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v112 = F_bms_union(m, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L25:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v101)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v105 = F_bms_union(m, v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)))
	if v89 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v92 = int32(1)
	goto L29
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	v92 = v91
	goto L29
L29:
	;
	v94 = v92 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v98 = F_bms_union(m, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v98
	goto L3
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v105
	goto L3
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v112
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v119
	goto L3
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v126
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass1_0), v8+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass1_1), int32(3480), int32(_a_F_reduce_outer_joins_pass1_2))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v142
	F_errmsg_internal(m, int32(_a_F_reduce_outer_joins_pass1_3), v8)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_reduce_outer_joins_pass1_1), int32(3486), int32(_a_F_reduce_outer_joins_pass1_2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v153
	goto L3
}
func F_regclassin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
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
	var v155 int32
	_ = v155
	var v159 int64
	_ = v159
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L29
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v159
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regclassin[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v159 = v6
	goto L2
L8:
	;
	v23 = int32(_a_F_regclassin_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regclassin[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regclassin[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v159 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v159 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v128 = F_makeRangeVarFromNameList(m, v122)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v130 = int32(0)
	v134 = F_RangeVarGetRelidExtended(m, v128, v130, int32(1), v130, v130)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	if v134 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v138 = F_errsave_start(m, v11)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v159 = base.I64_extend_i32_u(v134)
	goto L2
L41:
	;
	if v138 == int32(0) {
		v159 = v6
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v145 = F_NameListToString(m, v122)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v145
	F_errmsg(m, int32(_a_F_regclassin_1), v9)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, v11, int32(_a_F_regclassin_2), int32(922), int32(_a_F_regclassin_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	v159 = v6
	goto L2
L47:
	;
	F_errmsg_internal(m, int32(_a_F_regclassin_4), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_regclassin_2), int32(905), int32(_a_F_regclassin_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcollationout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	if v12 == int32(0) {
		v16 = F_pstrdup(m, int32(_a_F_regcollationout_0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v57 = v16
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v57)
		}
	} else {
		v23 = F_SearchSysCache1(m, int32(16), v11&int64(4294967295))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v27 = v25 + v26
				v29 = v27 + int32(4)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_regcollationout[0]))
				if v31 == int32(0) {
					v34 = F_pstrdup(m, v29)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v23)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							v57 = v34
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v57)
						}
					}
				} else {
					v38 = F_CollationIsVisible(m, v12)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						if v38 != 0 {
							v44 = int32(0)
							v45 = F_quote_qualified_identifier(m, v44, v29)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								F_ReleaseCatCache(m, v23)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v57 = v45
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_u(v57)
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
							v42 = F_get_namespace_name(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								v44 = v42
								v45 = F_quote_qualified_identifier(m, v44, v29)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int64(0)
								} else {
									F_ReleaseCatCache(m, v23)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int64(0)
									} else {
										v57 = v45
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_u(v57)
									}
								}
							}
						}
					}
				}
			} else {
				v50 = F_palloc(m, int32(64))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
					v55 = F_pg_snprintf(m, v50, int32(64), int32(_a_F_regcollationout_1), v9)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						v57 = v50
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v57)
					}
				}
			}
		}
	}
}
func F_register_ENR(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_lappend(m, v3, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
		return
	}
}
func F_regoperatorin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(432)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L25
	} else {
		goto L56
	}
L2:
	;
	m.G0 = v8 + int32(432)
	return v199
L3:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_regoperatorin[0]))
	if v115 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	v19 = int32(_a_F_regoperatorin_0)
	v23 = m.G0
	v25 = v23 - int32(32)
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v26
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperatorin[1])))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v103 = F_strlen(m, v11)
	mBase = m.M
	if v102 != v103 {
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v102 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperatorin[2])))
	if v38 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = v11
	goto L12
L10:
	;
	goto L11
L11:
	;
	v52 = v19
	v53 = v34
	goto L15
L12:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v48 == v34 {
		v42 = v42 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v102 = v42 - v11
	goto L5
L14:
	;
	goto L13
L15:
	;
	v60 = v25 + int32(base.Ui32(v53)>>(uint(int32(3))%32))&int32(28)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v61 | v62<<(uint(v53)%32)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v66 != 0 {
		v52 = v52 + v62
		v53 = v66
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v69 == int32(0) {
		v92 = v11
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v102 = v92 - v11
	goto L5
L19:
	;
	v73 = v11
	v74 = v69
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(base.Ui32(v74)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v82)>>(uint(v74)%32))&int32(1) == int32(0) {
		v92 = v73
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v92 = v90
	goto L18
L22:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v90 = v73 + int32(1)
	if v88 != 0 {
		v73 = v90
		v74 = v88
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v109 = F_DirectInputFunctionCallSafe(m, int32(588), v11, int32(-1), v10, v8+int32(16))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int64(0)
L26:
	;
	v113 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)))
	v199 = v113
	goto L2
L27:
	;
	v125 = F_parseNameAndArgTypes(m, v11, int32(1), v8+int32(428), v8+int32(424), v8+int32(16), v10)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v125 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v129 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v129)
	v199 = v5
	goto L2
L30:
	;
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+424))
	switch v131 - int32(1) {
	case 0:
		goto L34
	case 1:
		goto L32
	default:
		goto L33
	}
L32:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v8)+428))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v177 = F_OpernameGetOprid(m, v174, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L25
	} else {
		goto L47
	}
L33:
	;
	v154 = F_errsave_start(m, v10)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
	} else {
		goto L41
	}
L34:
	;
	v134 = F_errsave_start(m, v10)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	if v134 == int32(0) {
		v199 = v5
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(33685636))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_regoperatorin_1), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	F_errhint(m, int32(_a_F_regoperatorin_2), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v10, int32(_a_F_regoperatorin_3), int32(679), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v199 = v5
	goto L2
L41:
	;
	if v154 == int32(0) {
		v199 = v5
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_regoperatorin_5), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_regoperatorin_6), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, v10, int32(_a_F_regoperatorin_3), int32(684), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v199 = v5
	goto L2
L47:
	;
	if v177 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v181 = F_errsave_start(m, v10)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L25
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v199 = base.I64_extend_i32_u(v177)
	goto L2
L51:
	;
	if v181 == int32(0) {
		v199 = v5
		goto L2
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errmsg(m, int32(_a_F_regoperatorin_7), v8)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	F_errsave_finish(m, v10, int32(_a_F_regoperatorin_3), int32(691), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	v199 = v5
	goto L2
L56:
	;
	F_errmsg_internal(m, int32(_a_F_regoperatorin_8), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L25
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_regoperatorin_3), int32(662), int32(_a_F_regoperatorin_4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regoperin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L25
	} else {
		goto L49
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v171
L3:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_regoperin[0]))
	if v116 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	v20 = int32(_a_F_regoperin_0)
	v24 = m.G0
	v26 = v24 - int32(32)
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v27
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperin[1])))
	if v35 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v104 = F_strlen(m, v12)
	mBase = m.M
	if v103 != v104 {
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v103 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regoperin[2])))
	if v39 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = v12
	goto L12
L10:
	;
	goto L11
L11:
	;
	v53 = v20
	v54 = v35
	goto L15
L12:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v49 == v35 {
		v43 = v43 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v103 = v43 - v12
	goto L5
L14:
	;
	goto L13
L15:
	;
	v61 = v26 + int32(base.Ui32(v54)>>(uint(int32(3))%32))&int32(28)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v62 | v63<<(uint(v54)%32)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v67 != 0 {
		v53 = v53 + v63
		v54 = v67
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v70 == int32(0) {
		v93 = v12
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v103 = v93 - v12
	goto L5
L19:
	;
	v74 = v12
	v75 = v70
	goto L20
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(base.Ui32(v75)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v83)>>(uint(v75)%32))&int32(1) == int32(0) {
		v93 = v74
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v93 = v91
	goto L18
L22:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v91 = v74 + int32(1)
	if v89 != 0 {
		v74 = v91
		v75 = v89
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v110 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(24))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int64(0)
L26:
	;
	v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
	v171 = v114
	goto L2
L27:
	;
	v119 = F_stringToQualifiedNameList(m, v12, v11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v119 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v123)
	v171 = v6
	goto L2
L30:
	;
	goto L31
L31:
	;
	v129 = F_OpernameGetCandidates(m, v119, int32(0), int32(1), v9+int32(24))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if v129 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = F_errsave_start(m, v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v149 != 0 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	if v133 == int32(0) {
		v171 = v6
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg(m, int32(_a_F_regoperin_1), v9)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v11, int32(_a_F_regoperin_2), int32(516), int32(_a_F_regoperin_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v171 = v6
	goto L2
L41:
	;
	v150 = F_errsave_start(m, v11)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L25
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v168 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v129)+8)))
	v171 = v168
	goto L2
L44:
	;
	if v150 == int32(0) {
		v171 = v6
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	F_errmsg(m, int32(_a_F_regoperin_4), v9+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	F_errsave_finish(m, v11, int32(_a_F_regoperin_2), int32(521), int32(_a_F_regoperin_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v171 = v6
	goto L2
L49:
	;
	F_errmsg_internal(m, int32(_a_F_regoperin_5), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_regoperin_2), int32(501), int32(_a_F_regoperin_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocedureout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2 == int32(0) {
		v6 = F_pstrdup(m, int32(_a_F_regprocedureout_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v6)
		}
	} else {
		v13 = F_format_procedure_extended(m, v2, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v13)
		}
	}
}
func F_regprocin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L29
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v178
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regprocin[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v178 = v6
	goto L2
L8:
	;
	v23 = int32(_a_F_regprocin_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regprocin[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regprocin[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(24))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)))
	v178 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v178 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v129 = int32(0)
	v136 = F_FuncnameGetCandidates(m, v122, int32(-1), v129, v129, v129, v129, int32(1), v9+int32(24))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v136 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = F_errsave_start(m, v11)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	if v140 == int32(0) {
		v178 = v6
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg(m, int32(_a_F_regprocin_1), v9)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v11, int32(_a_F_regprocin_2), int32(103), int32(_a_F_regprocin_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v178 = v6
	goto L2
L45:
	;
	v157 = F_errsave_start(m, v11)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v175 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v136)+8)))
	v178 = v175
	goto L2
L48:
	;
	if v157 == int32(0) {
		v178 = v6
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	F_errmsg(m, int32(_a_F_regprocin_4), v9+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v11, int32(_a_F_regprocin_2), int32(108), int32(_a_F_regprocin_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v178 = v6
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(_a_F_regprocin_5), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_regprocin_2), int32(87), int32(_a_F_regprocin_3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_remove_redundant_nullability_quals(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L11
	} else {
		goto L99
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v15 - int32(63) {
	case 0:
		goto L2
	case 1:
		goto L4
	case 2:
		goto L5
	default:
		goto L1
	}
L4:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_remove_redundant_nullability_quals(m, v166, l1)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L57
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v51 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = v3
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v27<<(uint(int32(2))%32))))
	F_remove_redundant_nullability_quals(m, v36, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	return
L12:
	;
	v40 = v27 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v40 < v41 {
		v27 = v40
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L2
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if int32(0) < v56 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(0)
	v64 = v3
	goto L20
L18:
	;
	v161 = v3
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v161
	goto L2
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v63<<(uint(int32(2))%32))))
	v73 = int32(0)
	if v72 == v73 {
		v102 = v73
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v161 = v152
	goto L19
L22:
	;
	v154 = v63 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v154 < v155 {
		v63 = v154
		v64 = v152
		goto L20
	} else {
		goto L56
	}
L23:
	;
	if v102 != 0 {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	goto L23
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	switch v76 - int32(52) {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		v102 = v73
		goto L24
	}
L26:
	;
	v102 = int32(0)
	goto L24
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v88 != int32(4) {
		goto L26
	} else {
		goto L34
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v79 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+12)))
	if v80 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v81 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v84 != int32(6) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	if v87 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v102 = v81
	goto L24
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v91 == int32(0) {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v94 != int32(6) {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v97 == int32(0) {
		v102 = v91
		goto L24
	} else {
		goto L37
	}
L37:
	;
	goto L26
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	v104 = int32(0)
	if base.B2i32(v103 == v104)|base.B2i32(l1 == v104) != 0 {
		v149 = v104
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v150 = F_lappend(m, v64, v72)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L11
	} else {
		goto L55
	}
L41:
	;
	if v149 != 0 {
		v152 = v64
		goto L22
	} else {
		goto L54
	}
L42:
	;
	goto L41
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v114 < v115 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v117 = v114
	goto L46
L45:
	;
	v117 = v115
	goto L46
L46:
	;
	if v117 <= int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v120 = int32(1)
	goto L49
L48:
	;
	v120 = v117
	goto L49
L49:
	;
	v121 = int32(8)
	v126 = int32(0)
	goto L50
L50:
	;
	v133 = v126 << (uint(int32(2)) % 32)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1+v121+v133)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v103+v121+v133)))
	v138 = v135 & v137
	v140 = base.B2i32(v138 != int32(0))
	if v138 != 0 {
		v149 = v140
		goto L42
	} else {
		goto L52
	}
L51:
	;
	v149 = v140
	goto L42
L52:
	;
	v142 = v126 + int32(1)
	if v142 != v120 {
		v126 = v142
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L40
L55:
	;
	v152 = v150
	goto L22
L56:
	;
	goto L21
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_remove_redundant_nullability_quals(m, v169, l1)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v172 == int32(0) {
		v279 = v3
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v279
	goto L2
L60:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v175 <= int32(0) {
		v279 = v3
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v181 = v3
	v182 = v3
	goto L62
L62:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v181<<(uint(int32(2))%32))))
	v191 = int32(0)
	if v190 == v191 {
		v220 = v191
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v279 = v270
	goto L59
L64:
	;
	v272 = v181 + int32(1)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v272 < v273 {
		v181 = v272
		v182 = v270
		goto L62
	} else {
		goto L98
	}
L65:
	;
	if v220 != 0 {
		goto L80
	} else {
		goto L81
	}
L66:
	;
	goto L65
L67:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	switch v194 - int32(52) {
	case 0:
		goto L70
	case 1:
		goto L69
	default:
		v220 = v191
		goto L66
	}
L68:
	;
	v220 = int32(0)
	goto L66
L69:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	if v206 != int32(4) {
		goto L68
	} else {
		goto L76
	}
L70:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	if v197 != 0 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+12)))
	if v198 != 0 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v199 == int32(0) {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v202 != int32(6) {
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)+28))
	if v205 != 0 {
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v220 = v199
	goto L66
L76:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v209 == int32(0) {
		goto L68
	} else {
		goto L77
	}
L77:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v212 != int32(6) {
		goto L68
	} else {
		goto L78
	}
L78:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)+28))
	if v215 == int32(0) {
		v220 = v209
		goto L66
	} else {
		goto L79
	}
L79:
	;
	goto L68
L80:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+24))
	v222 = int32(0)
	if base.B2i32(v221 == v222)|base.B2i32(l1 == v222) != 0 {
		v267 = v222
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	v268 = F_lappend(m, v182, v190)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L11
	} else {
		goto L97
	}
L83:
	;
	if v267 != 0 {
		v270 = v182
		goto L64
	} else {
		goto L96
	}
L84:
	;
	goto L83
L85:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v232 < v233 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v235 = v232
	goto L88
L87:
	;
	v235 = v233
	goto L88
L88:
	;
	if v235 <= int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v238 = int32(1)
	goto L91
L90:
	;
	v238 = v235
	goto L91
L91:
	;
	v239 = int32(8)
	v244 = int32(0)
	goto L92
L92:
	;
	v251 = v244 << (uint(int32(2)) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1+v239+v251)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v221+v239+v251)))
	v256 = v253 & v255
	v258 = base.B2i32(v256 != int32(0))
	if v256 != 0 {
		v267 = v258
		goto L84
	} else {
		goto L94
	}
L93:
	;
	v267 = v258
	goto L84
L94:
	;
	v260 = v244 + int32(1)
	if v260 != v238 {
		v244 = v260
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	goto L82
L97:
	;
	v270 = v268
	goto L64
L98:
	;
	goto L63
L99:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v299
	F_errmsg_internal(m, int32(_a_F_remove_redundant_nullability_quals_0), v11)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_remove_redundant_nullability_quals_1), int32(3367), int32(_a_F_remove_redundant_nullability_quals_2))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_remove_self_joins_recurse(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v354 int32
	_ = v354
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v681 int32
	_ = v681
	var v700 int32
	_ = v700
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v925 int32
	_ = v925
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int64
	_ = v953
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int64
	_ = v972
	var v973 int32
	_ = v973
	var v974 int64
	_ = v974
	var v975 int32
	_ = v975
	var v976 int64
	_ = v976
	var v977 int32
	_ = v977
	var v978 int64
	_ = v978
	var v979 int32
	_ = v979
	var v980 int64
	_ = v980
	var v984 int64
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int64
	_ = v989
	var v992 int64
	_ = v992
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1166 int32
	_ = v1166
	var v1176 int32
	_ = v1176
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1227 int32
	_ = v1227
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1284 int32
	_ = v1284
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1543 int32
	_ = v1543
	var v1570 int32
	_ = v1570
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1628 int32
	_ = v1628
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1683 int32
	_ = v1683
	var v1689 int32
	_ = v1689
	var v1697 int32
	_ = v1697
	var v1713 int32
	_ = v1713
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(48)
	m.G0 = v30
	if l1 == v3 {
		v134 = v3
		v143 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v149 = int64(0)
	if v134 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v34 <= int32(0) {
		v134 = v3
		v143 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = v3
	v50 = v3
	v59 = v3
	goto L4
L4:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v39<<(uint(int32(2))%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v134 = v115
	v143 = v116
	goto L1
L6:
	;
	v118 = v39 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v118 < v119 {
		v39 = v118
		v50 = v115
		v59 = v116
		goto L4
	} else {
		goto L24
	}
L7:
	;
	if v69 == int32(63) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v110 = F_remove_self_joins_recurse(m, l0, v68)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L23
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74+v75<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v80 != 0 {
		v115 = v50
		v116 = v59
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L18
	} else {
		goto L20
	}
L13:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+21)))
	if v81 != int32(114) {
		v115 = v50
		v116 = v59
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	if v84 != 0 {
		v115 = v50
		v116 = v59
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	if v75 == v86 {
		v115 = v50
		v116 = v59
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
	if v75 == v88 {
		v115 = v50
		v116 = v59
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v90 = F_bms_add_member(m, v50, v75)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v115 = v90
	v116 = v59
	goto L6
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v98
	F_errmsg_internal(m, int32(_a_F_remove_self_joins_recurse_0), v30+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_remove_self_joins_recurse_1), int32(2004), int32(_a_F_remove_self_joins_recurse_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v115 = v50
	v116 = v110 | v59
	goto L6
L24:
	;
	goto L5
L25:
	;
	if int32(2) <= v193 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	v193 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v154 = v134 + int32(8)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v155 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v193 = base.I32_popcnt(v158)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v161 = v155 << (uint(int32(2)) % 32)
	if v161 <= int32(7) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v193 = base.I32_wrap_i64(v188)
	goto L25
L33:
	;
	if v161 == int32(0) {
		v188 = v149
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v185 = F_pg_popcount_optimized(m, v154, v161)
	mBase = m.M
	v188 = v185
	goto L32
L36:
	;
	v166 = v161
	v167 = v154
	v168 = v149
	goto L37
L37:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
	v170 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v169)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
	v172 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v171)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v176 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v180 = v170 + (v172 + (v174 + (v168 + v176)))
	v181 = int32(4)
	v184 = v166 - v181
	if v184 != 0 {
		v166 = v184
		v167 = v167 + v181
		v168 = v180
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v188 = v180
	goto L32
L39:
	;
	goto L38
L40:
	;
	v197 = F_palloc_mul(m, int32(8), v193)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L18
	} else {
		goto L43
	}
L41:
	;
	v1697 = v30
	v1713 = v143
	goto L42
L42:
	;
	m.G0 = v1697 + int32(48)
	return v1713 & int32(1)
L43:
	;
	if v134 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if int32(0) <= v255 {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v255 = base.I32_ctz(v241) | v242<<(uint(int32(5))%32)
	goto L44
L46:
	;
	v255 = int32(-2)
	goto L44
L47:
	;
	v206 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v209 <= v206 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v212 = v134 + int32(8)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v219 = v216 & int32(-1)
	if v219 != 0 {
		v241 = v219
		v242 = v206
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v220 = int32(1)
	if v220 == v209 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v224 = v220
	goto L51
L51:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v212+v224<<(uint(int32(2))%32))))
	if v231 != 0 {
		v241 = v231
		v242 = v224
		goto L45
	} else {
		goto L53
	}
L52:
	;
	goto L46
L53:
	;
	v233 = v224 + int32(1)
	if v233 != v209 {
		v224 = v233
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v260 = int32(0)
	v261 = v255
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_pg_qsort(m, v197, v193, int32(8), int32(875))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L18
	} else {
		goto L72
	}
L58:
	;
	v288 = v197 + v260<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v261
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v261<<(uint(int32(2))%32))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v295
	if v134 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L57
L60:
	;
	if int32(0) <= v354 {
		v260 = v260 + int32(1)
		v261 = v354
		goto L58
	} else {
		goto L71
	}
L61:
	;
	v354 = base.I32_ctz(v340) | v341<<(uint(int32(5))%32)
	goto L60
L62:
	;
	v354 = int32(-2)
	goto L60
L63:
	;
	v305 = v261 + int32(1)
	v307 = int32(base.Ui32(v305) >> (uint(int32(5)) % 32))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v308 <= v307 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v311 = v134 + int32(8)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311+v307<<(uint(int32(2))%32))))
	v318 = v315 & (int32(-1) << (uint(v305) % 32))
	if v318 != 0 {
		v340 = v318
		v341 = v307
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v320 = v307 + int32(1)
	if v320 == v308 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v323 = v320
	goto L67
L67:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v311+v323<<(uint(int32(2))%32))))
	if v330 != 0 {
		v340 = v330
		v341 = v323
		goto L61
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v332 = v323 + int32(1)
	if v332 != v308 {
		v323 = v332
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L59
L72:
	;
	v389 = l0
	v393 = int32(1)
	v395 = v30
	v397 = v3
	v402 = v134
	v409 = v197
	v411 = v143
	v412 = v193
	goto L73
L73:
	;
	if v393 != v412 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v1697 = v395
	v1713 = v1683
	goto L42
L75:
	;
	v1689 = v393 + int32(1)
	if v1689 <= v412 {
		v393 = v1689
		v397 = v1669
		v402 = v1674
		v411 = v1683
		goto L73
	} else {
		goto L344
	}
L76:
	;
	v417 = int32(3)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v409+v393<<(uint(v417)%32))+4))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v409+v397<<(uint(v417)%32))+4))
	if v420 == v424 {
		v1669 = v397
		v1674 = v402
		v1683 = v411
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v393-v397 <= int32(1) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	if v393 <= v397 {
		v1669 = v397
		v1674 = v402
		v1683 = v411
		goto L75
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v466 = int32(0)
	if v397 < v393 {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v438 = v397
	v443 = v402
	goto L84
L84:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v409+v438<<(uint(int32(3))%32))))
	v461 = F_bms_del_member(m, v443, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L18
	} else {
		goto L86
	}
L85:
	;
	v1669 = v393
	v1674 = v461
	v1683 = v411
	goto L75
L86:
	;
	v464 = v438 + int32(1)
	if v464 != v393 {
		v438 = v464
		v443 = v461
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v472 = v466
	v478 = v397
	goto L91
L89:
	;
	v514 = v397
	v521 = v466
	goto L90
L90:
	;
	v533 = F_bms_del_members(m, v402, v521)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L18
	} else {
		goto L95
	}
L91:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v409+v478<<(uint(int32(3))%32))))
	v501 = F_bms_add_member(m, v472, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L18
	} else {
		goto L93
	}
L92:
	;
	v514 = v393
	v521 = v501
	goto L90
L93:
	;
	v504 = v478 + int32(1)
	if v504 != v393 {
		v472 = v501
		v478 = v504
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	if v521 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	if int32(0) < v591 {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	v591 = base.I32_ctz(v577) | v578<<(uint(int32(5))%32)
	goto L96
L98:
	;
	v591 = int32(-2)
	goto L96
L99:
	;
	v542 = int32(0)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v545 <= v542 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v548 = v521 + int32(8)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v555 = v552 & int32(-1)
	if v555 != 0 {
		v577 = v555
		v578 = v542
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v556 = int32(1)
	if v556 == v545 {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v560 = v556
	goto L103
L103:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v548+v560<<(uint(int32(2))%32))))
	if v567 != 0 {
		v577 = v567
		v578 = v560
		goto L97
	} else {
		goto L105
	}
L104:
	;
	goto L98
L105:
	;
	v569 = v560 + int32(1)
	if v569 != v545 {
		v560 = v569
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v608 = v591
	v618 = v466
	goto L110
L108:
	;
	v1655 = v466
	goto L109
L109:
	;
	F_bms_free(m, v521)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L18
	} else {
		goto L343
	}
L110:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v389)+36))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v621+v608<<(uint(int32(2))%32))))
	if v521 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v1655 = v1570
	goto L109
L112:
	;
	if v521 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L113:
	;
	if v681 <= int32(0) {
		v1570 = v618
		goto L112
	} else {
		goto L124
	}
L114:
	;
	v681 = base.I32_ctz(v667) | v668<<(uint(int32(5))%32)
	goto L113
L115:
	;
	v681 = int32(-2)
	goto L113
L116:
	;
	v632 = v608 + int32(1)
	v634 = int32(base.Ui32(v632) >> (uint(int32(5)) % 32))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v635 <= v634 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v638 = v521 + int32(8)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v638+v634<<(uint(int32(2))%32))))
	v645 = v642 & (int32(-1) << (uint(v632) % 32))
	if v645 != 0 {
		v667 = v645
		v668 = v634
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v647 = v634 + int32(1)
	if v647 == v635 {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v650 = v647
	goto L120
L120:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v638+v650<<(uint(int32(2))%32))))
	if v657 != 0 {
		v667 = v657
		v668 = v650
		goto L114
	} else {
		goto L122
	}
L121:
	;
	goto L115
L122:
	;
	v659 = v650 + int32(1)
	if v659 != v635 {
		v650 = v659
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v700 = v681
	goto L125
L125:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v389)+36))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v711+v700<<(uint(int32(2))%32))))
	v716 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+28)) = v716
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v389)+120))
	if v718 == v716 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v1570 = v618
	goto L112
L127:
	;
	if v521 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L128:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v389)+144))
	if v799 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L129:
	;
	v724 = int32(0)
	goto L130
L130:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	if v749 <= v724 {
		goto L128
	} else {
		goto L132
	}
L131:
	;
	goto L127
L132:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v718)+12))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v751+v724<<(uint(int32(2))%32))))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)+12))
	v757 = F_bms_is_member(m, v700, v756)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v755)+12))
	v760 = F_bms_is_member(m, v608, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	if v757 != v760 {
		goto L127
	} else {
		goto L135
	}
L135:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v755)+16))
	v766 = F_bms_is_member(m, v700, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L18
	} else {
		goto L136
	}
L136:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v755)+16))
	v769 = F_bms_is_member(m, v608, v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	if v766 == v769 {
		v724 = v724 + int32(1)
		goto L130
	} else {
		goto L138
	}
L138:
	;
	goto L131
L139:
	;
	v900 = F_bms_add_member(m, int32(0), v608)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L18
	} else {
		goto L163
	}
L140:
	;
	v802 = int32(0)
	v882 = v802
	v889 = v802
	goto L139
L141:
	;
	goto L142
L142:
	;
	v804 = int32(0)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	if v804 < v805 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v809 = v805
	goto L145
L144:
	;
	v809 = v804
	goto L145
L145:
	;
	v810 = int32(0)
	v814 = v810
	v822 = v804
	v829 = v810
	goto L146
L146:
	;
	if v814 != v809 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v864 = int32(0)
	if base.B2i32(v863 == v864)|base.B2i32(v862 == v864) != 0 {
		v882 = v862
		v889 = v863
		goto L139
	} else {
		goto L161
	}
L148:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v799)+12))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v840+v814<<(uint(int32(2))%32))))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+4))
	v846 = base.B2i32(v845 == v608)
	if v845 == v608 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v862 = v822
	v863 = v829
	goto L150
L150:
	;
	goto L147
L151:
	;
	v847 = v844
	goto L153
L152:
	;
	v847 = v822
	goto L153
L153:
	;
	if v845 == v700 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v851 = v844
	goto L156
L155:
	;
	v851 = v829
	goto L156
L156:
	;
	if v845 == v608 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v852 = v829
	goto L159
L158:
	;
	v852 = v851
	goto L159
L159:
	;
	v853 = int32(0)
	if base.B2i32(v852 == v853)|base.B2i32(v847 == v853) != 0 {
		v814 = v814 + int32(1)
		v822 = v847
		v829 = v852
		goto L146
	} else {
		goto L160
	}
L160:
	;
	v862 = v847
	v863 = v852
	goto L150
L161:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v863)+16))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	if v869 != v870 {
		goto L127
	} else {
		goto L162
	}
L162:
	;
	v882 = v862
	v889 = v863
	goto L139
L163:
	;
	v902 = F_bms_add_member(m, v900, v700)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L18
	} else {
		goto L164
	}
L164:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v625)+8))
	v906 = F_generate_join_implied_equalities(m, v389, v902, v904, v715, int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	if v906 == int32(0) {
		goto L127
	} else {
		goto L166
	}
L166:
	;
	v910 = int32(0)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v906)+4))
	if v910 < v913 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v917 = v910
	v925 = v910
	v935 = v910
	goto L170
L168:
	;
	v1166 = v910
	v1176 = v910
	goto L169
L169:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v715)+204))
	v1185 = F_list_concat(m, v1176, v1184)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L18
	} else {
		goto L245
	}
L170:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v906)+12))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943+v917<<(uint(int32(2))%32))))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)+96))
	if v948 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v1166 = v1151
	v1176 = v1152
	goto L169
L172:
	;
	v1154 = v917 + int32(1)
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v906)+4))
	if v1154 < v1155 {
		v917 = v1154
		v925 = v1151
		v935 = v1152
		goto L170
	} else {
		goto L244
	}
L173:
	;
	v1146 = F_lappend(m, v925, v947)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L18
	} else {
		goto L243
	}
L174:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v947)+28))
	v953 = int64(0)
	if v951 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v997 != int32(2) {
		goto L173
	} else {
		goto L190
	}
L176:
	;
	v997 = int32(0)
	goto L175
L177:
	;
	goto L178
L178:
	;
	v958 = v951 + int32(8)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	if v959 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	v997 = base.I32_popcnt(v962)
	goto L175
L180:
	;
	goto L181
L181:
	;
	v965 = v959 << (uint(int32(2)) % 32)
	if v965 <= int32(7) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v997 = base.I32_wrap_i64(v992)
	goto L175
L183:
	;
	if v965 == int32(0) {
		v992 = v953
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v989 = F_pg_popcount_optimized(m, v958, v965)
	mBase = m.M
	v992 = v989
	goto L182
L186:
	;
	v970 = v965
	v971 = v958
	v972 = v953
	goto L187
L187:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+3)))
	v974 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+2)))
	v976 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v975)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+1)))
	v978 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v977)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v980 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v979)+uint32(_c_F_remove_self_joins_recurse[0]))))
	v984 = v974 + (v976 + (v978 + (v972 + v980)))
	v985 = int32(4)
	v988 = v970 - v985
	if v988 != 0 {
		v970 = v988
		v971 = v971 + v985
		v972 = v984
		goto L187
	} else {
		goto L189
	}
L188:
	;
	v992 = v984
	goto L182
L189:
	;
	goto L188
L190:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v947)+44))
	v1001 = int32(0)
	if v1000 == v1001 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v1046 != int32(1) {
		goto L173
	} else {
		goto L207
	}
L192:
	;
	v1046 = int32(0)
	goto L191
L193:
	;
	goto L194
L194:
	;
	v1009 = int32(1)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+4))
	if v1010 <= v1009 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1013 = v1009
	goto L197
L196:
	;
	v1013 = v1010
	goto L197
L197:
	;
	v1017 = int32(0)
	v1019 = v1001
	goto L198
L198:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1000+int32(8)+v1017<<(uint(int32(2))%32))))
	if v1026 != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1046 = v1038
	goto L191
L200:
	;
	goto L199
L201:
	;
	v1027 = int32(2)
	if v1019 != 0 {
		v1038 = v1027
		goto L200
	} else {
		goto L204
	}
L202:
	;
	v1033 = v1019
	goto L203
L203:
	;
	v1035 = v1017 + int32(1)
	if v1035 != v1013 {
		v1017 = v1035
		v1019 = v1033
		goto L198
	} else {
		goto L206
	}
L204:
	;
	v1028 = int32(1)
	if base.Ui32(v1028) < base.Ui32(base.I32_popcnt(v1026)) {
		v1038 = v1027
		goto L200
	} else {
		goto L205
	}
L205:
	;
	v1033 = v1028
	goto L203
L206:
	;
	v1038 = v1033
	goto L200
L207:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v947)+48))
	v1050 = int32(0)
	if v1049 == v1050 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v1095 != int32(1) {
		goto L173
	} else {
		goto L224
	}
L209:
	;
	v1095 = int32(0)
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1058 = int32(1)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1049)+4))
	if v1059 <= v1058 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1062 = v1058
	goto L214
L213:
	;
	v1062 = v1059
	goto L214
L214:
	;
	v1066 = int32(0)
	v1068 = v1050
	goto L215
L215:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1049+int32(8)+v1066<<(uint(int32(2))%32))))
	if v1075 != 0 {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	v1095 = v1087
	goto L208
L217:
	;
	goto L216
L218:
	;
	v1076 = int32(2)
	if v1068 != 0 {
		v1087 = v1076
		goto L217
	} else {
		goto L221
	}
L219:
	;
	v1082 = v1068
	goto L220
L220:
	;
	v1084 = v1066 + int32(1)
	if v1084 != v1062 {
		v1066 = v1084
		v1068 = v1082
		goto L215
	} else {
		goto L223
	}
L221:
	;
	v1077 = int32(1)
	if base.Ui32(v1077) < base.Ui32(base.I32_popcnt(v1075)) {
		v1087 = v1076
		goto L217
	} else {
		goto L222
	}
L222:
	;
	v1082 = v1077
	goto L220
L223:
	;
	v1087 = v1082
	goto L217
L224:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v947)+4))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	if v1099 != int32(17) {
		goto L173
	} else {
		goto L225
	}
L225:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+28))
	if v1102 == int32(0) {
		goto L173
	} else {
		goto L226
	}
L226:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+4))
	if v1105 != int32(2) {
		goto L173
	} else {
		goto L227
	}
L227:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+12))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1108)))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
	v1112 = F_copyObjectImpl(m, v1111)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L18
	} else {
		goto L228
	}
L228:
	;
	v1114 = int32(0)
	if v1109 == v1114 {
		v1121 = v1114
		goto L229
	} else {
		goto L230
	}
L229:
	;
	if v1112 == int32(0) {
		v1128 = int32(0)
		goto L232
	} else {
		goto L233
	}
L230:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1109)))
	if v1117 != int32(27) {
		v1121 = v1109
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1121 = v1120
	goto L229
L232:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v947)+48))
	v1130 = F_bms_singleton_member(m, v1129)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L18
	} else {
		goto L237
	}
L233:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	if v1124 != int32(27) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1128 = v1112
	goto L232
L235:
	;
	goto L236
L236:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1128 = v1127
	goto L232
L237:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v947)+44))
	v1133 = F_bms_singleton_member(m, v1132)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L18
	} else {
		goto L238
	}
L238:
	;
	F_ChangeVarNodes(m, v1128, v1130, v1133)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L18
	} else {
		goto L239
	}
L239:
	;
	v1137 = F_equal(m, v1121, v1128)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L18
	} else {
		goto L240
	}
L240:
	;
	if v1137 == int32(0) {
		goto L173
	} else {
		goto L241
	}
L241:
	;
	v1141 = F_lappend(m, v935, v947)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L18
	} else {
		goto L242
	}
L242:
	;
	v1151 = v925
	v1152 = v1141
	goto L172
L243:
	;
	v1151 = v1146
	v1152 = v935
	goto L172
L244:
	;
	goto L171
L245:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v625)+8))
	if v1166 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+4))
	v1191 = v1189
	goto L248
L247:
	;
	v1191 = int32(0)
	goto L248
L248:
	;
	v1196 = F_innerrel_is_unique_ext(m, v389, v902, v1187, v715, int32(0), v1185, base.B2i32(v1191 == int32(0)), v395+int32(28))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L18
	} else {
		goto L249
	}
L249:
	;
	if v1196 == int32(0) {
		goto L127
	} else {
		goto L250
	}
L250:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v395)+28))
	if v1200 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1392 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+44)) = v1392
	*(*int32)(unsafe.Add(mBase, uint32(v395)+40)) = v1392
	*(*int32)(unsafe.Add(mBase, uint32(v395)+36)) = v1392
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+60))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v625)+76))
	v1405 = F_remove_rel_from_jointree(m, v1399, v1400, v395+int32(44), v395+int32(40))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L18
	} else {
		goto L299
	}
L252:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+4))
	if v1203 <= int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v715)+76))
	v1227 = int32(0)
	goto L254
L254:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+12))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1235+v1227<<(uint(int32(2))%32))))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+4))
	v1241 = F_copyObjectImpl(m, v1240)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L18
	} else {
		goto L256
	}
L255:
	;
	goto L251
L256:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v625)+76))
	F_ChangeVarNodes(m, v1241, v1206, v1243)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L18
	} else {
		goto L257
	}
L257:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+28))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+44))
	if v1247 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v625)+204))
	if v1275 == int32(0) {
		goto L127
	} else {
		goto L272
	}
L259:
	;
	if v1246 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	v1262 = int32(0)
	if v1246 == v1262 {
		goto L268
	} else {
		goto L269
	}
L262:
	;
	v1252 = int32(0)
	v1273 = v1252
	v1274 = v1252
	goto L258
L263:
	;
	goto L264
L264:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+12))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+4))
	if int32(2) <= v1256 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+4))
	v1260 = v1259
	goto L267
L266:
	;
	v1260 = int32(0)
	goto L267
L267:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	v1273 = v1260
	v1274 = v1261
	goto L258
L268:
	;
	v1273 = int32(0)
	v1274 = v1262
	goto L258
L269:
	;
	goto L270
L270:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+12))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1266)))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+4))
	if v1268 < int32(2) {
		v1273 = v1267
		v1274 = v1262
		goto L258
	} else {
		goto L271
	}
L271:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	v1273 = v1267
	v1274 = v1271
	goto L258
L272:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	if v1278 <= int32(0) {
		goto L127
	} else {
		goto L273
	}
L273:
	;
	v1284 = int32(0)
	goto L274
L274:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+12))
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1309+v1284<<(uint(int32(2))%32))))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+96))
	if v1314 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v1362 = v1227 + int32(1)
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+4))
	if v1362 < v1363 {
		v1227 = v1362
		goto L254
	} else {
		goto L298
	}
L276:
	;
	goto L275
L277:
	;
	v1358 = v1284 + int32(1)
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	if v1358 < v1359 {
		v1284 = v1358
		goto L274
	} else {
		goto L297
	}
L278:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+28))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+44))
	if v1319 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v1347 = F_equal(m, v1273, v1344)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L18
	} else {
		goto L293
	}
L280:
	;
	if v1318 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	goto L282
L282:
	;
	v1334 = int32(0)
	if v1318 == v1334 {
		goto L289
	} else {
		goto L290
	}
L283:
	;
	v1324 = int32(0)
	v1344 = v1324
	v1345 = v1324
	goto L279
L284:
	;
	goto L285
L285:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+12))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+4))
	if int32(2) <= v1328 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+4))
	v1332 = v1331
	goto L288
L287:
	;
	v1332 = int32(0)
	goto L288
L288:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	v1344 = v1332
	v1345 = v1333
	goto L279
L289:
	;
	v1344 = int32(0)
	v1345 = v1334
	goto L279
L290:
	;
	goto L291
L291:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+12))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+4))
	if v1340 < int32(2) {
		v1344 = v1339
		v1345 = v1334
		goto L279
	} else {
		goto L292
	}
L292:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	v1344 = v1339
	v1345 = v1343
	goto L279
L293:
	;
	if v1347 == int32(0) {
		goto L277
	} else {
		goto L294
	}
L294:
	;
	v1351 = F_equal(m, v1274, v1345)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L18
	} else {
		goto L295
	}
L295:
	;
	if v1351 != 0 {
		goto L276
	} else {
		goto L296
	}
L296:
	;
	goto L277
L297:
	;
	goto L127
L298:
	;
	goto L255
L299:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1407)+60)) = v1405
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v395)+40))
	if v1409 == int32(1) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v625)+76))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v715)+76))
	F_ChangeVarNodes(m, v1412, v1413, v1414)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L18
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L18
	} else {
		goto L316
	}
L303:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v389)+284))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v625)+76))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v715)+76))
	F_ChangeVarNodes(m, v1417, v1418, v1419)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L18
	} else {
		goto L304
	}
L304:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v389)+136))
	if v1422 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v625)+76))
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v715)+76))
	F_ChangeVarNodes(m, v1422, v1423, v1424)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L18
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+60))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v715)+76))
	F_fixup_selfjoin_jointree(m, v389, v1428, v1429, v395+int32(36), v395+int32(35))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L18
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	if v882 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1570 = int32(1)
	goto L112
L311:
	;
	if v889 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v389)+144))
	v1439 = F_list_delete_ptr(m, v1438, v882)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L18
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v715)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v882)+4)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v882)+8)) = v1442
	goto L310
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389)+144)) = v1439
	goto L310
L316:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v625)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v1451
	F_errmsg_internal(m, int32(_a_F_remove_self_joins_recurse_3), v395)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L18
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_remove_self_joins_recurse_1), int32(1242), int32(_a_F_remove_self_joins_recurse_4))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L18
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	if int32(0) < v1543 {
		v700 = v1543
		goto L125
	} else {
		goto L330
	}
L320:
	;
	v1543 = base.I32_ctz(v1529) | v1530<<(uint(int32(5))%32)
	goto L319
L321:
	;
	v1543 = int32(-2)
	goto L319
L322:
	;
	v1494 = v700 + int32(1)
	v1496 = int32(base.Ui32(v1494) >> (uint(int32(5)) % 32))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v1497 <= v1496 {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1500 = v521 + int32(8)
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1500+v1496<<(uint(int32(2))%32))))
	v1507 = v1504 & (int32(-1) << (uint(v1494) % 32))
	if v1507 != 0 {
		v1529 = v1507
		v1530 = v1496
		goto L320
	} else {
		goto L324
	}
L324:
	;
	v1509 = v1496 + int32(1)
	if v1509 == v1497 {
		goto L321
	} else {
		goto L325
	}
L325:
	;
	v1512 = v1509
	goto L326
L326:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1500+v1512<<(uint(int32(2))%32))))
	if v1519 != 0 {
		v1529 = v1519
		v1530 = v1512
		goto L320
	} else {
		goto L328
	}
L327:
	;
	goto L321
L328:
	;
	v1521 = v1512 + int32(1)
	if v1521 != v1497 {
		v1512 = v1521
		goto L326
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	goto L126
L331:
	;
	if int32(0) < v1628 {
		v608 = v1628
		v618 = v1570
		goto L110
	} else {
		goto L342
	}
L332:
	;
	v1628 = base.I32_ctz(v1614) | v1615<<(uint(int32(5))%32)
	goto L331
L333:
	;
	v1628 = int32(-2)
	goto L331
L334:
	;
	v1579 = v608 + int32(1)
	v1581 = int32(base.Ui32(v1579) >> (uint(int32(5)) % 32))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v1582 <= v1581 {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1585 = v521 + int32(8)
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1585+v1581<<(uint(int32(2))%32))))
	v1592 = v1589 & (int32(-1) << (uint(v1579) % 32))
	if v1592 != 0 {
		v1614 = v1592
		v1615 = v1581
		goto L332
	} else {
		goto L336
	}
L336:
	;
	v1594 = v1581 + int32(1)
	if v1594 == v1582 {
		goto L333
	} else {
		goto L337
	}
L337:
	;
	v1597 = v1594
	goto L338
L338:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1585+v1597<<(uint(int32(2))%32))))
	if v1604 != 0 {
		v1614 = v1604
		v1615 = v1597
		goto L332
	} else {
		goto L340
	}
L339:
	;
	goto L333
L340:
	;
	v1606 = v1597 + int32(1)
	if v1606 != v1582 {
		v1597 = v1606
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	goto L111
L343:
	;
	v1669 = v514
	v1674 = v533
	v1683 = v411 | v1655
	goto L75
L344:
	;
	goto L74
}
func F_removeabbrev_cluster(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = int32(0)
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v16 = l1 + v12*int32(24)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+12)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v23 = F_heap_getattr_1(m, v17, v19, v20, v16+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v23
	v27 = v12 + int32(1)
	if v27 != l2 {
		v12 = v27
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_removecaptures(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v5&int32(32) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v13 = v5 & int32(215)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v13)
	v15 = v13
	goto L3
L2:
	;
	v15 = v5
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v16
	goto L7
L5:
	;
	v34 = v15
	goto L6
L6:
	;
	if v34&int32(24) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_removecaptures(m, l0, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v34 = v31
	goto L6
L9:
	;
	return
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v23&int32(8) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v28 = v26 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v28)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v30 != 0 {
		v20 = v30
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	return
L18:
	;
	v44 = v40
	goto L21
L19:
	;
	v51 = v34
	goto L20
L20:
	;
	v53 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
	v58 = v51 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v58)
	goto L17
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	F_freesubre(m, l0, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L23
	}
L22:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v51 = v48
	goto L20
L23:
	;
	if v45 != 0 {
		v44 = v45
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
}
func F_renameatt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v18 = F_relation_open(m, l0, int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_renameatt_check(m, l0, v22, l4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L64
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L60
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L52
	}
L8:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+119)))
	if v93 != int32(99) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v28 = F_find_all_inheritors(m, l0, int32(8), v13+int32(-4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L25
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v36 = int32(0)
	goto L13
L13:
	;
	v44 = int32(0)
	if v28 == v44 {
		v54 = v44
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v48 <= v36 {
		v54 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v54 = v50 + v36<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.B2i32(v54 == int32(0))|base.B2i32(v59 <= v36) != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v62 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if l0 != v65 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62+v36<<(uint(int32(2))%32))))
	v73 = F_renameatt_internal(m, v65, l1, l2, int32(0), int32(1), v72, l6)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v36 = v36 + int32(1)
	goto L13
L24:
	;
	goto L23
L25:
	;
	v78 = F_find_inheritance_children(m, l0, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v78 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v147 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L37
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+72))
	v99 = F_find_typed_table_dependencies(m, v96, v92+int32(4), l6)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v99 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v103 <= int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v111 = int32(0)
	goto L33
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v111<<(uint(int32(2))%32))))
	v124 = int32(1)
	v127 = F_renameatt_internal(m, v123, l1, l2, v124, v124, int32(0), l6)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v130 = v111 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v130 < v131 {
		v111 = v130
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v149 = F_SearchSysCacheCopyAttName(m, l0, l1)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v149 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+22)))
	v155 = v153 + v154
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+74)))
	if v156 <= int32(0) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v155)+94)))
	if l5 < v159 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v162 = F_check_for_column_name_collision(m, v18, l2, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v167 = F_strncpy(m, v155+int32(4), l2, int32(64))
	mBase = m.M
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+63)) = uint8(v168)
	goto L43
L43:
	;
	F_CatalogTupleUpdate(m, v147, v149+int32(4), v149)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_renameatt_internal[0]))
	if v175 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v156, v177, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v149)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	F_relation_close(m, v147, int32(3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	m.G0 = v15 - int32(-64)
	return v156
L52:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_0), v13+int32(-16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3950), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_3), v15)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3974), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_4), v13+int32(-48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3982), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg(m, int32(_a_F_renameatt_internal_5), v13+int32(-32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_renameatt_internal_1), int32(3997), int32(_a_F_renameatt_internal_2))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_renametrig_partition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = v10 + int32(-56)
	F_ScanKeyInit(m, v15, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = int32(1)
	v26 = F_systable_beginscan(m, l0, int32(2701), v23, int32(0), v23, v15)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L5
L4:
	;
	F_systable_endscan(m, v26)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v37 = F_systable_getnext(m, v26)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v47 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v37 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v44 != l2 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	F_renametrig_internal(m, l0, v47, v37, l3, l4)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
	if v52 != int32(112) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_relation_close(m, v47, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L20
	}
L13:
	;
	v56 = F_RelationGetPartitionDesc(m, v47, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v58 <= int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v66 = int32(0)
	goto L16
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v66<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	F_renametrig_partition(m, l0, v77, v78, l3, v43+int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L12
L18:
	;
	v82 = v66 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v82 < v83 {
		v66 = v82
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L4
L21:
	;
	m.G0 = v12 - int32(-64)
	return
}
func F_repack_is_permitted_for_relation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v4)
	v15 = F_pg_class_aclcheck_ext(m, l1, l2, int64(16384), v8+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v19 != 0 {
			v57 = v4
			m.G0 = v8 + int32(16)
			return v57
		} else {
			if v15 == int32(0) {
				v57 = int32(1)
				m.G0 = v8 + int32(16)
				return v57
			} else {
				v23 = F_get_rel_name(m, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						v57 = v4
						m.G0 = v8 + int32(16)
						return v57
					} else {
						v29 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = l0 - int32(1)
								if base.Ui32(v33) <= base.Ui32(int32(2)) {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v33<<(uint(int32(2))%32))+uint32(_c_F_repack_is_permitted_for_relation[0])))
									v39 = v38
								} else {
									v39 = int32(_a_F_repack_is_permitted_for_relation_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v23
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v39
								F_errmsg(m, int32(_a_F_repack_is_permitted_for_relation_1), v8)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_repack_is_permitted_for_relation_2), int32(2509), int32(_a_F_repack_is_permitted_for_relation_3))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v57 = v4
											m.G0 = v8 + int32(16)
											return v57
										}
									}
								}
							} else {
								F_pfree(m, v23)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v57 = v4
									m.G0 = v8 + int32(16)
									return v57
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_repalloc_mul(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14373(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_replace_correlation_vars_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(6) {
	case 0:
		goto L9
	case 1, 2, 5, 6:
		goto L4
	case 3:
		goto L8
	case 4:
		goto L7
	case 7:
		goto L6
	default:
		goto L10
	}
L4:
	;
	v781 = F_expression_tree_mutator_impl(m, l0, int32(894), l1)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L39
	} else {
		goto L195
	}
L5:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v671 <= int32(0) {
		goto L4
	} else {
		goto L169
	}
L6:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+4))
	if v602 == int32(5) {
		goto L4
	} else {
		goto L150
	}
L7:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v500 == int32(0) {
		goto L4
	} else {
		goto L126
	}
L8:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v397 == int32(0) {
		goto L4
	} else {
		goto L103
	}
L9:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v180 == int32(0) {
		goto L4
	} else {
		goto L53
	}
L10:
	;
	if v12 == int32(61) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v12 != int32(321) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v22 == int32(0) {
		v70 = l1
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	if v75 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v26 = v22 & int32(7)
	if v26 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if base.Ui32(v22) < base.Ui32(int32(8)) {
		v70 = v44
		goto L14
	} else {
		goto L23
	}
L17:
	;
	v42 = v22
	v44 = l1
	goto L16
L18:
	;
	goto L19
L19:
	;
	v29 = v22
	v31 = l1
	v33 = v3
	goto L20
L20:
	;
	v36 = int32(1)
	v37 = v29 - v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v40 = v33 + v36
	if v40 != v26 {
		v29 = v37
		v31 = v38
		v33 = v40
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v42 = v37
	v44 = v38
	goto L16
L22:
	;
	goto L21
L23:
	;
	v51 = v42
	v53 = v44
	goto L24
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v67 = v51 - int32(8)
	if v67 != 0 {
		v51 = v67
		v53 = v65
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v70 = v65
	goto L14
L26:
	;
	goto L25
L27:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v160 = F_palloc0(m, int32(28))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L39
	} else {
		goto L49
	}
L28:
	;
	v157 = v93 + int32(8)
	goto L27
L29:
	;
	v111 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v83 = int32(0)
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81+v83<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 == int32(321) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L29
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v98 == v99 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v102 = v83 + int32(1)
	if v78 != v102 {
		v83 = v102
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L33
L39:
	;
	return int32(0)
L40:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	F_IncrementVarSublevelsUp(m, v111, v115-v116, v115)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v122 = F_palloc0(m, int32(12))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(330)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+72))
	if v128 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v131 = v129
	goto L45
L44:
	;
	v131 = int32(0)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+72))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v136 = F_exprType(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v138 = F_lappend_oid(m, v134, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+72)) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	v143 = F_lappend(m, v142, v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = v143
	v157 = v122 + int32(8)
	goto L27
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = int64(4294967304)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = F_exprType(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v170 = F_exprTypmod(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = F_exprCollation(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+20)) = v174
	return v160
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 == int32(0) {
		v231 = l1
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	if v236 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L55:
	;
	v187 = v183 & int32(7)
	if v187 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if base.Ui32(v183) < base.Ui32(int32(8)) {
		v231 = v205
		goto L54
	} else {
		goto L63
	}
L57:
	;
	v203 = v183
	v205 = l1
	goto L56
L58:
	;
	goto L59
L59:
	;
	v190 = v183
	v192 = l1
	v194 = v3
	goto L60
L60:
	;
	v197 = int32(1)
	v198 = v190 - v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v201 = v194 + v197
	if v201 != v187 {
		v190 = v198
		v192 = v199
		v194 = v201
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v203 = v198
	v205 = v199
	goto L56
L62:
	;
	goto L61
L63:
	;
	v212 = v203
	v214 = v205
	goto L64
L64:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v228 = v212 - int32(8)
	if v228 != 0 {
		v212 = v228
		v214 = v226
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v231 = v226
	goto L54
L66:
	;
	goto L65
L67:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v383 = F_palloc0(m, int32(28))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L39
	} else {
		goto L102
	}
L68:
	;
	v380 = v254 + int32(8)
	goto L67
L69:
	;
	v342 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L39
	} else {
		goto L95
	}
L70:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v239 <= int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v243 = int32(0)
	v244 = v239
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v243<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v256 != int32(6) {
		v331 = v244
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L69
L74:
	;
	v333 = v243 + int32(1)
	if v333 < v331 {
		v243 = v333
		v244 = v331
		goto L72
	} else {
		goto L94
	}
L75:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v259 != v260 {
		v331 = v244
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+8)))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v262 != v263 {
		v331 = v244
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v265 != v266 {
		v331 = v244
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v268 != v269 {
		v331 = v244
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v271 != v272 {
		v331 = v244
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+32))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v274 != v275 {
		v331 = v244
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v255)+24))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v279 = int32(0)
	if base.B2i32(v277 == v279)|base.B2i32(v278 == v279) != 0 {
		v325 = base.B2i32(v277|v278 == v279)
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v325 != 0 {
		goto L68
	} else {
		goto L93
	}
L83:
	;
	goto L82
L84:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v293 != v294 {
		v325 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v296 = int32(1)
	if v293 <= v296 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v299 = v296
	goto L88
L87:
	;
	v299 = v293
	goto L88
L88:
	;
	v300 = int32(8)
	v305 = int32(0)
	goto L89
L89:
	;
	v313 = v305 << (uint(int32(2)) % 32)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v277+v300+v313)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v278+v300+v313)))
	v318 = base.B2i32(v315 == v317)
	if v315 != v317 {
		v325 = v318
		goto L83
	} else {
		goto L91
	}
L90:
	;
	v325 = v318
	goto L83
L91:
	;
	v321 = v305 + int32(1)
	if v321 != v299 {
		v305 = v321
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v331 = v330
	goto L74
L94:
	;
	goto L73
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+28)) = int32(0)
	v347 = F_palloc0(m, int32(12))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L39
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = int32(330)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+72))
	if v353 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v356 = v354
	goto L99
L98:
	;
	v356 = int32(0)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+8)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+72))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v361 = F_lappend_oid(m, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L39
	} else {
		goto L100
	}
L100:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v363)+72)) = v361
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	v366 = F_lappend(m, v365, v347)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L39
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = v366
	v380 = v347 + int32(8)
	goto L67
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+8)) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v383))) = int64(4294967304)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+12)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+16)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+24)) = v394
	return v383
L103:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v400 == int32(0) {
		v448 = l1
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v453 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L39
	} else {
		goto L117
	}
L105:
	;
	v404 = v400 & int32(7)
	if v404 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if base.Ui32(v400) < base.Ui32(int32(8)) {
		v448 = v422
		goto L104
	} else {
		goto L113
	}
L107:
	;
	v420 = v400
	v422 = l1
	goto L106
L108:
	;
	goto L109
L109:
	;
	v407 = v400
	v409 = l1
	v411 = v3
	goto L110
L110:
	;
	v414 = int32(1)
	v415 = v407 - v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v409)+16))
	v418 = v411 + v414
	if v418 != v404 {
		v407 = v415
		v409 = v416
		v411 = v418
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v420 = v415
	v422 = v416
	goto L106
L112:
	;
	goto L111
L113:
	;
	v429 = v420
	v431 = v422
	goto L114
L114:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v431)+16))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+16))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+16))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+16))
	v445 = v429 - int32(8)
	if v445 != 0 {
		v429 = v445
		v431 = v443
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v448 = v443
	goto L104
L116:
	;
	goto L115
L117:
	;
	v455 = int32(0)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v453)+52))
	F_IncrementVarSublevelsUp(m, v453, v455-v456, v455)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L39
	} else {
		goto L118
	}
L118:
	;
	v462 = F_palloc0(m, int32(12))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L39
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = int32(330)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+72))
	if v468 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	v471 = v469
	goto L122
L121:
	;
	v471 = int32(0)
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+8)) = v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+72))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v476 = F_lappend_oid(m, v474, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L39
	} else {
		goto L123
	}
L123:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v478)+72)) = v476
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v448)+28))
	v481 = F_lappend(m, v480, v462)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L39
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448)+28)) = v481
	v485 = F_palloc0(m, int32(28))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L39
	} else {
		goto L125
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v485))) = int64(4294967304)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v462)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+8)) = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+12)) = v491
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+20)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v453)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+24)) = v497
	return v485
L126:
	;
	v503 = F_exprType(m, l0)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L39
	} else {
		goto L127
	}
L127:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v505 == int32(0) {
		v553 = l1
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v558 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L39
	} else {
		goto L141
	}
L129:
	;
	v509 = v505 & int32(7)
	if v509 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if base.Ui32(v505) < base.Ui32(int32(8)) {
		v553 = v527
		goto L128
	} else {
		goto L137
	}
L131:
	;
	v525 = v505
	v527 = l1
	goto L130
L132:
	;
	goto L133
L133:
	;
	v512 = v505
	v514 = l1
	v516 = v3
	goto L134
L134:
	;
	v519 = int32(1)
	v520 = v512 - v519
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v514)+16))
	v523 = v516 + v519
	if v523 != v509 {
		v512 = v520
		v514 = v521
		v516 = v523
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v525 = v520
	v527 = v521
	goto L130
L136:
	;
	goto L135
L137:
	;
	v534 = v525
	v536 = v527
	goto L138
L138:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v536)+16))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+16))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+16))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+16))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+16))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+16))
	v550 = v534 - int32(8)
	if v550 != 0 {
		v534 = v550
		v536 = v548
		goto L138
	} else {
		goto L140
	}
L139:
	;
	v553 = v548
	goto L128
L140:
	;
	goto L139
L141:
	;
	v560 = int32(0)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v558)+16))
	F_IncrementVarSublevelsUp(m, v558, v560-v561, v560)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L39
	} else {
		goto L142
	}
L142:
	;
	v567 = F_palloc0(m, int32(12))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L39
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+4)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = int32(330)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+72))
	if v573 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v576 = v574
	goto L146
L145:
	;
	v576 = int32(0)
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+8)) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+72))
	v580 = F_lappend_oid(m, v579, v503)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L39
	} else {
		goto L147
	}
L147:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v582)+72)) = v580
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+28))
	v585 = F_lappend(m, v584, v567)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L39
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+28)) = v585
	v589 = F_palloc0(m, int32(28))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L39
	} else {
		goto L149
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v589))) = int64(4294967304)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v589)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v589)+12)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v589)+8)) = v593
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v558)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v589)+24)) = v598
	return v589
L150:
	;
	v605 = F_exprType(m, l0)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L39
	} else {
		goto L152
	}
L151:
	;
	return v646
L152:
	;
	v608 = l1
	goto L154
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L39
	} else {
		goto L166
	}
L154:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v608)+16))
	if v614 == int32(0) {
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v621 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L39
	} else {
		goto L158
	}
L156:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	if v618 != int32(5) {
		v608 = v614
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v624 = F_palloc0(m, int32(12))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L39
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624)+4)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = int32(330)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+72))
	if v630 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v633 = v631
	goto L162
L161:
	;
	v633 = int32(0)
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624)+8)) = v633
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)+72))
	v637 = F_lappend_oid(m, v636, v605)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L39
	} else {
		goto L163
	}
L163:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+72)) = v637
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v614)+28))
	v642 = F_lappend(m, v641, v624)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L39
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+28)) = v642
	v646 = F_palloc0(m, int32(28))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L39
	} else {
		goto L165
	}
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v646))) = int64(4294967304)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v624)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v646)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v646)+12)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(v646)+8)) = v650
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v621)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v646)+24)) = v655
	goto L151
L166:
	;
	F_errmsg_internal(m, int32(_a_F_replace_correlation_vars_mutator_0), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L39
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_replace_correlation_vars_mutator_1), int32(334), int32(_a_F_replace_correlation_vars_mutator_2))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L39
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v675 = F_exprType(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L39
	} else {
		goto L170
	}
L170:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v677 == int32(0) {
		v725 = l1
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v730 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L39
	} else {
		goto L184
	}
L172:
	;
	v681 = v677 & int32(7)
	if v681 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	if base.Ui32(v677) < base.Ui32(int32(8)) {
		v725 = v699
		goto L171
	} else {
		goto L180
	}
L174:
	;
	v697 = v677
	v699 = l1
	goto L173
L175:
	;
	goto L176
L176:
	;
	v684 = v677
	v686 = l1
	v688 = v3
	goto L177
L177:
	;
	v691 = int32(1)
	v692 = v684 - v691
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v686)+16))
	v695 = v688 + v691
	if v695 != v681 {
		v684 = v692
		v686 = v693
		v688 = v695
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v697 = v692
	v699 = v693
	goto L173
L179:
	;
	goto L178
L180:
	;
	v706 = v697
	v708 = v699
	goto L181
L181:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v708)+16))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+16))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+16))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+16))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+16))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+16))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+16))
	v722 = v706 - int32(8)
	if v722 != 0 {
		v706 = v722
		v708 = v720
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v725 = v720
	goto L171
L183:
	;
	goto L182
L184:
	;
	v732 = int32(0)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	F_IncrementVarSublevelsUp(m, v730, v732-v733, v732)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L39
	} else {
		goto L185
	}
L185:
	;
	v739 = F_palloc0(m, int32(12))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L39
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v739)+4)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v739))) = int32(330)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+72))
	if v745 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	v748 = v746
	goto L189
L188:
	;
	v748 = int32(0)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v739)+8)) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)+72))
	v752 = F_lappend_oid(m, v751, v675)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L39
	} else {
		goto L190
	}
L190:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v754)+72)) = v752
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v725)+28))
	v757 = F_lappend(m, v756, v739)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L39
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725)+28)) = v757
	v761 = F_palloc0(m, int32(28))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L39
	} else {
		goto L192
	}
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v761))) = int64(4294967304)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v761)+12)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v761)+8)) = v765
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v730)+12))
	v769 = F_exprTypmod(m, v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L39
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+16)) = v769
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v730)+12))
	v773 = F_exprCollation(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L39
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+20)) = v773
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v730)+12))
	v777 = F_exprLocation(m, v776)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v761)+24)) = v777
	return v761
L195:
	;
	return v781
}
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 != int32(321) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v316 = F_expression_tree_mutator_impl(m, l0, int32(876), l1)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L83
	}
L5:
	;
	if v12 != int32(6) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v129 = F_find_placeholder_info(m, l1, l0)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L37
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return l0
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+368))
	v22 = F_bms_is_member(m, v17, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v22 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return l0
L15:
	;
	goto L16
L16:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+372))
	if v30 == v29 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	return v127
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = F_palloc0(m, int32(28))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L29
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v38 = v29
	goto L21
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v38<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v49 = F_equal(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v58 = F_palloc0(m, int32(28))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L28
	}
L23:
	;
	if v49 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v54 = v38 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v54 < v55 {
		v38 = v54
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L22
L27:
	;
	goto L18
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(4294967304)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v70
	v127 = v58
	goto L17
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(4294967304)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+72))
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v91 = v89
	goto L32
L31:
	;
	v91 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+72))
	v95 = F_lappend_oid(m, v94, v81)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+72)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v81
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v104
	v107 = F_palloc0(m, int32(12))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(361)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v111
	v113 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+372))
	v117 = F_lappend(m, v116, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+372)) = v117
	v127 = v83
	goto L17
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+368))
	v133 = int32(0)
	if v131 == v133 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v186 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v186 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v132 == int32(0) {
		v179 = v133
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v186 = v179
	goto L38
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v143 < v142 {
		v179 = v133
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(1)
	if v142 <= v145 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v148 = v145
	goto L47
L46:
	;
	v148 = v142
	goto L47
L47:
	;
	v149 = int32(8)
	v154 = int32(0)
	goto L48
L48:
	;
	v161 = v154 << (uint(int32(2)) % 32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v131+v149+v161)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v132+v149+v161)))
	v168 = v163 & (v165 ^ int32(-1))
	v170 = base.B2i32(v168 == int32(0))
	if v168 != 0 {
		v179 = v170
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v179 = v170
	goto L42
L50:
	;
	v172 = v154 + int32(1)
	if v172 != v148 {
		v154 = v172
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v190 = F_palloc0(m, int32(24))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v205 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+372))
	if v206 == v205 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = int32(321)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+16)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v201 = F_replace_nestloop_params_mutator(m, v200, l1)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v201
	return v190
L57:
	;
	return v313
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v262 = F_exprType(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L12
	} else {
		goto L72
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v209 <= int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v214 = v205
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v214<<(uint(int32(2))%32))))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v225 = F_equal(m, l0, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
	} else {
		goto L63
	}
L62:
	;
	v234 = F_palloc0(m, int32(28))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L68
	}
L63:
	;
	if v225 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v230 = v214 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v230 < v231 {
		v214 = v230
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L62
L67:
	;
	goto L58
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = int64(4294967304)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+8)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = F_exprType(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v245 = F_exprTypmod(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+16)) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v249 = F_exprCollation(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+20)) = v249
	v313 = v234
	goto L57
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v265 = F_exprTypmod(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = F_exprCollation(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v271 = F_palloc0(m, int32(28))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v271))) = int64(4294967304)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+72))
	if v276 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v279 = v277
	goto L78
L77:
	;
	v279 = int32(0)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+8)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+72))
	v283 = F_lappend_oid(m, v282, v262)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+72)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v271)+16)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = v262
	v293 = F_palloc0(m, int32(12))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = int32(361)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = v297
	v299 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v299
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+372))
	v303 = F_lappend(m, v302, v293)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+372)) = v303
	v313 = v271
	goto L57
L83:
	;
	return v316
}
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_report_invalid_encoding_db[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	F_report_invalid_encoding_int(m, v6, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_reset_formatted_start_time(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_reset_formatted_start_time[0])) = uint8(v2)
	return
}
func F_resolve_anyrange_from_others(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = F_getBaseType(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = F_get_multirange_range(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = F_format_type_be(m, v10)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_resolve_anyrange_from_others_0)
								F_errmsg(m, int32(_a_F_resolve_anyrange_from_others_1), v7)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_resolve_anyrange_from_others_2), int32(701), int32(_a_F_resolve_anyrange_from_others_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_resolve_anyrange_from_others_4), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_resolve_anyrange_from_others_2), int32(705), int32(_a_F_resolve_anyrange_from_others_3))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
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
func F_romanian_UTF_8_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	v3 = F_SN_new_env(m, int32(44))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v7 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v3)+33)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(v3)+28)) = v7
		} else {
		}
		return v3
	}
}
func F_rowtype_field_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	if l0 == int32(2249) {
		v55 = int32(1)
		return v55
	} else {
		v11 = int32(0)
		v13 = F_lookup_rowtype_tupdesc_domain(m, l0, int32(-1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if int32(0) < l1 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if l1 <= v19 {
					v30 = v13 + v19<<(uint(int32(3))%32) + l1*int32(100)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)))
					if v31 != 0 {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						if int32(0) <= v41 {
							v49 = v11
							F_DecrTupleDescRefCount(m, v13)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v55 = v49
								return v55
							}
						} else {
							v55 = v11
							return v55
						}
					} else {
						v33 = v30 - int32(72)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
						if v34 != l2 {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
							if int32(0) <= v41 {
								v49 = v11
								F_DecrTupleDescRefCount(m, v13)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v55 = v49
									return v55
								}
							} else {
								v55 = v11
								return v55
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
							if v36 != l3 {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								if int32(0) <= v41 {
									v49 = v11
									F_DecrTupleDescRefCount(m, v13)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v55 = v49
										return v55
									}
								} else {
									v55 = v11
									return v55
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
								if v38 == l4 {
									v44 = int32(1)
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if v45 < int32(0) {
										v55 = v44
										return v55
									} else {
										v49 = v44
										F_DecrTupleDescRefCount(m, v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v55 = v49
											return v55
										}
									}
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if int32(0) <= v41 {
										v49 = v11
										F_DecrTupleDescRefCount(m, v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v55 = v49
											return v55
										}
									} else {
										v55 = v11
										return v55
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					if int32(0) <= v22 {
						v49 = v11
						F_DecrTupleDescRefCount(m, v13)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v55 = v49
							return v55
						}
					} else {
						v55 = v11
						return v55
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if int32(0) <= v22 {
					v49 = v11
					F_DecrTupleDescRefCount(m, v13)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = v49
						return v55
					}
				} else {
					v55 = v11
					return v55
				}
			}
		}
	}
}
func F_rt__int_size(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_ArrayGetNItemsSafe(m, v3, l0+int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*float32)(unsafe.Add(mBase, uint32(l1))) = base.F32_convert_i32_s(v6)
		return
	}
}
func F_rtrim1(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14240(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}

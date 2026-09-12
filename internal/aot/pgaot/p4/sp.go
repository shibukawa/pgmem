package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistGetBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v9 = F_spgGetCache(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if base.Ui32(l2) < base.Ui32(int32(8161)) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
			if v15 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v22 = base.I32_div_s(int32(819200)-v17<<(uint(int32(13))%32), int32(100))
				v24 = v22
			} else {
				v24 = int32(1638)
			}
			v31 = v9 + l1&int32(7)<<(uint(int32(3))%32) - int32(-64)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			if v32 == int32(-1) {
				v35 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v35)
				v37 = F_allocNewBuffer(m, l0, l1)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					return v37
				}
			} else {
				v40 = int32(8160)
				v41 = l2 + v24
				if base.Ui32(v40) <= base.Ui32(v41) {
					v44 = v40
				} else {
					v44 = v41
				}
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if v44 <= v45 {
					v47 = F_ReadBuffer(m, l0, v32)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_ConditionalLockBuffer(m, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								F_ReleaseBuffer(m, v47)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v55)
									v57 = F_allocNewBuffer(m, l0, l1)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										return v57
									}
								}
							} else {
								if v47 < int32(0) {
									v63 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+(v47^int32(-1))<<(uint(int32(2))%32))))
									v77 = v69
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v77 = v71 + v47<<(uint(int32(13))%32) + int32(-8192)
								}
								v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+14)))
								if v78 == int32(0) {
									v94 = int32(3)
									v100 = l1<<(uint(int32(1))%32)&int32(8) | base.B2i32(l1&v94 == v94)<<(uint(int32(2))%32)
									if v47 < int32(0) {
										v104 = *(*int32)(unsafe.Add(mBase, _consts[1]))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v47^int32(-1))<<(uint(int32(2))%32))))
										v118 = v110
									} else {
										v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										v118 = v112 + v47<<(uint(int32(13))%32) + int32(-8192)
									}
									if v118&int32(3) != 0 {
									} else {
									}
									v145 = F___memset(m, v118, int32(0), int32(8192))
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v118)+10)) = int32(1572864)
									v151 = int32(8196)
									*(*uint16)(unsafe.Add(mBase, uint32(v118)+18)) = uint16(v151)
									v157 = int32(8184)
									*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)) = uint16(v157)
									*(*uint16)(unsafe.Add(mBase, uint32(v118)+14)) = uint16(v157)
									v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)))
									v161 = v118 + v160
									v162 = int32(65410)
									*(*uint16)(unsafe.Add(mBase, uint32(v161)+6)) = uint16(v162)
									*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v100)
									v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+14)))
									v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+12)))
									v167 = v165 - v166
									v168 = int32(0)
									if v168 < v167 {
										v171 = v167
									} else {
										v171 = v168
									}
									*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v171 - v44
									v174 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v174)
									return v47
								} else {
									v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
									v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77+v81))))
									if v83&int32(2) != 0 {
										v94 = int32(3)
										v100 = l1<<(uint(int32(1))%32)&int32(8) | base.B2i32(l1&v94 == v94)<<(uint(int32(2))%32)
										if v47 < int32(0) {
											v104 = *(*int32)(unsafe.Add(mBase, _consts[1]))
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v47^int32(-1))<<(uint(int32(2))%32))))
											v118 = v110
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											v118 = v112 + v47<<(uint(int32(13))%32) + int32(-8192)
										}
										if v118&int32(3) != 0 {
										} else {
										}
										v145 = F___memset(m, v118, int32(0), int32(8192))
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v118)+10)) = int32(1572864)
										v151 = int32(8196)
										*(*uint16)(unsafe.Add(mBase, uint32(v118)+18)) = uint16(v151)
										v157 = int32(8184)
										*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)) = uint16(v157)
										*(*uint16)(unsafe.Add(mBase, uint32(v118)+14)) = uint16(v157)
										v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)))
										v161 = v118 + v160
										v162 = int32(65410)
										*(*uint16)(unsafe.Add(mBase, uint32(v161)+6)) = uint16(v162)
										*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v100)
										v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+14)))
										v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+12)))
										v167 = v165 - v166
										v168 = int32(0)
										if v168 < v167 {
											v171 = v167
										} else {
											v171 = v168
										}
										*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v171 - v44
										v174 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v174)
										return v47
									} else {
										v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+12)))
										if base.Ui32(int32(24)) < base.Ui32(v86) {
											v181 = int32(3)
											if base.B2i32(v83&int32(4) == int32(0)) == base.B2i32(l1&v181 == v181) {
												F_UnlockReleaseBuffer(m, v47)
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int32(0)
												} else {
													v214 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v214)
													v216 = F_allocNewBuffer(m, l0, l1)
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														return v216
													}
												}
											} else {
												v188 = int32(0)
												if base.B2i32(l1&int32(4) == v188)^base.B2i32(v83&int32(8) == v188) != 0 {
													F_UnlockReleaseBuffer(m, v47)
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int32(0)
													} else {
														v214 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v214)
														v216 = F_allocNewBuffer(m, l0, l1)
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return int32(0)
														} else {
															return v216
														}
													}
												} else {
													v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+14)))
													v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+12)))
													v197 = v195 - v196
													v198 = int32(0)
													if v198 < v197 {
														v201 = v197
													} else {
														v201 = v198
													}
													if v201 < v44 {
														F_UnlockReleaseBuffer(m, v47)
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v214 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v214)
															v216 = F_allocNewBuffer(m, l0, l1)
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return int32(0)
															} else {
																return v216
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v201 - v44
														v205 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v205)
														return v47
													}
												}
											}
										} else {
											v94 = int32(3)
											v100 = l1<<(uint(int32(1))%32)&int32(8) | base.B2i32(l1&v94 == v94)<<(uint(int32(2))%32)
											if v47 < int32(0) {
												v104 = *(*int32)(unsafe.Add(mBase, _consts[1]))
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v47^int32(-1))<<(uint(int32(2))%32))))
												v118 = v110
											} else {
												v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
												v118 = v112 + v47<<(uint(int32(13))%32) + int32(-8192)
											}
											if v118&int32(3) != 0 {
											} else {
											}
											v145 = F___memset(m, v118, int32(0), int32(8192))
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v118)+10)) = int32(1572864)
											v151 = int32(8196)
											*(*uint16)(unsafe.Add(mBase, uint32(v118)+18)) = uint16(v151)
											v157 = int32(8184)
											*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)) = uint16(v157)
											*(*uint16)(unsafe.Add(mBase, uint32(v118)+14)) = uint16(v157)
											v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)))
											v161 = v118 + v160
											v162 = int32(65410)
											*(*uint16)(unsafe.Add(mBase, uint32(v161)+6)) = uint16(v162)
											*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v100)
											v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+14)))
											v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+12)))
											v167 = v165 - v166
											v168 = int32(0)
											if v168 < v167 {
												v171 = v167
											} else {
												v171 = v168
											}
											*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v171 - v44
											v174 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v174)
											return v47
										}
									}
								}
							}
						}
					}
				} else {
					v214 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v214)
					v216 = F_allocNewBuffer(m, l0, l1)
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
						return int32(0)
					} else {
						return v216
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v222 = m.ExcPending
			if v222 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(351643), int32(0))
				mBase = m.M
				v226 = m.ExcPending
				if v226 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(513630), int32(576), int32(235971))
					mBase = m.M
					v231 = m.ExcPending
					if v231 != 0 {
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
func F_SpGistNewBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v85
L2:
	;
	return int32(0)
L3:
	;
	if v9 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v9
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v75
	v77 = int32(8)
	v79 = int32(0)
	v82 = F_ExtendBufferedRel(m, v7+v77, v79, v79, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L28
	}
L7:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v16) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v21 = F_ReadBuffer(m, l0, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v62 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L26
	}
L12:
	;
	v23 = F_ConditionalLockBuffer(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v23 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v21 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	F_ReleaseBuffer(m, v21)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L25
	}
L17:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+14)))
	if v43 == int32(0) {
		v85 = v21
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v21^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L17
L19:
	;
	goto L20
L20:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v42 = v36 + v21<<(uint(int32(13))%32) + int32(-8192)
	goto L17
L21:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v46))))
	if v48&int32(2) != 0 {
		v85 = v21
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+12)))
	if base.Ui32(v51) < base.Ui32(int32(25)) {
		v85 = v21
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_LockBuffer(m, v21, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	goto L11
L26:
	;
	if v62 != int32(-1) {
		v16 = v62
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v85 = v82
	goto L1
}

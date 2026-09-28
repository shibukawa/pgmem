package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multixact_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+48)))
	switch int32(base.Ui32(v17) >> (uint(int32(4)) % 32)) {
	case 0:
		v247 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
		v248 = *(*int64)(unsafe.Add(mBase, uint32(v247)))
		F_SimpleLruZeroAndWritePage(m, int32(_a_F_multixact_redo_0), v248)
		mBase = m.M
		v250 = m.ExcPending
		if v250 != 0 {
			return
		} else {
			m.G0 = v14 + int32(48)
			return
		}
	case 1:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
		F_SimpleLruZeroAndWritePage(m, int32(_a_F_multixact_redo_1), v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			m.G0 = v14 + int32(48)
			return
		}
	case 2:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
		v30 = v25 + int32(20)
		F_RecordNewMultiXact(m, v26, v27, v28, v30)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v34 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+16)))
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
			v41 = F_LWLockAcquire(m, v37+int32(1664), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[1]))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				v46 = int32(1)
				v47 = v33 + v46
				if v47 != 0 {
					v49 = v47
				} else {
					v49 = v46
				}
				if v45-v49 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = v49
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[1]))
					v56 = v55
				} else {
					v56 = v44
				}
				v57 = v34 + v35
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v56)+8))
				if base.Ui64(v58) < base.Ui64(v57) {
					*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v57
				} else {
				}
				v61 = int32(0)
				v63 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
				F_LWLockRelease(m, v63+int32(1664))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+36))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					if v70 <= int32(0) {
						v153 = v69
					} else {
						if v70 != int32(1) {
							v80 = v69
							v82 = v61
							v84 = int32(0)
							for {
								v91 = int32(3)
								v95 = v30 + v82<<(uint(v91)%32)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
								if base.B2i32(base.Ui32(v80) < base.Ui32(v91))|base.B2i32(base.Ui32(v96) < base.Ui32(v91)) == int32(0) {
									if v80-v96 < int32(0) {
										v106 = v96
									} else {
										v106 = v80
									}
								} else {
									if base.Ui32(v96) <= base.Ui32(v80) {
										v106 = v80
									} else {
										v106 = v96
									}
								}
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v107))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v106)) == int32(0) {
									if base.Ui32(v106) < base.Ui32(v107) {
										v119 = v107
									} else {
										v119 = v106
									}
								} else {
									if int32(0) <= v106-v107 {
										v119 = v106
									} else {
										v119 = v107
									}
								}
								v120 = int32(2)
								v121 = v82 + v120
								v123 = v84 + v120
								if v123 != v70&int32(2147483646) {
									v80 = v119
									v82 = v121
									v84 = v123
									continue
								} else {
									break
								}
								break
							}
							if v70&int32(1) == int32(0) {
								v153 = v119
							} else {
								v127 = v119
								v129 = v121
								v138 = int32(3)
								v141 = *(*int32)(unsafe.Add(mBase, uint32(v30+v129<<(uint(v138)%32))))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v141))&base.B2i32(base.Ui32(v138) <= base.Ui32(v127)) == int32(0) {
									if base.Ui32(v127) < base.Ui32(v141) {
										v153 = v141
									} else {
										v153 = v127
									}
								} else {
									if int32(0) <= v127-v141 {
										v153 = v127
									} else {
										v153 = v141
									}
								}
							}
						} else {
							v127 = v69
							v129 = v61
							v138 = int32(3)
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v30+v129<<(uint(v138)%32))))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v141))&base.B2i32(base.Ui32(v138) <= base.Ui32(v127)) == int32(0) {
								if base.Ui32(v127) < base.Ui32(v141) {
									v153 = v141
								} else {
									v153 = v127
								}
							} else {
								if int32(0) <= v127-v141 {
									v153 = v127
								} else {
									v153 = v141
								}
							}
						}
					}
					F_AdvanceNextFullTransactionIdPastXid(m, v153)
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return
					} else {
						m.G0 = v14 + int32(48)
						return
					}
				}
			}
		}
	case 3:
		v166 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
		v167 = *(*int64)(unsafe.Add(mBase, uint32(v166)+8))
		v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
		v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
		v172 = F_errstart(m, int32(14), int32(0))
		mBase = m.M
		v173 = m.ExcPending
		if v173 != 0 {
			return
		} else {
			if v172 != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v167
				v176 = base.I64_div_u_s(v167, int64(1636))
				*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(base.Ui64(v176) >> (uint(int64(5)) % 64))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v168
				*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = base.I64_extend_i32_u(int32(base.Ui32(v168) >> (uint(int32(15)) % 32)))
				F_errmsg_internal(m, int32(_a_F_multixact_redo_2), v14+int32(16))
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_multixact_redo_3), int32(2955), int32(_a_F_multixact_redo_4))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return
					} else {
						v196 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
						v200 = F_LWLockAcquire(m, v196+int32(_a_F_multixact_redo_5), int32(0))
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							F_SetMultiXactIdLimit(m, v168, v169)
							mBase = m.M
							v203 = m.ExcPending
							if v203 != 0 {
								return
							} else {
								if base.Ui64(int64(2)) <= base.Ui64(v167) {
									v210 = base.I64_div_u_s(v167-int64(1), int64(1636))
									F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_1), v210)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return
									} else {
										v215 = int32(1)
										if v168 == v215 {
											v221 = int32(_a_F_multixact_redo_6)
										} else {
											v221 = int32(base.Ui32(v168-v215) >> (uint(int32(10)) % 32))
										}
										F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_0), base.I64_extend_i32_u(v221))
										mBase = m.M
										v224 = m.ExcPending
										if v224 != 0 {
											return
										} else {
											v226 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
											F_LWLockRelease(m, v226+int32(_a_F_multixact_redo_5))
											mBase = m.M
											v230 = m.ExcPending
											if v230 != 0 {
												return
											} else {
												m.G0 = v14 + int32(48)
												return
											}
										}
									}
								} else {
									v215 = int32(1)
									if v168 == v215 {
										v221 = int32(_a_F_multixact_redo_6)
									} else {
										v221 = int32(base.Ui32(v168-v215) >> (uint(int32(10)) % 32))
									}
									F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_0), base.I64_extend_i32_u(v221))
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return
									} else {
										v226 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
										F_LWLockRelease(m, v226+int32(_a_F_multixact_redo_5))
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return
										} else {
											m.G0 = v14 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v196 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
				v200 = F_LWLockAcquire(m, v196+int32(_a_F_multixact_redo_5), int32(0))
				mBase = m.M
				v201 = m.ExcPending
				if v201 != 0 {
					return
				} else {
					F_SetMultiXactIdLimit(m, v168, v169)
					mBase = m.M
					v203 = m.ExcPending
					if v203 != 0 {
						return
					} else {
						if base.Ui64(int64(2)) <= base.Ui64(v167) {
							v210 = base.I64_div_u_s(v167-int64(1), int64(1636))
							F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_1), v210)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return
							} else {
								v215 = int32(1)
								if v168 == v215 {
									v221 = int32(_a_F_multixact_redo_6)
								} else {
									v221 = int32(base.Ui32(v168-v215) >> (uint(int32(10)) % 32))
								}
								F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_0), base.I64_extend_i32_u(v221))
								mBase = m.M
								v224 = m.ExcPending
								if v224 != 0 {
									return
								} else {
									v226 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
									F_LWLockRelease(m, v226+int32(_a_F_multixact_redo_5))
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										m.G0 = v14 + int32(48)
										return
									}
								}
							}
						} else {
							v215 = int32(1)
							if v168 == v215 {
								v221 = int32(_a_F_multixact_redo_6)
							} else {
								v221 = int32(base.Ui32(v168-v215) >> (uint(int32(10)) % 32))
							}
							F_SimpleLruTruncate(m, int32(_a_F_multixact_redo_0), base.I64_extend_i32_u(v221))
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
								return
							} else {
								v226 = *(*int32)(unsafe.Add(mBase, _c_F_multixact_redo[0]))
								F_LWLockRelease(m, v226+int32(_a_F_multixact_redo_5))
								mBase = m.M
								v230 = m.ExcPending
								if v230 != 0 {
									return
								} else {
									m.G0 = v14 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(24), int32(0))
		mBase = m.M
		v234 = m.ExcPending
		if v234 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17 & int32(240)
			F_errmsg_internal(m, int32(_a_F_multixact_redo_7), v14)
			mBase = m.M
			v240 = m.ExcPending
			if v240 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_multixact_redo_3), int32(2972), int32(_a_F_multixact_redo_4))
				mBase = m.M
				v245 = m.ExcPending
				if v245 != 0 {
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

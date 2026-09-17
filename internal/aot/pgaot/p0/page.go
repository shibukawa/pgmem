package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PageGetTempPage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v5 = F_palloc(m, v2<<(uint(int32(8))%32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_PageGetTempPageCopySpecial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v10 = v8 << (uint(int32(8)) % 32)
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v11&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v10)) == int32(0) {
			if v10 == int32(0) {
			} else {
				v28 = v11 + v10
				v30 = v11 + int32(4)
				if base.Ui32(v30) < base.Ui32(v28) {
					v32 = v28
				} else {
					v32 = v30
				}
				v38 = (v11^int32(-1)+v32)&int32(-4) + int32(4)
				if v38 == int32(0) {
				} else {
					base.MemoryFill(m, v11, int32(0), v38)
				}
			}
		} else {
			v38 = v10
			if v38 == int32(0) {
			} else {
				base.MemoryFill(m, v11, int32(0), v38)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+10)) = int32(_a_F_PageGetTempPageCopySpecial_0)
		v49 = v10 | int32(4)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)) = uint16(v49)
		v58 = v10 - (v16&int32(-256)-v15+int32(7))&int32(-8)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v58)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v58)
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
		v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v67 = (v61<<(uint(int32(8))%32) - v64) & int32(_a_F_PageGetTempPageCopySpecial_1)
		if v67 != 0 {
			base.MemoryCopy(m, v11+v58&int32(_a_F_PageGetTempPageCopySpecial_1), l0+v64, v67)
		} else {
		}
		return v11
	}
}
func F_PageIndexTupleDelete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v16) < base.Ui32(int32(24)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v211 = m.ExcPending
		if v211 != 0 {
			return
		} else {
			F_errcode(m, int32(16779816))
			mBase = m.M
			v214 = m.ExcPending
			if v214 != 0 {
				return
			} else {
				v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
				v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v217
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v216
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v215
				F_errmsg(m, int32(_a_F_PageIndexTupleDelete_0), v14)
				mBase = m.M
				v223 = m.ExcPending
				if v223 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PageIndexTupleDelete_1), int32(1073), int32(_a_F_PageIndexTupleDelete_2))
					mBase = m.M
					v228 = m.ExcPending
					if v228 != 0 {
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
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
		if base.Ui32(v19) < base.Ui32(v16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v211 = m.ExcPending
			if v211 != 0 {
				return
			} else {
				F_errcode(m, int32(16779816))
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return
				} else {
					v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
					v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v217
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v216
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v215
					F_errmsg(m, int32(_a_F_PageIndexTupleDelete_0), v14)
					mBase = m.M
					v223 = m.ExcPending
					if v223 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PageIndexTupleDelete_1), int32(1073), int32(_a_F_PageIndexTupleDelete_2))
						mBase = m.M
						v228 = m.ExcPending
						if v228 != 0 {
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
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			if base.B2i32(base.Ui32(v21) < base.Ui32(v19))|base.B2i32(base.Ui32(int32(_a_F_PageIndexTupleDelete_3)) < base.Ui32(v21))|base.B2i32((v21+int32(7))&int32(_a_F_PageIndexTupleDelete_4) != v21) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v211 = m.ExcPending
				if v211 != 0 {
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return
					} else {
						v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
						v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v217
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v216
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v215
						F_errmsg(m, int32(_a_F_PageIndexTupleDelete_0), v14)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PageIndexTupleDelete_1), int32(1073), int32(_a_F_PageIndexTupleDelete_2))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
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
				if v16 != int32(24) {
					v39 = int32(base.Ui32(v16+int32(_a_F_PageIndexTupleDelete_5)) >> (uint(int32(2)) % 32))
				} else {
					v39 = int32(0)
				}
				v40 = int32(_a_F_PageIndexTupleDelete_6)
				v41 = v39 & v40
				if base.Ui32(v41) <= base.Ui32((l1-int32(1))&v40) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v232 = m.ExcPending
					if v232 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l1
						F_errmsg_internal(m, int32(_a_F_PageIndexTupleDelete_7), v14+int32(32))
						mBase = m.M
						v238 = m.ExcPending
						if v238 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PageIndexTupleDelete_1), int32(1077), int32(_a_F_PageIndexTupleDelete_2))
							mBase = m.M
							v243 = m.ExcPending
							if v243 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v48 = l0 + int32(20)
					v50 = l1 << (uint(int32(2)) % 32)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50)))
					v54 = int32(base.Ui32(v52) >> (uint(int32(17)) % 32))
					v56 = v52 & int32(_a_F_PageIndexTupleDelete_8)
					if base.B2i32(base.Ui32(v56) < base.Ui32(v19))|base.B2i32(base.Ui32(v21) < base.Ui32(v56+v54))|base.B2i32(v56 != (v56+int32(7))&int32(_a_F_PageIndexTupleDelete_9)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v247 = m.ExcPending
						if v247 != 0 {
							return
						} else {
							F_errcode(m, int32(16779816))
							mBase = m.M
							v250 = m.ExcPending
							if v250 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v56
								F_errmsg(m, int32(_a_F_PageIndexTupleDelete_10), v14+int32(16))
								mBase = m.M
								v257 = m.ExcPending
								if v257 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_PageIndexTupleDelete_1), int32(1092), int32(_a_F_PageIndexTupleDelete_2))
									mBase = m.M
									v262 = m.ExcPending
									if v262 != 0 {
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
						v67 = v16 - v50
						if int32(25) <= v67 {
							v71 = v67 - int32(24)
							if v71 != 0 {
								v72 = l0 + v50
								base.MemoryCopy(m, v72+int32(20), v72+int32(24), v71)
							} else {
							}
							v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
							v81 = v79
						} else {
							v81 = v19
						}
						v85 = (v54 + int32(7)) & int32(_a_F_PageIndexTupleDelete_9)
						if base.Ui32(v81) < base.Ui32(v56) {
							v87 = v56 - v81
							if v87 != 0 {
								v88 = l0 + v81
								base.MemoryCopy(m, v88+v85, v88, v87)
							} else {
							}
							v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
							v95 = v92
						} else {
							v95 = v81
						}
						v96 = v85 + v95
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v96)
						v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v100 = v98 - int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v100)
						v102 = int32(_a_F_PageIndexTupleDelete_6)
						v107 = v39 & v102
						if base.B2i32(base.Ui32(v100&v102) < base.Ui32(int32(25)))|base.B2i32(base.Ui32(v107) < base.Ui32(int32(2))) != 0 {
						} else {
							if v107 != int32(2) {
								v114 = int32(1)
								v115 = v41 - v114
								v122 = v114
								v124 = int32(0)
								for {
									v135 = v48 + v122<<(uint(int32(2))%32)
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
									if base.Ui32(v136&int32(_a_F_PageIndexTupleDelete_8)) <= base.Ui32(v56) {
										*(*int32)(unsafe.Add(mBase, uint32(v135))) = (v85+v136)&int32(_a_F_PageIndexTupleDelete_8) | v136&int32(-32768)
									} else {
									}
									v148 = v135 + int32(4)
									v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
									if base.Ui32(v149&int32(_a_F_PageIndexTupleDelete_8)) <= base.Ui32(v56) {
										*(*int32)(unsafe.Add(mBase, uint32(v148))) = (v85+v149)&int32(_a_F_PageIndexTupleDelete_8) | v149&int32(-32768)
									} else {
									}
									v160 = int32(2)
									v161 = v122 + v160
									v163 = v124 + v160
									if v163 != v115&int32(-2) {
										v122 = v161
										v124 = v163
										continue
									} else {
										break
									}
									break
								}
								if v115&v114 == int32(0) {
								} else {
									v167 = v161
									v180 = v48 + v167<<(uint(int32(2))%32)
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									if base.Ui32(v56) < base.Ui32(v181&int32(_a_F_PageIndexTupleDelete_8)) {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v180))) = (v181+v85)&int32(_a_F_PageIndexTupleDelete_8) | v181&int32(-32768)
									}
								}
							} else {
								v167 = int32(1)
								v180 = v48 + v167<<(uint(int32(2))%32)
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								if base.Ui32(v56) < base.Ui32(v181&int32(_a_F_PageIndexTupleDelete_8)) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v180))) = (v181+v85)&int32(_a_F_PageIndexTupleDelete_8) | v181&int32(-32768)
								}
							}
						}
						m.G0 = v14 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_PageIndexTupleOverwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v22) < base.Ui32(int32(24)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v260 = m.ExcPending
		if v260 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16779816))
			mBase = m.M
			v263 = m.ExcPending
			if v263 != 0 {
				return int32(0)
			} else {
				v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
				v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v266
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v265
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v264
				F_errmsg(m, int32(_a_F_PageIndexTupleOverwrite_0), v20)
				mBase = m.M
				v272 = m.ExcPending
				if v272 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_PageIndexTupleOverwrite_1), int32(1426), int32(_a_F_PageIndexTupleOverwrite_2))
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
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
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
		if base.Ui32(v25) < base.Ui32(v22) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v260 = m.ExcPending
			if v260 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16779816))
				mBase = m.M
				v263 = m.ExcPending
				if v263 != 0 {
					return int32(0)
				} else {
					v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
					v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
					*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v266
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v265
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v264
					F_errmsg(m, int32(_a_F_PageIndexTupleOverwrite_0), v20)
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_PageIndexTupleOverwrite_1), int32(1426), int32(_a_F_PageIndexTupleOverwrite_2))
						mBase = m.M
						v277 = m.ExcPending
						if v277 != 0 {
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
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			if base.B2i32(base.Ui32(v27) < base.Ui32(v25))|base.B2i32(base.Ui32(int32(_a_F_PageIndexTupleOverwrite_3)) < base.Ui32(v27))|base.B2i32((v27+int32(7))&int32(_a_F_PageIndexTupleOverwrite_4) != v27) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v260 = m.ExcPending
				if v260 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
						return int32(0)
					} else {
						v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
						v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
						v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
						*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v266
						*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v265
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v264
						F_errmsg(m, int32(_a_F_PageIndexTupleOverwrite_0), v20)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_PageIndexTupleOverwrite_1), int32(1426), int32(_a_F_PageIndexTupleOverwrite_2))
							mBase = m.M
							v277 = m.ExcPending
							if v277 != 0 {
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
				if v22 != int32(24) {
					v45 = int32(base.Ui32(v22+int32(_a_F_PageIndexTupleOverwrite_5)) >> (uint(int32(2)) % 32))
				} else {
					v45 = int32(0)
				}
				v46 = int32(_a_F_PageIndexTupleOverwrite_6)
				v47 = v45 & v46
				if base.Ui32(v47) <= base.Ui32((l1-int32(1))&v46) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v281 = m.ExcPending
					if v281 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = l1
						F_errmsg_internal(m, int32(_a_F_PageIndexTupleOverwrite_7), v20+int32(32))
						mBase = m.M
						v287 = m.ExcPending
						if v287 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_PageIndexTupleOverwrite_1), int32(1430), int32(_a_F_PageIndexTupleOverwrite_2))
							mBase = m.M
							v292 = m.ExcPending
							if v292 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v54 = l0 + int32(20)
					v57 = v54 + l1<<(uint(int32(2))%32)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v60 = int32(base.Ui32(v58) >> (uint(int32(17)) % 32))
					v62 = v58 & int32(_a_F_PageIndexTupleOverwrite_8)
					if base.B2i32(base.Ui32(v62) < base.Ui32(v25))|base.B2i32(base.Ui32(v27) < base.Ui32(v62+v60))|base.B2i32(v62 != (v62+int32(7))&int32(_a_F_PageIndexTupleOverwrite_9)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v296 = m.ExcPending
						if v296 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16779816))
							mBase = m.M
							v299 = m.ExcPending
							if v299 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v60
								*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v62
								F_errmsg(m, int32(_a_F_PageIndexTupleOverwrite_10), v20+int32(16))
								mBase = m.M
								v306 = m.ExcPending
								if v306 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_PageIndexTupleOverwrite_1), int32(1442), int32(_a_F_PageIndexTupleOverwrite_2))
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
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
						v73 = int32(7)
						v76 = (l3 + v73) & int32(-8)
						v80 = (v60 + v73) & int32(_a_F_PageIndexTupleOverwrite_9)
						v82 = v80 + (v25 - v22)
						if base.Ui32(v82) < base.Ui32(v76) {
						} else {
							v84 = v80 - v76
							if v80 == v76 {
							} else {
								v86 = v62 - v25
								if v86 != 0 {
									v87 = l0 + v25
									base.MemoryCopy(m, v87+v84, v87, v86)
								} else {
								}
								v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
								v92 = v91 + v84
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v92)
								v94 = int32(1)
								if v45&int32(_a_F_PageIndexTupleOverwrite_6) != v94 {
									v109 = v94
									v114 = int32(0)
									for {
										v123 = v54 + v109<<(uint(int32(2))%32)
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
										if base.B2i32(base.Ui32(v124) < base.Ui32(int32(_a_F_PageIndexTupleOverwrite_11)))|base.B2i32(base.Ui32(v62) < base.Ui32(v124&int32(_a_F_PageIndexTupleOverwrite_8))) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v123))) = (v124+v84)&int32(_a_F_PageIndexTupleOverwrite_8) | v124&int32(-32768)
										} else {
										}
										v141 = v123 + int32(4)
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
										if base.B2i32(base.Ui32(v142) < base.Ui32(int32(_a_F_PageIndexTupleOverwrite_11)))|base.B2i32(base.Ui32(v62) < base.Ui32(v142&int32(_a_F_PageIndexTupleOverwrite_8))) == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v141))) = (v142+v84)&int32(_a_F_PageIndexTupleOverwrite_8) | v142&int32(-32768)
										} else {
										}
										v158 = int32(2)
										v159 = v109 + v158
										v161 = v114 + v158
										if v161 != v47&int32(_a_F_PageIndexTupleOverwrite_12) {
											v109 = v159
											v114 = v161
											continue
										} else {
											break
										}
										break
									}
									if v47&int32(1) == int32(0) {
									} else {
										v170 = v159
										v184 = v54 + v170<<(uint(int32(2))%32)
										v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
										if base.B2i32(base.Ui32(v185) < base.Ui32(int32(_a_F_PageIndexTupleOverwrite_11)))|base.B2i32(base.Ui32(v62) < base.Ui32(v185&int32(_a_F_PageIndexTupleOverwrite_8))) != 0 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v184))) = (v185+v84)&int32(_a_F_PageIndexTupleOverwrite_8) | v185&int32(-32768)
										}
									}
								} else {
									v170 = v94
									v184 = v54 + v170<<(uint(int32(2))%32)
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
									if base.B2i32(base.Ui32(v185) < base.Ui32(int32(_a_F_PageIndexTupleOverwrite_11)))|base.B2i32(base.Ui32(v62) < base.Ui32(v185&int32(_a_F_PageIndexTupleOverwrite_8))) != 0 {
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v184))) = (v185+v84)&int32(_a_F_PageIndexTupleOverwrite_8) | v185&int32(-32768)
									}
								}
							}
							v218 = (v84 + v58) & int32(_a_F_PageIndexTupleOverwrite_8)
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = v218 | (v219&int32(_a_F_PageIndexTupleOverwrite_13) | l3<<(uint(int32(17))%32))
							if l3 == int32(0) {
							} else {
								base.MemoryCopy(m, l0+v218, l2, l3)
							}
						}
						m.G0 = v20 + int32(48)
						return base.B2i32(base.Ui32(v76) <= base.Ui32(v82))
					}
				}
			}
		}
	}
}
func F_RecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = base.I32_div_u_s(l1, int32(4069))
	v15 = base.I64_extend_i32_u(v12) << (uint(int64(32)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v15
	v19 = F_fsm_readbuf(m, l0, v9, int32(1))
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
	F_LockBuffer(m, v19, int32(2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(_a_F_RecordPageWithFreeSpace_0)) < base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = int32(-1)
	goto L6
L5:
	;
	v29 = int32(base.Ui32(l2) >> (uint(int32(5)) % 32))
	goto L6
L6:
	;
	if v19 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v52 = v29 & int32(255)
	v58 = v50 + int32(28)
	v60 = l1 - v12*int32(4069) + int32(4095)
	v61 = v58 + v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v62 != v52 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RecordPageWithFreeSpace[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v19^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L7
L9:
	;
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_RecordPageWithFreeSpace[1]))
	v50 = v44 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L11:
	;
	if v154 != 0 {
		goto L42
	} else {
		goto L43
	}
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v52)
	v71 = v60
	goto L15
L13:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if base.Ui32(v64) < base.Ui32(v52) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v154 = int32(0)
	goto L11
L15:
	;
	v75 = int32(1)
	v76 = v71 - v75
	v77 = int32(2)
	v78 = base.I32_div_s(v76, v77)
	v80 = v78 << (uint(v75) % 32)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v80)+1)))
	v84 = v80 + v77
	if base.Ui32(v84) <= base.Ui32(int32(_a_F_RecordPageWithFreeSpace_1)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if base.Ui32(v103) < base.Ui32(v52) {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v88 = v82 & int32(255)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v84))))
	if base.Ui32(v90) < base.Ui32(v88) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v93 = v82
	goto L19
L19:
	;
	v95 = v78 + v58
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v96 != v93&int32(255) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v92 = v88
	goto L22
L21:
	;
	v92 = v90
	goto L22
L22:
	;
	v93 = v92
	goto L19
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v93)
	if int32(1) < v76 {
		v71 = v78
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L16
L26:
	;
	goto L25
L27:
	;
	v109 = int32(4094)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v154 = int32(1)
	goto L11
L30:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v109) {
		v130 = int32(0)
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v131 = v109 + v58
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v132 != v130&int32(255) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v119 = v109 << (uint(int32(1)) % 32)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+int32(29)+v119))))
	if v109 == int32(4081) {
		v130 = v121
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v58)+2)))
	if base.Ui32(v125) < base.Ui32(v121) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v127 = v121
	goto L37
L36:
	;
	v127 = v125
	goto L37
L37:
	;
	v130 = v127
	goto L32
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v130)
	goto L40
L39:
	;
	goto L40
L40:
	;
	if v109 != 0 {
		v109 = v109 - int32(1)
		goto L30
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	F_MarkBufferDirtyHint(m, v19, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_UnlockReleaseBuffer(m, v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	m.G0 = v9 + int32(16)
	return
}

package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_spgFormNodeTuple(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 != 0 {
		v62 = int32(8)
		v64 = F_palloc0(m, v62)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int32(0)
		} else {
			v68 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v68)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
			if l2 != 0 {
				v74 = v62 | int32(_a_F_spgFormNodeTuple_0)
			} else {
				v74 = v62
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v74)
			if l2 != 0 {
			} else {
				v77 = v64 + int32(8)
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
				if v78 == int32(1) {
					*(*int64)(unsafe.Add(mBase, uint32(v77))) = l1
				} else {
					v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
					if int32(0) < v82 {
						v111 = base.I32_wrap_i64(l1)
						v112 = v82
					} else {
						v86 = base.I32_wrap_i64(l1)
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
						if v87 == int32(1) {
							v91 = int32(18)
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
							if v93 == v91 {
								v96 = v91
							} else {
								v96 = int32(2)
							}
							if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v103 = int32(6)
							} else {
								v103 = v96
							}
							v111 = v86
							v112 = v103
						} else {
							if v87&int32(1) != 0 {
								v111 = v86
								v112 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
								v111 = v86
								v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
							}
						}
					}
					if v112 == int32(0) {
					} else {
						base.MemoryCopy(m, v77, v111, v112)
					}
				}
			}
			m.G0 = v9 + int32(16)
			return v64
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
		if v12 != 0 {
			v62 = int32(16)
			v64 = F_palloc0(m, v62)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				v68 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v68)
				*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
				if l2 != 0 {
					v74 = v62 | int32(_a_F_spgFormNodeTuple_0)
				} else {
					v74 = v62
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v74)
				if l2 != 0 {
				} else {
					v77 = v64 + int32(8)
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
					if v78 == int32(1) {
						*(*int64)(unsafe.Add(mBase, uint32(v77))) = l1
					} else {
						v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
						if int32(0) < v82 {
							v111 = base.I32_wrap_i64(l1)
							v112 = v82
						} else {
							v86 = base.I32_wrap_i64(l1)
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
							if v87 == int32(1) {
								v91 = int32(18)
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
								if v93 == v91 {
									v96 = v91
								} else {
									v96 = int32(2)
								}
								if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v103 = int32(6)
								} else {
									v103 = v96
								}
								v111 = v86
								v112 = v103
							} else {
								if v87&int32(1) != 0 {
									v111 = v86
									v112 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
									v111 = v86
									v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
								}
							}
						}
						if v112 == int32(0) {
						} else {
							base.MemoryCopy(m, v77, v111, v112)
						}
					}
				}
				m.G0 = v9 + int32(16)
				return v64
			}
		} else {
			v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
			if v14 <= int32(0) {
				v17 = base.I32_wrap_i64(l1)
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v18 == int32(1) {
					v22 = int32(18)
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					if v24 == v22 {
						v27 = v22
					} else {
						v27 = int32(2)
					}
					if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v34 = int32(6)
					} else {
						v34 = v27
					}
					v42 = v34
					v62 = (v42+int32(7))&int32(248) + int32(8)
					v64 = F_palloc0(m, v62)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v68)
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
						if l2 != 0 {
							v74 = v62 | int32(_a_F_spgFormNodeTuple_0)
						} else {
							v74 = v62
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v74)
						if l2 != 0 {
						} else {
							v77 = v64 + int32(8)
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
							if v78 == int32(1) {
								*(*int64)(unsafe.Add(mBase, uint32(v77))) = l1
							} else {
								v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
								if int32(0) < v82 {
									v111 = base.I32_wrap_i64(l1)
									v112 = v82
								} else {
									v86 = base.I32_wrap_i64(l1)
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
									if v87 == int32(1) {
										v91 = int32(18)
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
										if v93 == v91 {
											v96 = v91
										} else {
											v96 = int32(2)
										}
										if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v103 = int32(6)
										} else {
											v103 = v96
										}
										v111 = v86
										v112 = v103
									} else {
										if v87&int32(1) != 0 {
											v111 = v86
											v112 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v111 = v86
											v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
										}
									}
								}
								if v112 == int32(0) {
								} else {
									base.MemoryCopy(m, v77, v111, v112)
								}
							}
						}
						m.G0 = v9 + int32(16)
						return v64
					}
				} else {
					if v18&int32(1) == int32(0) {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v53 = int32(base.Ui32(v49) >> (uint(int32(2)) % 32))
						v59 = (v53+int32(7))&int32(2147483640) + int32(8)
						if base.Ui32(int32(_a_F_spgFormNodeTuple_1)) <= base.Ui32(v53) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_spgFormNodeTuple_2)
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
									F_errmsg(m, int32(_a_F_spgFormNodeTuple_3), v9)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_spgFormNodeTuple_4), int32(978), int32(_a_F_spgFormNodeTuple_5))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
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
							v62 = v59
							v64 = F_palloc0(m, v62)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v68)
								*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
								if l2 != 0 {
									v74 = v62 | int32(_a_F_spgFormNodeTuple_0)
								} else {
									v74 = v62
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v74)
								if l2 != 0 {
								} else {
									v77 = v64 + int32(8)
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
									if v78 == int32(1) {
										*(*int64)(unsafe.Add(mBase, uint32(v77))) = l1
									} else {
										v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
										if int32(0) < v82 {
											v111 = base.I32_wrap_i64(l1)
											v112 = v82
										} else {
											v86 = base.I32_wrap_i64(l1)
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
											if v87 == int32(1) {
												v91 = int32(18)
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
												if v93 == v91 {
													v96 = v91
												} else {
													v96 = int32(2)
												}
												if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
													v103 = int32(6)
												} else {
													v103 = v96
												}
												v111 = v86
												v112 = v103
											} else {
												if v87&int32(1) != 0 {
													v111 = v86
													v112 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
												} else {
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
													v111 = v86
													v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
												}
											}
										}
										if v112 == int32(0) {
										} else {
											base.MemoryCopy(m, v77, v111, v112)
										}
									}
								}
								m.G0 = v9 + int32(16)
								return v64
							}
						}
					} else {
						v42 = int32(base.Ui32(v18) >> (uint(int32(1)) % 32))
						v62 = (v42+int32(7))&int32(248) + int32(8)
						v64 = F_palloc0(m, v62)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v68)
							*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
							if l2 != 0 {
								v74 = v62 | int32(_a_F_spgFormNodeTuple_0)
							} else {
								v74 = v62
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v74)
							if l2 != 0 {
							} else {
								v77 = v64 + int32(8)
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
								if v78 == int32(1) {
									*(*int64)(unsafe.Add(mBase, uint32(v77))) = l1
								} else {
									v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
									if int32(0) < v82 {
										v111 = base.I32_wrap_i64(l1)
										v112 = v82
									} else {
										v86 = base.I32_wrap_i64(l1)
										v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
										if v87 == int32(1) {
											v91 = int32(18)
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
											if v93 == v91 {
												v96 = v91
											} else {
												v96 = int32(2)
											}
											if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v103 = int32(6)
											} else {
												v103 = v96
											}
											v111 = v86
											v112 = v103
										} else {
											if v87&int32(1) != 0 {
												v111 = v86
												v112 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
											} else {
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
												v111 = v86
												v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
											}
										}
									}
									if v112 == int32(0) {
									} else {
										base.MemoryCopy(m, v77, v111, v112)
									}
								}
							}
							m.G0 = v9 + int32(16)
							return v64
						}
					}
				}
			} else {
				v53 = v14
				v59 = (v53+int32(7))&int32(2147483640) + int32(8)
				if base.Ui32(int32(_a_F_spgFormNodeTuple_1)) <= base.Ui32(v53) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_spgFormNodeTuple_2)
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v59
							F_errmsg(m, int32(_a_F_spgFormNodeTuple_3), v9)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_spgFormNodeTuple_4), int32(978), int32(_a_F_spgFormNodeTuple_5))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
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
					v62 = v59
					v64 = F_palloc0(m, v62)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)) = uint16(v68)
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(-1)
						if l2 != 0 {
							v74 = v62 | int32(_a_F_spgFormNodeTuple_0)
						} else {
							v74 = v62
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)) = uint16(v74)
						if l2 != 0 {
						} else {
							v77 = v64 + int32(8)
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
							if v78 == int32(1) {
								*(*int64)(unsafe.Add(mBase, uint32(v77))) = l1
							} else {
								v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+60)))
								if int32(0) < v82 {
									v111 = base.I32_wrap_i64(l1)
									v112 = v82
								} else {
									v86 = base.I32_wrap_i64(l1)
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
									if v87 == int32(1) {
										v91 = int32(18)
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
										if v93 == v91 {
											v96 = v91
										} else {
											v96 = int32(2)
										}
										if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v103 = int32(6)
										} else {
											v103 = v96
										}
										v111 = v86
										v112 = v103
									} else {
										if v87&int32(1) != 0 {
											v111 = v86
											v112 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v111 = v86
											v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
										}
									}
								}
								if v112 == int32(0) {
								} else {
									base.MemoryCopy(m, v77, v111, v112)
								}
							}
						}
						m.G0 = v9 + int32(16)
						return v64
					}
				}
			}
		}
	}
}
func F_spg_box_quad_inner_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v128 float64
	_ = v128
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v140 float64
	_ = v140
	var v146 float64
	_ = v146
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v163 int32
	_ = v163
	var v164 float64
	_ = v164
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v170 float64
	_ = v170
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v186 float64
	_ = v186
	var v194 float64
	_ = v194
	var v199 float64
	_ = v199
	var v200 float64
	_ = v200
	var v201 float64
	_ = v201
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v222 float64
	_ = v222
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 float64
	_ = v298
	var v300 float64
	_ = v300
	var v302 float64
	_ = v302
	var v304 float64
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 float64
	_ = v360
	var v362 float64
	_ = v362
	var v364 float64
	_ = v364
	var v366 float64
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v443 int64
	_ = v443
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v455 float64
	_ = v455
	var v464 int32
	_ = v464
	var v466 float64
	_ = v466
	var v472 int32
	_ = v472
	var v474 float64
	_ = v474
	var v480 int32
	_ = v480
	var v482 float64
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 float64
	_ = v515
	var v516 float64
	_ = v516
	var v522 float64
	_ = v522
	var v523 float64
	_ = v523
	var v529 float64
	_ = v529
	var v530 float64
	_ = v530
	var v536 float64
	_ = v536
	var v537 float64
	_ = v537
	var v546 int32
	_ = v546
	var v547 float64
	_ = v547
	var v548 float64
	_ = v548
	var v554 float64
	_ = v554
	var v555 float64
	_ = v555
	var v561 float64
	_ = v561
	var v562 float64
	_ = v562
	var v568 float64
	_ = v568
	var v569 float64
	_ = v569
	var v578 int32
	_ = v578
	var v579 float64
	_ = v579
	var v581 float64
	_ = v581
	var v582 float64
	_ = v582
	var v586 float64
	_ = v586
	var v587 float64
	_ = v587
	var v593 float64
	_ = v593
	var v597 float64
	_ = v597
	var v603 float64
	_ = v603
	var v605 float64
	_ = v605
	var v606 float64
	_ = v606
	var v610 float64
	_ = v610
	var v611 float64
	_ = v611
	var v617 float64
	_ = v617
	var v621 float64
	_ = v621
	var v630 int32
	_ = v630
	var v631 float64
	_ = v631
	var v632 float64
	_ = v632
	var v638 float64
	_ = v638
	var v647 int32
	_ = v647
	var v648 float64
	_ = v648
	var v650 float64
	_ = v650
	var v651 float64
	_ = v651
	var v655 float64
	_ = v655
	var v662 int32
	_ = v662
	var v663 float64
	_ = v663
	var v665 float64
	_ = v665
	var v666 float64
	_ = v666
	var v670 float64
	_ = v670
	var v677 int32
	_ = v677
	var v678 float64
	_ = v678
	var v679 float64
	_ = v679
	var v685 float64
	_ = v685
	var v694 int32
	_ = v694
	var v695 float64
	_ = v695
	var v697 float64
	_ = v697
	var v698 float64
	_ = v698
	var v702 float64
	_ = v702
	var v709 int32
	_ = v709
	var v710 float64
	_ = v710
	var v711 float64
	_ = v711
	var v717 float64
	_ = v717
	var v726 int32
	_ = v726
	var v727 float64
	_ = v727
	var v728 float64
	_ = v728
	var v734 float64
	_ = v734
	var v741 int32
	_ = v741
	var v742 float64
	_ = v742
	var v744 float64
	_ = v744
	var v745 float64
	_ = v745
	var v749 float64
	_ = v749
	var v757 int32
	_ = v757
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 float64
	_ = v827
	var v828 float64
	_ = v828
	var v832 float64
	_ = v832
	var v838 float64
	_ = v838
	var v839 float64
	_ = v839
	var v840 float64
	_ = v840
	var v844 float64
	_ = v844
	var v850 float64
	_ = v850
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 float64
	_ = v863
	var v864 float64
	_ = v864
	var v867 int32
	_ = v867
	var v868 float64
	_ = v868
	var v869 int64
	_ = v869
	var v871 int64
	_ = v871
	var v874 float64
	_ = v874
	var v877 int64
	_ = v877
	var v879 int64
	_ = v879
	var v890 float64
	_ = v890
	var v898 float64
	_ = v898
	var v903 float64
	_ = v903
	var v904 float64
	_ = v904
	var v905 float64
	_ = v905
	var v914 float64
	_ = v914
	var v915 float64
	_ = v915
	var v917 float64
	_ = v917
	var v919 float64
	_ = v919
	var v926 float64
	_ = v926
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v23 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = F_palloc(m, int32(64))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v47 = v23
	goto L3
L3:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)))
	if v48 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int64(0)
L5:
	;
	v31 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+56)) = v31
	v33 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v33
	v47 = v27
	goto L3
L6:
	;
	m.G0 = v19 + int32(32)
	return int64(0)
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v51
	v54 = F_palloc_mul(m, int32(4), v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v296 = F_palloc(m, int32(32))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L63
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v57 <= int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v61 = int32(0)
	goto L12
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v77+v61<<(uint(int32(2))%32)))) = v61
	v83 = v61 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v83 < v84 {
		v61 = v83
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v84 <= int32(0) {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v88 <= int32(0) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v92 = F_palloc_mul(m, int32(8), v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v95 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v99 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v251 = F_palloc_mul(m, int32(4), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L54
	}
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v99*int32(56))+48))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*float64)(unsafe.Add(mBase, uint32(v47)))
	if base.F64_lt(v123, v124) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L20
L23:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v122)+8))
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v47)+32))
	if base.F64_lt(v135, v136) != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v134 = base.F64_sub(v124, v123)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v47)+24))
	if base.F64_gt(v123, v128) == int32(0) {
		v134 = float64(0)
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v134 = base.F64_sub(v123, v128)
	goto L23
L28:
	;
	v155 = m.G0
	v157 = v155 - int32(32)
	m.G0 = v157
	v159 = base.F64_abs(v134)
	v160 = base.F64_abs(v146)
	v163 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v159)) < base.Ui64(base.I64_reinterpret_f64(v160)))
	if base.Ui64(base.I64_reinterpret_f64(v159)) < base.Ui64(base.I64_reinterpret_f64(v160)) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v146 = base.F64_sub(v136, v135)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v47)+56))
	if base.F64_gt(v135, v140) == int32(0) {
		v146 = float64(0)
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v146 = base.F64_sub(v135, v140)
	goto L28
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v92+v99<<(uint(int32(3))%32)))) = v222
	v230 = v99 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v230 < v231 {
		v99 = v230
		goto L21
	} else {
		goto L53
	}
L34:
	;
	m.G0 = v157 + int32(32)
	goto L33
L35:
	;
	v164 = v159
	goto L37
L36:
	;
	v164 = v160
	goto L37
L37:
	;
	v165 = base.I64_reinterpret_f64(v164)
	v167 = int64(base.Ui64(v165) >> (uint(int64(52)) % 64))
	if v167 == int64(2047) {
		v222 = v164
		goto L34
	} else {
		goto L38
	}
L38:
	;
	if base.Ui64(base.I64_reinterpret_f64(v159)) < base.Ui64(base.I64_reinterpret_f64(v160)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v170 = v160
	goto L41
L40:
	;
	v170 = v159
	goto L41
L41:
	;
	if v165 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v222 = v170
	goto L34
L43:
	;
	v173 = base.I64_reinterpret_f64(v170)
	v175 = int64(base.Ui64(v173) >> (uint(int64(52)) % 64))
	if v175 == int64(2047) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if int32(65) <= base.I32_wrap_i64(v175)-base.I32_wrap_i64(v167) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v222 = base.F64_add(v159, v160)
	goto L34
L46:
	;
	goto L47
L47:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v173) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	F_sq(m, v157+int32(24), v157+int32(16), v199)
	mBase = m.M
	F_sq(m, v157+int32(8), v157, v200)
	mBase = m.M
	v210 = *(*float64)(unsafe.Add(mBase, uint32(v157)))
	v211 = *(*float64)(unsafe.Add(mBase, uint32(v157)+16))
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v157)+8))
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v157)+24))
	v222 = base.F64_mul(v201, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v210, v211), v213), v215)))
	goto L34
L49:
	;
	v186 = float64(1.90109156629516e-211)
	v199 = base.F64_mul(v170, v186)
	v200 = base.F64_mul(v164, v186)
	v201 = float64(5.260135901548374e+210)
	goto L48
L50:
	;
	goto L51
L51:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v165) {
		v199 = v170
		v200 = v164
		v201 = float64(1)
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v194 = float64(5.260135901548374e+210)
	v199 = base.F64_mul(v170, v194)
	v200 = base.F64_mul(v164, v194)
	v201 = float64(1.90109156629516e-211)
	goto L48
L53:
	;
	goto L22
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v92
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v255 < int32(2) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v261 = int32(1)
	goto L56
L56:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v276 = F_palloc_mul(m, int32(8), v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	goto L6
L58:
	;
	v279 = v261 << (uint(int32(2)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v279+v280))) = v276
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v285 = v283 << (uint(int32(3)) % 32)
	if v285 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286+v279)))
	base.MemoryCopy(m, v288, v92, v285)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v291 = v261 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v291 < v292 {
		v261 = v291
		goto L56
	} else {
		goto L62
	}
L62:
	;
	goto L57
L63:
	;
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v294)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v296))) = v298
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v294)))
	*(*float64)(unsafe.Add(mBase, uint32(v296)+8)) = v300
	v302 = *(*float64)(unsafe.Add(mBase, uint32(v294)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v296)+16)) = v302
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v294)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v296)+24)) = v304
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v308 = F_palloc_mul(m, int32(4), v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v310 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v319 = v4
	goto L68
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v396 = F_palloc_mul(m, int32(4), v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L80
	}
L68:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v332 = v329 + v319*int32(56)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+8))
	switch v333 - int32(603) {
	case 0:
		goto L71
	case 1:
		goto L73
	default:
		goto L72
	}
L69:
	;
	goto L67
L70:
	;
	v358 = F_palloc(m, int32(32))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L78
	}
L71:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	v356 = v355
	goto L70
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	v337 = F_pg_detoast_datum(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v356 = v337 + int32(8)
	goto L70
L75:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v332)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v345
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_inner_consistent_0), v19)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_inner_consistent_1), int32(544), int32(_a_F_spg_box_quad_inner_consistent_2))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v360 = *(*float64)(unsafe.Add(mBase, uint32(v356)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v358))) = v360
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v356)))
	*(*float64)(unsafe.Add(mBase, uint32(v358)+8)) = v362
	v364 = *(*float64)(unsafe.Add(mBase, uint32(v356)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v358)+16)) = v364
	v366 = *(*float64)(unsafe.Add(mBase, uint32(v356)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v358)+24)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v308+v319<<(uint(int32(2))%32)))) = v358
	v373 = v319 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v373 < v374 {
		v319 = v373
		goto L68
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v396
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v401 = F_palloc_mul(m, int32(4), v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v401
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v404 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v409 = F_palloc_mul(m, int32(4), v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v412 = int32(_a_F_spg_box_quad_inner_consistent_3)
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_spg_box_quad_inner_consistent[0]))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_box_quad_inner_consistent[0])) = v415
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if int32(0) < v417 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v409
	goto L84
L86:
	;
	v430 = v4
	v435 = v4
	goto L89
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_box_quad_inner_consistent[0])) = v413
	goto L6
L89:
	;
	v437 = F_palloc(m, int32(64))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v47)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+56)) = v439
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v47)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+48)) = v441
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v47)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+40)) = v443
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v47)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+32)) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v47)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+24)) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v47)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+16)) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+8)) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v437))) = v453
	v455 = *(*float64)(unsafe.Add(mBase, uint32(v296)))
	if v430&int32(8) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v430&int32(4) != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v437))) = v455
	goto L92
L94:
	;
	goto L95
L95:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v437)+8)) = v455
	goto L92
L96:
	;
	v464 = int32(16)
	goto L98
L97:
	;
	v464 = int32(24)
	goto L98
L98:
	;
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v296)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v437+v464))) = v466
	if v430&int32(2) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v472 = int32(32)
	goto L101
L100:
	;
	v472 = int32(40)
	goto L101
L101:
	;
	v474 = *(*float64)(unsafe.Add(mBase, uint32(v296)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v437+v472))) = v474
	if v430&int32(1) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v480 = int32(48)
	goto L104
L103:
	;
	v480 = int32(56)
	goto L104
L104:
	;
	v482 = *(*float64)(unsafe.Add(mBase, uint32(v296)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v437+v480))) = v482
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v484 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v994 = v435 + int32(1)
	v996 = v994 & int32(255)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v996 < v997 {
		v430 = v996
		v435 = v994
		goto L89
	} else {
		goto L199
	}
L106:
	;
	F_pfree(m, v437)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L4
	} else {
		goto L198
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L4
	} else {
		goto L195
	}
L108:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v492 = int32(0)
	goto L111
L109:
	;
	goto L110
L110:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v777 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v775+v776<<(uint(v777)%32)))) = v437
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v781+v782<<(uint(v777)%32)))) = v430
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v787 <= int32(0) {
		goto L158
	} else {
		goto L159
	}
L111:
	;
	v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487+v492*int32(56))+6)))
	switch v508 - int32(1) {
	case 0:
		goto L121
	case 1:
		goto L120
	case 2:
		goto L124
	case 3:
		goto L118
	case 4:
		goto L119
	case 5, 7:
		goto L122
	case 6:
		goto L123
	case 8:
		goto L114
	case 9:
		goto L115
	case 10:
		goto L117
	case 11:
		goto L116
	default:
		goto L107
	}
L112:
	;
	goto L110
L113:
	;
	v757 = v492 + int32(1)
	if v757 != v484 {
		v492 = v757
		goto L111
	} else {
		goto L157
	}
L114:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v742 = *(*float64)(unsafe.Add(mBase, uint32(v741)+24))
	v744 = base.F64_add(v742, float64(1e-06))
	v745 = *(*float64)(unsafe.Add(mBase, uint32(v437)+32))
	if base.F64_ge(v744, v745) == int32(0) {
		goto L106
	} else {
		goto L155
	}
L115:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v727 = *(*float64)(unsafe.Add(mBase, uint32(v726)+16))
	v728 = *(*float64)(unsafe.Add(mBase, uint32(v437)+32))
	if base.F64_gt(v727, base.F64_add(v728, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L153
	}
L116:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v710 = *(*float64)(unsafe.Add(mBase, uint32(v709)+16))
	v711 = *(*float64)(unsafe.Add(mBase, uint32(v437)+40))
	if base.F64_le(v710, base.F64_add(v711, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L151
	}
L117:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v695 = *(*float64)(unsafe.Add(mBase, uint32(v694)+24))
	v697 = base.F64_add(v695, float64(1e-06))
	v698 = *(*float64)(unsafe.Add(mBase, uint32(v437)+40))
	if base.F64_lt(v697, v698) == int32(0) {
		goto L106
	} else {
		goto L149
	}
L118:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v678 = *(*float64)(unsafe.Add(mBase, uint32(v677)))
	v679 = *(*float64)(unsafe.Add(mBase, uint32(v437)+8))
	if base.F64_le(v678, base.F64_add(v679, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L147
	}
L119:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v663 = *(*float64)(unsafe.Add(mBase, uint32(v662)+8))
	v665 = base.F64_add(v663, float64(1e-06))
	v666 = *(*float64)(unsafe.Add(mBase, uint32(v437)+8))
	if base.F64_lt(v665, v666) == int32(0) {
		goto L106
	} else {
		goto L145
	}
L120:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v648 = *(*float64)(unsafe.Add(mBase, uint32(v647)+8))
	v650 = base.F64_add(v648, float64(1e-06))
	v651 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	if base.F64_ge(v650, v651) == int32(0) {
		goto L106
	} else {
		goto L143
	}
L121:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v630)))
	v632 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	if base.F64_gt(v631, base.F64_add(v632, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L141
	}
L122:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v579 = *(*float64)(unsafe.Add(mBase, uint32(v578)+8))
	v581 = base.F64_add(v579, float64(1e-06))
	v582 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	if base.F64_ge(v581, v582) == int32(0) {
		goto L106
	} else {
		goto L133
	}
L123:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v547 = *(*float64)(unsafe.Add(mBase, uint32(v546)+8))
	v548 = *(*float64)(unsafe.Add(mBase, uint32(v437)+24))
	if base.F64_le(v547, base.F64_add(v548, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L129
	}
L124:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v308+v492<<(uint(int32(2))%32))))
	v515 = *(*float64)(unsafe.Add(mBase, uint32(v514)))
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v437)+24))
	if base.F64_le(v515, base.F64_add(v516, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L125
	}
L125:
	;
	v522 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	v523 = *(*float64)(unsafe.Add(mBase, uint32(v514)+8))
	if base.F64_le(v522, base.F64_add(v523, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L126
	}
L126:
	;
	v529 = *(*float64)(unsafe.Add(mBase, uint32(v514)+16))
	v530 = *(*float64)(unsafe.Add(mBase, uint32(v437)+56))
	if base.F64_le(v529, base.F64_add(v530, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L127
	}
L127:
	;
	v536 = *(*float64)(unsafe.Add(mBase, uint32(v437)+32))
	v537 = *(*float64)(unsafe.Add(mBase, uint32(v514)+24))
	if base.F64_le(v536, base.F64_add(v537, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L128
	}
L128:
	;
	goto L113
L129:
	;
	v554 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	v555 = *(*float64)(unsafe.Add(mBase, uint32(v546)))
	if base.F64_le(v554, base.F64_add(v555, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L130
	}
L130:
	;
	v561 = *(*float64)(unsafe.Add(mBase, uint32(v546)+24))
	v562 = *(*float64)(unsafe.Add(mBase, uint32(v437)+56))
	if base.F64_le(v561, base.F64_add(v562, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L131
	}
L131:
	;
	v568 = *(*float64)(unsafe.Add(mBase, uint32(v437)+32))
	v569 = *(*float64)(unsafe.Add(mBase, uint32(v546)+16))
	if base.F64_le(v568, base.F64_add(v569, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L132
	}
L132:
	;
	goto L113
L133:
	;
	v586 = *(*float64)(unsafe.Add(mBase, uint32(v578)))
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v437)+8))
	if base.F64_le(v586, base.F64_add(v587, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L134
	}
L134:
	;
	v593 = *(*float64)(unsafe.Add(mBase, uint32(v437)+16))
	if base.F64_le(v593, v581) == int32(0) {
		goto L106
	} else {
		goto L135
	}
L135:
	;
	v597 = *(*float64)(unsafe.Add(mBase, uint32(v437)+24))
	if base.F64_ge(base.F64_add(v597, float64(1e-06)), v586) == int32(0) {
		goto L106
	} else {
		goto L136
	}
L136:
	;
	v603 = *(*float64)(unsafe.Add(mBase, uint32(v578)+24))
	v605 = base.F64_add(v603, float64(1e-06))
	v606 = *(*float64)(unsafe.Add(mBase, uint32(v437)+32))
	if base.F64_ge(v605, v606) == int32(0) {
		goto L106
	} else {
		goto L137
	}
L137:
	;
	v610 = *(*float64)(unsafe.Add(mBase, uint32(v578)+16))
	v611 = *(*float64)(unsafe.Add(mBase, uint32(v437)+40))
	if base.F64_le(v610, base.F64_add(v611, float64(1e-06))) == int32(0) {
		goto L106
	} else {
		goto L138
	}
L138:
	;
	v617 = *(*float64)(unsafe.Add(mBase, uint32(v437)+48))
	if base.F64_le(v617, v605) == int32(0) {
		goto L106
	} else {
		goto L139
	}
L139:
	;
	v621 = *(*float64)(unsafe.Add(mBase, uint32(v437)+56))
	if base.F64_ge(base.F64_add(v621, float64(1e-06)), v610) == int32(0) {
		goto L106
	} else {
		goto L140
	}
L140:
	;
	goto L113
L141:
	;
	v638 = *(*float64)(unsafe.Add(mBase, uint32(v437)+16))
	if base.F64_lt(base.F64_add(v638, float64(1e-06)), v631) == int32(0) {
		goto L106
	} else {
		goto L142
	}
L142:
	;
	goto L113
L143:
	;
	v655 = *(*float64)(unsafe.Add(mBase, uint32(v437)+16))
	if base.F64_le(v655, v650) == int32(0) {
		goto L106
	} else {
		goto L144
	}
L144:
	;
	goto L113
L145:
	;
	v670 = *(*float64)(unsafe.Add(mBase, uint32(v437)+24))
	if base.F64_gt(v670, v665) == int32(0) {
		goto L106
	} else {
		goto L146
	}
L146:
	;
	goto L113
L147:
	;
	v685 = *(*float64)(unsafe.Add(mBase, uint32(v437)+24))
	if base.F64_ge(base.F64_add(v685, float64(1e-06)), v678) == int32(0) {
		goto L106
	} else {
		goto L148
	}
L148:
	;
	goto L113
L149:
	;
	v702 = *(*float64)(unsafe.Add(mBase, uint32(v437)+56))
	if base.F64_gt(v702, v697) == int32(0) {
		goto L106
	} else {
		goto L150
	}
L150:
	;
	goto L113
L151:
	;
	v717 = *(*float64)(unsafe.Add(mBase, uint32(v437)+56))
	if base.F64_ge(base.F64_add(v717, float64(1e-06)), v710) == int32(0) {
		goto L106
	} else {
		goto L152
	}
L152:
	;
	goto L113
L153:
	;
	v734 = *(*float64)(unsafe.Add(mBase, uint32(v437)+48))
	if base.F64_lt(base.F64_add(v734, float64(1e-06)), v727) != 0 {
		goto L113
	} else {
		goto L154
	}
L154:
	;
	goto L106
L155:
	;
	v749 = *(*float64)(unsafe.Add(mBase, uint32(v437)+48))
	if base.F64_le(v749, v744) == int32(0) {
		goto L106
	} else {
		goto L156
	}
L156:
	;
	goto L113
L157:
	;
	goto L112
L158:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v953 + int32(1)
	goto L105
L159:
	;
	v791 = F_palloc_mul(m, int32(8), v787)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v793+v794<<(uint(int32(2))%32)))) = v791
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v799 <= int32(0) {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v806 = int32(0)
	goto L162
L162:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v822+v806*int32(56))+48))
	v827 = *(*float64)(unsafe.Add(mBase, uint32(v826)))
	v828 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	if base.F64_lt(v827, v828) != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L158
L164:
	;
	v839 = *(*float64)(unsafe.Add(mBase, uint32(v826)+8))
	v840 = *(*float64)(unsafe.Add(mBase, uint32(v437)+32))
	if base.F64_lt(v839, v840) != 0 {
		goto L170
	} else {
		goto L171
	}
L165:
	;
	v838 = base.F64_sub(v828, v827)
	goto L164
L166:
	;
	goto L167
L167:
	;
	v832 = *(*float64)(unsafe.Add(mBase, uint32(v437)+24))
	if base.F64_gt(v827, v832) == int32(0) {
		v838 = float64(0)
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v838 = base.F64_sub(v827, v832)
	goto L164
L169:
	;
	v859 = m.G0
	v861 = v859 - int32(32)
	m.G0 = v861
	v863 = base.F64_abs(v838)
	v864 = base.F64_abs(v850)
	v867 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v863)) < base.Ui64(base.I64_reinterpret_f64(v864)))
	if base.Ui64(base.I64_reinterpret_f64(v863)) < base.Ui64(base.I64_reinterpret_f64(v864)) {
		goto L176
	} else {
		goto L177
	}
L170:
	;
	v850 = base.F64_sub(v840, v839)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v844 = *(*float64)(unsafe.Add(mBase, uint32(v437)+56))
	if base.F64_gt(v839, v844) == int32(0) {
		v850 = float64(0)
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v850 = base.F64_sub(v839, v844)
	goto L169
L174:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v791+v806<<(uint(int32(3))%32)))) = v926
	v934 = v806 + int32(1)
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v934 < v935 {
		v806 = v934
		goto L162
	} else {
		goto L194
	}
L175:
	;
	m.G0 = v861 + int32(32)
	goto L174
L176:
	;
	v868 = v863
	goto L178
L177:
	;
	v868 = v864
	goto L178
L178:
	;
	v869 = base.I64_reinterpret_f64(v868)
	v871 = int64(base.Ui64(v869) >> (uint(int64(52)) % 64))
	if v871 == int64(2047) {
		v926 = v868
		goto L175
	} else {
		goto L179
	}
L179:
	;
	if base.Ui64(base.I64_reinterpret_f64(v863)) < base.Ui64(base.I64_reinterpret_f64(v864)) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v874 = v864
	goto L182
L181:
	;
	v874 = v863
	goto L182
L182:
	;
	if v869 == int64(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v926 = v874
	goto L175
L184:
	;
	v877 = base.I64_reinterpret_f64(v874)
	v879 = int64(base.Ui64(v877) >> (uint(int64(52)) % 64))
	if v879 == int64(2047) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	if int32(65) <= base.I32_wrap_i64(v879)-base.I32_wrap_i64(v871) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v926 = base.F64_add(v863, v864)
	goto L175
L187:
	;
	goto L188
L188:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v877) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	F_sq(m, v861+int32(24), v861+int32(16), v903)
	mBase = m.M
	F_sq(m, v861+int32(8), v861, v904)
	mBase = m.M
	v914 = *(*float64)(unsafe.Add(mBase, uint32(v861)))
	v915 = *(*float64)(unsafe.Add(mBase, uint32(v861)+16))
	v917 = *(*float64)(unsafe.Add(mBase, uint32(v861)+8))
	v919 = *(*float64)(unsafe.Add(mBase, uint32(v861)+24))
	v926 = base.F64_mul(v905, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v914, v915), v917), v919)))
	goto L175
L190:
	;
	v890 = float64(1.90109156629516e-211)
	v903 = base.F64_mul(v874, v890)
	v904 = base.F64_mul(v868, v890)
	v905 = float64(5.260135901548374e+210)
	goto L189
L191:
	;
	goto L192
L192:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v869) {
		v903 = v874
		v904 = v868
		v905 = float64(1)
		goto L189
	} else {
		goto L193
	}
L193:
	;
	v898 = float64(5.260135901548374e+210)
	v903 = base.F64_mul(v874, v898)
	v904 = base.F64_mul(v868, v898)
	v905 = float64(1.90109156629516e-211)
	goto L189
L194:
	;
	goto L163
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v508
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_inner_consistent_4), v19+int32(16))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_inner_consistent_1), int32(691), int32(_a_F_spg_box_quad_inner_consistent_5))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	goto L105
L199:
	;
	goto L90
}
func F_spg_box_quad_leaf_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int64
	_ = v204
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v4)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+32)))
	if v18 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v14
	goto L3
L2:
	;
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if int32(0) < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(32)
	return v204
L5:
	;
	v32 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	v190 = int64(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v191 <= int32(0) {
		v204 = v190
		goto L4
	} else {
		goto L67
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v36 = v33 + v32*int32(56)
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	switch v38 - int32(603) {
	case 0:
		goto L11
	case 1:
		goto L13
	default:
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v76 = base.I64_extend_i32_u(v75)
	switch v37 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L38
	case 3:
		goto L32
	case 4:
		goto L33
	case 5:
		goto L36
	case 6:
		goto L26
	case 7:
		goto L37
	case 8:
		goto L28
	case 9:
		goto L29
	case 10:
		goto L31
	case 11:
		goto L30
	default:
		goto L27
	}
L11:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v75 = v74
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L22
	}
L13:
	;
	if int32(1)<<(uint(v37)%32)&int32(_a_F_spg_box_quad_leaf_consistent_0) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = base.B2i32(base.Ui32(v37) <= base.Ui32(int32(12)))
	goto L16
L15:
	;
	v48 = int32(0)
	goto L16
L16:
	;
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v51)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v54 = F_pg_detoast_datum(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int64(0)
L21:
	;
	v75 = v54 + int32(8)
	goto L10
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v64
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_leaf_consistent_1), v11)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_leaf_consistent_2), int32(544), int32(_a_F_spg_box_quad_leaf_consistent_3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v179 = v32 + int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v179 < v180 {
		v32 = v179
		goto L8
	} else {
		goto L66
	}
L26:
	;
	v173 = F_DirectFunctionCall2Coll(m, int32(101), int32(0), v14, v76)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L64
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L20
	} else {
		goto L61
	}
L28:
	;
	v151 = F_DirectFunctionCall2Coll(m, int32(104), int32(0), v14, v76)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L20
	} else {
		goto L59
	}
L29:
	;
	v144 = F_DirectFunctionCall2Coll(m, int32(105), int32(0), v14, v76)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L20
	} else {
		goto L57
	}
L30:
	;
	v137 = F_DirectFunctionCall2Coll(m, int32(103), int32(0), v14, v76)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L20
	} else {
		goto L55
	}
L31:
	;
	v130 = F_DirectFunctionCall2Coll(m, int32(102), int32(0), v14, v76)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L53
	}
L32:
	;
	v123 = F_DirectFunctionCall2Coll(m, int32(106), int32(0), v14, v76)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L51
	}
L33:
	;
	v116 = F_DirectFunctionCall2Coll(m, int32(97), int32(0), v14, v76)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L20
	} else {
		goto L49
	}
L34:
	;
	v109 = F_DirectFunctionCall2Coll(m, int32(100), int32(0), v14, v76)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L47
	}
L35:
	;
	v102 = F_DirectFunctionCall2Coll(m, int32(99), int32(0), v14, v76)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L45
	}
L36:
	;
	v95 = F_DirectFunctionCall2Coll(m, int32(119), int32(0), v14, v76)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L20
	} else {
		goto L43
	}
L37:
	;
	v88 = F_DirectFunctionCall2Coll(m, int32(120), int32(0), v14, v76)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L20
	} else {
		goto L41
	}
L38:
	;
	v81 = F_DirectFunctionCall2Coll(m, int32(98), int32(0), v14, v76)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v81 != int64(0) {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v204 = int64(0)
	goto L4
L41:
	;
	if v88 != int64(0) {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v204 = int64(0)
	goto L4
L43:
	;
	if v95 != int64(0) {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v204 = int64(0)
	goto L4
L45:
	;
	if v102 != int64(0) {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v204 = int64(0)
	goto L4
L47:
	;
	if v109 != int64(0) {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	v204 = int64(0)
	goto L4
L49:
	;
	if v116 != int64(0) {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v204 = int64(0)
	goto L4
L51:
	;
	if v123 != int64(0) {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	v204 = int64(0)
	goto L4
L53:
	;
	if v130 != int64(0) {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v204 = int64(0)
	goto L4
L55:
	;
	if v137 != int64(0) {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	v204 = int64(0)
	goto L4
L57:
	;
	if v144 != int64(0) {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	v204 = int64(0)
	goto L4
L59:
	;
	if v151 != int64(0) {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	v204 = int64(0)
	goto L4
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v37
	F_errmsg_internal(m, int32(_a_F_spg_box_quad_leaf_consistent_4), v11+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_spg_box_quad_leaf_consistent_2), int32(831), int32(_a_F_spg_box_quad_leaf_consistent_5))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L20
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
	if v173 != int64(0) {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	v204 = int64(0)
	goto L4
L66:
	;
	goto L9
L67:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+20))
	v197 = F_spg_key_orderbys_distances(m, v14, int32(0), v194, v191)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+9)) = uint8(base.B2i32(v195 == int32(3292)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v197
	v204 = v190
	goto L4
}
func F_spg_kd_choose(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 float64
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
	if v7 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_spg_kd_choose_0), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_spg_kd_choose_1), int32(64), int32(_a_F_spg_kd_choose_2))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v28 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v27))) = v28
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		v39 = *(*float64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v26)+(v31^int32(-1))<<(uint(int32(3))%32)&int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v26 & int64(4294967295)
		*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v28
		*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = base.B2i32(base.F64_gt(v25, v39) == int32(0))
		return int64(0)
	}
}
func F_spg_kd_inner_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 float64
	_ = v71
	var v79 int32
	_ = v79
	var v84 float64
	_ = v84
	var v86 int32
	_ = v86
	var v91 float64
	_ = v91
	var v100 float64
	_ = v100
	var v107 int32
	_ = v107
	var v110 float64
	_ = v110
	var v118 int32
	_ = v118
	var v121 float64
	_ = v121
	var v125 int32
	_ = v125
	var v130 float64
	_ = v130
	var v134 float64
	_ = v134
	var v140 float64
	_ = v140
	var v144 float64
	_ = v144
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 float64
	_ = v246
	var v247 float64
	_ = v247
	var v250 float64
	_ = v250
	var v252 float64
	_ = v252
	var v256 float64
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+37)))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return int64(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v194 = F_palloc_mul(m, int32(4), int32(2))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L51
	}
L3:
	;
	v43 = base.F64_add(v22, float64(1e-06))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v47 = int32(0)
	v52 = int32(6)
	goto L12
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v17)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if int32(0) < v23 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v183 = int32(6)
	goto L2
L8:
	;
	return int64(0)
L9:
	;
	F_errmsg_internal(m, int32(_a_F_spg_kd_inner_consistent_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_spg_kd_inner_consistent_1), int32(175), int32(_a_F_spg_kd_inner_consistent_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v61 = v44 + v47*int32(56)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+6)))
	switch v63 - int32(1) {
	case 0:
		goto L23
	default:
		goto L15
	case 4:
		goto L22
	case 5:
		goto L21
	case 7:
		goto L18
	case 9, 28:
		goto L20
	case 10, 29:
		goto L19
	}
L13:
	;
	v183 = v173
	goto L2
L14:
	;
	v176 = v47 + int32(1)
	if v176 != v23 {
		v47 = v176
		v52 = v173
		goto L12
	} else {
		goto L50
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L47
	}
L16:
	;
	if v152 != 0 {
		v173 = v152
		goto L14
	} else {
		goto L46
	}
L17:
	;
	v152 = v52 & int32(4)
	goto L16
L18:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v125&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v118&int32(1) != 0 {
		v173 = v52
		goto L14
	} else {
		goto L37
	}
L20:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v107&int32(1) != 0 {
		v173 = v52
		goto L14
	} else {
		goto L35
	}
L21:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v86&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v79&int32(1) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L26
	}
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v66&int32(1) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
	if base.F64_lt(base.F64_add(v71, float64(1e-06)), v22) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v152 = v52 & int32(2)
	goto L16
L26:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
	if base.F64_gt(v84, v43) != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v173 = v52
	goto L14
L28:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
	if base.F64_gt(v22, base.F64_add(v91, float64(1e-06))) != 0 {
		v152 = v52 & int32(2)
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v62)+8))
	if base.F64_gt(v22, base.F64_add(v100, float64(1e-06))) != 0 {
		v152 = v52 & int32(2)
		goto L16
	} else {
		goto L33
	}
L31:
	;
	if base.F64_lt(v43, v91) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L32
	}
L32:
	;
	goto L17
L33:
	;
	if base.F64_lt(v43, v100) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L34
	}
L34:
	;
	goto L17
L35:
	;
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v62)+8))
	if base.F64_lt(base.F64_add(v110, float64(1e-06)), v22) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v152 = v52 & int32(2)
	goto L16
L37:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v62)+8))
	if base.F64_gt(v121, v43) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L17
L39:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
	if base.F64_gt(v22, base.F64_add(v130, float64(1e-06))) != 0 {
		v152 = v52 & int32(2)
		goto L16
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v62)+8))
	if base.F64_gt(v22, base.F64_add(v140, float64(1e-06))) != 0 {
		v152 = v52 & int32(2)
		goto L16
	} else {
		goto L44
	}
L42:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v62)+16))
	if base.F64_gt(v134, v43) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L43
	}
L43:
	;
	goto L17
L44:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v62)+24))
	if base.F64_gt(v144, v43) == int32(0) {
		v173 = v52
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L17
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	goto L1
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159+v47*int32(56))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v163
	F_errmsg_internal(m, int32(_a_F_spg_kd_inner_consistent_3), v15)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_spg_kd_inner_consistent_1), int32(249), int32(_a_F_spg_kd_inner_consistent_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	goto L13
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v194
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if int32(0) < v197 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v202 = F_palloc_mul(m, int32(4), v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v183&int32(2) != 0 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v202
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v207 = F_palloc_mul(m, int32(4), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v210 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v256
	goto L54
L58:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+80)) = v22
	v250 = *(*float64)(unsafe.Add(mBase, uint32(v229)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v250
	v252 = *(*float64)(unsafe.Add(mBase, uint32(v229)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = v252
	v256 = v22
	goto L57
L59:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v22
	*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = v22
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v242)))
	v247 = *(*float64)(unsafe.Add(mBase, uint32(v243)))
	*(*float64)(unsafe.Add(mBase, uint32(v15)+80)) = v247
	v256 = v246
	goto L57
L60:
	;
	v213 = int64(-4503599627370496)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v213
	v215 = int64(9218868437227405312)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v215
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v215
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v215
	v242 = v15 + int32(24)
	v243 = v15 + int32(16)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v229)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v230
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v229)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v232
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v234
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v229)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v236
	if v210&int32(1) != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v242 = v229
	v243 = v229 + int32(16)
	goto L59
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v268 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v263+v264<<(uint(int32(2))%32)))) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v268 < v270 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if v183&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v273 = int32(_a_F_spg_kd_inner_consistent_4)
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v276
	v280 = F_box_copy(m, v15+int32(32))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v304 + int32(1)
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v274
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(int32(2))%32)))) = v280
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v294 = F_spg_key_orderbys_distances(m, base.I64_extend_i32_u(v280), int32(0), v292, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v296+v297<<(uint(int32(2))%32)))) = v294
	goto L69
L72:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v312+v313<<(uint(int32(2))%32)))) = int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if int32(0) < v319 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v361 = F_palloc_mul(m, int32(4), int32(2))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L8
	} else {
		goto L80
	}
L75:
	;
	v322 = int32(_a_F_spg_kd_inner_consistent_4)
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v325
	v329 = F_box_copy(m, v15-int32(-64))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L8
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v353 + int32(1)
	goto L74
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spg_kd_inner_consistent[0])) = v323
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v333+v334<<(uint(int32(2))%32)))) = v329
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v343 = F_spg_key_orderbys_distances(m, base.I64_extend_i32_u(v329), int32(0), v341, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v346<<(uint(int32(2))%32)))) = v343
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v361
	v364 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v364
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = v364
	goto L1
}
func F_spg_key_orderbys_distances(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v78 float64
	_ = v78
	var v84 float64
	_ = v84
	var v88 float64
	_ = v88
	var v94 float64
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v118 float64
	_ = v118
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v134 float64
	_ = v134
	var v142 float64
	_ = v142
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v170 float64
	_ = v170
	var v177 float64
	_ = v177
	var v188 int32
	_ = v188
	v17 = F_palloc_mul(m, int32(8), l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < l3 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = base.I32_wrap_i64(l0)
	v28 = l2
	v36 = v17
	v39 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	return v17
L6:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v36))) = v177
	v188 = v39 + int32(1)
	if v188 != l3 {
		v28 = v28 + int32(56)
		v36 = v36 + int32(8)
		v39 = v188
		goto L6
	} else {
		goto L43
	}
L9:
	;
	v46 = F_DirectFunctionCall2Coll(m, int32(114), int32(0), v41&int64(4294967295), l0&int64(4294967295))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v49 = math.Float64frombits(uint64(0x7ff8000000000000))
	v50 = base.I32_wrap_i64(v41)
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v50)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v51)&int64(9223372036854775807)) {
		v177 = v49
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v177 = base.F64_reinterpret_i64(v46)
	goto L8
L13:
	;
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v25)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) {
		v177 = v49
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v50)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v63)&int64(9223372036854775807)) {
		v177 = v49
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v25)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v69)&int64(9223372036854775807)) {
		v177 = v49
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if base.F64_lt(v51, v57) != 0 {
		v84 = base.F64_sub(v57, v51)
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if base.F64_lt(v63, v69) != 0 {
		v94 = base.F64_sub(v69, v63)
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v78 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
	if base.F64_gt(v51, v78) == int32(0) {
		v84 = float64(0)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v84 = base.F64_sub(v51, v78)
	goto L17
L20:
	;
	v103 = m.G0
	v105 = v103 - int32(32)
	m.G0 = v105
	v107 = base.F64_abs(v84)
	v108 = base.F64_abs(v94)
	v111 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v107)) < base.Ui64(base.I64_reinterpret_f64(v108)))
	if base.Ui64(base.I64_reinterpret_f64(v107)) < base.Ui64(base.I64_reinterpret_f64(v108)) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v88 = *(*float64)(unsafe.Add(mBase, uint32(v25)+8))
	if base.F64_gt(v63, v88) == int32(0) {
		v94 = float64(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v94 = base.F64_sub(v63, v88)
	goto L20
L23:
	;
	v177 = v170
	goto L8
L24:
	;
	m.G0 = v105 + int32(32)
	goto L23
L25:
	;
	v112 = v107
	goto L27
L26:
	;
	v112 = v108
	goto L27
L27:
	;
	v113 = base.I64_reinterpret_f64(v112)
	v115 = int64(base.Ui64(v113) >> (uint(int64(52)) % 64))
	if v115 == int64(2047) {
		v170 = v112
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if base.Ui64(base.I64_reinterpret_f64(v107)) < base.Ui64(base.I64_reinterpret_f64(v108)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v118 = v108
	goto L31
L30:
	;
	v118 = v107
	goto L31
L31:
	;
	if v113 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v170 = v118
	goto L24
L33:
	;
	v121 = base.I64_reinterpret_f64(v118)
	v123 = int64(base.Ui64(v121) >> (uint(int64(52)) % 64))
	if v123 == int64(2047) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if int32(65) <= base.I32_wrap_i64(v123)-base.I32_wrap_i64(v115) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v170 = base.F64_add(v107, v108)
	goto L24
L36:
	;
	goto L37
L37:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v121) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	F_sq(m, v105+int32(24), v105+int32(16), v147)
	mBase = m.M
	F_sq(m, v105+int32(8), v105, v148)
	mBase = m.M
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v105)))
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v105)+16))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v105)+8))
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v105)+24))
	v170 = base.F64_mul(v149, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v158, v159), v161), v163)))
	goto L24
L39:
	;
	v134 = float64(1.90109156629516e-211)
	v147 = base.F64_mul(v118, v134)
	v148 = base.F64_mul(v112, v134)
	v149 = float64(5.260135901548374e+210)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v113) {
		v147 = v118
		v148 = v112
		v149 = float64(1)
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v142 = float64(5.260135901548374e+210)
	v147 = base.F64_mul(v118, v142)
	v148 = base.F64_mul(v112, v142)
	v149 = float64(1.90109156629516e-211)
	goto L38
L43:
	;
	goto L7
}
func F_spg_range_quad_leaf_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
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
	var v87 int32
	_ = v87
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
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int64
	_ = v132
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v20)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v25 = F_range_get_typcache(m, l0, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int64(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v28 <= int32(0) {
		v132 = v27
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v11 + int32(16)
	return v132
L5:
	;
	v36 = int32(0)
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v42 = v39 + v36*int32(56)
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+48))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+6)))
	switch v44 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L9
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	case 5:
		goto L15
	case 6:
		goto L14
	case 7:
		goto L13
	default:
		goto L10
	case 15:
		goto L12
	case 17:
		goto L11
	}
L7:
	;
	v132 = v27
	goto L4
L8:
	;
	v123 = v36 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v123 < v124 {
		v36 = v123
		goto L6
	} else {
		goto L52
	}
L9:
	;
	v117 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L49
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L46
	}
L11:
	;
	v93 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L43
	}
L12:
	;
	v89 = F_range_contains_elem_internal(m, v25, v16, v43)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L41
	}
L13:
	;
	v84 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L38
	}
L14:
	;
	v78 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L35
	}
L15:
	;
	v72 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L32
	}
L16:
	;
	v66 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v60 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v54 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v48 = F_pg_detoast_datum(m, base.I32_wrap_i64(v43))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v50 = F_range_before_internal(m, v25, v16, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v50 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v132 = int64(0)
	goto L4
L23:
	;
	v56 = F_range_overlaps_internal(m, v25, v16, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v56 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v132 = int64(0)
	goto L4
L26:
	;
	v62 = F_range_overright_internal(m, v25, v16, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v62 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v132 = int64(0)
	goto L4
L29:
	;
	v68 = F_range_after_internal(m, v25, v16, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v68 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v132 = int64(0)
	goto L4
L32:
	;
	v74 = F_range_adjacent_internal(m, v25, v16, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v74 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v132 = int64(0)
	goto L4
L35:
	;
	v80 = F_range_contains_internal(m, v25, v16, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v80 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v132 = int64(0)
	goto L4
L38:
	;
	v86 = F_range_contained_by_internal(m, v25, v16, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v86 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v132 = int64(0)
	goto L4
L41:
	;
	if v89 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v132 = int64(0)
	goto L4
L43:
	;
	v95 = F_range_eq_internal(m, v25, v16, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v95 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v132 = int64(0)
	goto L4
L46:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v36*int32(56))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v106
	F_errmsg_internal(m, int32(_a_F_spg_range_quad_leaf_consistent_0), v11)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_spg_range_quad_leaf_consistent_1), int32(987), int32(_a_F_spg_range_quad_leaf_consistent_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v119 = F_range_overleft_internal(m, v25, v16, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v119 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v132 = int64(0)
	goto L4
L52:
	;
	goto L7
}
func F_spg_range_quad_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v26 = F_range_get_typcache(m, l0, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v30 = F_palloc_mul(m, int32(16), v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v34 = F_palloc_mul(m, int32(16), v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v15 + int32(80)
	return int64(0)
L7:
	;
	F_qsort_arg(m, v30, v70, int32(16), int32(1685), v26)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	v40 = int32(0)
	v45 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v89)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = int64(2)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v96 = F_palloc_mul(m, int32(4), v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v45<<(uint(int32(3))%32))))
	v57 = F_pg_detoast_datum(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	if v70 != 0 {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v60 = v40 << (uint(int32(4)) % 32)
	F_range_deserialize(m, v26, v57, v30+v60, v60+v34, v15+int32(6))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+6)))
	v68 = int32(1)
	v70 = v40 + (v67 ^ v68)
	v72 = v45 + v68
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v72 < v73 {
		v40 = v70
		v45 = v72
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v96
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v101 = F_palloc_mul(m, int32(8), v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v101
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v104 <= int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v107 = v89
	goto L20
L20:
	;
	v120 = v107 << (uint(int32(3)) % 32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121)))
	v124 = F_pg_detoast_datum(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L6
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v120))) = base.I64_extend_i32_u(v124)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v130+v107<<(uint(int32(2))%32)))) = int32(0)
	v137 = v107 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v137 < v138 {
		v107 = v137
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_qsort_arg(m, v34, v70, int32(16), int32(1685), v26)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v148 = int32(0)
	v150 = base.I32_div_s(v70, int32(2))
	v152 = v150 << (uint(int32(4)) % 32)
	v157 = F_range_serialize(m, v26, v30+v152, v152+v34, v148, v148)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v159)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = base.I64_extend_i32_u(v157)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
	if v163 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v168 = int32(4)
	goto L29
L28:
	;
	v168 = int32(5)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v168
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v172 = F_palloc_mul(m, int32(4), v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v172
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v177 = F_palloc_mul(m, int32(8), v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v180 <= int32(0) {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v188 = v148
	goto L33
L33:
	;
	v196 = v188 << (uint(int32(3)) % 32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196+v197)))
	v200 = F_pg_detoast_datum(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L6
L35:
	;
	v203 = v15 - int32(-64)
	v205 = v15 + int32(48)
	F_range_deserialize(m, v26, v157, v203, v205, v15+int32(47))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v211 = v15 + int32(24)
	v213 = v15 + int32(8)
	F_range_deserialize(m, v26, v200, v211, v213, v15+int32(7))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v219 != 0 {
		v238 = int32(5)
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v239+v196))) = base.I64_extend_i32_u(v200)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v247 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v243+v188<<(uint(int32(2))%32)))) = v238 - v247
	v251 = v188 + v247
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v251 < v252 {
		v188 = v251
		goto L33
	} else {
		goto L49
	}
L39:
	;
	v220 = F_range_cmp_bounds(m, v26, v211, v203)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v224 = F_range_cmp_bounds(m, v26, v213, v205)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if int32(0) <= v224 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v228 = int32(1)
	goto L44
L43:
	;
	v228 = int32(2)
	goto L44
L44:
	;
	if int32(0) <= v220 {
		v238 = v228
		goto L38
	} else {
		goto L45
	}
L45:
	;
	if int32(0) <= v224 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v235 = int32(4)
	goto L48
L47:
	;
	v235 = int32(3)
	goto L48
L48:
	;
	v238 = v235
	goto L38
L49:
	;
	goto L34
}
func F_spg_text_inner_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int64
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v308 int32
	_ = v308
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = F_pg_newlocale_from_collation(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v31 = int32(1)
	v32 = v30 + v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+38)))
	if v36 == v31 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	v40 = F_pg_detoast_datum_packed(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v73 = int32(0)
	v74 = v2
	v75 = v32
	goto L5
L5:
	;
	v77 = v75 + int32(4)
	v78 = F_palloc(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v73 = v71
	v74 = v40
	v75 = v32 + v71
	goto L5
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v48 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v59 = int32(1)
	if v42&v59 != 0 {
		v71 = int32(base.Ui32(v42)>>(uint(v59)%32)) - v59
		goto L6
	} else {
		goto L17
	}
L11:
	;
	v51 = int32(16)
	goto L13
L12:
	;
	v51 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v48-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = int32(4)
	goto L16
L15:
	;
	v58 = v51
	goto L16
L16:
	;
	v71 = v58
	goto L6
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v77 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v84 = int32(0)
	v85 = base.B2i32(v83 == v84)
	if v85|v85 == v84 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v91 = int32(4)
	base.MemoryCopy(m, v78+v91, v33+v91, v83)
	goto L21
L20:
	;
	goto L21
L21:
	;
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v98 = int32(4)
	v100 = int32(1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v102&v100 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v110 = F_palloc_mul(m, int32(4), v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	v105 = v100
	goto L27
L26:
	;
	v105 = v98
	goto L27
L27:
	;
	base.MemoryCopy(m, v78+v96+v98, v74+v105, v73)
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v110
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v115 = F_palloc_mul(m, int32(4), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v115
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v120 = F_palloc_mul(m, int32(8), v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v120
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v122 < v125 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v129 = v78 + int32(4)
	v130 = int32(1)
	v150 = v2
	goto L34
L32:
	;
	goto L33
L33:
	;
	m.G0 = v21 + int32(16)
	return int64(0)
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v156+v150<<(uint(int32(3))%32))))
	if int32(0) < base.I32_extend16_s(base.I32_wrap_i64(v160)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v78+int32(3)))) = uint8(v160)
	v166 = v75
	goto L38
L37:
	;
	v166 = v75 - v130
	goto L38
L38:
	;
	v167 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v167 < v168 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v415 = v150 + int32(1)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v415 < v416 {
		v150 = v415
		goto L34
	} else {
		goto L101
	}
L40:
	;
	v171 = v167
	goto L43
L41:
	;
	goto L42
L42:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v365 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v363+v364<<(uint(v365)%32)))) = v150
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v369+v370<<(uint(v365)%32)))) = v166 - v374
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v166<<(uint(v365)%32) + int32(16)
	v384 = F_datumCopy(m, base.I64_extend_i32_u(v78), int32(0), int32(-1))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L100
	}
L43:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v192 = v189 + v171*int32(56)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+6)))
	if base.B2i32(base.Ui32(v193) < base.Ui32(int32(11)))|base.B2i32(v193 == int32(28)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L42
L45:
	;
	v342 = v171 + int32(1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v342 < v343 {
		v171 = v342
		goto L43
	} else {
		goto L99
	}
L46:
	;
	if v34&v130 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v205 = v193
	goto L48
L48:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v192)+48))
	v207 = F_pg_detoast_datum_packed(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v205 = v193 - int32(10)
	goto L48
L50:
	;
	v239 = int32(1)
	if v209&v239 != 0 {
		goto L62
	} else {
		goto L63
	}
L51:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v209 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	if v215 == int32(18) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v226 = int32(1)
	if v209&v226 != 0 {
		v238 = int32(base.Ui32(v209)>>(uint(v226)%32)) - v226
		goto L50
	} else {
		goto L61
	}
L55:
	;
	v218 = int32(16)
	goto L57
L56:
	;
	v218 = int32(0)
	goto L57
L57:
	;
	if base.Ui32((v215-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v225 = int32(4)
	goto L60
L59:
	;
	v225 = v218
	goto L60
L60:
	;
	v238 = v225
	goto L50
L61:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	v238 = int32(base.Ui32(v232)>>(uint(int32(2))%32)) - int32(4)
	goto L50
L62:
	;
	v243 = v239
	goto L64
L63:
	;
	v243 = int32(4)
	goto L64
L64:
	;
	v244 = v207 + v243
	v245 = base.B2i32(v238 < v166)
	if v238 < v166 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v246 = v238
	goto L67
L66:
	;
	v246 = v166
	goto L67
L67:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v246) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	switch v205&int32(_a_F_spg_text_inner_consistent_0) - int32(1) {
	case 0, 1:
		goto L90
	case 2:
		goto L89
	case 3, 4:
		goto L88
	default:
		goto L87
	case 27:
		goto L86
	}
L69:
	;
	v308 = int32(0)
	goto L68
L70:
	;
	v282 = v277
	v283 = v278
	v284 = v279
	goto L80
L71:
	;
	if (v129|v244)&int32(3) != 0 {
		v277 = v129
		v278 = v244
		v279 = v246
		goto L70
	} else {
		goto L74
	}
L72:
	;
	v270 = v129
	v271 = v244
	v272 = v246
	goto L73
L73:
	;
	if v272 == int32(0) {
		goto L69
	} else {
		goto L79
	}
L74:
	;
	v254 = v129
	v255 = v244
	v256 = v246
	goto L75
L75:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v259 != v260 {
		v277 = v254
		v278 = v255
		v279 = v256
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v270 = v265
	v271 = v263
	v272 = v267
	goto L73
L77:
	;
	v262 = int32(4)
	v263 = v255 + v262
	v265 = v254 + v262
	v267 = v256 - v262
	if base.Ui32(int32(3)) < base.Ui32(v267) {
		v254 = v265
		v255 = v263
		v256 = v267
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v277 = v270
	v278 = v271
	v279 = v272
	goto L70
L80:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v287 == v288 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v308 = v287 - v288
	goto L68
L82:
	;
	v290 = int32(1)
	v295 = v284 - v290
	if v295 != 0 {
		v282 = v282 + v290
		v283 = v283 + v290
		v284 = v295
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	goto L69
L86:
	;
	if v308 != 0 {
		goto L39
	} else {
		goto L98
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L95
	}
L88:
	;
	if int32(0) <= v308 {
		goto L45
	} else {
		goto L94
	}
L89:
	;
	if v238 < v166 {
		goto L39
	} else {
		goto L92
	}
L90:
	;
	if v308 <= int32(0) {
		goto L45
	} else {
		goto L91
	}
L91:
	;
	goto L39
L92:
	;
	if v308 == int32(0) {
		goto L45
	} else {
		goto L93
	}
L93:
	;
	goto L39
L94:
	;
	goto L39
L95:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323+v171*int32(56))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v327
	F_errmsg_internal(m, int32(_a_F_spg_text_inner_consistent_1), v21)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_spg_text_inner_consistent_2), int32(552), int32(_a_F_spg_text_inner_consistent_3))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	goto L45
L99:
	;
	goto L44
L100:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v386+v387<<(uint(int32(3))%32)))) = v384
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v392 + int32(1)
	goto L39
L101:
	;
	goto L35
}

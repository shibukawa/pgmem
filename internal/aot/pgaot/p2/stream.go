package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_begin_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
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
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	v5 = l4
	v19 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v19 == int32(0) {
		v33 = *(*int32)(unsafe.Add(mBase, _consts[736]))
		v51 = v33
		v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
		if l1 == int32(0) {
			v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
			v66 = v57
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v59 != int32(1) {
				v63 = base.I32_div_s(v58, int32(2))
				v64 = v63
			} else {
				v64 = v58
			}
			v66 = v64
		}
		v67 = int32(32767)
		if v67 <= v51 {
			v70 = v67
		} else {
			v70 = v51
		}
		v73 = v53 * (v70 + int32(1))
		v75 = int32(16)
		v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
		v81 = int32(32766) - v80
		if base.Ui32(v73) < base.Ui32(v81) {
			v83 = v73
		} else {
			v83 = v81
		}
		if base.Ui32(v66) < base.Ui32(v83) {
			v85 = v66
		} else {
			v85 = v83
		}
		v86 = int32(1)
		if v70 <= v86 {
			v89 = v86
		} else {
			v89 = v70
		}
		v91 = v89 * int32(84)
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		if v93 != int32(-1) {
			v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
			v100 = v97
		} else {
			v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
			v100 = v99
		}
		if base.Ui32(v85) < base.Ui32(v100) {
			v102 = v85
		} else {
			v102 = v100
		}
		if base.Ui32(v102) <= base.Ui32(int32(1)) {
			v105 = int32(1)
		} else {
			v105 = v102
		}
		v107 = v105 + int32(1)
		v108 = base.I32_extend16_s(v107)
		v113 = (v108 + v80) << (uint(int32(2)) % 32)
		v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
			mBase = m.M
			v127 = (v122 + v113 + int32(87)) & int32(-8)
			*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
			if l8 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
			} else {
			}
			v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
			v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
			*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
			if v136 != 0 {
				if v51 != 0 {
					v158 = v70
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
					v158 = int32(1)
				}
			} else {
				v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
				if v146&int32(1) != 0 {
					if v51 != 0 {
						v158 = v70
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
						v158 = int32(1)
					}
				} else {
					if l0&int32(2) != 0 {
						if v51 != 0 {
							v158 = v70
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
							v158 = int32(1)
						}
					} else {
						if v51 <= int32(0) {
							if v51 != 0 {
								v158 = v70
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							v153 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
							v158 = v70
						}
					}
				}
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
			v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
			*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
			*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
			*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
			*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
			*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
			*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
			*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
			v171 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
			v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
			if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
				v179 = v105
			} else {
				v179 = v162
			}
			if l0&int32(4) != 0 {
				v183 = v179
			} else {
				v183 = int32(1)
			}
			*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
			if int32(0) < v158 {
				v197 = int32(0)
				for {
					v205 = v197 * int32(84)
					v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
					v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
					v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
					*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
					v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
					v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
					v222 = v197 + int32(1)
					if v222 != v158 {
						v197 = v222
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return v122
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if l2 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
			if base.Ui32(v23) < base.Ui32(int32(12000)) {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[736]))
				v51 = v33
				v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
				if l1 == int32(0) {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					v66 = v57
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v59 != int32(1) {
						v63 = base.I32_div_s(v58, int32(2))
						v64 = v63
					} else {
						v64 = v58
					}
					v66 = v64
				}
				v67 = int32(32767)
				if v67 <= v51 {
					v70 = v67
				} else {
					v70 = v51
				}
				v73 = v53 * (v70 + int32(1))
				v75 = int32(16)
				v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
				v81 = int32(32766) - v80
				if base.Ui32(v73) < base.Ui32(v81) {
					v83 = v73
				} else {
					v83 = v81
				}
				if base.Ui32(v66) < base.Ui32(v83) {
					v85 = v66
				} else {
					v85 = v83
				}
				v86 = int32(1)
				if v70 <= v86 {
					v89 = v86
				} else {
					v89 = v70
				}
				v91 = v89 * int32(84)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
				if v93 != int32(-1) {
					v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
					v100 = v97
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
					v100 = v99
				}
				if base.Ui32(v85) < base.Ui32(v100) {
					v102 = v85
				} else {
					v102 = v100
				}
				if base.Ui32(v102) <= base.Ui32(int32(1)) {
					v105 = int32(1)
				} else {
					v105 = v102
				}
				v107 = v105 + int32(1)
				v108 = base.I32_extend16_s(v107)
				v113 = (v108 + v80) << (uint(int32(2)) % 32)
				v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
					mBase = m.M
					v127 = (v122 + v113 + int32(87)) & int32(-8)
					*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
					if l8 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
					} else {
					}
					v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
					v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
					if v136 != 0 {
						if v51 != 0 {
							v158 = v70
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
							v158 = int32(1)
						}
					} else {
						v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
						if v146&int32(1) != 0 {
							if v51 != 0 {
								v158 = v70
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							if l0&int32(2) != 0 {
								if v51 != 0 {
									v158 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								if v51 <= int32(0) {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v153 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
									v158 = v70
								}
							}
						}
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
					v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
					*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
					*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
					v171 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
					v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
					if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
						v179 = v105
					} else {
						v179 = v162
					}
					if l0&int32(4) != 0 {
						v183 = v179
					} else {
						v183 = int32(1)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
					if int32(0) < v158 {
						v197 = int32(0)
						for {
							v205 = v197 * int32(84)
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
							v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
							v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
							v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
							v222 = v197 + int32(1)
							if v222 != v158 {
								v197 = v222
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v122
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				if base.B2i32(base.Ui32(v26) < base.Ui32(int32(12000))) == int32(0) {
					if l0&int32(1) != 0 {
						v36 = F_get_tablespace_maintenance_io_concurrency(m, v22)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v51 = v36
							v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
							if l1 == int32(0) {
								v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
								v66 = v57
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v59 != int32(1) {
									v63 = base.I32_div_s(v58, int32(2))
									v64 = v63
								} else {
									v64 = v58
								}
								v66 = v64
							}
							v67 = int32(32767)
							if v67 <= v51 {
								v70 = v67
							} else {
								v70 = v51
							}
							v73 = v53 * (v70 + int32(1))
							v75 = int32(16)
							v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
							v81 = int32(32766) - v80
							if base.Ui32(v73) < base.Ui32(v81) {
								v83 = v73
							} else {
								v83 = v81
							}
							if base.Ui32(v66) < base.Ui32(v83) {
								v85 = v66
							} else {
								v85 = v83
							}
							v86 = int32(1)
							if v70 <= v86 {
								v89 = v86
							} else {
								v89 = v70
							}
							v91 = v89 * int32(84)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							if v93 != int32(-1) {
								v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
								v100 = v97
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
								v100 = v99
							}
							if base.Ui32(v85) < base.Ui32(v100) {
								v102 = v85
							} else {
								v102 = v100
							}
							if base.Ui32(v102) <= base.Ui32(int32(1)) {
								v105 = int32(1)
							} else {
								v105 = v102
							}
							v107 = v105 + int32(1)
							v108 = base.I32_extend16_s(v107)
							v113 = (v108 + v80) << (uint(int32(2)) % 32)
							v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
								mBase = m.M
								v127 = (v122 + v113 + int32(87)) & int32(-8)
								*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
								if l8 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
								} else {
								}
								v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
								v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
								if v136 != 0 {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
									if v146&int32(1) != 0 {
										if v51 != 0 {
											v158 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										if l0&int32(2) != 0 {
											if v51 != 0 {
												v158 = v70
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
												v158 = int32(1)
											}
										} else {
											if v51 <= int32(0) {
												if v51 != 0 {
													v158 = v70
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
													v158 = int32(1)
												}
											} else {
												v153 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
												v158 = v70
											}
										}
									}
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
								v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
								*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
								*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
								*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
								v171 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
								if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
									v179 = v105
								} else {
									v179 = v162
								}
								if l0&int32(4) != 0 {
									v183 = v179
								} else {
									v183 = int32(1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
								if int32(0) < v158 {
									v197 = int32(0)
									for {
										v205 = v197 * int32(84)
										v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
										v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
										v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
										v222 = v197 + int32(1)
										if v222 != v158 {
											v197 = v222
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								return v122
							}
						}
					} else {
						v40 = F_get_tablespace(m, v22)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							if v42 != 0 {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
								if int32(0) <= v43 {
									v49 = v43
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, _consts[736]))
									v49 = v48
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, _consts[736]))
								v49 = v48
							}
							v51 = v49
							v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
							if l1 == int32(0) {
								v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
								v66 = v57
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v59 != int32(1) {
									v63 = base.I32_div_s(v58, int32(2))
									v64 = v63
								} else {
									v64 = v58
								}
								v66 = v64
							}
							v67 = int32(32767)
							if v67 <= v51 {
								v70 = v67
							} else {
								v70 = v51
							}
							v73 = v53 * (v70 + int32(1))
							v75 = int32(16)
							v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
							v81 = int32(32766) - v80
							if base.Ui32(v73) < base.Ui32(v81) {
								v83 = v73
							} else {
								v83 = v81
							}
							if base.Ui32(v66) < base.Ui32(v83) {
								v85 = v66
							} else {
								v85 = v83
							}
							v86 = int32(1)
							if v70 <= v86 {
								v89 = v86
							} else {
								v89 = v70
							}
							v91 = v89 * int32(84)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							if v93 != int32(-1) {
								v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
								v100 = v97
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
								v100 = v99
							}
							if base.Ui32(v85) < base.Ui32(v100) {
								v102 = v85
							} else {
								v102 = v100
							}
							if base.Ui32(v102) <= base.Ui32(int32(1)) {
								v105 = int32(1)
							} else {
								v105 = v102
							}
							v107 = v105 + int32(1)
							v108 = base.I32_extend16_s(v107)
							v113 = (v108 + v80) << (uint(int32(2)) % 32)
							v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
								mBase = m.M
								v127 = (v122 + v113 + int32(87)) & int32(-8)
								*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
								if l8 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
								} else {
								}
								v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
								v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
								if v136 != 0 {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
									if v146&int32(1) != 0 {
										if v51 != 0 {
											v158 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										if l0&int32(2) != 0 {
											if v51 != 0 {
												v158 = v70
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
												v158 = int32(1)
											}
										} else {
											if v51 <= int32(0) {
												if v51 != 0 {
													v158 = v70
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
													v158 = int32(1)
												}
											} else {
												v153 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
												v158 = v70
											}
										}
									}
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
								v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
								*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
								*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
								*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
								v171 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
								*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
								if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
									v179 = v105
								} else {
									v179 = v162
								}
								if l0&int32(4) != 0 {
									v183 = v179
								} else {
									v183 = int32(1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
								if int32(0) < v158 {
									v197 = int32(0)
									for {
										v205 = v197 * int32(84)
										v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
										v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
										v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
										*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
										v222 = v197 + int32(1)
										if v222 != v158 {
											v197 = v222
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								return v122
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[736]))
					v51 = v33
					v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
					if l1 == int32(0) {
						v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						v66 = v57
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v59 != int32(1) {
							v63 = base.I32_div_s(v58, int32(2))
							v64 = v63
						} else {
							v64 = v58
						}
						v66 = v64
					}
					v67 = int32(32767)
					if v67 <= v51 {
						v70 = v67
					} else {
						v70 = v51
					}
					v73 = v53 * (v70 + int32(1))
					v75 = int32(16)
					v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
					v81 = int32(32766) - v80
					if base.Ui32(v73) < base.Ui32(v81) {
						v83 = v73
					} else {
						v83 = v81
					}
					if base.Ui32(v66) < base.Ui32(v83) {
						v85 = v66
					} else {
						v85 = v83
					}
					v86 = int32(1)
					if v70 <= v86 {
						v89 = v86
					} else {
						v89 = v70
					}
					v91 = v89 * int32(84)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					if v93 != int32(-1) {
						v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
						v100 = v97
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
						v100 = v99
					}
					if base.Ui32(v85) < base.Ui32(v100) {
						v102 = v85
					} else {
						v102 = v100
					}
					if base.Ui32(v102) <= base.Ui32(int32(1)) {
						v105 = int32(1)
					} else {
						v105 = v102
					}
					v107 = v105 + int32(1)
					v108 = base.I32_extend16_s(v107)
					v113 = (v108 + v80) << (uint(int32(2)) % 32)
					v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
						mBase = m.M
						v127 = (v122 + v113 + int32(87)) & int32(-8)
						*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
						if l8 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
						} else {
						}
						v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
						v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
						*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
						if v136 != 0 {
							if v51 != 0 {
								v158 = v70
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
							if v146&int32(1) != 0 {
								if v51 != 0 {
									v158 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								if l0&int32(2) != 0 {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									if v51 <= int32(0) {
										if v51 != 0 {
											v158 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										v153 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
										v158 = v70
									}
								}
							}
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
						v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
						*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
						*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
						*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
						*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
						*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
						*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
						*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
						v171 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
						if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
							v179 = v105
						} else {
							v179 = v162
						}
						if l0&int32(4) != 0 {
							v183 = v179
						} else {
							v183 = int32(1)
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
						if int32(0) < v158 {
							v197 = int32(0)
							for {
								v205 = v197 * int32(84)
								v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
								v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
								*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
								v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
								v222 = v197 + int32(1)
								if v222 != v158 {
									v197 = v222
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						return v122
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if base.B2i32(base.Ui32(v26) < base.Ui32(int32(12000))) == int32(0) {
				if l0&int32(1) != 0 {
					v36 = F_get_tablespace_maintenance_io_concurrency(m, v22)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v51 = v36
						v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
						if l1 == int32(0) {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							v66 = v57
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v59 != int32(1) {
								v63 = base.I32_div_s(v58, int32(2))
								v64 = v63
							} else {
								v64 = v58
							}
							v66 = v64
						}
						v67 = int32(32767)
						if v67 <= v51 {
							v70 = v67
						} else {
							v70 = v51
						}
						v73 = v53 * (v70 + int32(1))
						v75 = int32(16)
						v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
						v81 = int32(32766) - v80
						if base.Ui32(v73) < base.Ui32(v81) {
							v83 = v73
						} else {
							v83 = v81
						}
						if base.Ui32(v66) < base.Ui32(v83) {
							v85 = v66
						} else {
							v85 = v83
						}
						v86 = int32(1)
						if v70 <= v86 {
							v89 = v86
						} else {
							v89 = v70
						}
						v91 = v89 * int32(84)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v93 != int32(-1) {
							v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
							v100 = v97
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
							v100 = v99
						}
						if base.Ui32(v85) < base.Ui32(v100) {
							v102 = v85
						} else {
							v102 = v100
						}
						if base.Ui32(v102) <= base.Ui32(int32(1)) {
							v105 = int32(1)
						} else {
							v105 = v102
						}
						v107 = v105 + int32(1)
						v108 = base.I32_extend16_s(v107)
						v113 = (v108 + v80) << (uint(int32(2)) % 32)
						v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
							mBase = m.M
							v127 = (v122 + v113 + int32(87)) & int32(-8)
							*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
							if l8 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
							} else {
							}
							v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
							v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
							if v136 != 0 {
								if v51 != 0 {
									v158 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
								if v146&int32(1) != 0 {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									if l0&int32(2) != 0 {
										if v51 != 0 {
											v158 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										if v51 <= int32(0) {
											if v51 != 0 {
												v158 = v70
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
												v158 = int32(1)
											}
										} else {
											v153 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
											v158 = v70
										}
									}
								}
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
							v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
							*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
							*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
							*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
							v171 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
							if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
								v179 = v105
							} else {
								v179 = v162
							}
							if l0&int32(4) != 0 {
								v183 = v179
							} else {
								v183 = int32(1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
							if int32(0) < v158 {
								v197 = int32(0)
								for {
									v205 = v197 * int32(84)
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
									v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
									v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
									v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
									v222 = v197 + int32(1)
									if v222 != v158 {
										v197 = v222
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							return v122
						}
					}
				} else {
					v40 = F_get_tablespace(m, v22)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						if v42 != 0 {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
							if int32(0) <= v43 {
								v49 = v43
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, _consts[736]))
								v49 = v48
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, _consts[736]))
							v49 = v48
						}
						v51 = v49
						v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
						if l1 == int32(0) {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							v66 = v57
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v59 != int32(1) {
								v63 = base.I32_div_s(v58, int32(2))
								v64 = v63
							} else {
								v64 = v58
							}
							v66 = v64
						}
						v67 = int32(32767)
						if v67 <= v51 {
							v70 = v67
						} else {
							v70 = v51
						}
						v73 = v53 * (v70 + int32(1))
						v75 = int32(16)
						v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
						v81 = int32(32766) - v80
						if base.Ui32(v73) < base.Ui32(v81) {
							v83 = v73
						} else {
							v83 = v81
						}
						if base.Ui32(v66) < base.Ui32(v83) {
							v85 = v66
						} else {
							v85 = v83
						}
						v86 = int32(1)
						if v70 <= v86 {
							v89 = v86
						} else {
							v89 = v70
						}
						v91 = v89 * int32(84)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v93 != int32(-1) {
							v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
							v100 = v97
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
							v100 = v99
						}
						if base.Ui32(v85) < base.Ui32(v100) {
							v102 = v85
						} else {
							v102 = v100
						}
						if base.Ui32(v102) <= base.Ui32(int32(1)) {
							v105 = int32(1)
						} else {
							v105 = v102
						}
						v107 = v105 + int32(1)
						v108 = base.I32_extend16_s(v107)
						v113 = (v108 + v80) << (uint(int32(2)) % 32)
						v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
							mBase = m.M
							v127 = (v122 + v113 + int32(87)) & int32(-8)
							*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
							if l8 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
							} else {
							}
							v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
							v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
							if v136 != 0 {
								if v51 != 0 {
									v158 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
								if v146&int32(1) != 0 {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									if l0&int32(2) != 0 {
										if v51 != 0 {
											v158 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
											v158 = int32(1)
										}
									} else {
										if v51 <= int32(0) {
											if v51 != 0 {
												v158 = v70
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
												v158 = int32(1)
											}
										} else {
											v153 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
											v158 = v70
										}
									}
								}
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
							v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
							*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
							*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
							*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
							v171 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
							if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
								v179 = v105
							} else {
								v179 = v162
							}
							if l0&int32(4) != 0 {
								v183 = v179
							} else {
								v183 = int32(1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
							if int32(0) < v158 {
								v197 = int32(0)
								for {
									v205 = v197 * int32(84)
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
									v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
									v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
									v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
									v222 = v197 + int32(1)
									if v222 != v158 {
										v197 = v222
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							return v122
						}
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[736]))
				v51 = v33
				v53 = *(*int32)(unsafe.Add(mBase, _consts[432]))
				if l1 == int32(0) {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					v66 = v57
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v59 != int32(1) {
						v63 = base.I32_div_s(v58, int32(2))
						v64 = v63
					} else {
						v64 = v58
					}
					v66 = v64
				}
				v67 = int32(32767)
				if v67 <= v51 {
					v70 = v67
				} else {
					v70 = v51
				}
				v73 = v53 * (v70 + int32(1))
				v75 = int32(16)
				v80 = (v53<<(uint(v75)%32) - int32(65536)) >> (uint(v75) % 32)
				v81 = int32(32766) - v80
				if base.Ui32(v73) < base.Ui32(v81) {
					v83 = v73
				} else {
					v83 = v81
				}
				if base.Ui32(v66) < base.Ui32(v83) {
					v85 = v66
				} else {
					v85 = v83
				}
				v86 = int32(1)
				if v70 <= v86 {
					v89 = v86
				} else {
					v89 = v70
				}
				v91 = v89 * int32(84)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
				if v93 != int32(-1) {
					v97 = *(*int32)(unsafe.Add(mBase, _consts[737]))
					v100 = v97
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, _consts[738]))
					v100 = v99
				}
				if base.Ui32(v85) < base.Ui32(v100) {
					v102 = v85
				} else {
					v102 = v100
				}
				if base.Ui32(v102) <= base.Ui32(int32(1)) {
					v105 = int32(1)
				} else {
					v105 = v102
				}
				v107 = v105 + int32(1)
				v108 = base.I32_extend16_s(v107)
				v113 = (v108 + v80) << (uint(int32(2)) % 32)
				v117 = F_palloc(m, v91+l8*v108+v113+int32(96))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					v122 = F__emscripten_memset_bulkmem(m, v117, base.I32_extend8_s(int32(0)), int32(80))
					mBase = m.M
					v127 = (v122 + v113 + int32(87)) & int32(-8)
					*(*int32)(unsafe.Add(mBase, uint32(v122)+64)) = v127
					if l8 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v122)+60)) = (v127 + v91 + int32(7)) & int32(-8)
					} else {
					}
					v136 = *(*int32)(unsafe.Add(mBase, _consts[739]))
					v140 = int32(base.Ui32(l0)>>(uint(int32(3))%32)) & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+25)) = uint8(v140)
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)) = uint8(base.B2i32(v136 == int32(0)))
					if v136 != 0 {
						if v51 != 0 {
							v158 = v70
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
							v158 = int32(1)
						}
					} else {
						v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[230])))
						if v146&int32(1) != 0 {
							if v51 != 0 {
								v158 = v70
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
								v158 = int32(1)
							}
						} else {
							if l0&int32(2) != 0 {
								if v51 != 0 {
									v158 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
									v158 = int32(1)
								}
							} else {
								if v51 <= int32(0) {
									if v51 != 0 {
										v158 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(8)
										v158 = int32(1)
									}
								} else {
									v153 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v122)+26)) = uint8(v153)
									v158 = v70
								}
							}
						}
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v122))) = uint16(v158)
					v162 = *(*int32)(unsafe.Add(mBase, _consts[432]))
					*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = l8
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)) = uint16(v162)
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+8)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v122)+32)) = l6
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v107)
					*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(-1)
					v171 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v171
					v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*uint8)(unsafe.Add(mBase, uint32(v122)+27)) = uint8(base.B2i32(v173 != v171))
					if base.Ui32(v105) < base.Ui32(base.I32_extend16_s(v162)) {
						v179 = v105
					} else {
						v179 = v162
					}
					if l0&int32(4) != 0 {
						v183 = v179
					} else {
						v183 = int32(1)
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v122)+14)) = uint16(v183)
					if int32(0) < v158 {
						v197 = int32(0)
						for {
							v205 = v197 * int32(84)
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v205+v206)+4)) = l2
							v209 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v209+v205)+8)) = l3
							v212 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*uint8)(unsafe.Add(mBase, uint32(v212+v205)+12)) = uint8(v5)
							v215 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v215+v205)+16)) = l5
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+64))
							*(*int32)(unsafe.Add(mBase, uint32(v218+v205)+20)) = l1
							v222 = v197 + int32(1)
							if v222 != v158 {
								v197 = v222
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v122
				}
			}
		}
	}
}

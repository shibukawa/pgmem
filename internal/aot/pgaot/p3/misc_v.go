package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___vfprintf_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = l2
	v20 = F__emscripten_memset_bulkmem(m, v12+int32(160), base.I32_extend8_s(v6), int32(40))
	mBase = m.M
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v21
	v30 = F_printf_core(m, int32(0), l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		if v30 < int32(0) {
			v122 = int32(-1)
			m.G0 = v12 + int32(208)
			return v122
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v38 = int32(0)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40 & int32(-33)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v44 == v38 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(80)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v12
				v56 = v53
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v59 - int32(1) | v59
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v64&int32(8) != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 | int32(32)
					v81 = int32(-1)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v73
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73 + v76
					v81 = int32(0)
				}
				if v81 != 0 {
					v92 = int32(-1)
					v93 = v56
					if v93 != 0 {
						v96 = int32(0)
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v101 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							if v106 != 0 {
								v110 = v92
							} else {
								v110 = int32(-1)
							}
							v111 = v110
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					} else {
						v111 = v92
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
						if v113&int32(32) != 0 {
							v119 = int32(-1)
						} else {
							v119 = v111
						}
						if v37 < v38 {
							v122 = v119
						} else {
							v122 = v119
						}
						m.G0 = v12 + int32(208)
						return v122
					}
				} else {
					v83 = v56
					v90 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = v90
						v93 = v83
						if v93 != 0 {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v106 != 0 {
									v110 = v92
								} else {
									v110 = int32(-1)
								}
								v111 = v110
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						} else {
							v111 = v92
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					}
				}
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v55 != 0 {
					v83 = v6
					v90 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = v90
						v93 = v83
						if v93 != 0 {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v106 != 0 {
									v110 = v92
								} else {
									v110 = int32(-1)
								}
								v111 = v110
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						} else {
							v111 = v92
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					}
				} else {
					v56 = v6
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v59 - int32(1) | v59
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v64&int32(8) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64 | int32(32)
						v81 = int32(-1)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v73
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73 + v76
						v81 = int32(0)
					}
					if v81 != 0 {
						v92 = int32(-1)
						v93 = v56
						if v93 != 0 {
							v96 = int32(0)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								if v106 != 0 {
									v110 = v92
								} else {
									v110 = int32(-1)
								}
								v111 = v110
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						} else {
							v111 = v92
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
							if v113&int32(32) != 0 {
								v119 = int32(-1)
							} else {
								v119 = v111
							}
							if v37 < v38 {
								v122 = v119
							} else {
								v122 = v119
							}
							m.G0 = v12 + int32(208)
							return v122
						}
					} else {
						v83 = v56
						v90 = F_printf_core(m, l0, l1, v12+int32(200), v12+int32(80), v12+int32(160), l3, l4)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = v90
							v93 = v83
							if v93 != 0 {
								v96 = int32(0)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v96, v96)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v101
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v93
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
									if v106 != 0 {
										v110 = v92
									} else {
										v110 = int32(-1)
									}
									v111 = v110
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
									if v113&int32(32) != 0 {
										v119 = int32(-1)
									} else {
										v119 = v111
									}
									if v37 < v38 {
										v122 = v119
									} else {
										v122 = v119
									}
									m.G0 = v12 + int32(208)
									return v122
								}
							} else {
								v111 = v92
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113 | v40&int32(32)
								if v113&int32(32) != 0 {
									v119 = int32(-1)
								} else {
									v119 = v111
								}
								if v37 < v38 {
									v122 = v119
								} else {
									v122 = v119
								}
								m.G0 = v12 + int32(208)
								return v122
							}
						}
					}
				}
			}
		}
	}
}
func F_varbit_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(8)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v19 = F_palloc(m, v16+int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = v16 - int32(8)
			if v22 < int32(0) {
				v89 = v19
				v91 = v15
				v92 = v2
			} else {
				v25 = v19
				v27 = v15
				v28 = v2
				for {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
					v36 = int32(48)
					v37 = v33&int32(1) | v36
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+7)) = uint8(v37)
					if v33&int32(2) != 0 {
						v43 = int32(49)
					} else {
						v43 = v36
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v43)
					if v33&int32(4) != 0 {
						v49 = int32(49)
					} else {
						v49 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v49)
					if v33&int32(8) != 0 {
						v55 = int32(49)
					} else {
						v55 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)) = uint8(v55)
					if v33&int32(16) != 0 {
						v61 = int32(49)
					} else {
						v61 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)) = uint8(v61)
					if v33&int32(32) != 0 {
						v67 = int32(49)
					} else {
						v67 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)) = uint8(v67)
					if v33&int32(64) != 0 {
						v73 = int32(49)
					} else {
						v73 = int32(48)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v73)
					if int32(0) <= base.I32_extend8_s(v33) {
						v80 = int32(48)
					} else {
						v80 = int32(49)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v80)
					v83 = v27 + int32(1)
					v84 = int32(8)
					v85 = v25 + v84
					v87 = v28 + v84
					if v87 <= v22 {
						v25 = v85
						v27 = v83
						v28 = v87
						continue
					} else {
						break
					}
					break
				}
				v89 = v85
				v91 = v83
				v92 = v87
			}
			if v16 <= v92 {
				v180 = v89
			} else {
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v101 = (v16 - v92) & int32(3)
				if v101 == int32(0) {
					v129 = v89
					v130 = v98
					v131 = v92
				} else {
					v105 = v89
					v106 = v98
					v107 = v92
					v110 = int32(0)
					for {
						if int32(0) <= base.I32_extend8_s(v106) {
							v118 = int32(48)
						} else {
							v118 = int32(49)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v118)
						v120 = int32(1)
						v121 = v107 + v120
						v123 = v106 << (uint(v120) % 32)
						v125 = v105 + v120
						v127 = v110 + v120
						if v127 != v101 {
							v105 = v125
							v106 = v123
							v107 = v121
							v110 = v127
							continue
						} else {
							break
						}
						break
					}
					v129 = v125
					v130 = v123
					v131 = v121
				}
				if base.Ui32(int32(-4)) < base.Ui32(v92-v16) {
					v180 = v129
				} else {
					v140 = v129
					v141 = v130
					v142 = v131
					for {
						if v141&int32(16) != 0 {
							v152 = int32(49)
						} else {
							v152 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+3)) = uint8(v152)
						if v141&int32(32) != 0 {
							v158 = int32(49)
						} else {
							v158 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+2)) = uint8(v158)
						if v141&int32(64) != 0 {
							v164 = int32(49)
						} else {
							v164 = int32(48)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)) = uint8(v164)
						if int32(0) <= base.I32_extend8_s(v141) {
							v171 = int32(48)
						} else {
							v171 = int32(49)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v171)
						v173 = int32(4)
						v176 = v140 + v173
						v178 = v142 + v173
						if v178 != v16 {
							v140 = v176
							v141 = v141 << (uint(v173) % 32)
							v142 = v178
							continue
						} else {
							break
						}
						break
					}
					v180 = v176
				}
			}
			v188 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v188)
			return v19
		}
	}
}
func F_varbit_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pq_getmsgint(m, v12, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if base.Ui32(v14) < base.Ui32(int32(2147483641)) {
			if base.B2i32(v11 < v14)&base.B2i32(int32(0) < v11) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16777346))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg(m, int32(710139), v9)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(517075), int32(662), int32(37783))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
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
				v27 = int32(base.Ui32(v14+int32(7)) >> (uint(int32(3)) % 32))
				v29 = v27 + int32(8)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v29 << (uint(int32(2)) % 32)
					F_pq_copymsgbytes(m, v12, v30+int32(8), v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v42 = int32(base.Ui32(v40) >> (uint(int32(2)) % 32))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						v48 = v42<<(uint(int32(3))%32) - v45 + int32(-64)
						if int32(0) < v48 {
							v53 = v30 + v42 - int32(1)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
							v57 = v54 & (int32(255) << (uint(v48) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v57)
						} else {
						}
						m.G0 = v9 + int32(16)
						return v30
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(347245), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(517075), int32(652), int32(37783))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
func F_varchar_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v36 = v2
		return v36
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v36 = v2
			return v36
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v36 = v2
				return v36
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if int32(0) <= v22 {
						if v18 < int32(0) {
							v36 = v2
							return v36
						} else {
							v27 = int32(4)
							if v22-v27 < v18-v27 {
								v36 = v2
								return v36
							} else {
								v32 = F_relabel_to_typmod(m, v17, v22)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v36 = v32
									return v36
								}
							}
						}
					} else {
						v32 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v36 = v32
							return v36
						}
					}
				}
			}
		}
	}
}
func F_varstr_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
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
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
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
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = F_pg_detoast_datum_packed(m, l0)
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v24 = int32(1)
	v25 = v15 + v24
	v27 = v19 & v24
	if v19 == v24 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v30 = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v32&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v19)>>(uint(v45)%32)) - v45
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v41 = v30
	goto L9
L8:
	;
	v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
	goto L9
L9:
	;
	if v32 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = v30
	goto L12
L11:
	;
	v44 = v41
	goto L12
L12:
	;
	v55 = v44
	goto L3
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v56 = v25
	goto L16
L15:
	;
	v56 = v15 + int32(4)
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v57 == int32(1042) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v62 = int32(-1)
	v64 = v55 - int32(1)
	if v62 <= v64 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v83 = v55
	goto L19
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+29)))
	if v86 != 0 {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v83 = v82
	goto L19
L21:
	;
	v67 = v62
	goto L23
L22:
	;
	v67 = v64
	goto L23
L23:
	;
	v71 = v55
	goto L24
L24:
	;
	v75 = v71 - int32(1)
	if v75 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v82 = v71
	goto L20
L26:
	;
	v82 = v67 + int32(1)
	goto L20
L27:
	;
	goto L28
L28:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v75))))
	if v79 == int32(32) {
		v71 = v75
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if l0 != v15 {
		goto L181
	} else {
		goto L182
	}
L31:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v688 = int32(4)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if base.Ui32(v688) <= base.Ui32(v689) {
		goto L174
	} else {
		goto L175
	}
L32:
	;
	v268 = v83
	v275 = v56
	goto L34
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v87 <= v83 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v276 = int32(4)
	if base.Ui32(v276) <= base.Ui32(v268) {
		goto L99
	} else {
		goto L100
	}
L35:
	;
	v89 = int32(1)
	v90 = v83 + v89
	v91 = int32(1073741823)
	v93 = v87 << (uint(v89) % 32)
	if base.Ui32(v91) <= base.Ui32(v93) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v83 != v106 {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v96 = v91
	goto L40
L39:
	;
	v96 = v93
	goto L40
L40:
	;
	if base.Ui32(v96) < base.Ui32(v90) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v98 = v90
	goto L43
L42:
	;
	v98 = v96
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v101 = F_repalloc(m, v100, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v101
	goto L37
L45:
	;
	if v83 != 0 {
		goto L70
	} else {
		goto L71
	}
L46:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v177 = v108
	goto L45
L47:
	;
	goto L48
L48:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	if v110 != int32(1) {
		v177 = v109
		goto L45
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v83) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if v174 == int32(0) {
		goto L31
	} else {
		goto L68
	}
L51:
	;
	v174 = int32(0)
	goto L50
L52:
	;
	v148 = v143
	v149 = v144
	v150 = v145
	goto L62
L53:
	;
	if (v109|v56)&int32(3) != 0 {
		v143 = v109
		v144 = v56
		v145 = v83
		goto L52
	} else {
		goto L56
	}
L54:
	;
	v136 = v109
	v137 = v56
	v138 = v83
	goto L55
L55:
	;
	if v138 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L56:
	;
	v120 = v109
	v121 = v56
	v122 = v83
	goto L57
L57:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v125 != v126 {
		v143 = v120
		v144 = v121
		v145 = v122
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v136 = v131
	v137 = v129
	v138 = v133
	goto L55
L59:
	;
	v128 = int32(4)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v143 = v136
	v144 = v137
	v145 = v138
	goto L52
L62:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 == v154 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v174 = v153 - v154
	goto L50
L64:
	;
	v156 = int32(1)
	v161 = v150 - v156
	if v161 != 0 {
		v148 = v148 + v156
		v149 = v149 + v156
		v150 = v161
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L51
L68:
	;
	v177 = v109
	goto L45
L69:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v180+v83))) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v83
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v187 == v182 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v178 = F__emscripten_memcpy_bulkmem(m, v177, v56, v83)
	mBase = m.M
	goto L72
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v268 = v258
	v275 = v265
	goto L34
L74:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v194 = F_pg_strxfrm(m, v190, v191, v192, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v231) {
		goto L91
	} else {
		goto L92
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v194
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v194) < base.Ui32(v197) {
		v258 = v194
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v201 = v194
	v203 = v197
	goto L79
L79:
	;
	v208 = int32(1)
	v209 = v201 + v208
	v210 = int32(1073741823)
	v212 = v203 << (uint(v208) % 32)
	if base.Ui32(v210) <= base.Ui32(v212) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v258 = v226
	goto L73
L81:
	;
	v215 = v210
	goto L83
L82:
	;
	v215 = v212
	goto L83
L83:
	;
	if base.Ui32(v215) < base.Ui32(v209) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v217 = v209
	goto L86
L85:
	;
	v217 = v215
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v220 = F_repalloc(m, v219, v217)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v220
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v226 = F_pg_strxfrm(m, v220, v223, v224, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v226
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v229) <= base.Ui32(v226) {
		v201 = v226
		v203 = v229
		goto L79
	} else {
		goto L89
	}
L89:
	;
	goto L80
L90:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v253 = m.T0[v252].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v246, int32(4), v248, int32(-1), v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L98
	}
L91:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v246 = v234
	goto L90
L92:
	;
	goto L93
L93:
	;
	v235 = int32(2)
	if base.Ui32(v231) <= base.Ui32(v235) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v238 = v235
	goto L96
L95:
	;
	v238 = v231
	goto L96
L96:
	;
	v240 = v238 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v243 = F_repalloc(m, v242, v240)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v243
	v246 = v243
	goto L90
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v253
	v258 = v253
	goto L73
L99:
	;
	v279 = v276
	goto L101
L100:
	;
	v279 = v268
	goto L101
L101:
	;
	if v279 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v282 = int32(128)
	if v282 <= v83 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v280 = F__emscripten_memcpy_bulkmem(m, v12+int32(12), v275, v279)
	mBase = m.M
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v285 = v282
	goto L108
L107:
	;
	v285 = v83
	goto L108
L108:
	;
	v291 = v285 - int32(1636608432)
	if v56&int32(3) != 0 {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v551 = v14 - int32(-64)
	if int32(129) <= v83 {
		goto L149
	} else {
		goto L150
	}
L110:
	;
	v523 = int32(14)
	v525 = v519 ^ v520 - base.I32_rotl(v519, v523)
	v529 = v525 ^ v518 - base.I32_rotl(v525, int32(11))
	v533 = v529 ^ v519 - base.I32_rotl(v529, int32(25))
	v537 = v533 ^ v525 - base.I32_rotl(v533, int32(16))
	v541 = v537 ^ v529 - base.I32_rotl(v537, int32(4))
	v545 = v541 ^ v533 - base.I32_rotl(v541, v523)
	v549 = v545 ^ v537 - base.I32_rotl(v545, int32(24))
	goto L109
L111:
	;
	switch v449 - int32(1) {
	case 0:
		v511 = v450
		v512 = v451
		v513 = v452
		goto L138
	case 1:
		v504 = v450
		v505 = v451
		v506 = v452
		goto L139
	case 2:
		v497 = v450
		v498 = v451
		v499 = v452
		goto L140
	case 3:
		v491 = v451
		v492 = v452
		goto L141
	case 4:
		v487 = v451
		v488 = v452
		goto L142
	case 5:
		v481 = v451
		v482 = v452
		goto L143
	case 6:
		v475 = v451
		v476 = v452
		goto L144
	case 7:
		v470 = v452
		goto L145
	case 8:
		v465 = v452
		goto L146
	case 9:
		v460 = v452
		goto L147
	case 10:
		goto L148
	default:
		v518 = v450
		v519 = v451
		v520 = v452
		goto L110
	}
L112:
	;
	v400 = v56
	v401 = v285
	v402 = v291
	v403 = v291
	v404 = v291
	goto L135
L113:
	;
	if base.Ui32(int32(11)) < base.Ui32(v285) {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(v285) < base.Ui32(int32(12)) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v448 = v56
	v449 = v285
	v450 = v291
	v451 = v291
	v452 = v291
	goto L111
L117:
	;
	switch v347 - int32(1) {
	case 0:
		v397 = v348
		goto L124
	case 1:
		v392 = v348
		goto L125
	case 2:
		goto L126
	case 3:
		v385 = v349
		goto L127
	case 4:
		v382 = v349
		goto L128
	case 5:
		v377 = v349
		goto L129
	case 6:
		goto L130
	case 7:
		v368 = v350
		goto L131
	case 8:
		v363 = v350
		goto L132
	case 9:
		v358 = v350
		goto L133
	case 10:
		goto L134
	default:
		v518 = v348
		v519 = v349
		v520 = v350
		goto L110
	}
L118:
	;
	v346 = v56
	v347 = v285
	v348 = v291
	v349 = v291
	v350 = v291
	goto L117
L119:
	;
	goto L120
L120:
	;
	v298 = v56
	v299 = v285
	v300 = v291
	v301 = v291
	v302 = v291
	goto L121
L121:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v305 = v304 + v301
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
	v309 = v308 + v302
	v311 = int32(4)
	v313 = v306 + v300 - v309 ^ base.I32_rotl(v309, v311)
	v317 = v305 - v313 ^ base.I32_rotl(v313, int32(6))
	v318 = v309 + v305
	v319 = v313 + v318
	v320 = v317 + v319
	v324 = v318 - v317 ^ base.I32_rotl(v317, int32(8))
	v328 = v319 - v324 ^ base.I32_rotl(v324, int32(16))
	v332 = v320 - v328 ^ base.I32_rotl(v328, int32(19))
	v333 = v324 + v320
	v334 = v328 + v333
	v335 = v332 + v334
	v339 = v333 - v332 ^ base.I32_rotl(v332, v311)
	v340 = int32(12)
	v341 = v298 + v340
	v343 = v299 - v340
	if base.Ui32(int32(11)) < base.Ui32(v343) {
		v298 = v341
		v299 = v343
		v300 = v334
		v301 = v335
		v302 = v339
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v346 = v341
	v347 = v343
	v348 = v334
	v349 = v335
	v350 = v339
	goto L117
L123:
	;
	goto L122
L124:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	v518 = v397 + v398
	v519 = v349
	v520 = v350
	goto L110
L125:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	v397 = v393<<(uint(int32(8))%32) + v392
	goto L124
L126:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+2)))
	v392 = v388<<(uint(int32(16))%32) + v348
	goto L125
L127:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v518 = v386 + v348
	v519 = v385
	v520 = v350
	goto L110
L128:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+4)))
	v385 = v382 + v383
	goto L127
L129:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+5)))
	v382 = v378<<(uint(int32(8))%32) + v377
	goto L128
L130:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+6)))
	v377 = v373<<(uint(int32(16))%32) + v349
	goto L129
L131:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	v518 = v369 + v348
	v519 = v371 + v349
	v520 = v368
	goto L110
L132:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+8)))
	v368 = v364<<(uint(int32(8))%32) + v363
	goto L131
L133:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+9)))
	v363 = v359<<(uint(int32(16))%32) + v358
	goto L132
L134:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+10)))
	v358 = v354<<(uint(int32(24))%32) + v350
	goto L133
L135:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	v407 = v406 + v403
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	v411 = v410 + v404
	v413 = int32(4)
	v415 = v408 + v402 - v411 ^ base.I32_rotl(v411, v413)
	v419 = v407 - v415 ^ base.I32_rotl(v415, int32(6))
	v420 = v411 + v407
	v421 = v415 + v420
	v422 = v419 + v421
	v426 = v420 - v419 ^ base.I32_rotl(v419, int32(8))
	v430 = v421 - v426 ^ base.I32_rotl(v426, int32(16))
	v434 = v422 - v430 ^ base.I32_rotl(v430, int32(19))
	v435 = v426 + v422
	v436 = v430 + v435
	v437 = v434 + v436
	v441 = v435 - v434 ^ base.I32_rotl(v434, v413)
	v442 = int32(12)
	v443 = v400 + v442
	v445 = v401 - v442
	if base.Ui32(int32(11)) < base.Ui32(v445) {
		v400 = v443
		v401 = v445
		v402 = v436
		v403 = v437
		v404 = v441
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v448 = v443
	v449 = v445
	v450 = v436
	v451 = v437
	v452 = v441
	goto L111
L137:
	;
	goto L136
L138:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v518 = v511 + v514
	v519 = v512
	v520 = v513
	goto L110
L139:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+1)))
	v511 = v507<<(uint(int32(8))%32) + v504
	v512 = v505
	v513 = v506
	goto L138
L140:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+2)))
	v504 = v500<<(uint(int32(16))%32) + v497
	v505 = v498
	v506 = v499
	goto L139
L141:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+3)))
	v497 = v493<<(uint(int32(24))%32) + v450
	v498 = v491
	v499 = v492
	goto L140
L142:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+4)))
	v491 = v487 + v489
	v492 = v488
	goto L141
L143:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+5)))
	v487 = v483<<(uint(int32(8))%32) + v481
	v488 = v482
	goto L142
L144:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+6)))
	v481 = v477<<(uint(int32(16))%32) + v475
	v482 = v476
	goto L143
L145:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+7)))
	v475 = v471<<(uint(int32(24))%32) + v451
	v476 = v470
	goto L144
L146:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)))
	v470 = v466<<(uint(int32(8))%32) + v465
	goto L145
L147:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+9)))
	v465 = v461<<(uint(int32(16))%32) + v460
	goto L146
L148:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+10)))
	v460 = v456<<(uint(int32(24))%32) + v452
	goto L147
L149:
	;
	v558 = int32(711645284)
	v561 = v83 - int32(1636608428) ^ v558 - int32(1455628627)
	v566 = v561 ^ int32(-1636608428) - base.I32_rotl(v561, int32(25))
	v571 = v566 ^ v558 - base.I32_rotl(v566, int32(16))
	v575 = v571 ^ v561 - base.I32_rotl(v571, int32(4))
	v579 = v575 ^ v566 - base.I32_rotl(v575, int32(14))
	goto L152
L150:
	;
	v585 = v549
	goto L151
L151:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v590 = int32(32) - v589
	v591 = v585 << (uint(v589) % 32)
	if v591 != 0 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v585 = v579 ^ v571 - base.I32_rotl(v579, int32(24)) ^ v549
	goto L151
L153:
	;
	v619 = v14 + int32(40)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v625 = int32(711645284)
	v628 = v620 - int32(1636608428) ^ v625 - int32(1455628627)
	v633 = v628 ^ int32(-1636608428) - base.I32_rotl(v628, int32(25))
	v638 = v633 ^ v625 - base.I32_rotl(v633, int32(16))
	v642 = v638 ^ v628 - base.I32_rotl(v638, int32(4))
	v646 = v642 ^ v633 - base.I32_rotl(v642, int32(14))
	v650 = v646 ^ v638 - base.I32_rotl(v646, int32(24))
	goto L163
L154:
	;
	v598 = int32(32) - (base.I32_clz(v591) ^ int32(31))
	v599 = int32(255)
	if base.Ui32(v590&v599) < base.Ui32(v598&v599) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v608 = v590 + int32(1)
	goto L156
L156:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v551)+16))
	v611 = v609 + int32(base.Ui32(v585)>>(uint(v590)%32))
	v613 = v608 & int32(255)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	if base.Ui32(v614) < base.Ui32(v613) {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v604 = v590 + int32(1)
	goto L159
L158:
	;
	v604 = v598
	goto L159
L159:
	;
	v608 = v604
	goto L156
L160:
	;
	v616 = v613
	goto L162
L161:
	;
	v616 = v614
	goto L162
L162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v611))) = uint8(v616)
	goto L153
L163:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	v655 = int32(32) - v654
	v656 = v650 << (uint(v654) % 32)
	if v656 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v683 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v683)
	v698 = v620
	goto L30
L165:
	;
	v663 = int32(32) - (base.I32_clz(v656) ^ int32(31))
	v664 = int32(255)
	if base.Ui32(v655&v664) < base.Ui32(v663&v664) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	v673 = v655 + int32(1)
	goto L167
L167:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	v676 = v674 + int32(base.Ui32(v650)>>(uint(v655)%32))
	v678 = v673 & int32(255)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if base.Ui32(v679) < base.Ui32(v678) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v669 = v655 + int32(1)
	goto L170
L169:
	;
	v669 = v663
	goto L170
L170:
	;
	v673 = v669
	goto L167
L171:
	;
	v681 = v678
	goto L173
L172:
	;
	v681 = v679
	goto L173
L173:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v681)
	goto L164
L174:
	;
	v692 = v688
	goto L176
L175:
	;
	v692 = v689
	goto L176
L176:
	;
	if v692 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v698 = v695
	goto L30
L178:
	;
	v693 = F__emscripten_memcpy_bulkmem(m, v12+int32(12), v687, v692)
	mBase = m.M
	goto L180
L179:
	;
	goto L180
L180:
	;
	goto L177
L181:
	;
	F_pfree(m, v15)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	m.G0 = v12 + int32(16)
	v711 = int32(24)
	v713 = int32(65280)
	v715 = int32(8)
	return v698<<(uint(v711)%32) | v698&v713<<(uint(v715)%32) | (int32(base.Ui32(v698)>>(uint(v715)%32))&v713 | int32(base.Ui32(v698)>>(uint(v711)%32)))
L184:
	;
	goto L183
}
func F_varstr_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = F_pg_newlocale_from_collation(m, l4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L80
	}
L4:
	;
	return v219
L5:
	;
	return int32(0)
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)))
	if v11 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = base.B2i32(l1 < l3)
	if l1 < l3 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if l1 != l3 {
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v15 = l1
	goto L12
L11:
	;
	v15 = l3
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v15) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v77 != 0 {
		v219 = v77
		goto L4
	} else {
		goto L31
	}
L14:
	;
	v77 = int32(0)
	goto L13
L15:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L25
L16:
	;
	if (l0|l2)&int32(3) != 0 {
		v46 = l0
		v47 = l2
		v48 = v15
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v39 = l0
	v40 = l2
	v41 = v15
	goto L18
L18:
	;
	if v41 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v23 = l0
	v24 = l2
	v25 = v15
	goto L20
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L18
L22:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L15
L25:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v77 = v56 - v57
	goto L13
L27:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	return base.B2i32(l3 < l1) - v14
L32:
	;
	v146 = F_pg_strncoll(m, l0, l1, l2, l3, v7)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L53
	}
L33:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v143 != 0 {
		goto L32
	} else {
		goto L52
	}
L35:
	;
	v143 = int32(0)
	goto L34
L36:
	;
	v117 = v112
	v118 = v113
	v119 = v114
	goto L46
L37:
	;
	if (l0|l2)&int32(3) != 0 {
		v112 = l0
		v113 = l2
		v114 = l1
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v105 = l0
	v106 = l2
	v107 = l1
	goto L39
L39:
	;
	if v107 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v89 = l0
	v90 = l2
	v91 = l1
	goto L41
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v94 != v95 {
		v112 = v89
		v113 = v90
		v114 = v91
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v105 = v100
	v106 = v98
	v107 = v102
	goto L39
L43:
	;
	v97 = int32(4)
	v98 = v90 + v97
	v100 = v89 + v97
	v102 = v91 - v97
	if base.Ui32(int32(3)) < base.Ui32(v102) {
		v89 = v100
		v90 = v98
		v91 = v102
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v112 = v105
	v113 = v106
	v114 = v107
	goto L36
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 == v123 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v143 = v122 - v123
	goto L34
L48:
	;
	v125 = int32(1)
	v130 = v119 - v125
	if v130 != 0 {
		v117 = v117 + v125
		v118 = v118 + v125
		v119 = v130
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	return int32(0)
L53:
	;
	if v146 != 0 {
		v219 = v146
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v148 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	return int32(0)
L56:
	;
	goto L57
L57:
	;
	v153 = base.B2i32(l1 < l3)
	if l1 < l3 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = l1
	goto L60
L59:
	;
	v154 = l3
	goto L60
L60:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v154) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v216 != 0 {
		v219 = v216
		goto L4
	} else {
		goto L79
	}
L62:
	;
	v216 = int32(0)
	goto L61
L63:
	;
	v190 = v185
	v191 = v186
	v192 = v187
	goto L73
L64:
	;
	if (l0|l2)&int32(3) != 0 {
		v185 = l0
		v186 = l2
		v187 = v154
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v178 = l0
	v179 = l2
	v180 = v154
	goto L66
L66:
	;
	if v180 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v162 = l0
	v163 = l2
	v164 = v154
	goto L68
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v167 != v168 {
		v185 = v162
		v186 = v163
		v187 = v164
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v178 = v173
	v179 = v171
	v180 = v175
	goto L66
L70:
	;
	v170 = int32(4)
	v171 = v163 + v170
	v173 = v162 + v170
	v175 = v164 - v170
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v162 = v173
		v163 = v171
		v164 = v175
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v185 = v178
	v186 = v179
	v187 = v180
	goto L63
L73:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v195 == v196 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v216 = v195 - v196
	goto L61
L75:
	;
	v198 = int32(1)
	v203 = v192 - v198
	if v203 != 0 {
		v190 = v190 + v198
		v191 = v191 + v198
		v192 = v203
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L62
L79:
	;
	v219 = base.B2i32(l3 < l1) - v153
	goto L4
L80:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(258179), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	F_errhint(m, int32(602565), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(525792), int32(1648), int32(113268))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_varstrfastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if l1 != l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v74 == int32(1042) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v71 = int32(0)
	goto L3
L5:
	;
	v45 = v40
	v46 = v41
	v47 = v42
	goto L15
L6:
	;
	if (l0|l2)&int32(3) != 0 {
		v40 = l0
		v41 = l2
		v42 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v33 = l0
	v34 = l2
	v35 = l1
	goto L8
L8:
	;
	if v35 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v17 = l0
	v18 = l2
	v19 = l1
	goto L10
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != v23 {
		v40 = v17
		v41 = v18
		v42 = v19
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v33 = v28
	v34 = v26
	v35 = v30
	goto L8
L12:
	;
	v25 = int32(4)
	v26 = v18 + v25
	v28 = v17 + v25
	v30 = v19 - v25
	if base.Ui32(int32(3)) < base.Ui32(v30) {
		v17 = v28
		v18 = v26
		v19 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v40 = v33
	v41 = v34
	v42 = v35
	goto L5
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 == v51 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v71 = v50 - v51
	goto L3
L17:
	;
	v53 = int32(1)
	v58 = v47 - v53
	if v58 != 0 {
		v45 = v45 + v53
		v46 = v46 + v53
		v47 = v58
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	return int32(0)
L22:
	;
	v79 = int32(-1)
	v81 = l1 - int32(1)
	if v79 <= v81 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v123 = l1
	v124 = l3
	goto L24
L24:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v125 <= v123 {
		goto L45
	} else {
		goto L46
	}
L25:
	;
	v102 = int32(-1)
	v104 = l3 - int32(1)
	if v102 <= v104 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v84 = v79
	goto L28
L27:
	;
	v84 = v81
	goto L28
L28:
	;
	v88 = l1
	goto L29
L29:
	;
	v92 = v88 - int32(1)
	if v92 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v99 = v88
	goto L25
L31:
	;
	v99 = v84 + int32(1)
	goto L25
L32:
	;
	goto L33
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92))))
	if v96 == int32(32) {
		v88 = v92
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v123 = v99
	v124 = v122
	goto L24
L36:
	;
	v107 = v102
	goto L38
L37:
	;
	v107 = v104
	goto L38
L38:
	;
	v111 = l3
	goto L39
L39:
	;
	v115 = v111 - int32(1)
	if v115 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v122 = v111
	goto L35
L41:
	;
	v122 = v107 + int32(1)
	goto L35
L42:
	;
	goto L43
L43:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v115))))
	if v119 == int32(32) {
		v111 = v115
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v127 = int32(1)
	v128 = v123 + v127
	v129 = int32(1073741823)
	v131 = v125 << (uint(v127) % 32)
	if base.Ui32(v129) <= base.Ui32(v131) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v146 <= v124 {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v134 = v129
	goto L50
L49:
	;
	v134 = v131
	goto L50
L50:
	;
	if base.Ui32(v134) < base.Ui32(v128) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v136 = v128
	goto L53
L52:
	;
	v136 = v134
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v139 = F_repalloc(m, v138, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return int32(0)
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v139
	goto L47
L56:
	;
	v148 = int32(1)
	v149 = v124 + v148
	v150 = int32(1073741823)
	v152 = v146 << (uint(v148) % 32)
	if base.Ui32(v150) <= base.Ui32(v152) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v123 != v166 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v155 = v150
	goto L61
L60:
	;
	v155 = v152
	goto L61
L61:
	;
	if base.Ui32(v155) < base.Ui32(v149) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v157 = v149
	goto L64
L63:
	;
	v157 = v155
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v160 = F_repalloc(m, v159, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v160
	goto L58
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v241 == v124 {
		goto L94
	} else {
		goto L95
	}
L67:
	;
	if v123 != 0 {
		goto L89
	} else {
		goto L90
	}
L68:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v123) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if v229 != 0 {
		goto L67
	} else {
		goto L87
	}
L70:
	;
	v229 = int32(0)
	goto L69
L71:
	;
	v203 = v198
	v204 = v199
	v205 = v200
	goto L81
L72:
	;
	if (v165|l0)&int32(3) != 0 {
		v198 = v165
		v199 = l0
		v200 = v123
		goto L71
	} else {
		goto L75
	}
L73:
	;
	v191 = v165
	v192 = l0
	v193 = v123
	goto L74
L74:
	;
	if v193 == int32(0) {
		goto L70
	} else {
		goto L80
	}
L75:
	;
	v175 = v165
	v176 = l0
	v177 = v123
	goto L76
L76:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v180 != v181 {
		v198 = v175
		v199 = v176
		v200 = v177
		goto L71
	} else {
		goto L78
	}
L77:
	;
	v191 = v186
	v192 = v184
	v193 = v188
	goto L74
L78:
	;
	v183 = int32(4)
	v184 = v176 + v183
	v186 = v175 + v183
	v188 = v177 - v183
	if base.Ui32(int32(3)) < base.Ui32(v188) {
		v175 = v186
		v176 = v184
		v177 = v188
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v198 = v191
	v199 = v192
	v200 = v193
	goto L71
L81:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v208 == v209 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v229 = v208 - v209
	goto L69
L83:
	;
	v211 = int32(1)
	v216 = v205 - v211
	if v216 != 0 {
		v203 = v203 + v211
		v204 = v204 + v211
		v205 = v216
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	goto L70
L87:
	;
	v239 = int32(1)
	goto L66
L88:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v233+v123))) = uint8(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v123
	v239 = v235
	goto L66
L89:
	;
	v231 = F__emscripten_memcpy_bulkmem(m, v165, l0, v123)
	mBase = m.M
	goto L91
L90:
	;
	goto L91
L91:
	;
	goto L88
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v322 = int32(-1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v327 = m.T0[v326].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v321, v322, v320, v322, v324)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L54
	} else {
		goto L123
	}
L93:
	;
	if v239 == int32(0) {
		v320 = v240
		goto L92
	} else {
		goto L120
	}
L94:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v124) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	goto L96
L96:
	;
	if v124 != 0 {
		goto L117
	} else {
		goto L118
	}
L97:
	;
	if v304 == int32(0) {
		goto L93
	} else {
		goto L115
	}
L98:
	;
	v304 = int32(0)
	goto L97
L99:
	;
	v278 = v273
	v279 = v274
	v280 = v275
	goto L109
L100:
	;
	if (v240|l2)&int32(3) != 0 {
		v273 = v240
		v274 = l2
		v275 = v124
		goto L99
	} else {
		goto L103
	}
L101:
	;
	v266 = v240
	v267 = l2
	v268 = v124
	goto L102
L102:
	;
	if v268 == int32(0) {
		goto L98
	} else {
		goto L108
	}
L103:
	;
	v250 = v240
	v251 = l2
	v252 = v124
	goto L104
L104:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v255 != v256 {
		v273 = v250
		v274 = v251
		v275 = v252
		goto L99
	} else {
		goto L106
	}
L105:
	;
	v266 = v261
	v267 = v259
	v268 = v263
	goto L102
L106:
	;
	v258 = int32(4)
	v259 = v251 + v258
	v261 = v250 + v258
	v263 = v252 - v258
	if base.Ui32(int32(3)) < base.Ui32(v263) {
		v250 = v261
		v251 = v259
		v252 = v263
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v273 = v266
	v274 = v267
	v275 = v268
	goto L99
L109:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v283 == v284 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v304 = v283 - v284
	goto L97
L111:
	;
	v286 = int32(1)
	v291 = v280 - v286
	if v291 != 0 {
		v278 = v278 + v286
		v279 = v279 + v286
		v280 = v291
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L110
L114:
	;
	goto L98
L115:
	;
	goto L96
L116:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v311 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v309+v124))) = uint8(v311)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v124
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v320 = v314
	goto L92
L117:
	;
	v307 = F__emscripten_memcpy_bulkmem(m, v240, l2, v124)
	mBase = m.M
	goto L119
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
	if v317 != 0 {
		v320 = v240
		goto L92
	} else {
		goto L121
	}
L121:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	return v318
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v361
	v363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v363)
	return v361
L123:
	;
	if v327 != 0 {
		v361 = v327
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v331 != int32(1) {
		v361 = int32(0)
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v339 == int32(0) {
		v358 = v338
		v359 = v339
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v361 = v359 - v358
	goto L122
L127:
	;
	goto L126
L128:
	;
	if v338 != v339 {
		v358 = v338
		v359 = v339
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v343 = v334
	v344 = v335
	goto L130
L130:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v348 == int32(0) {
		v358 = v347
		v359 = v348
		goto L127
	} else {
		goto L132
	}
L131:
	;
	v358 = v347
	v359 = v348
	goto L127
L132:
	;
	v351 = int32(1)
	if v347 == v348 {
		v343 = v343 + v351
		v344 = v344 + v351
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
}
func F_void_out(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pstrdup(m, int32(791891))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}

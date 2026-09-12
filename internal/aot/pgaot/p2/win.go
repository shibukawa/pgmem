package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinGetFuncArgInFrame(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+400))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	switch l2 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(396072), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523219), int32(3482), int32(396113))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		if l1 < int32(0) {
			v217 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
			v225 = int32(0)
			m.G0 = v16 + int32(48)
			return v225
		} else {
			F_update_frameheadpos(m, v18)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
				v42 = v40 + base.I64_extend_i32_u(l1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
				switch int32(base.Ui32(v43)>>(uint(int32(15))%32)) & int32(7) {
				case 0:
					v192 = v42
					v193 = v42
					v196 = F_window_gettupleslot(m, l0, v192, v19)
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
						return int32(0)
					} else {
						if v196 == int32(0) {
							v217 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
							v225 = int32(0)
							m.G0 = v16 + int32(48)
							return v225
						} else {
							v200 = F_row_is_in_frame(m, v18, v192, v19)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return int32(0)
							} else {
								if v200 <= int32(0) {
									v217 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
									v225 = int32(0)
									m.G0 = v16 + int32(48)
									return v225
								} else {
									if l3 != 0 {
										F_WinSetMarkPosition(m, l0, v193)
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
											v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
											v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												v225 = v211
												m.G0 = v16 + int32(48)
												return v225
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
										v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
										v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
										v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return int32(0)
										} else {
											v225 = v211
											m.G0 = v16 + int32(48)
											return v225
										}
									}
								}
							}
						}
					}
				case 1:
					v176 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
					v192 = v42 + base.I64_extend_i32_u(base.B2i32(v176 <= v42)&base.B2i32(v40 <= v176))
					v193 = v42
					v196 = F_window_gettupleslot(m, l0, v192, v19)
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
						return int32(0)
					} else {
						if v196 == int32(0) {
							v217 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
							v225 = int32(0)
							m.G0 = v16 + int32(48)
							return v225
						} else {
							v200 = F_row_is_in_frame(m, v18, v192, v19)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return int32(0)
							} else {
								if v200 <= int32(0) {
									v217 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
									v225 = int32(0)
									m.G0 = v16 + int32(48)
									return v225
								} else {
									if l3 != 0 {
										F_WinSetMarkPosition(m, l0, v193)
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
											v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
											v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												v225 = v211
												m.G0 = v16 + int32(48)
												return v225
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
										v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
										v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
										v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return int32(0)
										} else {
											v225 = v211
											m.G0 = v16 + int32(48)
											return v225
										}
									}
								}
							}
						}
					}
				case 2:
					F_update_grouptailpos(m, v18)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v18)+344))
						if v42 < v50 {
							v192 = v42
							v193 = v42
						} else {
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
							v53 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							if v52 <= v53 {
								v192 = v42
								v193 = v42
							} else {
								if v53 < v50 {
									v57 = v50
								} else {
									v57 = v53
								}
								v192 = v42 + v52 - v57
								v193 = v42
							}
						}
						v196 = F_window_gettupleslot(m, l0, v192, v19)
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return int32(0)
						} else {
							if v196 == int32(0) {
								v217 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
								v225 = int32(0)
								m.G0 = v16 + int32(48)
								return v225
							} else {
								v200 = F_row_is_in_frame(m, v18, v192, v19)
								mBase = m.M
								v201 = m.ExcPending
								if v201 != 0 {
									return int32(0)
								} else {
									if v200 <= int32(0) {
										v217 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
										v225 = int32(0)
										m.G0 = v16 + int32(48)
										return v225
									} else {
										if l3 != 0 {
											F_WinSetMarkPosition(m, l0, v193)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
												v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
												v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
												v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													v225 = v211
													m.G0 = v16 + int32(48)
													return v225
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
											v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
											v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												v225 = v211
												m.G0 = v16 + int32(48)
												return v225
											}
										}
									}
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v78
						F_errmsg_internal(m, int32(31108), v16+int32(16))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523219), int32(3542), int32(396113))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 4:
					F_update_grouptailpos(m, v18)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = *(*int64)(unsafe.Add(mBase, uint32(v18)+344))
						if v42 < v61 {
							v192 = v42
							v193 = v42
						} else {
							v63 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
							v64 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							if v63 <= v64 {
								v192 = v42
								v193 = v42
							} else {
								if v64 < v61 {
									v67 = v61
								} else {
									v67 = v64
								}
								if v67 == v42 {
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
									v192 = v69
									v193 = v42
								} else {
									v192 = v42 + v63 + (v67 ^ int64(-1))
									v193 = v42
								}
							}
						}
						v196 = F_window_gettupleslot(m, l0, v192, v19)
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return int32(0)
						} else {
							if v196 == int32(0) {
								v217 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
								v225 = int32(0)
								m.G0 = v16 + int32(48)
								return v225
							} else {
								v200 = F_row_is_in_frame(m, v18, v192, v19)
								mBase = m.M
								v201 = m.ExcPending
								if v201 != 0 {
									return int32(0)
								} else {
									if v200 <= int32(0) {
										v217 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
										v225 = int32(0)
										m.G0 = v16 + int32(48)
										return v225
									} else {
										if l3 != 0 {
											F_WinSetMarkPosition(m, l0, v193)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
												v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
												v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
												v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													v225 = v211
													m.G0 = v16 + int32(48)
													return v225
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
											v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
											v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												v225 = v211
												m.G0 = v16 + int32(48)
												return v225
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
	case 2:
		if int32(0) < l1 {
			v217 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
			v225 = int32(0)
			m.G0 = v16 + int32(48)
			return v225
		} else {
			F_update_frametailpos(m, v18)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
				v98 = v94 + base.I64_extend_i32_s(l1) - int64(1)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
				switch int32(base.Ui32(v99)>>(uint(int32(15))%32)) & int32(7) {
				case 0:
					v192 = v98
					v193 = v98
					v196 = F_window_gettupleslot(m, l0, v192, v19)
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
						return int32(0)
					} else {
						if v196 == int32(0) {
							v217 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
							v225 = int32(0)
							m.G0 = v16 + int32(48)
							return v225
						} else {
							v200 = F_row_is_in_frame(m, v18, v192, v19)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return int32(0)
							} else {
								if v200 <= int32(0) {
									v217 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
									v225 = int32(0)
									m.G0 = v16 + int32(48)
									return v225
								} else {
									if l3 != 0 {
										F_WinSetMarkPosition(m, l0, v193)
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
											v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
											v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												v225 = v211
												m.G0 = v16 + int32(48)
												return v225
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
										v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
										v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
										v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return int32(0)
										} else {
											v225 = v211
											m.G0 = v16 + int32(48)
											return v225
										}
									}
								}
							}
						}
					}
				case 1:
					v182 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
					F_update_frameheadpos(m, v18)
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						v189 = v98 - base.I64_extend_i32_u(base.B2i32(v182 < v94)&base.B2i32(v98 <= v182))
						v190 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
						if v189 < v190 {
							v217 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
							v225 = int32(0)
							m.G0 = v16 + int32(48)
							return v225
						} else {
							v192 = v189
							v193 = v190
							v196 = F_window_gettupleslot(m, l0, v192, v19)
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
								return int32(0)
							} else {
								if v196 == int32(0) {
									v217 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
									v225 = int32(0)
									m.G0 = v16 + int32(48)
									return v225
								} else {
									v200 = F_row_is_in_frame(m, v18, v192, v19)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int32(0)
									} else {
										if v200 <= int32(0) {
											v217 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
											v225 = int32(0)
											m.G0 = v16 + int32(48)
											return v225
										} else {
											if l3 != 0 {
												F_WinSetMarkPosition(m, l0, v193)
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
													v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
													v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														v225 = v211
														m.G0 = v16 + int32(48)
														return v225
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
												v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
												v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
												v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													v225 = v211
													m.G0 = v16 + int32(48)
													return v225
												}
											}
										}
									}
								}
							}
						}
					}
				case 2:
					F_update_grouptailpos(m, v18)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
						if v106 <= v98 {
							v117 = v98
						} else {
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v18)+344))
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
							if v109 <= v108 {
								v117 = v98
							} else {
								if v106 < v109 {
									v113 = v106
								} else {
									v113 = v109
								}
								v117 = v98 + v108 - v113
							}
						}
						F_update_frameheadpos(m, v18)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							if v120 <= v117 {
								v192 = v117
								v193 = v120
								v196 = F_window_gettupleslot(m, l0, v192, v19)
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									if v196 == int32(0) {
										v217 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
										v225 = int32(0)
										m.G0 = v16 + int32(48)
										return v225
									} else {
										v200 = F_row_is_in_frame(m, v18, v192, v19)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											if v200 <= int32(0) {
												v217 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
												v225 = int32(0)
												m.G0 = v16 + int32(48)
												return v225
											} else {
												if l3 != 0 {
													F_WinSetMarkPosition(m, l0, v193)
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
														v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
														v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
														v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return int32(0)
														} else {
															v225 = v211
															m.G0 = v16 + int32(48)
															return v225
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
													v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
													v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														v225 = v211
														m.G0 = v16 + int32(48)
														return v225
													}
												}
											}
										}
									}
								}
							} else {
								v217 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
								v225 = int32(0)
								m.G0 = v16 + int32(48)
								return v225
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int32(0)
					} else {
						v151 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v151
						F_errmsg_internal(m, int32(31108), v16+int32(32))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523219), int32(3612), int32(396113))
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 4:
					F_update_grouptailpos(m, v18)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v18)+352))
						if v124 <= v98 {
							v142 = v98
						} else {
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v18)+344))
							v127 = *(*int64)(unsafe.Add(mBase, uint32(v18)+192))
							if v127 <= v126 {
								v142 = v98
							} else {
								if v124 < v127 {
									v130 = v124
								} else {
									v130 = v127
								}
								if v130-int64(1) == v98 {
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
									v142 = v134
								} else {
									v142 = v98 + v126 - v130 + int64(1)
								}
							}
						}
						F_update_frameheadpos(m, v18)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							v145 = *(*int64)(unsafe.Add(mBase, uint32(v18)+184))
							if v145 <= v142 {
								v192 = v142
								v193 = v145
								v196 = F_window_gettupleslot(m, l0, v192, v19)
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									if v196 == int32(0) {
										v217 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
										v225 = int32(0)
										m.G0 = v16 + int32(48)
										return v225
									} else {
										v200 = F_row_is_in_frame(m, v18, v192, v19)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											if v200 <= int32(0) {
												v217 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
												v225 = int32(0)
												m.G0 = v16 + int32(48)
												return v225
											} else {
												if l3 != 0 {
													F_WinSetMarkPosition(m, l0, v193)
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
														v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
														v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
														v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return int32(0)
														} else {
															v225 = v211
															m.G0 = v16 + int32(48)
															return v225
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v19
													v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
													v211 = m.T0[v210].(func(*base.Module, int32, int32, int32) int32)(m, v209, v20, l4)
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														v225 = v211
														m.G0 = v16 + int32(48)
														return v225
													}
												}
											}
										}
									}
								}
							} else {
								v217 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v217)
								v225 = int32(0)
								m.G0 = v16 + int32(48)
								return v225
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v166 = m.ExcPending
		if v166 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
			F_errmsg_internal(m, int32(508498), v16)
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523219), int32(3618), int32(396113))
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
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
func F_WinSetMarkPosition(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 <= l1 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		F_tuplestore_select_read_pointer(m, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			if v12 < l1 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
				v17 = F_tuplestore_skiptuples(m, v14, l1-v12, int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					F_tuplestore_select_read_pointer(m, v20, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						if v24 < l1 {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
							v29 = F_tuplestore_skiptuples(m, v26, l1-v24, int32(1))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
								return
							}
						} else {
							return
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_tuplestore_select_read_pointer(m, v20, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					if v24 < l1 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+144))
						v29 = F_tuplestore_skiptuples(m, v26, l1-v24, int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = l1
							return
						}
					} else {
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(442436), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errfinish(m, int32(523219), int32(3292), int32(262465))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
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

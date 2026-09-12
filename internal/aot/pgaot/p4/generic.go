package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_Generic_Text_IC_like(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	if l2 != 0 {
		v10 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v227 = m.ExcPending
				if v227 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v230 = m.ExcPending
					if v230 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(521187), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(479844), int32(202), int32(382879))
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
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
				v18 = *(*int32)(unsafe.Add(mBase, _consts[485]))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_consts[1294])))
				if v24 <= int32(1) {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
					if v27 != int32(105) {
						v127 = int32(1)
						v128 = l1 + v127
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v131 = v129 & v127
						if v129 == v127 {
							v134 = int32(4)
							v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
							if v136&int32(254) == int32(2) {
								v145 = v134
							} else {
								v145 = base.B2i32(v136 == int32(18)) << (uint(v134) % 32)
							}
							if v136 == int32(1) {
								v148 = v134
							} else {
								v148 = v145
							}
							v159 = v148
						} else {
							v149 = int32(1)
							if v131 != 0 {
								v159 = int32(base.Ui32(v129)>>(uint(v149)%32)) - v149
							} else {
								v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v159 = int32(base.Ui32(v153)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v131 != 0 {
							v160 = v128
						} else {
							v160 = l1 + int32(4)
						}
						v161 = int32(1)
						v162 = l0 + v161
						v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						v167 = v165 & v161
						if v167 != 0 {
							v168 = v162
						} else {
							v168 = l0 + int32(4)
						}
						if v165 == int32(1) {
							v171 = int32(4)
							v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
							if v173&int32(254) == int32(2) {
								v182 = v171
							} else {
								v182 = base.B2i32(v173 == int32(18)) << (uint(v171) % 32)
							}
							if v173 == int32(1) {
								v185 = v171
							} else {
								v185 = v182
							}
							v186 = F_SB_IMatchText(m, v168, v185, v160, v159, v10)
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int32(0)
							} else {
								return v186
							}
						} else {
							if v167 != 0 {
								v189 = int32(1)
								v193 = F_SB_IMatchText(m, v168, int32(base.Ui32(v165)>>(uint(v189)%32))-v189, v160, v159, v10)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									return v193
								}
							} else {
								v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v201 = F_SB_IMatchText(m, v168, int32(base.Ui32(v196)>>(uint(int32(2))%32))-int32(4), v160, v159, v10)
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return int32(0)
								} else {
									return v201
								}
							}
						}
					} else {
						v31 = F_DirectFunctionCall1Coll(m, int32(1452), l2, l1)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_pg_detoast_datum_packed(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = int32(1)
								v36 = v33 + v35
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
								v39 = v37 & v35
								if v37 == v35 {
									v42 = int32(4)
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
									if v44&int32(254) == int32(2) {
										v53 = v42
									} else {
										v53 = base.B2i32(v44 == int32(18)) << (uint(v42) % 32)
									}
									if v44 == int32(1) {
										v56 = v42
									} else {
										v56 = v53
									}
									v67 = v56
								} else {
									v57 = int32(1)
									if v39 != 0 {
										v67 = int32(base.Ui32(v37)>>(uint(v57)%32)) - v57
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								v71 = F_DirectFunctionCall1Coll(m, int32(1452), l2, l0)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v73 = F_pg_detoast_datum_packed(m, v71)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v77 = int32(1)
										v78 = v73 + v77
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
										v81 = v79 & v77
										if v79 == v77 {
											v84 = int32(4)
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
											if v86&int32(254) == int32(2) {
												v95 = v84
											} else {
												v95 = base.B2i32(v86 == int32(18)) << (uint(v84) % 32)
											}
											if v86 == int32(1) {
												v98 = v84
											} else {
												v98 = v95
											}
											v109 = v98
										} else {
											v99 = int32(1)
											if v81 != 0 {
												v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
												v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										if v39 != 0 {
											v110 = v36
										} else {
											v110 = v33 + int32(4)
										}
										if v81 != 0 {
											v111 = v78
										} else {
											v111 = v73 + int32(4)
										}
										v113 = *(*int32)(unsafe.Add(mBase, _consts[485]))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
										if v114 == int32(6) {
											v118 = F_UTF8_MatchText(m, v111, v109, v110, v67, int32(0))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												return v118
											}
										} else {
											v122 = F_MB_MatchText(m, v111, v109, v110, v67, int32(0))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												return v122
											}
										}
									}
								}
							}
						}
					}
				} else {
					v31 = F_DirectFunctionCall1Coll(m, int32(1452), l2, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_pg_detoast_datum_packed(m, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = int32(1)
							v36 = v33 + v35
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
							v39 = v37 & v35
							if v37 == v35 {
								v42 = int32(4)
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
								if v44&int32(254) == int32(2) {
									v53 = v42
								} else {
									v53 = base.B2i32(v44 == int32(18)) << (uint(v42) % 32)
								}
								if v44 == int32(1) {
									v56 = v42
								} else {
									v56 = v53
								}
								v67 = v56
							} else {
								v57 = int32(1)
								if v39 != 0 {
									v67 = int32(base.Ui32(v37)>>(uint(v57)%32)) - v57
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v71 = F_DirectFunctionCall1Coll(m, int32(1452), l2, l0)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = F_pg_detoast_datum_packed(m, v71)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v77 = int32(1)
									v78 = v73 + v77
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
									v81 = v79 & v77
									if v79 == v77 {
										v84 = int32(4)
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
										if v86&int32(254) == int32(2) {
											v95 = v84
										} else {
											v95 = base.B2i32(v86 == int32(18)) << (uint(v84) % 32)
										}
										if v86 == int32(1) {
											v98 = v84
										} else {
											v98 = v95
										}
										v109 = v98
									} else {
										v99 = int32(1)
										if v81 != 0 {
											v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
											v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									if v39 != 0 {
										v110 = v36
									} else {
										v110 = v33 + int32(4)
									}
									if v81 != 0 {
										v111 = v78
									} else {
										v111 = v73 + int32(4)
									}
									v113 = *(*int32)(unsafe.Add(mBase, _consts[485]))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
									if v114 == int32(6) {
										v118 = F_UTF8_MatchText(m, v111, v109, v110, v67, int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											return v118
										}
									} else {
										v122 = F_MB_MatchText(m, v111, v109, v110, v67, int32(0))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											return v122
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
		v207 = m.ExcPending
		if v207 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v210 = m.ExcPending
			if v210 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(521134), int32(0))
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(538820), int32(0))
					mBase = m.M
					v218 = m.ExcPending
					if v218 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(479844), int32(194), int32(382879))
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
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

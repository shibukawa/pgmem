package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v11 != 0 {
			v57 = v11
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v58 != 0 {
			} else {
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
				*(*int64)(unsafe.Add(mBase, uint32(v57))) = v59 + int64(1)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v63 == int32(0) {
				} else {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v66 + int64(1)
				}
			}
			m.G0 = v6 + int32(16)
			return v57
		} else {
			v14 = v6 + int32(12)
			v15 = int32(0)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v16 == v15 {
				v33 = int32(0)
				if v14 == v33 {
					v41 = v33
				} else {
					v36 = v33
					v37 = v15
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
				}
				v44 = v41
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				switch v19 - int32(429) {
				case 0:
					if v14 == int32(0) {
						v44 = int32(1)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						v36 = v26
						v37 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
						v41 = v37
						v44 = v41
					}
				case 1:
					if v14 == int32(0) {
						v44 = int32(2)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
						v36 = v31
						v37 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
						v41 = v37
						v44 = v41
					}
				default:
					v33 = int32(0)
					if v14 == v33 {
						v41 = v33
					} else {
						v36 = v33
						v37 = v15
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
						v41 = v37
					}
					v44 = v41
				}
			}
			if v44 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(66285), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(517423), int32(330), int32(368459))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v49 = F_MemoryContextAlloc(m, v47, int32(16))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53
					*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v53
					v57 = v49
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v58 != 0 {
					} else {
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
						*(*int64)(unsafe.Add(mBase, uint32(v57))) = v59 + int64(1)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v63 == int32(0) {
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v66 + int64(1)
						}
					}
					m.G0 = v6 + int32(16)
					return v57
				}
			}
		}
	} else {
		v14 = v6 + int32(12)
		v15 = int32(0)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v16 == v15 {
			v33 = int32(0)
			if v14 == v33 {
				v41 = v33
			} else {
				v36 = v33
				v37 = v15
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
			}
			v44 = v41
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			switch v19 - int32(429) {
			case 0:
				if v14 == int32(0) {
					v44 = int32(1)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
					v36 = v26
					v37 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
					v44 = v41
				}
			case 1:
				if v14 == int32(0) {
					v44 = int32(2)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
					v36 = v31
					v37 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
					v44 = v41
				}
			default:
				v33 = int32(0)
				if v14 == v33 {
					v41 = v33
				} else {
					v36 = v33
					v37 = v15
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
					v41 = v37
				}
				v44 = v41
			}
		}
		if v44 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(66285), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(517423), int32(330), int32(368459))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v49 = F_MemoryContextAlloc(m, v47, int32(16))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v53 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v53
				v57 = v49
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v58 != 0 {
				} else {
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = v59 + int64(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v63 == int32(0) {
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v66 + int64(1)
					}
				}
				m.G0 = v6 + int32(16)
				return v57
			}
		}
	}
}
func F_bool_alltrue(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 != 0 {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		return int32(0)
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v5 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v8 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int64(0) {
				v17 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
				return base.B2i32(v17 == v8)
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		}
	}
}
func F_executeBoolItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	F_check_stack_depth(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l3 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if int32(0) < v17 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v213 = m.ExcPending
				if v213 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(302479), int32(0))
					mBase = m.M
					v217 = m.ExcPending
					if v217 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(520430), int32(1788), int32(302873))
						mBase = m.M
						v222 = m.ExcPending
						if v222 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				switch v20 - int32(4) {
				case 0:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = int32(0)
						v31 = F_executeBoolItem(m, l0, v7+int32(-28), l2, v27)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 == int32(0) {
								v204 = v27
								m.G0 = v9 - int32(-64)
								return v204
							} else {
								F_jspGetRightArg(m, l1, v7+int32(-56))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									v42 = F_executeBoolItem(m, l0, v7+int32(-56), l2, int32(0))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										if v42 == int32(1) {
											v46 = v31
										} else {
											v46 = v42
										}
										v204 = v46
										m.G0 = v9 - int32(-64)
										return v204
									}
								}
							}
						}
					}
				case 1:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v55 = F_executeBoolItem(m, l0, v7+int32(-28), l2, int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 == int32(1) {
								v204 = int32(1)
								m.G0 = v9 - int32(-64)
								return v204
							} else {
								F_jspGetRightArg(m, l1, v7+int32(-56))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v66 = F_executeBoolItem(m, l0, v7+int32(-56), l2, int32(0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										if v66 != 0 {
											v68 = v66
										} else {
											v68 = v55
										}
										v204 = v68
										m.G0 = v9 - int32(-64)
										return v204
									}
								}
							}
						}
					}
				case 2:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v77 = F_executeBoolItem(m, l0, v7+int32(-28), l2, int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							if v77 == int32(2) {
								v204 = int32(2)
							} else {
								v204 = base.B2i32(v77 != int32(1))
							}
							m.G0 = v9 - int32(-64)
							return v204
						}
					}
				case 3:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						v198 = F_executeBoolItem(m, l0, v7+int32(-28), l2, int32(0))
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
							return int32(0)
						} else {
							v204 = base.B2i32(v198 == int32(2))
							m.G0 = v9 - int32(-64)
							return v204
						}
					}
				case 4, 5, 6, 7, 8, 9:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_jspGetRightArg(m, l1, v7+int32(-56))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v97 = F_executePredicate(m, l0, l1, v7+int32(-28), v7+int32(-56), l2, int32(1), int32(1439), l0)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								v204 = v97
								m.G0 = v9 - int32(-64)
								return v204
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return int32(0)
					} else {
						v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v181
						F_errmsg_internal(m, int32(503557), v9)
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(520430), int32(1902), int32(302873))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 26:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v137 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
							v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
							v143 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v143)
							v150 = F_executeItemOptUnwrapTarget(m, l0, v7+int32(-28), l2, v7+int32(-56), v143)
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v142)
								v153 = int32(2)
								if v150 == v153 {
									v204 = v153
								} else {
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									v204 = base.B2i32(v156|v157 != int32(0))
								}
								m.G0 = v9 - int32(-64)
								return v204
							}
						} else {
							v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
							v162 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v162)
							v168 = F_executeItemOptUnwrapTarget(m, l0, v7+int32(-28), l2, v162, int32(1))
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v161)
								v171 = int32(2)
								if v168 == v171 {
									v176 = v171
								} else {
									v176 = base.B2i32(v168 == int32(0))
								}
								v204 = v176
								m.G0 = v9 - int32(-64)
								return v204
							}
						}
					}
				case 37:
					F_jspGetArg(m, l1, v7+int32(-28))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						F_jspGetRightArg(m, l1, v7+int32(-56))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v111 = int32(0)
							v114 = F_executePredicate(m, l0, l1, v7+int32(-28), v7+int32(-56), l2, v111, int32(1440), v111)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								v204 = v114
								m.G0 = v9 - int32(-64)
								return v204
							}
						}
					}
				case 38:
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
					v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					F_jspInitByBuffer(m, v7+int32(-28), v120, v121)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						v126 = int32(0)
						v131 = F_executePredicate(m, l0, l1, v7+int32(-28), v126, l2, v126, int32(1441), v7+int32(-56))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							v204 = v131
							m.G0 = v9 - int32(-64)
							return v204
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v20 - int32(4) {
			case 0:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					v31 = F_executeBoolItem(m, l0, v7+int32(-28), l2, v27)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							v204 = v27
							m.G0 = v9 - int32(-64)
							return v204
						} else {
							F_jspGetRightArg(m, l1, v7+int32(-56))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v42 = F_executeBoolItem(m, l0, v7+int32(-56), l2, int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									if v42 == int32(1) {
										v46 = v31
									} else {
										v46 = v42
									}
									v204 = v46
									m.G0 = v9 - int32(-64)
									return v204
								}
							}
						}
					}
				}
			case 1:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v55 = F_executeBoolItem(m, l0, v7+int32(-28), l2, int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 == int32(1) {
							v204 = int32(1)
							m.G0 = v9 - int32(-64)
							return v204
						} else {
							F_jspGetRightArg(m, l1, v7+int32(-56))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v66 = F_executeBoolItem(m, l0, v7+int32(-56), l2, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 != 0 {
										v68 = v66
									} else {
										v68 = v55
									}
									v204 = v68
									m.G0 = v9 - int32(-64)
									return v204
								}
							}
						}
					}
				}
			case 2:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v77 = F_executeBoolItem(m, l0, v7+int32(-28), l2, int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						if v77 == int32(2) {
							v204 = int32(2)
						} else {
							v204 = base.B2i32(v77 != int32(1))
						}
						m.G0 = v9 - int32(-64)
						return v204
					}
				}
			case 3:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					v198 = F_executeBoolItem(m, l0, v7+int32(-28), l2, int32(0))
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return int32(0)
					} else {
						v204 = base.B2i32(v198 == int32(2))
						m.G0 = v9 - int32(-64)
						return v204
					}
				}
			case 4, 5, 6, 7, 8, 9:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_jspGetRightArg(m, l1, v7+int32(-56))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						v97 = F_executePredicate(m, l0, l1, v7+int32(-28), v7+int32(-56), l2, int32(1), int32(1439), l0)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							v204 = v97
							m.G0 = v9 - int32(-64)
							return v204
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int32(0)
				} else {
					v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v181
					F_errmsg_internal(m, int32(503557), v9)
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(520430), int32(1902), int32(302873))
						mBase = m.M
						v190 = m.ExcPending
						if v190 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 26:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v137 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						v143 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v143)
						v150 = F_executeItemOptUnwrapTarget(m, l0, v7+int32(-28), l2, v7+int32(-56), v143)
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v142)
							v153 = int32(2)
							if v150 == v153 {
								v204 = v153
							} else {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								v204 = base.B2i32(v156|v157 != int32(0))
							}
							m.G0 = v9 - int32(-64)
							return v204
						}
					} else {
						v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						v162 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v162)
						v168 = F_executeItemOptUnwrapTarget(m, l0, v7+int32(-28), l2, v162, int32(1))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v161)
							v171 = int32(2)
							if v168 == v171 {
								v176 = v171
							} else {
								v176 = base.B2i32(v168 == int32(0))
							}
							v204 = v176
							m.G0 = v9 - int32(-64)
							return v204
						}
					}
				}
			case 37:
				F_jspGetArg(m, l1, v7+int32(-28))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					F_jspGetRightArg(m, l1, v7+int32(-56))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v111 = int32(0)
						v114 = F_executePredicate(m, l0, l1, v7+int32(-28), v7+int32(-56), l2, v111, int32(1440), v111)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							v204 = v114
							m.G0 = v9 - int32(-64)
							return v204
						}
					}
				}
			case 38:
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
				v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F_jspInitByBuffer(m, v7+int32(-28), v120, v121)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					v126 = int32(0)
					v131 = F_executePredicate(m, l0, l1, v7+int32(-28), v126, l2, v126, int32(1441), v7+int32(-56))
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return int32(0)
					} else {
						v204 = v131
						m.G0 = v9 - int32(-64)
						return v204
					}
				}
			}
		}
	}
}

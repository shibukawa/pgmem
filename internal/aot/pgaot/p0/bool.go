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
					F_errmsg_internal(m, int32(_a_F_bool_accum_0), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bool_accum_1), int32(330), int32(_a_F_bool_accum_2))
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
					*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v53
					*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53
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
				F_errmsg_internal(m, int32(_a_F_bool_accum_0), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bool_accum_1), int32(330), int32(_a_F_bool_accum_2))
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
				*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v49))) = v53
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
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
				v189 = m.ExcPending
				if v189 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_executeBoolItem_0), int32(0))
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_executeBoolItem_1), int32(1788), int32(_a_F_executeBoolItem_2))
						mBase = m.M
						v198 = m.ExcPending
						if v198 != 0 {
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
					v24 = v7 + int32(-28)
					F_jspGetArg(m, l1, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = int32(0)
						v29 = F_executeBoolItem(m, l0, v24, l2, v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 == int32(0) {
								v180 = v27
								m.G0 = v9 - int32(-64)
								return v180
							} else {
								v34 = v7 + int32(-56)
								F_jspGetRightArg(m, l1, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v38 = F_executeBoolItem(m, l0, v34, l2, int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										if v38 == int32(1) {
											v42 = v29
										} else {
											v42 = v38
										}
										v180 = v42
										m.G0 = v9 - int32(-64)
										return v180
									}
								}
							}
						}
					}
				case 1:
					v44 = v7 + int32(-28)
					F_jspGetArg(m, l1, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v49 = F_executeBoolItem(m, l0, v44, l2, int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(1) {
								v180 = int32(1)
								m.G0 = v9 - int32(-64)
								return v180
							} else {
								v54 = v7 + int32(-56)
								F_jspGetRightArg(m, l1, v54)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v58 = F_executeBoolItem(m, l0, v54, l2, int32(0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										if v58 != 0 {
											v60 = v58
										} else {
											v60 = v49
										}
										v180 = v60
										m.G0 = v9 - int32(-64)
										return v180
									}
								}
							}
						}
					}
				case 2:
					v62 = v7 + int32(-28)
					F_jspGetArg(m, l1, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v67 = F_executeBoolItem(m, l0, v62, l2, int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							if v67 == int32(2) {
								v180 = int32(2)
							} else {
								v180 = base.B2i32(v67 != int32(1))
							}
							m.G0 = v9 - int32(-64)
							return v180
						}
					}
				case 3:
					v170 = v7 + int32(-28)
					F_jspGetArg(m, l1, v170)
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return int32(0)
					} else {
						v174 = F_executeBoolItem(m, l0, v170, l2, int32(0))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							v180 = base.B2i32(v174 == int32(2))
							m.G0 = v9 - int32(-64)
							return v180
						}
					}
				case 4, 5, 6, 7, 8, 9:
					v74 = v7 + int32(-28)
					F_jspGetArg(m, l1, v74)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v78 = v7 + int32(-56)
						F_jspGetRightArg(m, l1, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v83 = F_executePredicate(m, l0, l1, v74, v78, l2, int32(1), int32(1423), l0)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								v180 = v83
								m.G0 = v9 - int32(-64)
								return v180
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int32(0)
					} else {
						v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v159
						F_errmsg_internal(m, int32(_a_F_executeBoolItem_3), v9)
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_executeBoolItem_1), int32(1902), int32(_a_F_executeBoolItem_2))
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 26:
					v114 = v7 + int32(-28)
					F_jspGetArg(m, l1, v114)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v117 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
							v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
							v123 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v123)
							v128 = F_executeItemOptUnwrapTarget(m, l0, v114, l2, v7+int32(-56), v123)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v122)
								v131 = int32(2)
								if v128 == v131 {
									v180 = v131
								} else {
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									v180 = base.B2i32(v134|v135 != int32(0))
								}
								m.G0 = v9 - int32(-64)
								return v180
							}
						} else {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
							v140 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v140)
							v146 = F_executeItemOptUnwrapTarget(m, l0, v7+int32(-28), l2, v140, int32(1))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v139)
								v149 = int32(2)
								if v146 == v149 {
									v154 = v149
								} else {
									v154 = base.B2i32(v146 == int32(0))
								}
								v180 = v154
								m.G0 = v9 - int32(-64)
								return v180
							}
						}
					}
				case 37:
					v86 = v7 + int32(-28)
					F_jspGetArg(m, l1, v86)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v90 = v7 + int32(-56)
						F_jspGetRightArg(m, l1, v90)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v93 = int32(0)
							v96 = F_executePredicate(m, l0, l1, v86, v90, l2, v93, int32(1424), v93)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v180 = v96
								m.G0 = v9 - int32(-64)
								return v180
							}
						}
					}
				case 38:
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
					v101 = v7 + int32(-28)
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					F_jspInitByBuffer(m, v101, v102, v103)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v106 = int32(0)
						v111 = F_executePredicate(m, l0, l1, v101, v106, l2, v106, int32(1425), v7+int32(-56))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v180 = v111
							m.G0 = v9 - int32(-64)
							return v180
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v20 - int32(4) {
			case 0:
				v24 = v7 + int32(-28)
				F_jspGetArg(m, l1, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					v29 = F_executeBoolItem(m, l0, v24, l2, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 == int32(0) {
							v180 = v27
							m.G0 = v9 - int32(-64)
							return v180
						} else {
							v34 = v7 + int32(-56)
							F_jspGetRightArg(m, l1, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v38 = F_executeBoolItem(m, l0, v34, l2, int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									if v38 == int32(1) {
										v42 = v29
									} else {
										v42 = v38
									}
									v180 = v42
									m.G0 = v9 - int32(-64)
									return v180
								}
							}
						}
					}
				}
			case 1:
				v44 = v7 + int32(-28)
				F_jspGetArg(m, l1, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v49 = F_executeBoolItem(m, l0, v44, l2, int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 == int32(1) {
							v180 = int32(1)
							m.G0 = v9 - int32(-64)
							return v180
						} else {
							v54 = v7 + int32(-56)
							F_jspGetRightArg(m, l1, v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v58 = F_executeBoolItem(m, l0, v54, l2, int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									if v58 != 0 {
										v60 = v58
									} else {
										v60 = v49
									}
									v180 = v60
									m.G0 = v9 - int32(-64)
									return v180
								}
							}
						}
					}
				}
			case 2:
				v62 = v7 + int32(-28)
				F_jspGetArg(m, l1, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v67 = F_executeBoolItem(m, l0, v62, l2, int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						if v67 == int32(2) {
							v180 = int32(2)
						} else {
							v180 = base.B2i32(v67 != int32(1))
						}
						m.G0 = v9 - int32(-64)
						return v180
					}
				}
			case 3:
				v170 = v7 + int32(-28)
				F_jspGetArg(m, l1, v170)
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return int32(0)
				} else {
					v174 = F_executeBoolItem(m, l0, v170, l2, int32(0))
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int32(0)
					} else {
						v180 = base.B2i32(v174 == int32(2))
						m.G0 = v9 - int32(-64)
						return v180
					}
				}
			case 4, 5, 6, 7, 8, 9:
				v74 = v7 + int32(-28)
				F_jspGetArg(m, l1, v74)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v78 = v7 + int32(-56)
					F_jspGetRightArg(m, l1, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v83 = F_executePredicate(m, l0, l1, v74, v78, l2, int32(1), int32(1423), l0)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							v180 = v83
							m.G0 = v9 - int32(-64)
							return v180
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int32(0)
				} else {
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v159
					F_errmsg_internal(m, int32(_a_F_executeBoolItem_3), v9)
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_executeBoolItem_1), int32(1902), int32(_a_F_executeBoolItem_2))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 26:
				v114 = v7 + int32(-28)
				F_jspGetArg(m, l1, v114)
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v117 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						v123 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v123)
						v128 = F_executeItemOptUnwrapTarget(m, l0, v114, l2, v7+int32(-56), v123)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v122)
							v131 = int32(2)
							if v128 == v131 {
								v180 = v131
							} else {
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								v180 = base.B2i32(v134|v135 != int32(0))
							}
							m.G0 = v9 - int32(-64)
							return v180
						}
					} else {
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
						v140 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v140)
						v146 = F_executeItemOptUnwrapTarget(m, l0, v7+int32(-28), l2, v140, int32(1))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v139)
							v149 = int32(2)
							if v146 == v149 {
								v154 = v149
							} else {
								v154 = base.B2i32(v146 == int32(0))
							}
							v180 = v154
							m.G0 = v9 - int32(-64)
							return v180
						}
					}
				}
			case 37:
				v86 = v7 + int32(-28)
				F_jspGetArg(m, l1, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					v90 = v7 + int32(-56)
					F_jspGetRightArg(m, l1, v90)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = int32(0)
						v96 = F_executePredicate(m, l0, l1, v86, v90, l2, v93, int32(1424), v93)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v180 = v96
							m.G0 = v9 - int32(-64)
							return v180
						}
					}
				}
			case 38:
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
				v101 = v7 + int32(-28)
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F_jspInitByBuffer(m, v101, v102, v103)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					v106 = int32(0)
					v111 = F_executePredicate(m, l0, l1, v101, v106, l2, v106, int32(1425), v7+int32(-56))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						v180 = v111
						m.G0 = v9 - int32(-64)
						return v180
					}
				}
			}
		}
	}
}

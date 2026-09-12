package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetNewObjectId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v5 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+316))
		v13 = base.B2i32(v11 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v13)
		v15 = v13
	} else {
		v15 = int32(0)
	}
	if v15 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v23 = F_LWLockAcquire(m, v19+int32(256), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _consts[67]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			if base.Ui32(int32(16383)) < base.Ui32(v29) {
				v45 = v28
			} else {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[151])))
				if base.B2i32(v33 == int32(0))&base.B2i32(base.Ui32(int32(9999)) < base.Ui32(v29)) != 0 {
					v45 = v28
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(16384)
					v42 = *(*int32)(unsafe.Add(mBase, _consts[67]))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
					v45 = v42
				}
			}
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			if v46 == int32(0) {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
				v52 = m.G0
				v54 = v52 - int32(16)
				m.G0 = v54
				*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v49 - int32(-8192)
				F_XLogBeginInsert(m)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_XLogRegisterData(m, v54+int32(12), int32(4))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v66 = F_XLogInsert(m, int32(0), int32(48))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							m.G0 = v54 + int32(16)
							v72 = *(*int32)(unsafe.Add(mBase, _consts[67]))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = int32(8192)
							v75 = v72
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
							v78 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77 + v78
							v82 = *(*int32)(unsafe.Add(mBase, _consts[67]))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v83 - v78
							v88 = *(*int32)(unsafe.Add(mBase, _consts[29]))
							F_LWLockRelease(m, v88+int32(256))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								return v77
							}
						}
					}
				}
			} else {
				v75 = v45
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
				v78 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77 + v78
				v82 = *(*int32)(unsafe.Add(mBase, _consts[67]))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v83 - v78
				v88 = *(*int32)(unsafe.Add(mBase, _consts[29]))
				F_LWLockRelease(m, v88+int32(256))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					return v77
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(14546), int32(0))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(488068), int32(561), int32(458280))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
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
func F_check_object_ownership(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	switch l1 {
	case 0, 27, 47, 48:
		v207 = F_superuser_arg(m, l0)
		mBase = m.M
		v208 = m.ExcPending
		if v208 != 0 {
			return
		} else {
			if v207 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v212 = m.ExcPending
				if v212 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
						return
					} else {
						F_errmsg(m, int32(214738), int32(0))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return
						} else {
							F_errfinish(m, int32(486136), int32(2551), int32(234346))
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
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
		}
	case 1, 19, 25, 29, 34:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v48 = F_object_ownercheck(m, v46, v47, l0)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			if v48 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v52 = F_NameListToString(m, v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_aclcheck_error(m, int32(2), l1, v52)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						m.G0 = v9 + int32(112)
						return
					}
				}
			}
		}
	case 2, 3, 10, 11, 31, 32, 50:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v228 = m.ExcPending
		if v228 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = l1
			F_errmsg_internal(m, int32(477412), v9+int32(96))
			mBase = m.M
			v234 = m.ExcPending
			if v234 != 0 {
				return
			} else {
				F_errfinish(m, int32(486136), int32(2561), int32(234346))
				mBase = m.M
				v239 = m.ExcPending
				if v239 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4, 12, 49:
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v23 = F_object_ownercheck(m, v21, v22, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				F_aclcheck_error_type(m, int32(2), v22)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v9 + int32(112)
					return
				}
			}
		}
	case 5:
		v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
		v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
		v101 = F_typenameTypeId(m, int32(0), v100)
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return
		} else {
			v104 = F_typenameTypeId(m, int32(0), v98)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return
			} else {
				v107 = F_object_ownercheck(m, int32(1247), v101, l0)
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					if v107 != 0 {
						m.G0 = v9 + int32(112)
						return
					} else {
						v110 = F_object_ownercheck(m, int32(1247), v104, l0)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							if v110 != 0 {
								m.G0 = v9 + int32(112)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										v119 = F_format_type_be(m, v101)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											v121 = F_format_type_be(m, v104)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v121
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v119
												F_errmsg(m, int32(187652), v9+int32(32))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return
												} else {
													F_errfinish(m, int32(486136), int32(2496), int32(234346))
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
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
								}
							}
						}
					}
				}
			}
		}
	case 6, 18, 20, 23, 28, 35, 37, 40, 41, 44, 51:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
		v13 = F_object_ownercheck(m, int32(1259), v12, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			if v13 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
				F_aclcheck_error(m, int32(2), l1, v16+int32(4))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					m.G0 = v9 + int32(112)
					return
				}
			}
		}
	case 7, 8, 24, 26, 39, 45, 46:
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v66 = F_object_ownercheck(m, v64, v65, l0)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			if v66 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v69 = F_NameListToString(m, l3)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_aclcheck_error(m, int32(2), l1, v69)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						m.G0 = v9 + int32(112)
						return
					}
				}
			}
		}
	case 9, 14, 15, 16, 17, 21, 30, 36, 38, 42:
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v58 = F_object_ownercheck(m, v56, v57, l0)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			if v58 != 0 {
				m.G0 = v9 + int32(112)
				return
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				F_aclcheck_error(m, int32(2), l1, v61)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					m.G0 = v9 + int32(112)
					return
				}
			}
		}
	case 13:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v30 = F_SearchSysCache1(m, int32(19), v29)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			if v30 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
					F_errmsg_internal(m, int32(68118), v9)
					mBase = m.M
					v247 = m.ExcPending
					if v247 != 0 {
						return
					} else {
						F_errfinish(m, int32(486136), int32(2426), int32(234346))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35)+84))
				F_ReleaseCatCache(m, v30)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v41 = F_object_ownercheck(m, int32(1247), v37, l0)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 != 0 {
							m.G0 = v9 + int32(112)
							return
						} else {
							F_aclcheck_error_type(m, int32(2), v37)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								m.G0 = v9 + int32(112)
								return
							}
						}
					}
				}
			}
		}
	case 22:
		v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[223])))
		if v74 != 0 {
			m.G0 = v9 + int32(112)
			return
		} else {
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v77 = F_object_ownercheck(m, v75, v76, l0)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				if v77 != 0 {
					m.G0 = v9 + int32(112)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
							F_errmsg(m, int32(41655), v9+int32(16))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								F_errfinish(m, int32(486136), int32(2480), int32(234346))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
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
			}
		}
	case 33:
		v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v147 = F_superuser_arg(m, v146)
		mBase = m.M
		v148 = m.ExcPending
		if v148 != 0 {
			return
		} else {
			if v147 != 0 {
				v149 = F_superuser_arg(m, l0)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return
				} else {
					if v149 != 0 {
						m.G0 = v9 + int32(112)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								F_errmsg(m, int32(449687), int32(0))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(517425)
									F_errdetail(m, int32(606301), v9+int32(48))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return
									} else {
										F_errfinish(m, int32(486136), int32(2523), int32(234346))
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
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
					}
				}
			} else {
				v174 = F_has_createrole_privilege(m, l0)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return
				} else {
					if v174 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v256 = m.ExcPending
						if v256 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return
							} else {
								F_errmsg(m, int32(449687), int32(0))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(531957)
									F_errdetail(m, int32(606301), v9+int32(80))
									mBase = m.M
									v270 = m.ExcPending
									if v270 != 0 {
										return
									} else {
										F_errfinish(m, int32(486136), int32(2532), int32(234346))
										mBase = m.M
										v275 = m.ExcPending
										if v275 != 0 {
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
					} else {
						v178 = F_is_admin_of_role(m, l0, v146)
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							if v178 != 0 {
								m.G0 = v9 + int32(112)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v186 = m.ExcPending
									if v186 != 0 {
										return
									} else {
										F_errmsg(m, int32(449687), int32(0))
										mBase = m.M
										v190 = m.ExcPending
										if v190 != 0 {
											return
										} else {
											v192 = F_GetUserNameFromId(m, v146, int32(1))
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v192
												*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = int32(522250)
												F_errdetail(m, int32(643230), v9-int32(-64))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return
												} else {
													F_errfinish(m, int32(486136), int32(2540), int32(234346))
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
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
								}
							}
						}
					}
				}
			}
		}
	case 43:
		v137 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
		v139 = F_typenameTypeId(m, int32(0), v138)
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return
		} else {
			v141 = F_object_ownercheck(m, int32(1247), v139, l0)
			mBase = m.M
			v142 = m.ExcPending
			if v142 != 0 {
				return
			} else {
				if v141 != 0 {
					m.G0 = v9 + int32(112)
					return
				} else {
					F_aclcheck_error_type(m, int32(2), v139)
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return
					} else {
						m.G0 = v9 + int32(112)
						return
					}
				}
			}
		}
	default:
		m.G0 = v9 + int32(112)
		return
	}
}
func F_get_object_address_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(1247)
	v16 = F_LookupTypeName(m, v5, l2, l3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			if l3 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = F_TypeNameToString(m, l2)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
							F_errmsg(m, int32(71432), v9)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(486136), int32(1624), int32(360890))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
			}
		} else {
			v38 = F_typeTypeId(m, v16)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v38
				if l1 == int32(12) {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v44)+79)))
					if v46 != int32(100) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v63 = F_TypeNameToString(m, l2)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v63
									F_errmsg(m, int32(273876), v9+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										F_errfinish(m, int32(486136), int32(1635), int32(360890))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
					} else {
						F_ReleaseCatCache(m, v16)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				} else {
					F_ReleaseCatCache(m, v16)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_get_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = int32(1)
	v15 = v10 - v14
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
	if v17 != v14 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+v15<<(uint(int32(2))%32))))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 == int32(0) {
		v51 = v31
		v52 = v32
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v52-v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v31 != v32 {
		v51 = v31
		v52 = v32
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v36 = l1
	v37 = v26
	goto L10
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v40
		v52 = v41
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v51 = v40
	v52 = v41
	goto L7
L12:
	;
	v44 = int32(1)
	if v40 == v41 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v10 < v11 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10+v13))) = uint8(v56)
	return int32(0)
L16:
	;
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v62 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
	goto L1
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v65 != int32(1) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v68)
	return int32(0)
}

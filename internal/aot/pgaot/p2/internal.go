package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginInternalSubTransaction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(4474999)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[165])))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v13)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if base.Ui32(int32(19)) < base.Ui32(v17) {
		F_CommitTransactionCommand(m)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, _consts[61]))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
			if base.Ui32(int32(19)) < base.Ui32(v41) {
				*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
				v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
				m.G0 = v8 + int32(32)
				return
			} else {
				if v41 != 0 {
					if int32(1)<<(uint(v41)%32)&int32(1011558) == int32(0) {
						*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
						v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
						m.G0 = v8 + int32(32)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
							if base.Ui32(v54) <= base.Ui32(int32(19)) {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_consts[161])))
								v64 = v63
							} else {
								v64 = int32(538282)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v64
							F_errmsg_internal(m, int32(185892), v8)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_errfinish(m, int32(488504), int32(3114), int32(424856))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
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
					F_StartTransaction(m)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
						v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		}
	} else {
		if int32(1)<<(uint(v17)%32)&int32(5242) == int32(0) {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				if base.Ui32(v91) <= base.Ui32(int32(19)) {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v91<<(uint(int32(2))%32))+uint32(_consts[161])))
					v101 = v100
				} else {
					v101 = int32(538282)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v101
				F_errmsg_internal(m, int32(185712), v8+int32(16))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					F_errfinish(m, int32(488504), int32(4750), int32(255940))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			F_PushTransaction(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				if l0 == int32(0) {
					F_CommitTransactionCommand(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _consts[61]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
						if base.Ui32(int32(19)) < base.Ui32(v41) {
							*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
							v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
							m.G0 = v8 + int32(32)
							return
						} else {
							if v41 != 0 {
								if int32(1)<<(uint(v41)%32)&int32(1011558) == int32(0) {
									*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
									v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
									m.G0 = v8 + int32(32)
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
										if base.Ui32(v54) <= base.Ui32(int32(19)) {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_consts[161])))
											v64 = v63
										} else {
											v64 = int32(538282)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v64
										F_errmsg_internal(m, int32(185892), v8)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errfinish(m, int32(488504), int32(3114), int32(424856))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
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
								F_StartTransaction(m)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
									v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[61]))
					v33 = *(*int32)(unsafe.Add(mBase, _consts[68]))
					v34 = F_MemoryContextStrdup(m, v33, l0)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v34
						F_CommitTransactionCommand(m)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, _consts[61]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
							if base.Ui32(int32(19)) < base.Ui32(v41) {
								*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
								v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
								m.G0 = v8 + int32(32)
								return
							} else {
								if v41 != 0 {
									if int32(1)<<(uint(v41)%32)&int32(1011558) == int32(0) {
										*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
										v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
										m.G0 = v8 + int32(32)
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
											if base.Ui32(v54) <= base.Ui32(int32(19)) {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_consts[161])))
												v64 = v63
											} else {
												v64 = int32(538282)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v64
											F_errmsg_internal(m, int32(185892), v8)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(488504), int32(3114), int32(424856))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
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
									F_StartTransaction(m)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[165])) = uint8(v11)
										v82 = *(*int32)(unsafe.Add(mBase, _consts[166]))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
										m.G0 = v8 + int32(32)
										return
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
func F_internal_inetpl(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	v3 = int32(0)
	v18 = F_palloc0(m, int32(22))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(1)
		v23 = v18 + v22
		v25 = v18 + int32(4)
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v26&v22 != 0 {
			v29 = v23
		} else {
			v29 = v25
		}
		v31 = v29 + int32(2)
		v32 = int32(1)
		v33 = l0 + v32
		v35 = l0 + int32(4)
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v38 = v36 & v32
		if v38 != 0 {
			v39 = v33
		} else {
			v39 = v35
		}
		v41 = v39 + int32(2)
		if v38 != 0 {
			v46 = int32(1)
		} else {
			v46 = int32(4)
		}
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v46))))
		if v48 == int32(2) {
			v51 = int32(3)
		} else {
			v51 = int32(15)
		}
		v53 = l1
		v54 = v51
		v56 = v3
		v58 = v3
		for {
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v41))))
			v72 = int32(255)
			v75 = v70 + (base.I32_wrap_i64(v53)&v72 + v56)
			*(*uint8)(unsafe.Add(mBase, uint32(v54+v31))) = uint8(v75)
			v77 = int32(1)
			v78 = v54 - v77
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v41))))
			v83 = v53 >> (uint(int64(8)) % 64)
			v87 = int32(8)
			v90 = v81 + (base.I32_wrap_i64(v83)&v72 + int32(base.Ui32(v75)>>(uint(v87)%32)))
			*(*uint8)(unsafe.Add(mBase, uint32(v31+v78))) = uint8(v90)
			v93 = int32(base.Ui32(v90) >> (uint(v87) % 32))
			v94 = int32(2)
			v97 = v53 >> (uint(int64(16)) % 64)
			if v51 != v58|v77 {
				v53 = v97
				v54 = v54 - v94
				v56 = v93
				v58 = v58 + v94
				continue
			} else {
				break
			}
			break
		}
		if base.B2i32(base.Ui32(v90) < base.Ui32(int32(256)))&base.B2i32(base.Ui64(v83) < base.Ui64(int64(256))) != 0 {
			v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			if v129&int32(1) != 0 {
				v132 = v23
			} else {
				v132 = v25
			}
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v133&int32(1) != 0 {
				v136 = v33
			} else {
				v136 = v35
			}
			v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)) = uint8(v137)
			v139 = int32(1)
			v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			if v141&v139 != 0 {
				v144 = v139
			} else {
				v144 = int32(4)
			}
			v146 = int32(1)
			v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v148&v146 != 0 {
				v151 = v146
			} else {
				v151 = int32(4)
			}
			v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v151))))
			*(*uint8)(unsafe.Add(mBase, uint32(v18+v144))) = uint8(v153)
			if v153 == int32(2) {
				v159 = int32(40)
			} else {
				v159 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v159
			return v18
		} else {
			if base.B2i32(v93 == int32(1))&base.B2i32(v97 == int64(-1)) != 0 {
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v129&int32(1) != 0 {
					v132 = v23
				} else {
					v132 = v25
				}
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v133&int32(1) != 0 {
					v136 = v33
				} else {
					v136 = v35
				}
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
				*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)) = uint8(v137)
				v139 = int32(1)
				v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v141&v139 != 0 {
					v144 = v139
				} else {
					v144 = int32(4)
				}
				v146 = int32(1)
				v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v148&v146 != 0 {
					v151 = v146
				} else {
					v151 = int32(4)
				}
				v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v151))))
				*(*uint8)(unsafe.Add(mBase, uint32(v18+v144))) = uint8(v153)
				if v153 == int32(2) {
					v159 = int32(40)
				} else {
					v159 = int32(88)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v159
				return v18
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(398294), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492620), int32(1951), int32(298164))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
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
}
func F_internal_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(309540)
			F_errmsg(m, int32(190934), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(489163), int32(373), int32(66716))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_internal_putbytes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L35
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L31
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return v102
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	v16 = l0
	v17 = l1
	v18 = v15
	v20 = v13
	goto L7
L5:
	;
	goto L6
L6:
	;
	v102 = int32(0)
	goto L3
L7:
	;
	if base.Ui32(v20) <= base.Ui32(v18) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[472]))
	if v24 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	v42 = v18
	v43 = v20
	goto L11
L11:
	;
	if base.Ui32(v17) < base.Ui32(v43) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v27)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	v34 = F_internal_flush_buffer(m, v31, int32(4379224), int32(4379228))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v34 != 0 {
		v102 = int32(-1)
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	v42 = v41
	v43 = v39
	goto L11
L16:
	;
	if v85 != 0 {
		v16 = v84
		v17 = v85
		v18 = v86
		v20 = v87
		goto L7
	} else {
		goto L30
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	v71 = v43 - v42
	if base.Ui32(v71) < base.Ui32(v17) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	if v46 != v42 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, _consts[472]))
	if v51 == v48 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v54)
	v61 = F_internal_flush_buffer(m, v16, v9+int32(8), v9+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if v61 != 0 {
		v102 = int32(-1)
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	v66 = *(*int32)(unsafe.Add(mBase, _consts[474]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v84 = v16
	v85 = v67
	v86 = v64
	v87 = v66
	goto L16
L23:
	;
	v73 = v71
	goto L25
L24:
	;
	v73 = v17
	goto L25
L25:
	;
	if v73 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v76 = int32(4379228)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	v79 = v78 + v73
	*(*int32)(unsafe.Add(mBase, _consts[475])) = v79
	v81 = v17 - v73
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v81
	v84 = v16 + v73
	v85 = v81
	v86 = v79
	v87 = v43
	goto L16
L27:
	;
	v74 = F__emscripten_memcpy_bulkmem(m, v69+v42, v16, v73)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L8
L31:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(252794), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(492147), int32(886), int32(332311))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(252794), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(492147), int32(886), int32(332311))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_internal_text_pattern_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v38 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v10 = int32(4)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v12&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = int32(1)
	if v7&v25 != 0 {
		v37 = int32(base.Ui32(v7)>>(uint(v25)%32)) - v25
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v21 = v10
	goto L7
L6:
	;
	v21 = base.B2i32(v12 == int32(18)) << (uint(v10) % 32)
	goto L7
L7:
	;
	if v12 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = v10
	goto L10
L9:
	;
	v24 = v21
	goto L10
L10:
	;
	v37 = v24
	goto L1
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	v69 = int32(1)
	if v7&v69 != 0 {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	v41 = int32(4)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v43&int32(254) == int32(2) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v56 = int32(1)
	if v38&v56 != 0 {
		v68 = int32(base.Ui32(v38)>>(uint(v56)%32)) - v56
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v52 = v41
	goto L18
L17:
	;
	v52 = base.B2i32(v43 == int32(18)) << (uint(v41) % 32)
	goto L18
L18:
	;
	if v43 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v55 = v41
	goto L21
L20:
	;
	v55 = v52
	goto L21
L21:
	;
	v68 = v55
	goto L12
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v68 = int32(base.Ui32(v62)>>(uint(int32(2))%32)) - int32(4)
	goto L12
L23:
	;
	return v147
L24:
	;
	v73 = v69
	goto L26
L25:
	;
	v73 = int32(4)
	goto L26
L26:
	;
	v74 = l0 + v73
	v75 = int32(1)
	if v38&v75 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = v75
	goto L29
L28:
	;
	v79 = int32(4)
	goto L29
L29:
	;
	v80 = l1 + v79
	v81 = base.B2i32(v37 < v68)
	if v37 < v68 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v82 = v37
	goto L32
L31:
	;
	v82 = v68
	goto L32
L32:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v82) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if v144 != 0 {
		v147 = v144
		goto L23
	} else {
		goto L51
	}
L34:
	;
	v144 = int32(0)
	goto L33
L35:
	;
	v118 = v113
	v119 = v114
	v120 = v115
	goto L45
L36:
	;
	if (v74|v80)&int32(3) != 0 {
		v113 = v74
		v114 = v80
		v115 = v82
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v106 = v74
	v107 = v80
	v108 = v82
	goto L38
L38:
	;
	if v108 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v90 = v74
	v91 = v80
	v92 = v82
	goto L40
L40:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v95 != v96 {
		v113 = v90
		v114 = v91
		v115 = v92
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v106 = v101
	v107 = v99
	v108 = v103
	goto L38
L42:
	;
	v98 = int32(4)
	v99 = v91 + v98
	v101 = v90 + v98
	v103 = v92 - v98
	if base.Ui32(int32(3)) < base.Ui32(v103) {
		v90 = v101
		v91 = v99
		v92 = v103
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v113 = v106
	v114 = v107
	v115 = v108
	goto L35
L45:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v123 == v124 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v144 = v123 - v124
	goto L33
L47:
	;
	v126 = int32(1)
	v131 = v120 - v126
	if v131 != 0 {
		v118 = v118 + v126
		v119 = v119 + v126
		v120 = v131
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	goto L34
L51:
	;
	if v37 < v68 {
		v147 = int32(-1)
		goto L23
	} else {
		goto L52
	}
L52:
	;
	v147 = base.B2i32(v68 < v37)
	goto L23
}

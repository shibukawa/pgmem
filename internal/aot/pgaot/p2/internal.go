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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(_a_F_BeginInternalSubTransaction_0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v13)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[1]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if base.Ui32(int32(19)) < base.Ui32(v17) {
		F_CommitTransactionCommand(m)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[1]))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
			if base.Ui32(int32(19)) < base.Ui32(v41) {
				*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
				v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
				m.G0 = v8 + int32(32)
				return
			} else {
				if v41 != 0 {
					if int32(1)<<(uint(v41)%32)&int32(_a_F_BeginInternalSubTransaction_1) == int32(0) {
						*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
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
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_c_F_BeginInternalSubTransaction[4])))
								v61 = v59
							} else {
								v61 = int32(_a_F_BeginInternalSubTransaction_2)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
							F_errmsg_internal(m, int32(_a_F_BeginInternalSubTransaction_3), v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_BeginInternalSubTransaction_4), int32(3114), int32(_a_F_BeginInternalSubTransaction_5))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		}
	} else {
		if int32(1)<<(uint(v17)%32)&int32(_a_F_BeginInternalSubTransaction_6) == int32(0) {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				if base.Ui32(v88) <= base.Ui32(int32(19)) {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(int32(2))%32))+uint32(_c_F_BeginInternalSubTransaction[4])))
					v95 = v93
				} else {
					v95 = int32(_a_F_BeginInternalSubTransaction_2)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v95
				F_errmsg_internal(m, int32(_a_F_BeginInternalSubTransaction_7), v8+int32(16))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_BeginInternalSubTransaction_4), int32(_a_F_BeginInternalSubTransaction_8), int32(_a_F_BeginInternalSubTransaction_9))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
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
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[1]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
						if base.Ui32(int32(19)) < base.Ui32(v41) {
							*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
							v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
							m.G0 = v8 + int32(32)
							return
						} else {
							if v41 != 0 {
								if int32(1)<<(uint(v41)%32)&int32(_a_F_BeginInternalSubTransaction_1) == int32(0) {
									*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
									v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
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
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_c_F_BeginInternalSubTransaction[4])))
											v61 = v59
										} else {
											v61 = int32(_a_F_BeginInternalSubTransaction_2)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
										F_errmsg_internal(m, int32(_a_F_BeginInternalSubTransaction_3), v8)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_BeginInternalSubTransaction_4), int32(3114), int32(_a_F_BeginInternalSubTransaction_5))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
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
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
									v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
									*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[1]))
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[5]))
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
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[1]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
							if base.Ui32(int32(19)) < base.Ui32(v41) {
								*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
								v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
								m.G0 = v8 + int32(32)
								return
							} else {
								if v41 != 0 {
									if int32(1)<<(uint(v41)%32)&int32(_a_F_BeginInternalSubTransaction_1) == int32(0) {
										*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
										v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
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
												v59 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_c_F_BeginInternalSubTransaction[4])))
												v61 = v59
											} else {
												v61 = int32(_a_F_BeginInternalSubTransaction_2)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
											F_errmsg_internal(m, int32(_a_F_BeginInternalSubTransaction_3), v8)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_BeginInternalSubTransaction_4), int32(3114), int32(_a_F_BeginInternalSubTransaction_5))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
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
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[0])) = uint8(v11)
										v79 = *(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_BeginInternalSubTransaction[3])) = v79
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v100 int32
	_ = v100
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
	v17 = F_palloc0(m, int32(22))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(1)
		v22 = v17 + v21
		v24 = v17 + int32(4)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		if v25&v21 != 0 {
			v28 = v22
		} else {
			v28 = v24
		}
		v30 = v28 + int32(2)
		v31 = int32(1)
		v32 = l0 + v31
		v34 = l0 + int32(4)
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v37 = v35 & v31
		if v37 != 0 {
			v38 = v32
		} else {
			v38 = v34
		}
		v40 = v38 + int32(2)
		if v37 != 0 {
			v45 = int32(1)
		} else {
			v45 = int32(4)
		}
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v45))))
		if v47 == int32(2) {
			v50 = int32(3)
		} else {
			v50 = int32(15)
		}
		v52 = l1
		v54 = v50
		v55 = v3
		v59 = v3
		for {
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v40))))
			v69 = base.I32_wrap_i64(v52)
			v70 = int32(255)
			v73 = v68 + (v55 + v69&v70)
			*(*uint8)(unsafe.Add(mBase, uint32(v54+v30))) = uint8(v73)
			v75 = int32(1)
			v76 = v54 - v75
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v76))))
			v80 = int32(8)
			v87 = v79 + (int32(base.Ui32(v69)>>(uint(v80)%32))&v70 + int32(base.Ui32(v73)>>(uint(v80)%32)))
			*(*uint8)(unsafe.Add(mBase, uint32(v30+v76))) = uint8(v87)
			v90 = int32(base.Ui32(v87) >> (uint(v80) % 32))
			v91 = int32(2)
			v94 = v52 >> (uint(int64(16)) % 64)
			if v59|v75 != v50 {
				v52 = v94
				v54 = v54 - v91
				v55 = v90
				v59 = v59 + v91
				continue
			} else {
				break
			}
			break
		}
		v100 = int32(0)
		if base.B2i32(v90 == v100)&base.B2i32(v94 == int64(0))|base.B2i32(v90 == int32(1))&base.B2i32(v94 == int64(-1)) == v100 {
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
					F_errmsg(m, int32(_a_F_internal_inetpl_0), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_internal_inetpl_1), int32(1951), int32(_a_F_internal_inetpl_2))
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
		} else {
			v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v129&int32(1) != 0 {
				v132 = v22
			} else {
				v132 = v24
			}
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v133&int32(1) != 0 {
				v136 = v32
			} else {
				v136 = v34
			}
			v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)) = uint8(v137)
			v139 = int32(1)
			v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
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
			*(*uint8)(unsafe.Add(mBase, uint32(v17+v144))) = uint8(v153)
			if v153 == int32(2) {
				v159 = int32(40)
			} else {
				v159 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v159
			return v17
		}
	}
}
func F_internal_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_internal_out_0), int32(373), int32(_a_F_internal_out_1), int32(_a_F_internal_out_2), int32(_a_F_internal_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
	var v21 int32
	_ = v21
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
	var v68 int32
	_ = v68
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
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L29
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v101
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[1]))
	v16 = l0
	v17 = l1
	v18 = v15
	v21 = v13
	goto L6
L4:
	;
	goto L5
L5:
	;
	v101 = int32(0)
	goto L2
L6:
	;
	if base.Ui32(v21) <= base.Ui32(v18) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[2]))
	if v24 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v42 = v18
	v43 = v21
	goto L10
L10:
	;
	if base.Ui32(v17) < base.Ui32(v43) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v27)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[3]))
	v34 = F_internal_flush_buffer(m, v31, int32(_a_F_internal_putbytes_0), int32(_a_F_internal_putbytes_1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v34 != 0 {
		v101 = int32(-1)
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[0]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[1]))
	v42 = v41
	v43 = v39
	goto L10
L15:
	;
	if v84 != 0 {
		v16 = v83
		v17 = v84
		v18 = v85
		v21 = v87
		goto L6
	} else {
		goto L28
	}
L16:
	;
	v68 = v43 - v42
	if base.Ui32(v68) < base.Ui32(v17) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[4]))
	if v46 != v42 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[2]))
	if v51 == v48 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v54)
	v61 = F_internal_flush_buffer(m, v16, v9+int32(8), v9+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v61 != 0 {
		v101 = int32(-1)
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[1]))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v83 = v16
	v84 = v67
	v85 = v64
	v87 = v66
	goto L15
L22:
	;
	v70 = v68
	goto L24
L23:
	;
	v70 = v17
	goto L24
L24:
	;
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[3]))
	base.MemoryCopy(m, v72+v42, v16, v70)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v75 = int32(_a_F_internal_putbytes_1)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[1]))
	v78 = v77 + v70
	*(*int32)(unsafe.Add(mBase, _c_F_internal_putbytes[1])) = v78
	v80 = v17 - v70
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v80
	v83 = v16 + v70
	v84 = v80
	v85 = v78
	v87 = v43
	goto L15
L28:
	;
	goto L7
L29:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_internal_putbytes_2), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_internal_putbytes_3), int32(886), int32(_a_F_internal_putbytes_4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v37 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v13 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = int32(1)
	if v7&v24 != 0 {
		v36 = int32(base.Ui32(v7)>>(uint(v24)%32)) - v24
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v16 = int32(16)
	goto L7
L6:
	;
	v16 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v13-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v23 = int32(4)
	goto L10
L9:
	;
	v23 = v16
	goto L10
L10:
	;
	v36 = v23
	goto L1
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	v67 = int32(1)
	if v7&v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v43 == int32(18) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(1)
	if v37&v54 != 0 {
		v66 = int32(base.Ui32(v37)>>(uint(v54)%32)) - v54
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v46 = int32(16)
	goto L18
L17:
	;
	v46 = int32(0)
	goto L18
L18:
	;
	if base.Ui32((v43-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v53 = int32(4)
	goto L21
L20:
	;
	v53 = v46
	goto L21
L21:
	;
	v66 = v53
	goto L12
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
	goto L12
L23:
	;
	return v145
L24:
	;
	v71 = v67
	goto L26
L25:
	;
	v71 = int32(4)
	goto L26
L26:
	;
	v72 = l0 + v71
	v73 = int32(1)
	if v37&v73 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v77 = v73
	goto L29
L28:
	;
	v77 = int32(4)
	goto L29
L29:
	;
	v78 = l1 + v77
	v79 = base.B2i32(v36 < v66)
	if v36 < v66 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v80 = v36
	goto L32
L31:
	;
	v80 = v66
	goto L32
L32:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v80) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	if v142 != 0 {
		v145 = v142
		goto L23
	} else {
		goto L51
	}
L34:
	;
	v142 = int32(0)
	goto L33
L35:
	;
	v116 = v111
	v117 = v112
	v118 = v113
	goto L45
L36:
	;
	if (v72|v78)&int32(3) != 0 {
		v111 = v72
		v112 = v78
		v113 = v80
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v104 = v72
	v105 = v78
	v106 = v80
	goto L38
L38:
	;
	if v106 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v88 = v72
	v89 = v78
	v90 = v80
	goto L40
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v93 != v94 {
		v111 = v88
		v112 = v89
		v113 = v90
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v104 = v99
	v105 = v97
	v106 = v101
	goto L38
L42:
	;
	v96 = int32(4)
	v97 = v89 + v96
	v99 = v88 + v96
	v101 = v90 - v96
	if base.Ui32(int32(3)) < base.Ui32(v101) {
		v88 = v99
		v89 = v97
		v90 = v101
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v111 = v104
	v112 = v105
	v113 = v106
	goto L35
L45:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v121 == v122 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v142 = v121 - v122
	goto L33
L47:
	;
	v124 = int32(1)
	v129 = v118 - v124
	if v129 != 0 {
		v116 = v116 + v124
		v117 = v117 + v124
		v118 = v129
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
	if v36 < v66 {
		v145 = int32(-1)
		goto L23
	} else {
		goto L52
	}
L52:
	;
	v145 = base.B2i32(v66 < v36)
	goto L23
}

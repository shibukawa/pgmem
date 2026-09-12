package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaGetByte(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v18 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v48 = v35
		} else {
			v36 = int32(1)
			if v18&v36 != 0 {
				v48 = int32(base.Ui32(v18)>>(uint(v36)%32)) - v36
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.B2i32(int32(0) <= v15)&base.B2i32(v15 < v48) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v48 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
					F_errmsg(m, int32(487632), v8)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524313), int32(3325), int32(365356))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
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
			v72 = int32(1)
			if v18&v72 != 0 {
				v76 = v72
			} else {
				v76 = int32(4)
			}
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v76+v15))))
			m.G0 = v8 + int32(16)
			return v79
		}
	}
}
func F_byteaSetByte(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v17 = int32(base.Ui32(v15) >> (uint(int32(2)) % 32))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = int32(0)
		if base.B2i32(v19 <= v18)&base.B2i32(v18 < v17-int32(4)) == v19 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17 - int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
					F_errmsg(m, int32(487632), v8)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524313), int32(3392), int32(365343))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*uint8)(unsafe.Add(mBase, uint32(v18+v11)+4)) = uint8(v47)
			m.G0 = v8 + int32(16)
			return v11
		}
	}
}
func F_bytea_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if base.Ui32((v12-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(420348), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524313), int32(4149), int32(586334))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
				v35 = base.B2i32(v12 == int32(18)) << (uint(int32(4)) % 32)
				if base.Ui32(int32(2)) < base.Ui32(v35) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(420348), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524313), int32(4149), int32(586334))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
					if v35 == int32(0) {
						return int32(0)
					} else {
						v42 = int32(1)
						if v9&v42 != 0 {
							v46 = v42
						} else {
							v46 = int32(4)
						}
						v47 = v5 + v46
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
						if v35 == int32(1) {
							return base.I32_extend16_s(v48)
						} else {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
							return base.I32_extend16_s(v53 | v48<<(uint(int32(8))%32))
						}
					}
				}
			}
		} else {
			v23 = int32(1)
			if v9&v23 != 0 {
				v35 = int32(base.Ui32(v9)>>(uint(v23)%32)) - v23
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v35 = int32(base.Ui32(v29)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(2)) < base.Ui32(v35) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(420348), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524313), int32(4149), int32(586334))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
				if v35 == int32(0) {
					return int32(0)
				} else {
					v42 = int32(1)
					if v9&v42 != 0 {
						v46 = v42
					} else {
						v46 = int32(4)
					}
					v47 = v5 + v46
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
					if v35 == int32(1) {
						return base.I32_extend16_s(v48)
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
						return base.I32_extend16_s(v53 | v48<<(uint(int32(8))%32))
					}
				}
			}
		}
	}
}
func F_bytea_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v13 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v46 = int32(4)
				v49 = v46 & int32(3)
				v50 = int32(1)
				if v13&v50 != 0 {
					v54 = v50
				} else {
					v54 = int32(4)
				}
				v55 = v9 + v54
				v56 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v46) {
					v63 = v56
					v64 = v56
					for {
						v70 = int32(4)
						v71 = v63 + v70
						v73 = v64 + v70
						if v73 != v46&int32(4) {
							v63 = v71
							v64 = v73
							continue
						} else {
							break
						}
						break
					}
					v75 = v55 + v63
					v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
					v80 = int32(8)
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)))
					v89 = v71
					v90 = (v76<<(uint(int32(16))%32)|v79<<(uint(v80)%32)|v83)<<(uint(v80)%32) | v87
				} else {
					v89 = v56
					v90 = v56
				}
				if v49 != 0 {
					v96 = v89
					v97 = v90
					v99 = v56
					for {
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v55))))
						v107 = v104 | v97<<(uint(int32(8))%32)
						v108 = int32(1)
						v111 = v99 + v108
						if v111 != v49 {
							v96 = v96 + v108
							v97 = v107
							v99 = v111
							continue
						} else {
							break
						}
						break
					}
					v114 = v107
				} else {
					v114 = v90
				}
				return v114
			} else {
				v41 = base.B2i32(v17 == int32(18)) << (uint(int32(4)) % 32)
				if base.Ui32(int32(4)) < base.Ui32(v41) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(420694), int32(0))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524313), int32(4174), int32(584043))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
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
					if v41 != 0 {
						v46 = v41
						v49 = v46 & int32(3)
						v50 = int32(1)
						if v13&v50 != 0 {
							v54 = v50
						} else {
							v54 = int32(4)
						}
						v55 = v9 + v54
						v56 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v46) {
							v63 = v56
							v64 = v56
							for {
								v70 = int32(4)
								v71 = v63 + v70
								v73 = v64 + v70
								if v73 != v46&int32(4) {
									v63 = v71
									v64 = v73
									continue
								} else {
									break
								}
								break
							}
							v75 = v55 + v63
							v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
							v80 = int32(8)
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)))
							v89 = v71
							v90 = (v76<<(uint(int32(16))%32)|v79<<(uint(v80)%32)|v83)<<(uint(v80)%32) | v87
						} else {
							v89 = v56
							v90 = v56
						}
						if v49 != 0 {
							v96 = v89
							v97 = v90
							v99 = v56
							for {
								v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v55))))
								v107 = v104 | v97<<(uint(int32(8))%32)
								v108 = int32(1)
								v111 = v99 + v108
								if v111 != v49 {
									v96 = v96 + v108
									v97 = v107
									v99 = v111
									continue
								} else {
									break
								}
								break
							}
							v114 = v107
						} else {
							v114 = v90
						}
						return v114
					} else {
						return int32(0)
					}
				}
			}
		} else {
			v28 = int32(1)
			if v13&v28 != 0 {
				v41 = int32(base.Ui32(v13)>>(uint(v28)%32)) - v28
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v41 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(4)) < base.Ui32(v41) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(420694), int32(0))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524313), int32(4174), int32(584043))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
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
				if v41 != 0 {
					v46 = v41
					v49 = v46 & int32(3)
					v50 = int32(1)
					if v13&v50 != 0 {
						v54 = v50
					} else {
						v54 = int32(4)
					}
					v55 = v9 + v54
					v56 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v46) {
						v63 = v56
						v64 = v56
						for {
							v70 = int32(4)
							v71 = v63 + v70
							v73 = v64 + v70
							if v73 != v46&int32(4) {
								v63 = v71
								v64 = v73
								continue
							} else {
								break
							}
							break
						}
						v75 = v55 + v63
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
						v80 = int32(8)
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)))
						v89 = v71
						v90 = (v76<<(uint(int32(16))%32)|v79<<(uint(v80)%32)|v83)<<(uint(v80)%32) | v87
					} else {
						v89 = v56
						v90 = v56
					}
					if v49 != 0 {
						v96 = v89
						v97 = v90
						v99 = v56
						for {
							v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v55))))
							v107 = v104 | v97<<(uint(int32(8))%32)
							v108 = int32(1)
							v111 = v99 + v108
							if v111 != v49 {
								v96 = v96 + v108
								v97 = v107
								v99 = v111
								continue
							} else {
								break
							}
							break
						}
						v114 = v107
					} else {
						v114 = v90
					}
					return v114
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_bytea_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v14 == int32(1) {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v49 = int32(4)
				v52 = v49 & int32(3)
				v53 = int32(1)
				if v14&v53 != 0 {
					v57 = v53
				} else {
					v57 = int32(4)
				}
				v58 = v10 + v57
				if base.Ui32(v49) < base.Ui32(int32(4)) {
					v112 = int32(0)
					v119 = v8
				} else {
					v66 = int32(0)
					v71 = int32(0)
					v73 = v8
					for {
						v74 = int64(16)
						v76 = v66 + v58
						v77 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
						v78 = int64(8)
						v81 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
						v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
						v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+3)))
						v90 = (v73<<(uint(v74)%64)|v77<<(uint(v78)%64)|v81)<<(uint(v74)%64) | v85<<(uint(v78)%64) | v89
						v91 = int32(4)
						v92 = v66 + v91
						v94 = v71 + v91
						if v94 != v49&int32(12) {
							v66 = v92
							v71 = v94
							v73 = v90
							continue
						} else {
							break
						}
						break
					}
					v112 = v92
					v119 = v90
				}
				if v52 != 0 {
					v120 = v112
					v122 = int32(0)
					v127 = v119
					for {
						v129 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120+v58))))
						v132 = v129 | v127<<(uint(int64(8))%64)
						v133 = int32(1)
						v136 = v122 + v133
						if v136 != v52 {
							v120 = v120 + v133
							v122 = v136
							v127 = v132
							continue
						} else {
							break
						}
						break
					}
					v145 = v132
				} else {
					v145 = v119
				}
				v146 = F_Int64GetDatum(m, v145)
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return int32(0)
				} else {
					return v146
				}
			} else {
				v42 = base.B2i32(v18 == int32(18)) << (uint(int32(4)) % 32)
				if base.Ui32(int32(8)) < base.Ui32(v42) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(420370), int32(0))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524313), int32(4199), int32(580123))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
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
					if v42 != 0 {
						v49 = v42
						v52 = v49 & int32(3)
						v53 = int32(1)
						if v14&v53 != 0 {
							v57 = v53
						} else {
							v57 = int32(4)
						}
						v58 = v10 + v57
						if base.Ui32(v49) < base.Ui32(int32(4)) {
							v112 = int32(0)
							v119 = v8
						} else {
							v66 = int32(0)
							v71 = int32(0)
							v73 = v8
							for {
								v74 = int64(16)
								v76 = v66 + v58
								v77 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
								v78 = int64(8)
								v81 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
								v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
								v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+3)))
								v90 = (v73<<(uint(v74)%64)|v77<<(uint(v78)%64)|v81)<<(uint(v74)%64) | v85<<(uint(v78)%64) | v89
								v91 = int32(4)
								v92 = v66 + v91
								v94 = v71 + v91
								if v94 != v49&int32(12) {
									v66 = v92
									v71 = v94
									v73 = v90
									continue
								} else {
									break
								}
								break
							}
							v112 = v92
							v119 = v90
						}
						if v52 != 0 {
							v120 = v112
							v122 = int32(0)
							v127 = v119
							for {
								v129 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120+v58))))
								v132 = v129 | v127<<(uint(int64(8))%64)
								v133 = int32(1)
								v136 = v122 + v133
								if v136 != v52 {
									v120 = v120 + v133
									v122 = v136
									v127 = v132
									continue
								} else {
									break
								}
								break
							}
							v145 = v132
						} else {
							v145 = v119
						}
						v146 = F_Int64GetDatum(m, v145)
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							return v146
						}
					} else {
						v46 = F_Int64GetDatum(m, int64(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							return v46
						}
					}
				}
			}
		} else {
			v29 = int32(1)
			if v14&v29 != 0 {
				v42 = int32(base.Ui32(v14)>>(uint(v29)%32)) - v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v42 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(8)) < base.Ui32(v42) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(420370), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524313), int32(4199), int32(580123))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
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
				if v42 != 0 {
					v49 = v42
					v52 = v49 & int32(3)
					v53 = int32(1)
					if v14&v53 != 0 {
						v57 = v53
					} else {
						v57 = int32(4)
					}
					v58 = v10 + v57
					if base.Ui32(v49) < base.Ui32(int32(4)) {
						v112 = int32(0)
						v119 = v8
					} else {
						v66 = int32(0)
						v71 = int32(0)
						v73 = v8
						for {
							v74 = int64(16)
							v76 = v66 + v58
							v77 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
							v78 = int64(8)
							v81 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
							v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
							v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+3)))
							v90 = (v73<<(uint(v74)%64)|v77<<(uint(v78)%64)|v81)<<(uint(v74)%64) | v85<<(uint(v78)%64) | v89
							v91 = int32(4)
							v92 = v66 + v91
							v94 = v71 + v91
							if v94 != v49&int32(12) {
								v66 = v92
								v71 = v94
								v73 = v90
								continue
							} else {
								break
							}
							break
						}
						v112 = v92
						v119 = v90
					}
					if v52 != 0 {
						v120 = v112
						v122 = int32(0)
						v127 = v119
						for {
							v129 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120+v58))))
							v132 = v129 | v127<<(uint(int64(8))%64)
							v133 = int32(1)
							v136 = v122 + v133
							if v136 != v52 {
								v120 = v120 + v133
								v122 = v136
								v127 = v132
								continue
							} else {
								break
							}
							break
						}
						v145 = v132
					} else {
						v145 = v119
					}
					v146 = F_Int64GetDatum(m, v145)
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						return v146
					}
				} else {
					v46 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						return v46
					}
				}
			}
		}
	}
}
func F_bytea_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v15 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v18 = int32(4)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v20&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v15&v33 != 0 {
		v45 = int32(base.Ui32(v15)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v29 = v18
	goto L10
L9:
	;
	v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
	goto L10
L10:
	;
	if v20 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = v18
	goto L13
L12:
	;
	v32 = v29
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v76 < v45 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v49 = int32(4)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v51&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v64 = int32(1)
	if v46&v64 != 0 {
		v76 = int32(base.Ui32(v46)>>(uint(v64)%32)) - v64
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v60 = v49
	goto L21
L20:
	;
	v60 = base.B2i32(v51 == int32(18)) << (uint(v49) % 32)
	goto L21
L21:
	;
	if v51 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v63 = v49
	goto L24
L23:
	;
	v63 = v60
	goto L24
L24:
	;
	v76 = v63
	goto L15
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v78 = v8
	goto L28
L27:
	;
	v78 = v13
	goto L28
L28:
	;
	v79 = int32(1)
	if v15&v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v83 = v79
	goto L31
L30:
	;
	v83 = int32(4)
	goto L31
L31:
	;
	v84 = v8 + v83
	v85 = int32(1)
	if v46&v85 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v85
	goto L34
L33:
	;
	v89 = int32(4)
	goto L34
L34:
	;
	v90 = v13 + v89
	if v45 < v76 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v92 = v45
	goto L37
L36:
	;
	v92 = v76
	goto L37
L37:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v92) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v154 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v154 = int32(0)
	goto L38
L40:
	;
	v128 = v123
	v129 = v124
	v130 = v125
	goto L50
L41:
	;
	if (v84|v90)&int32(3) != 0 {
		v123 = v84
		v124 = v90
		v125 = v92
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v116 = v84
	v117 = v90
	v118 = v92
	goto L43
L43:
	;
	if v118 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v100 = v84
	v101 = v90
	v102 = v92
	goto L45
L45:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v105 != v106 {
		v123 = v100
		v124 = v101
		v125 = v102
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v116 = v111
	v117 = v109
	v118 = v113
	goto L43
L47:
	;
	v108 = int32(4)
	v109 = v101 + v108
	v111 = v100 + v108
	v113 = v102 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v100 = v111
		v101 = v109
		v102 = v113
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L40
L50:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v133 == v134 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v154 = v133 - v134
	goto L38
L52:
	;
	v136 = int32(1)
	v141 = v130 - v136
	if v141 != 0 {
		v128 = v128 + v136
		v129 = v129 + v136
		v130 = v141
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L39
L56:
	;
	v155 = v13
	goto L58
L57:
	;
	v155 = v78
	goto L58
L58:
	;
	if int32(0) < v154 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v158 = v8
	goto L61
L60:
	;
	v158 = v155
	goto L61
L61:
	;
	return v158
}
func F_bytea_substr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = F_bytea_substring(m, v2, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_bytea_substring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = int32(1)
	if l1 <= v6 {
		v9 = v6
	} else {
		v9 = l1
	}
	if int32(0) <= l2 {
		v13 = int32(0)
		v15 = l1 + l2
		if base.B2i32(l2 < v13)^base.B2i32(v15 < l1) == v13 {
			if v15 <= int32(0) {
				v25 = F_DirectFunctionCall1Coll(m, int32(579), int32(0), int32(790160))
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_pg_detoast_datum_packed(m, v25)
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						return v29
					}
				}
			} else {
				v33 = v15 - v9
				v36 = F_pg_detoast_datum_slice(m, l0, v9-int32(1), v33)
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					return v36
				}
			}
		} else {
			v33 = int32(-1)
			v36 = F_pg_detoast_datum_slice(m, l0, v9-int32(1), v33)
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				return v36
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(459814), int32(0))
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524313), int32(3158), int32(345808))
					v54 = m.ExcPending
					if v54 != 0 {
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

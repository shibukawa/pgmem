package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckRelationTableSpaceMove(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if l1 == v12 {
		v34 = v3
		m.G0 = v9 + int32(16)
		return v34
	} else {
		if v12 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[109]))
			if l1 == v17 {
				v34 = v3
				m.G0 = v9 + int32(16)
				return v34
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
				switch v19 - int32(83) {
				case 0, 22, 26, 31, 33:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
					if v22 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48 + int32(4)
								F_errmsg(m, int32(681384), v9)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486034), int32(3714), int32(336783))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
						if l1 == int32(1664) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(412120), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486034), int32(3720), int32(336783))
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
							v27 = int32(1)
							v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
							if v28 != int32(116) {
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							} else {
								v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v31 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(141674), int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(486034), int32(3729), int32(336783))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
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
									v34 = v27
									m.G0 = v9 + int32(16)
									return v34
								}
							}
						}
					}
				default:
					if l1 == int32(1664) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(412120), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486034), int32(3720), int32(336783))
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
						v27 = int32(1)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
						if v28 != int32(116) {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(141674), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(486034), int32(3729), int32(336783))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
			switch v19 - int32(83) {
			case 0, 22, 26, 31, 33:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48 + int32(4)
							F_errmsg(m, int32(681384), v9)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486034), int32(3714), int32(336783))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
					if l1 == int32(1664) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(412120), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486034), int32(3720), int32(336783))
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
						v27 = int32(1)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
						if v28 != int32(116) {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(141674), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(486034), int32(3729), int32(336783))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			default:
				if l1 == int32(1664) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(412120), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486034), int32(3720), int32(336783))
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
					v27 = int32(1)
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
					if v28 != int32(116) {
						v34 = v27
						m.G0 = v9 + int32(16)
						return v34
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(141674), int32(0))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486034), int32(3729), int32(336783))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	switch l2 - int32(1663) {
	case 0:
		if l4 == int32(-1) {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = l1
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_consts[1327])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v40
				v45 = F_pg_sprintf(m, l0, int32(173076), v10+int32(176))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = l1
				v52 = F_pg_sprintf(m, l0, int32(38224), v10+int32(160))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		} else {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+216)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+212)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = l1
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_consts[1327])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+220)) = v61
				v66 = F_pg_sprintf(m, l0, int32(173040), v10+int32(208))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = l1
				v74 = F_pg_sprintf(m, l0, int32(37939), v10+int32(192))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		}
	case 1:
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = l3
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_consts[1327])))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v19
			v24 = F_pg_sprintf(m, l0, int32(173090), v10+int32(144))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v10 + int32(224)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = l3
			v30 = F_pg_sprintf(m, l0, int32(38540), v10+int32(128))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v10 + int32(224)
				return
			}
		}
	default:
		if l4 == int32(-1) {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l3
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_consts[1327])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(544687)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(480448)
				v94 = F_pg_sprintf(m, l0, int32(173058), v10+int32(32))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(544687)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(480448)
				v104 = F_pg_sprintf(m, l0, int32(38209), v10)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		} else {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l4
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_consts[1327])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v112
				*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(544687)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = int32(480448)
				v123 = F_pg_sprintf(m, l0, int32(173018), v10+int32(96))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(544687)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(480448)
				v136 = F_pg_sprintf(m, l0, int32(37920), v10-int32(-64))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		}
	}
}
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v51 int32
	_ = v51
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v79 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v79 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v79 = int32(0)
				} else {
					v79 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v23) {
				if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
					v79 = v11
				} else {
					if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						v79 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0) {
					if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
						v79 = v11
					} else {
						if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				} else {
					v79 = v11
				}
			}
		}
	} else {
		if l0 <= int32(5999) {
			v35 = l0 - int32(4177)
			if base.Ui32(int32(9)) < base.Ui32(v35) {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v79 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v79 = int32(0)
					} else {
						v79 = v11
					}
				}
			} else {
				if int32(1)<<(uint(v35)%32)&int32(963) == int32(0) {
					if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
							v79 = int32(0)
						} else {
							v79 = v11
						}
					}
				} else {
					v79 = v11
				}
			}
		} else {
			switch l0 - int32(6243) {
			case 0, 1, 2, 3, 4, 59, 60:
				v79 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v79 = int32(0)
			default:
				if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
					v79 = v11
				} else {
					v51 = l0 - int32(6100)
					if base.Ui32(int32(15)) < base.Ui32(v51) {
						v79 = int32(0)
					} else {
						if int32(1)<<(uint(v51)%32)&int32(49153) != 0 {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v85 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	if v79 != 0 {
		v86 = int32(0)
	} else {
		v86 = v85
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v86
	v90 = int32(0)
	v95 = F_LockAcquireExtended(m, v7+int32(16), l1, v90, v90, v7+int32(12), v90)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		return
	} else {
		if v95 != int32(3) {
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v102 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v101)+53)) = uint8(v102)
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_RelationBuildDesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
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
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v785 int32
	_ = v785
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v903 int32
	_ = v903
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v944 int32
	_ = v944
	var v950 int64
	_ = v950
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1113 int32
	_ = v1113
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	v20 = m.G0
	v22 = v20 - int32(320)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	v27 = *(*int32)(unsafe.Add(mBase, _consts[1123]))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1124]))
	if v29 <= v27 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = F_repalloc(m, v25, v29<<(uint(int32(4))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v45 = v25
	v46 = v27
	goto L3
L3:
	;
	v48 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1123])) = v46 + v48
	v52 = v46 << (uint(int32(3)) % 32)
	v53 = v45 + v52
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = l0
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)) = uint8(v55)
	v59 = F_ScanPgRelation(m, l0, v48, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1124])) = v29 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _consts[1122])) = v33
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1123]))
	v45 = v33
	v46 = v44
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L4
	} else {
		goto L269
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L4
	} else {
		goto L266
	}
L8:
	;
	m.G0 = v22 + int32(320)
	return v1113
L9:
	;
	v1055 = int32(4458560)
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[1123]))
	*(*int32)(unsafe.Add(mBase, _consts[1123])) = v1057 - int32(1)
	if l1 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L10:
	;
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v78 = v59
	goto L14
L12:
	;
	goto L13
L13:
	;
	v1049 = int32(4458560)
	v1051 = *(*int32)(unsafe.Add(mBase, _consts[1123]))
	*(*int32)(unsafe.Add(mBase, _consts[1123])) = v1051 - int32(1)
	v1113 = int32(0)
	goto L8
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v84 = v82 + v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(4470560)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v90 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v90
	v93 = F_palloc0(m, int32(276))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = int32(0)
	v98 = F_palloc(m, int32(144))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L19
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+48)) = v101
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+120)))
	v105 = F_CreateTemplateTupleDesc(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v101 = F__emscripten_memcpy_bulkmem(m, v98, v84, int32(144))
	mBase = m.M
	goto L21
L21:
	;
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+52)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v87
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+32)) = v112
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+25)) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v93)+56)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v93)+40)) = v112
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+118)))
	switch v122 - int32(112) {
	case 0, 5:
		goto L24
	default:
		goto L25
	case 4:
		goto L26
	}
L23:
	;
	v177 = v93 + int32(56)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+72))
	if v180 != 0 {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	v170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v170)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(-1)
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L41
	}
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+68))
	v129 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	if v129 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v137 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	v130 = int32(1)
	if v125 == v129 {
		v137 = v130
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v137 = int32(0)
	goto L28
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	if v133 == v125 {
		v137 = v130
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v141 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v142)
	if v141 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+68))
	v150 = F_GetTempNamespaceProcNumber(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L40
	}
L37:
	;
	v146 = v139
	goto L39
L38:
	;
	v146 = v141
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v146
	goto L23
L40:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v152)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v150
	goto L23
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v160 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v160
	F_errmsg_internal(m, int32(493881), v22)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(490658), int32(1197), int32(480388))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v182 = v180
	goto L46
L45:
	;
	v182 = int32(2249)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = int32(-1)
	v189 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v191 = F_MemoryContextAllocZero(m, v189, int32(20))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	F_ScanKeyInit(m, v22+int32(160), int32(1), int32(3), int32(184), v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v201 = int32(5)
	F_ScanKeyInit(m, v22+int32(208), v201, v201, int32(146), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v209 = F_table_open(m, int32(1249), int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1125])))
	v218 = F_systable_beginscan(m, v209, int32(2659), v213, int32(0), int32(2), v22+int32(160))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220)+120)))
	v230 = int32(0)
	v231 = v221
	v236 = int32(0)
	goto L53
L52:
	;
	F_systable_endscan(m, v218)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L93
	}
L53:
	;
	v242 = F_systable_getnext(m, v218)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	v379 = v372
	v380 = int32(0)
	v383 = v371
	goto L52
L55:
	;
	if v242 == int32(0) {
		v379 = v230
		v380 = v231
		v383 = v236
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+22)))
	v248 = v246 + v247
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+74)))
	if v249 <= int32(0) {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+120)))
	if v253 < v249 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v263 = (v249 - int32(1)) & int32(65535)
	v264 = int32(100)
	goto L60
L59:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	F_populate_compact_attribute(m, v272, v263)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L63
	}
L60:
	;
	v270 = F__emscripten_memcpy_bulkmem(m, v255+v256<<(uint(int32(4))%32)+v263*v264+int32(20), v248, v264)
	mBase = m.M
	goto L62
L62:
	;
	goto L59
L63:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+86)))
	if v275 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v278 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+16)) = uint8(v278)
	goto L66
L65:
	;
	goto L66
L66:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+90)))
	if v280 == int32(115) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v283 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+17)) = uint8(v283)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+90)))
	v286 = v285
	goto L69
L68:
	;
	v286 = v280
	goto L69
L69:
	;
	if v286&int32(255) == int32(118) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v291 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+18)) = uint8(v291)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+87)))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+88)))
	if v294 != int32(1) {
		v371 = v236
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v372 = v230 + v293
	v374 = v231 - int32(1)
	if v374 != 0 {
		v230 = v372
		v231 = v374
		v236 = v371
		goto L53
	} else {
		goto L92
	}
L74:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v209)+52))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298)+18)))
	if base.Ui32(v299&int32(2047)) <= base.Ui32(int32(24)) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+271)))
	if v315 != 0 {
		v371 = v236
		goto L73
	} else {
		goto L81
	}
L76:
	;
	v307 = F_getmissingattr(m, v297, int32(25), v22+int32(271))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v312 = F_fastgetattr_3(m, v242, int32(25), v297, v22+int32(271))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L80
	}
L79:
	;
	v314 = v307
	goto L75
L80:
	;
	v314 = v312
	goto L75
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(1)
	if v236 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v323 = int32(*(*int16)(unsafe.Add(mBase, uint32(v322)+120)))
	v326 = F_MemoryContextAllocZero(m, v321, v323<<(uint(int32(3))%32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	v328 = v236
	goto L84
L84:
	;
	v333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+72)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	v335 = int32(*(*int8)(unsafe.Add(mBase, uint32(v248)+83)))
	v338 = F_array_get_element(m, v314, int32(1), v22+int32(272), int32(-1), v333, v334, v335, v22+int32(159))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L86
	}
L85:
	;
	v328 = v326
	goto L84
L86:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	if v340 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v367 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v328+v263<<(uint(int32(3))%32)))) = uint8(v367)
	v371 = v328
	goto L73
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328+v263<<(uint(int32(3))%32))+4)) = v338
	goto L87
L89:
	;
	goto L90
L90:
	;
	v347 = int32(4470560)
	v348 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v351 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v351
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	v357 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+72)))
	v358 = F_datumCopy(m, v338, v356, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328+v263<<(uint(int32(3))%32))+4)) = v358
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v348
	goto L87
L92:
	;
	goto L54
L93:
	;
	F_sequence_close(m, v209, int32(1))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	if v380 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v390 = int32(*(*int16)(unsafe.Add(mBase, uint32(v389)+120)))
	if int32(0) < v390 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v393)+20)) = int32(0)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+16)))
	if v396 != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v944 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+128)) = v944
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+88)) = uint8(v944)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+84)) = v944
	v950 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+92)) = v950
	*(*int64)(unsafe.Add(mBase, uint32(v93)+100)) = v950
	*(*int64)(unsafe.Add(mBase, uint32(v93)+108)) = v950
	*(*int64)(unsafe.Add(mBase, uint32(v93)+116)) = v950
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+124)) = uint8(v944)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960)+119)))
	switch v961 - int32(73) {
	case 0, 32:
		goto L222
	default:
		goto L220
	case 10, 36, 41, 43:
		goto L221
	}
L100:
	;
	F_pfree(m, v191)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L4
	} else {
		goto L219
	}
L101:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	v407 = base.B2i32(base.Ui32(v405) < base.Ui32(int32(12000)))
	goto L108
L102:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+17)))
	if v397 != 0 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+18)))
	if v398 != 0 {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	if int32(0) < v379 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	if v383 != 0 {
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v401)+122)))
	if v402 <= int32(0) {
		goto L100
	} else {
		goto L107
	}
L107:
	;
	goto L101
L108:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v408)+16)) = v191
	if int32(0) < v379 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v383
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v623 = int32(*(*int16)(unsafe.Add(mBase, uint32(v622)+122)))
	if v623 <= int32(0) {
		goto L158
	} else {
		goto L159
	}
L110:
	;
	v412 = int32(0)
	v414 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v417 = F_MemoryContextAllocZero(m, v414, v379<<(uint(int32(3))%32))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v600 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+12)) = uint16(v600)
	goto L109
L113:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	F_ScanKeyInit(m, v22+int32(272), int32(2), int32(3), int32(184), v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v429 = F_table_open(m, int32(2604), int32(1))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	F_systable_endscan(m, v437)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L143
	}
L116:
	;
	v432 = int32(1)
	v437 = F_systable_beginscan(m, v429, int32(2656), v432, int32(0), v432, v22+int32(272))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	v439 = F_systable_getnext(m, v437)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	if v439 == int32(0) {
		v545 = v412
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v445 = v439
	v448 = v412
	goto L120
L120:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v445)+16))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+22)))
	v464 = v462 + v463
	if v379 <= v448 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v545 = v536
	goto L115
L122:
	;
	v468 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v429)+52))
	v492 = F_fastgetattr_3(m, v445, int32(4), v489, v22+int32(271))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L129
	}
L125:
	;
	if v468 == int32(0) {
		v545 = v448
		goto L115
	} else {
		goto L126
	}
L126:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v473 = int32(*(*int16)(unsafe.Add(mBase, uint32(v464)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v472 + int32(4)
	F_errmsg_internal(m, int32(682419), v22+int32(112))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(490658), int32(4529), int32(319059))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v545 = v448
	goto L115
L129:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+271)))
	if v494 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v538 = F_systable_getnext(m, v437)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L141
	}
L131:
	;
	v499 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v519 = F_text_to_cstring(m, v492)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L138
	}
L134:
	;
	if v499 == int32(0) {
		v536 = v448
		goto L130
	} else {
		goto L135
	}
L135:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v504 = int32(*(*int16)(unsafe.Add(mBase, uint32(v464)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v503 + int32(4)
	F_errmsg_internal(m, int32(682311), v22+int32(96))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(490658), int32(4538), int32(319059))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v536 = v448
	goto L130
L138:
	;
	v523 = v417 + v448<<(uint(int32(3))%32)
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v464)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v523))) = uint16(v524)
	v527 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v528 = F_MemoryContextStrdup(m, v527, v519)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v523)+4)) = v528
	F_pfree(m, v519)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v536 = v448 + int32(1)
	goto L130
L141:
	;
	if v538 != 0 {
		v445 = v538
		v448 = v536
		goto L120
	} else {
		goto L142
	}
L142:
	;
	goto L121
L143:
	;
	F_sequence_close(m, v429, int32(1))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	if v545 == v379 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if int32(2) <= v545 {
		goto L151
	} else {
		goto L152
	}
L146:
	;
	v567 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v567 == int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v379 - v545
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v571 + int32(4)
	F_errmsg_internal(m, int32(679320), v22+int32(80))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(490658), int32(4556), int32(319059))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	goto L145
L151:
	;
	F_pg_qsort(m, v417, v545, int32(8), int32(1605))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v595))) = v417
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v598)+12)) = uint16(v545)
	goto L109
L154:
	;
	goto L153
L155:
	;
	v917 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v903)+122)))
	if v917 != 0 {
		goto L99
	} else {
		goto L218
	}
L156:
	;
	v858 = int32(0)
	v859 = int32(*(*int16)(unsafe.Add(mBase, uint32(v844)+120)))
	if v859 <= v858 {
		v903 = v844
		goto L155
	} else {
		goto L211
	}
L157:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	F_ScanKeyInit(m, v22+int32(272), int32(9), int32(3), int32(184), v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L164
	}
L158:
	;
	if base.Ui32(v405) < base.Ui32(int32(12000)) {
		v903 = v622
		goto L155
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v634 = F_MemoryContextAllocZero(m, v631, v623*int32(12))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L163
	}
L161:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+16)))
	if v626 != int32(1) {
		v844 = v622
		goto L156
	} else {
		goto L162
	}
L162:
	;
	v636 = int32(0)
	goto L157
L163:
	;
	v636 = v634
	goto L157
L164:
	;
	v645 = int32(0)
	v648 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	F_systable_endscan(m, v656)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L4
	} else {
		goto L198
	}
L166:
	;
	v651 = int32(1)
	v656 = F_systable_beginscan(m, v648, int32(2665), v651, int32(0), v651, v22+int32(272))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v658 = F_systable_getnext(m, v656)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	if v658 == int32(0) {
		v785 = v645
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v667 = v658
	v669 = v645
	goto L170
L170:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v667)+16))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+22)))
	v683 = v681 + v682
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+72)))
	switch v684 - int32(99) {
	case 0:
		goto L173
	default:
		v775 = v669
		goto L172
	case 11:
		goto L174
	}
L171:
	;
	v785 = v775
	goto L165
L172:
	;
	v776 = F_systable_getnext(m, v656)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L4
	} else {
		goto L196
	}
L173:
	;
	if v623 <= v669 {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+76)))
	if v687 != 0 {
		v775 = v669
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v688 = F_extractNotNullColumn(m, v667)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v694 = int32(105)
	*(*uint8)(unsafe.Add(mBase, uint32(v690+v688<<(uint(int32(4))%32))+15)) = uint8(v694)
	v775 = v669
	goto L172
L177:
	;
	v699 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L4
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v719 = v636 + v669*int32(12)
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v719)+8)) = uint8(v720)
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v719)+9)) = uint8(v722)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v719)+10)) = uint8(v724)
	v727 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v730 = F_MemoryContextStrdup(m, v727, v683+int32(4))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L184
	}
L180:
	;
	if v699 == int32(0) {
		v785 = v669
		goto L165
	} else {
		goto L181
	}
L181:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v703 + int32(4)
	F_errmsg_internal(m, int32(679424), v22-int32(-64))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(490658), int32(4657), int32(319036))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v785 = v669
	goto L165
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v719))) = v730
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v648)+52))
	v737 = F_fastgetattr_3(m, v667, int32(28), v734, v22+int32(271))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+271)))
	if v739 == int32(1) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v744 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v762 = F_text_to_cstring(m, v737)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L193
	}
L189:
	;
	if v744 == int32(0) {
		v775 = v669
		goto L172
	} else {
		goto L190
	}
L190:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v748 + int32(4)
	F_errmsg_internal(m, int32(679237), v22+int32(48))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(490658), int32(4673), int32(319036))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	v775 = v669
	goto L172
L193:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v766 = F_MemoryContextStrdup(m, v765, v762)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v719)+4)) = v766
	F_pfree(m, v762)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	v775 = v669 + int32(1)
	goto L172
L196:
	;
	if v776 != 0 {
		v667 = v776
		v669 = v775
		goto L170
	} else {
		goto L197
	}
L197:
	;
	goto L171
L198:
	;
	F_sequence_close(m, v648, int32(1))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	if v785 == v623 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	if int32(2) <= v785 {
		goto L206
	} else {
		goto L207
	}
L201:
	;
	v805 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	if v805 == int32(0) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v623 - v785
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v809 + int32(4)
	F_errmsg_internal(m, int32(679267), v22+int32(32))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(490658), int32(4690), int32(319036))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	goto L200
L206:
	;
	F_pg_qsort(m, v636, v785, int32(12), int32(1606))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v833)+4)) = v636
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v836)+14)) = uint16(v785)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	if base.Ui32(v405) < base.Ui32(int32(12000)) {
		v903 = v838
		goto L155
	} else {
		goto L210
	}
L209:
	;
	goto L208
L210:
	;
	v844 = v838
	goto L156
L211:
	;
	v864 = v858
	v867 = v844
	goto L212
L212:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v886 = v881 + v864<<(uint(int32(4))%32) + int32(31)
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	if v887 == int32(117) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v903 = v893
	goto L155
L214:
	;
	v890 = int32(118)
	*(*uint8)(unsafe.Add(mBase, uint32(v886))) = uint8(v890)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v893 = v892
	goto L216
L215:
	;
	v893 = v867
	goto L216
L216:
	;
	v895 = v864 + int32(1)
	v896 = int32(*(*int16)(unsafe.Add(mBase, uint32(v893)+120)))
	if v895 < v896 {
		v864 = v895
		v867 = v893
		goto L212
	} else {
		goto L217
	}
L217:
	;
	goto L213
L218:
	;
	v918 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+14)) = uint16(v918)
	goto L99
L219:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v922)+16)) = int32(0)
	goto L99
L220:
	;
	F_RelationParseRelOptions(m, v93, v78)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L4
	} else {
		goto L225
	}
L221:
	;
	F_RelationInitTableAccessMethod(m, v93)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L4
	} else {
		goto L224
	}
L222:
	;
	F_RelationInitIndexAccessInfo(m, v93)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	goto L220
L224:
	;
	goto L220
L225:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+124)))
	if v971 == int32(1) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+125)))
	if v980 == int32(1) {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	F_RelationBuildRuleLock(m, v93)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L4
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v93)+68)) = int64(0)
	v979 = v970
	goto L226
L230:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v979 = v976
	goto L226
L231:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988)+127)))
	if v989 == int32(1) {
		goto L237
	} else {
		goto L238
	}
L232:
	;
	F_RelationBuildTriggers(m, v93)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L4
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+76)) = int32(0)
	v988 = v979
	goto L231
L235:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v988 = v985
	goto L231
L236:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+60)) = v996
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+117)))
	if v1002 != 0 {
		goto L242
	} else {
		goto L243
	}
L237:
	;
	F_RelationBuildRowSecurity(m, v93)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L4
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = int32(0)
	goto L236
L240:
	;
	goto L236
L241:
	;
	F_RelationInitPhysicalAddr(m, v93)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L4
	} else {
		goto L245
	}
L242:
	;
	v1003 = int32(0)
	goto L244
L243:
	;
	v1003 = v1000
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v1003
	goto L241
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = int32(0)
	F_pfree(m, v78)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012+v52)+4)))
	if v1014 != int32(1) {
		goto L9
	} else {
		goto L247
	}
L247:
	;
	F_RelationDestroyRelation(m, v93, int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L4
	} else {
		goto L248
	}
L248:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	v1023 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1021+v52)+4)) = uint8(v1023)
	v1027 = F_ScanPgRelation(m, l0, int32(1), v1023)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	if v1027 != 0 {
		v78 = v1027
		goto L14
	} else {
		goto L250
	}
L250:
	;
	goto L15
L251:
	;
	v1108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+26)) = uint8(v1108)
	v1113 = v93
	goto L8
L252:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
	v1068 = F_hash_search(m, v1064, v177, int32(1), v22+int32(160))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+160)))
	if v1070 == int32(1) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+4)) = v93
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+16))
	if v1075 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+4)) = v93
	goto L251
L257:
	;
	F_RelationDestroyRelation(m, v1073, int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L4
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v1082 == int32(0) {
		goto L251
	} else {
		goto L261
	}
L260:
	;
	goto L251
L261:
	;
	v1087 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	if v1087 == int32(0) {
		goto L251
	} else {
		goto L263
	}
L263:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v1091 + int32(4)
	F_errmsg_internal(m, int32(674722), v22+int32(16))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(490658), int32(1314), int32(480388))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	goto L251
L266:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v1138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v1138
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v1137 + int32(4)
	F_errmsg_internal(m, int32(679562), v22+int32(144))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L4
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(490658), int32(585), int32(480356))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v1157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v380
	F_errmsg_internal(m, int32(55325), v22+int32(128))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(490658), int32(667), int32(480356))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCacheInitFilePostInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v2+int32(2048))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_RelationCacheInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	v1 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1127]))
	if v13 == int32(5842711) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_read_relmap_file(m, int32(4459356), int32(309148), int32(0), int32(22))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[1128]))
	if v23 == int32(5842711) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[1121]))
	F_read_relmap_file(m, int32(4460404), v28, int32(0), int32(22))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
	F_hash_seq_init(m, v10+int32(12), v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v41 = F_hash_seq_search(m, v10+int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	F_list_free(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L129
	}
L12:
	;
	F_list_free(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L96
	}
L13:
	;
	v238 = v138
	goto L12
L14:
	;
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = v41
	v45 = v1
	v46 = v1
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_smgrreleaseall(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L94
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	if v51 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F_smgrreleaseall(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L63
	}
L20:
	;
	goto L19
L21:
	;
	v61 = int32(4458556)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	*(*int32)(unsafe.Add(mBase, _consts[1129])) = v63 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v67 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	if v54 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v59 = F_hash_seq_search(m, v10+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	if v59 != 0 {
		v43 = v59
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v138 = v45
	v139 = v46
	goto L20
L28:
	;
	v134 = F_hash_seq_search(m, v10+int32(12))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L61
	}
L29:
	;
	F_RelationClearRelation(m, v50)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+119)))
	switch v73 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L34
	default:
		goto L33
	}
L32:
	;
	v130 = v45
	v131 = v46
	goto L28
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
	if v113 != int32(2662) {
		goto L50
	} else {
		goto L51
	}
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+88))
	if v76 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v77 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
	v82 = v80 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+72)) = v82
	if v82 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	F_RelationInitPhysicalAddr(m, v50)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L48
	}
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	F_smgrclose(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L47
	}
L40:
	;
	v87 = v77 + int32(76)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	if v89 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+76)) = v96
	v98 = int32(4394040)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+80)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v87
	*(*int32)(unsafe.Add(mBase, _consts[799])) = v87
	goto L42
L44:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[799]))
	v96 = v91
	goto L43
L45:
	;
	goto L46
L46:
	;
	v93 = int32(4394040)
	*(*int32)(unsafe.Add(mBase, _consts[798])) = v93
	v96 = v93
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = int32(0)
	goto L38
L48:
	;
	goto L33
L49:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+25)))
	if v122 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	if v113 != int32(1259) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v120 = F_lappend(m, v45, v50)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	v118 = F_lcons(m, v50, v45)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v130 = v118
	v131 = v46
	goto L28
L55:
	;
	v130 = v120
	v131 = v46
	goto L28
L56:
	;
	v125 = F_lcons(m, v50, v46)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v127 = F_lappend(m, v46, v50)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	v130 = v45
	v131 = v125
	goto L28
L60:
	;
	v130 = v45
	v131 = v127
	goto L28
L61:
	;
	if v134 != 0 {
		v43 = v134
		v45 = v130
		v46 = v131
		goto L18
	} else {
		goto L62
	}
L62:
	;
	v138 = v130
	v139 = v131
	goto L20
L63:
	;
	v142 = int32(0)
	if v138 == v142 {
		v238 = v142
		goto L12
	} else {
		goto L64
	}
L64:
	;
	v145 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v146 <= v145 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v150 = v145
	goto L66
L66:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v157 = int32(2)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v150<<(uint(v157)%32))))
	v162 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	goto L70
L67:
	;
	goto L13
L68:
	;
	v216 = v150 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v216 < v217 {
		v150 = v216
		goto L66
	} else {
		goto L93
	}
L69:
	;
	F_RelationRebuildRelation(m, v160)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L92
	}
L70:
	;
	if v163 == v157 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+25)))
	if v166 != int32(1) {
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v172 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	if v169 != int32(1) {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+72))
	v177 = v175 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+72)) = v177
	if v177 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v160)+256))
	if v205 != 0 {
		goto L88
	} else {
		goto L89
	}
L79:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	F_smgrclose(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L87
	}
L80:
	;
	v182 = v172 + int32(76)
	v184 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	if v184 != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+76)) = v191
	v193 = int32(4394040)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+80)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v182
	*(*int32)(unsafe.Add(mBase, _consts[799])) = v182
	goto L82
L84:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[799]))
	v191 = v186
	goto L83
L85:
	;
	goto L86
L86:
	;
	v188 = int32(4394040)
	*(*int32)(unsafe.Add(mBase, _consts[798])) = v188
	v191 = v188
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = int32(0)
	goto L78
L88:
	;
	F_pfree(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+26)) = uint8(v208)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+256)) = v208
	goto L68
L91:
	;
	goto L90
L92:
	;
	goto L68
L93:
	;
	goto L67
L94:
	;
	F_list_free(m, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v323 = v1
	goto L11
L96:
	;
	if v139 == int32(0) {
		v323 = v1
		goto L11
	} else {
		goto L97
	}
L97:
	;
	v243 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v244 <= v243 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v323 = v139
	goto L11
L99:
	;
	goto L100
L100:
	;
	v248 = v243
	goto L101
L101:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v255 = int32(2)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254+v248<<(uint(v255)%32))))
	v260 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	goto L105
L102:
	;
	v323 = v139
	goto L11
L103:
	;
	v314 = v248 + int32(1)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v314 < v315 {
		v248 = v314
		goto L101
	} else {
		goto L128
	}
L104:
	;
	F_RelationRebuildRelation(m, v258)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L127
	}
L105:
	;
	if v261 == v255 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+25)))
	if v264 != int32(1) {
		goto L104
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	if v270 != 0 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v258)+16))
	if v267 != int32(1) {
		goto L104
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+72))
	v275 = v273 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v270)+72)) = v275
	if v275 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v258)+256))
	if v303 != 0 {
		goto L123
	} else {
		goto L124
	}
L114:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	F_smgrclose(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L122
	}
L115:
	;
	v280 = v270 + int32(76)
	v282 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	if v282 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L117
L117:
	;
	goto L114
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270)+76)) = v289
	v291 = int32(4394040)
	*(*int32)(unsafe.Add(mBase, uint32(v270)+80)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v280
	*(*int32)(unsafe.Add(mBase, _consts[799])) = v280
	goto L117
L119:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _consts[799]))
	v289 = v284
	goto L118
L120:
	;
	goto L121
L121:
	;
	v286 = int32(4394040)
	*(*int32)(unsafe.Add(mBase, _consts[798])) = v286
	v289 = v286
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = int32(0)
	goto L113
L123:
	;
	F_pfree(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+26)) = uint8(v306)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+256)) = v306
	goto L103
L126:
	;
	goto L125
L127:
	;
	goto L103
L128:
	;
	goto L102
L129:
	;
	v326 = int32(0)
	v328 = *(*int32)(unsafe.Add(mBase, _consts[1123]))
	if v328 <= v326 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	m.G0 = v10 + int32(32)
	return
L131:
	;
	v332 = v328 & int32(7)
	v334 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	if base.Ui32(int32(8)) <= base.Ui32(v328) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v341 = v326
	v343 = int32(0)
	goto L135
L133:
	;
	v372 = v326
	goto L134
L134:
	;
	if v332 == int32(0) {
		goto L130
	} else {
		goto L138
	}
L135:
	;
	v349 = v334 + v341<<(uint(int32(3))%32)
	v350 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+12)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+4)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+20)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+28)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+36)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+44)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+52)) = uint8(v350)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+60)) = uint8(v350)
	v366 = int32(8)
	v367 = v341 + v366
	v369 = v343 + v366
	if v369 != v328&int32(2147483640) {
		v341 = v367
		v343 = v369
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v372 = v367
	goto L134
L137:
	;
	goto L136
L138:
	;
	v381 = int32(0)
	v382 = v372
	goto L139
L139:
	;
	v391 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v334+v382<<(uint(int32(3))%32))+4)) = uint8(v391)
	v396 = v381 + v391
	if v396 != v332 {
		v381 = v396
		v382 = v382 + v391
		goto L139
	} else {
		goto L141
	}
L140:
	;
	goto L130
L141:
	;
	goto L140
}
func F_RelationClearMissing(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+120)))
	v20 = F__emscripten_memset_bulkmem(m, v10+int32(80), base.I32_extend8_s(int32(0)), int32(100))
	mBase = m.M
	goto L1
L1:
	;
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v23
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v27)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+72)) = uint8(v27)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v23
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v27)
	v43 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	if int32(0) < v14 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L21
	}
L5:
	;
	v48 = int32(1)
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_sequence_close(m, v43, int32(3))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L20
	}
L8:
	;
	v57 = F_SearchSysCache2(m, int32(7), v12, base.I32_extend16_s(v48))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	if v57 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+22)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v62)+88)))
	if v64 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v74 = F_heap_modify_tuple(m, v57, v67, v10+int32(80), v10+int32(48), v10+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_ReleaseCatCache(m, v57)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L18
	}
L15:
	;
	F_CatalogTupleUpdate(m, v43, v74+int32(4), v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	F_pfree(m, v74)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	if v48 != v14 {
		v48 = v48 + int32(1)
		goto L8
	} else {
		goto L19
	}
L19:
	;
	goto L9
L20:
	;
	m.G0 = v10 + int32(192)
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v48
	F_errmsg_internal(m, int32(45850), v10)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(487561), int32(1999), int32(323839))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationForgetRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1126]))
	v15 = F_hash_search(m, v10, v6+int32(12), v2, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 == int32(0) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v19 == int32(0) {
				m.G0 = v6 + int32(16)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				if v22 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v84
						F_errmsg_internal(m, int32(277141), v6)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errfinish(m, int32(490658), int32(2903), int32(259943))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
					if v23 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						if v26 == int32(0) {
							F_RelationClearRelation(m, v19)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _consts[39]))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v31
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							if v33 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
								v38 = v36 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v38
								if v38 == int32(0) {
									v43 = v33 + int32(76)
									v45 = *(*int32)(unsafe.Add(mBase, _consts[798]))
									if v45 != 0 {
										v47 = *(*int32)(unsafe.Add(mBase, _consts[799]))
										v52 = v47
									} else {
										v49 = int32(4394040)
										*(*int32)(unsafe.Add(mBase, _consts[798])) = v49
										v52 = v49
									}
									*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v52
									v54 = int32(4394040)
									*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v54
									*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v43
									*(*int32)(unsafe.Add(mBase, _consts[799])) = v43
								} else {
								}
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
								F_smgrclose(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
									if v66 != 0 {
										F_pfree(m, v66)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
											*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
											m.G0 = v6 + int32(16)
											return
										}
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
								if v66 != 0 {
									F_pfree(m, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v31
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if v33 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
							v38 = v36 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v38
							if v38 == int32(0) {
								v43 = v33 + int32(76)
								v45 = *(*int32)(unsafe.Add(mBase, _consts[798]))
								if v45 != 0 {
									v47 = *(*int32)(unsafe.Add(mBase, _consts[799]))
									v52 = v47
								} else {
									v49 = int32(4394040)
									*(*int32)(unsafe.Add(mBase, _consts[798])) = v49
									v52 = v49
								}
								*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v52
								v54 = int32(4394040)
								*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v43
								*(*int32)(unsafe.Add(mBase, _consts[799])) = v43
							} else {
							}
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							F_smgrclose(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
								if v66 != 0 {
									F_pfree(m, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
							if v66 != 0 {
								F_pfree(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v69 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationGetFKeyList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v13 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v115
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v115 = v16
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v11, int32(9), int32(3), int32(184), v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v27 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v30 = int32(1)
	v33 = F_systable_beginscan(m, v27, int32(2665), v30, int32(0), v30, v11)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v35 = F_systable_getnext(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v2
	v41 = v35
	goto L13
L11:
	;
	v86 = v2
	goto L12
L12:
	;
	F_systable_endscan(m, v33)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L23
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v47 = v45 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+72)))
	if v48 == int32(102) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v86 = v81
	goto L12
L15:
	;
	v52 = F_palloc0(m, int32(280))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	v81 = v39
	goto L17
L17:
	;
	v82 = F_systable_getnext(m, v33)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(470)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+20)) = uint8(v62)
	v72 = int32(0)
	F_DeconstructFkConstraintRow(m, v41, v52+int32(16), v52+int32(22), v52+int32(86), v52+int32(152), v72, v72, v72, v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v78 = F_lappend(m, v39, v52)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v81 = v78
	goto L17
L21:
	;
	if v82 != 0 {
		v39 = v81
		v41 = v82
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F_sequence_close(m, v27, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v97 = int32(4470560)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v101 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v104 = F_copyObjectImpl(m, v86)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v104
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v98
	F_list_free_deep(m, v103)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v115 = v86
	goto L1
}
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	v7 = F_palloc(m, int32(92))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		if v11 < l1 {
			v24 = F_palloc(m, l1*int32(48))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = v24
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v27
				if int32(0) < l2 {
					v33 = F_palloc(m, l2*int32(48))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v33
						v36 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
						v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
						v44 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
						*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
						v52 = v43 ^ int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
						return v7
					}
				} else {
					v35 = int32(0)
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
					v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
					v52 = v43 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
					return v7
				}
			}
		} else {
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v27
			if int32(0) < l2 {
				v33 = F_palloc(m, l2*int32(48))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = v33
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
					v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
					v52 = v43 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
					return v7
				}
			} else {
				v35 = int32(0)
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
				v42 = *(*int32)(unsafe.Add(mBase, _consts[39]))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
				v44 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
				v52 = v43 ^ int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
				return v7
			}
		}
	}
}
func F_RelationIncrementReferenceCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_ResourceOwnerEnlarge(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6 + int32(1)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[231]))
		if v11 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			F_ResourceOwnerRemember(m, v13, l0, int32(1714624))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_RelationInitLockInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v2
	v6 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+117)))
	if v8 != 0 {
		v9 = int32(0)
	} else {
		v9 = v6
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v9
	return
}
func F_RelationIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
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
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v133
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v133 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(45662), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(490958), int32(940), int32(63625))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L40
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v87 == int32(0) {
		v123 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v123 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= v90 {
		v123 = v3
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v96 = v90
	goto L34
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v109 = base.B2i32(v108 == v40)
	if v108 == v40 {
		v123 = v109
		goto L14
	} else {
		goto L36
	}
L35:
	;
	v123 = v109
	goto L14
L36:
	;
	v110 = F_get_relname_relid(m, v39+int32(4), v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v110 != 0 {
		v123 = v109
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v113 = v96 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v113 < v114 {
		v96 = v113
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v133 = v123
	goto L1
}
func F_RelationMapInvalidate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, _consts[1127]))
		if v3 != int32(5842711) {
			return
		} else {
			F_read_relmap_file(m, int32(4459356), int32(309148), int32(0), int32(22))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[1128]))
		if v13 != int32(5842711) {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[1121]))
			F_read_relmap_file(m, int32(4460404), v18, int32(0), int32(22))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_relation_needs_vacanalyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 float64
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 float32
	_ = v184
	var v185 float32
	_ = v185
	var v188 float32
	_ = v188
	var v189 float32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v204 float32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v212 float32
	_ = v212
	var v215 float32
	_ = v215
	var v219 float32
	_ = v219
	var v221 float32
	_ = v221
	var v222 float32
	_ = v222
	var v223 float32
	_ = v223
	var v226 float32
	_ = v226
	var v230 float32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v257 int32
	_ = v257
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	v30 = m.G0
	v32 = v30 - int32(112)
	m.G0 = v32
	if l1 != 0 {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		if v34 < l4 {
			v36 = v34
		} else {
			v36 = l4
		}
		if v34 < int32(0) {
			v39 = l4
		} else {
			v39 = v36
		}
		v41 = *(*int32)(unsafe.Add(mBase, _consts[47]))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		if v42 < v41 {
			v44 = v42
		} else {
			v44 = v41
		}
		if v42 < int32(0) {
			v47 = v41
		} else {
			v47 = v44
		}
		v49 = *(*int32)(unsafe.Add(mBase, _consts[402]))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		if v50 < int32(0) {
			v53 = v49
		} else {
			v53 = v50
		}
		v54 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
		v56 = *(*float64)(unsafe.Add(mBase, _consts[403]))
		if base.F64_ge(v54, float64(0)) != 0 {
			v59 = v54
		} else {
			v59 = v56
		}
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v62 = *(*int32)(unsafe.Add(mBase, _consts[404]))
		if int32(-2) < v60 {
			v65 = v60
		} else {
			v65 = v62
		}
		v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+72))
		v68 = *(*float64)(unsafe.Add(mBase, _consts[405]))
		if base.F64_ge(v66, float64(0)) != 0 {
			v71 = v66
		} else {
			v71 = v68
		}
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v74 = *(*int32)(unsafe.Add(mBase, _consts[406]))
		if int32(-2) < v72 {
			v77 = v72
		} else {
			v77 = v74
		}
		v79 = *(*int32)(unsafe.Add(mBase, _consts[407]))
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v80 < int32(0) {
			v83 = v79
		} else {
			v83 = v80
		}
		v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
		v86 = *(*float64)(unsafe.Add(mBase, _consts[408]))
		if base.F64_ge(v84, float64(0)) != 0 {
			v89 = v84
		} else {
			v89 = v86
		}
		v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v108 = v39
		v110 = v47
		v111 = v65
		v112 = v77
		v113 = v53
		v114 = v83
		v115 = v89
		v116 = v59
		v117 = v71
		v118 = v90
	} else {
		v92 = *(*int32)(unsafe.Add(mBase, _consts[47]))
		v94 = *(*int32)(unsafe.Add(mBase, _consts[402]))
		v96 = *(*float64)(unsafe.Add(mBase, _consts[403]))
		v98 = *(*int32)(unsafe.Add(mBase, _consts[404]))
		v100 = *(*float64)(unsafe.Add(mBase, _consts[405]))
		v102 = *(*int32)(unsafe.Add(mBase, _consts[406]))
		v104 = *(*int32)(unsafe.Add(mBase, _consts[407]))
		v106 = *(*float64)(unsafe.Add(mBase, _consts[408]))
		v108 = l4
		v110 = v92
		v111 = v98
		v112 = v102
		v113 = v94
		v114 = v104
		v115 = v106
		v116 = v96
		v117 = v100
		v118 = int32(1)
	}
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+136))
	if base.Ui32(v119) < base.Ui32(int32(3)) {
		v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+140))
		if v148 != 0 {
			v151 = *(*int32)(unsafe.Add(mBase, _consts[409]))
			if v108 == v151 {
				v154 = int32(-1)
			} else {
				v154 = v151 - v108
			}
			v160 = int32(base.Ui32(v148-v154) >> (uint(int32(31)) % 32))
		} else {
			v160 = int32(0)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v160)
		if (v160|v118)&int32(1) != 0 {
			v169 = v160
			if l3 == int32(0) {
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
				v307 = int32(0)
				if l0 != int32(2619) {
					v311 = v307
				} else {
					v311 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
				m.G0 = v32 + int32(112)
				return
			} else {
				v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[410])))
				if v174 != int32(1) {
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
					v307 = int32(0)
					if l0 != int32(2619) {
						v311 = v307
					} else {
						v311 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
					m.G0 = v32 + int32(112)
					return
				} else {
					v178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
					if v178&int32(1) == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
						v307 = int32(0)
						if l0 != int32(2619) {
							v311 = v307
						} else {
							v311 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
						m.G0 = v32 + int32(112)
						return
					} else {
						v184 = float32(0)
						v185 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
						if base.F32_lt(v185, v184) != 0 {
							v188 = v184
						} else {
							v188 = v185
						}
						v189 = float32(1)
						v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
						if v190 <= int32(0) {
							v204 = v189
						} else {
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
							if v193 <= int32(0) {
								v204 = v189
							} else {
								if base.Ui32(v193) < base.Ui32(v190) {
									v198 = v193
								} else {
									v198 = v190
								}
								v204 = base.F32_sub(float32(1), base.F32_div(base.F32_convert_i32_u(v198), base.F32_convert_i32_u(v190)))
							}
						}
						v205 = *(*int64)(unsafe.Add(mBase, uint32(l3)+88))
						v206 = *(*int64)(unsafe.Add(mBase, uint32(l3)+96))
						v207 = *(*int64)(unsafe.Add(mBase, uint32(l3)+80))
						v212 = base.F32_add(base.F32_mul(base.F32_demote_f64(v115), v188), base.F32_convert_i32_s(v114))
						if v112 < int32(0) {
							v219 = v212
						} else {
							v215 = base.F32_convert_i32_u(v112)
							if base.F32_gt(v212, v215) == int32(0) {
								v219 = v212
							} else {
								v219 = v215
							}
						}
						v221 = base.F32_convert_i64_s(v205)
						v222 = base.F32_convert_i64_s(v206)
						v223 = base.F32_convert_i64_s(v207)
						v226 = base.F32_add(base.F32_mul(base.F32_demote_f64(v116), v188), base.F32_convert_i32_s(v113))
						v230 = base.F32_add(base.F32_mul(base.F32_mul(v188, base.F32_demote_f64(v117)), v204), base.F32_convert_i32_s(v111))
						v233 = F_errstart(m, int32(12), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return
						} else {
							if int32(0) <= v111 {
								if v233 == int32(0) {
									if v169|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v307 = base.F32_lt(v226, v221)
									if l0 != int32(2619) {
										v311 = v307
									} else {
										v311 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+40)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = base.F64_promote_f32(v230)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+24)) = base.F64_promote_f32(v222)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+16)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2 + int32(4)
									F_errmsg_internal(m, int32(650856), v32)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										v278 = int32(3138)
										F_errfinish(m, int32(488518), v278, int32(334470))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v169|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v307 = base.F32_lt(v226, v221)
											if l0 != int32(2619) {
												v311 = v307
											} else {
												v311 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
											m.G0 = v32 + int32(112)
											return
										}
									}
								}
							} else {
								if v233 == int32(0) {
									if v169|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v307 = base.F32_lt(v226, v221)
									if l0 != int32(2619) {
										v311 = v307
									} else {
										v311 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+96)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+88)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+80)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+72)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l2 + int32(4)
									F_errmsg_internal(m, int32(650943), v32-int32(-64))
									mBase = m.M
									v276 = m.ExcPending
									if v276 != 0 {
										return
									} else {
										v278 = int32(3142)
										F_errfinish(m, int32(488518), v278, int32(334470))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v169|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v307 = base.F32_lt(v226, v221)
											if l0 != int32(2619) {
												v311 = v307
											} else {
												v311 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
											m.G0 = v32 + int32(112)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v165 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v165)
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v165)
			m.G0 = v32 + int32(112)
			return
		}
	} else {
		v123 = *(*int32)(unsafe.Add(mBase, _consts[411]))
		v124 = v123 - v110
		v125 = int32(3)
		if base.Ui32(v124) < base.Ui32(v125) {
			v129 = v124 - v125
		} else {
			v129 = v124
		}
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v129))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v119)) == int32(0) {
			v141 = base.B2i32(base.Ui32(v119) < base.Ui32(v129))
		} else {
			v141 = int32(base.Ui32(v119-v129) >> (uint(int32(31)) % 32))
		}
		if v141 == int32(0) {
			v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+140))
			if v148 != 0 {
				v151 = *(*int32)(unsafe.Add(mBase, _consts[409]))
				if v108 == v151 {
					v154 = int32(-1)
				} else {
					v154 = v151 - v108
				}
				v160 = int32(base.Ui32(v148-v154) >> (uint(int32(31)) % 32))
			} else {
				v160 = int32(0)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v160)
			if (v160|v118)&int32(1) != 0 {
				v169 = v160
				if l3 == int32(0) {
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
					v307 = int32(0)
					if l0 != int32(2619) {
						v311 = v307
					} else {
						v311 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
					m.G0 = v32 + int32(112)
					return
				} else {
					v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[410])))
					if v174 != int32(1) {
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
						v307 = int32(0)
						if l0 != int32(2619) {
							v311 = v307
						} else {
							v311 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
						m.G0 = v32 + int32(112)
						return
					} else {
						v178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
						if v178&int32(1) == int32(0) {
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
							v307 = int32(0)
							if l0 != int32(2619) {
								v311 = v307
							} else {
								v311 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
							m.G0 = v32 + int32(112)
							return
						} else {
							v184 = float32(0)
							v185 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
							if base.F32_lt(v185, v184) != 0 {
								v188 = v184
							} else {
								v188 = v185
							}
							v189 = float32(1)
							v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
							if v190 <= int32(0) {
								v204 = v189
							} else {
								v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
								if v193 <= int32(0) {
									v204 = v189
								} else {
									if base.Ui32(v193) < base.Ui32(v190) {
										v198 = v193
									} else {
										v198 = v190
									}
									v204 = base.F32_sub(float32(1), base.F32_div(base.F32_convert_i32_u(v198), base.F32_convert_i32_u(v190)))
								}
							}
							v205 = *(*int64)(unsafe.Add(mBase, uint32(l3)+88))
							v206 = *(*int64)(unsafe.Add(mBase, uint32(l3)+96))
							v207 = *(*int64)(unsafe.Add(mBase, uint32(l3)+80))
							v212 = base.F32_add(base.F32_mul(base.F32_demote_f64(v115), v188), base.F32_convert_i32_s(v114))
							if v112 < int32(0) {
								v219 = v212
							} else {
								v215 = base.F32_convert_i32_u(v112)
								if base.F32_gt(v212, v215) == int32(0) {
									v219 = v212
								} else {
									v219 = v215
								}
							}
							v221 = base.F32_convert_i64_s(v205)
							v222 = base.F32_convert_i64_s(v206)
							v223 = base.F32_convert_i64_s(v207)
							v226 = base.F32_add(base.F32_mul(base.F32_demote_f64(v116), v188), base.F32_convert_i32_s(v113))
							v230 = base.F32_add(base.F32_mul(base.F32_mul(v188, base.F32_demote_f64(v117)), v204), base.F32_convert_i32_s(v111))
							v233 = F_errstart(m, int32(12), int32(0))
							mBase = m.M
							v234 = m.ExcPending
							if v234 != 0 {
								return
							} else {
								if int32(0) <= v111 {
									if v233 == int32(0) {
										if v169|base.F32_lt(v219, v223) != 0 {
											v289 = int32(1)
										} else {
											v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
										v307 = base.F32_lt(v226, v221)
										if l0 != int32(2619) {
											v311 = v307
										} else {
											v311 = int32(0)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
										m.G0 = v32 + int32(112)
										return
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = base.F64_promote_f32(v226)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+40)) = base.F64_promote_f32(v221)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = base.F64_promote_f32(v230)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+24)) = base.F64_promote_f32(v222)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+16)) = base.F64_promote_f32(v219)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = base.F64_promote_f32(v223)
										*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2 + int32(4)
										F_errmsg_internal(m, int32(650856), v32)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											v278 = int32(3138)
											F_errfinish(m, int32(488518), v278, int32(334470))
											mBase = m.M
											v281 = m.ExcPending
											if v281 != 0 {
												return
											} else {
												if v169|base.F32_lt(v219, v223) != 0 {
													v289 = int32(1)
												} else {
													v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
												v307 = base.F32_lt(v226, v221)
												if l0 != int32(2619) {
													v311 = v307
												} else {
													v311 = int32(0)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
												m.G0 = v32 + int32(112)
												return
											}
										}
									}
								} else {
									if v233 == int32(0) {
										if v169|base.F32_lt(v219, v223) != 0 {
											v289 = int32(1)
										} else {
											v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
										v307 = base.F32_lt(v226, v221)
										if l0 != int32(2619) {
											v311 = v307
										} else {
											v311 = int32(0)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
										m.G0 = v32 + int32(112)
										return
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v32)+96)) = base.F64_promote_f32(v226)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+88)) = base.F64_promote_f32(v221)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+80)) = base.F64_promote_f32(v219)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+72)) = base.F64_promote_f32(v223)
										*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l2 + int32(4)
										F_errmsg_internal(m, int32(650943), v32-int32(-64))
										mBase = m.M
										v276 = m.ExcPending
										if v276 != 0 {
											return
										} else {
											v278 = int32(3142)
											F_errfinish(m, int32(488518), v278, int32(334470))
											mBase = m.M
											v281 = m.ExcPending
											if v281 != 0 {
												return
											} else {
												if v169|base.F32_lt(v219, v223) != 0 {
													v289 = int32(1)
												} else {
													v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
												v307 = base.F32_lt(v226, v221)
												if l0 != int32(2619) {
													v311 = v307
												} else {
													v311 = int32(0)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
												m.G0 = v32 + int32(112)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v165 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v165)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v165)
				m.G0 = v32 + int32(112)
				return
			}
		} else {
			v144 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v144)
			v169 = v144
			if l3 == int32(0) {
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
				v307 = int32(0)
				if l0 != int32(2619) {
					v311 = v307
				} else {
					v311 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
				m.G0 = v32 + int32(112)
				return
			} else {
				v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[410])))
				if v174 != int32(1) {
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
					v307 = int32(0)
					if l0 != int32(2619) {
						v311 = v307
					} else {
						v311 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
					m.G0 = v32 + int32(112)
					return
				} else {
					v178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
					if v178&int32(1) == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v169)
						v307 = int32(0)
						if l0 != int32(2619) {
							v311 = v307
						} else {
							v311 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
						m.G0 = v32 + int32(112)
						return
					} else {
						v184 = float32(0)
						v185 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
						if base.F32_lt(v185, v184) != 0 {
							v188 = v184
						} else {
							v188 = v185
						}
						v189 = float32(1)
						v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
						if v190 <= int32(0) {
							v204 = v189
						} else {
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
							if v193 <= int32(0) {
								v204 = v189
							} else {
								if base.Ui32(v193) < base.Ui32(v190) {
									v198 = v193
								} else {
									v198 = v190
								}
								v204 = base.F32_sub(float32(1), base.F32_div(base.F32_convert_i32_u(v198), base.F32_convert_i32_u(v190)))
							}
						}
						v205 = *(*int64)(unsafe.Add(mBase, uint32(l3)+88))
						v206 = *(*int64)(unsafe.Add(mBase, uint32(l3)+96))
						v207 = *(*int64)(unsafe.Add(mBase, uint32(l3)+80))
						v212 = base.F32_add(base.F32_mul(base.F32_demote_f64(v115), v188), base.F32_convert_i32_s(v114))
						if v112 < int32(0) {
							v219 = v212
						} else {
							v215 = base.F32_convert_i32_u(v112)
							if base.F32_gt(v212, v215) == int32(0) {
								v219 = v212
							} else {
								v219 = v215
							}
						}
						v221 = base.F32_convert_i64_s(v205)
						v222 = base.F32_convert_i64_s(v206)
						v223 = base.F32_convert_i64_s(v207)
						v226 = base.F32_add(base.F32_mul(base.F32_demote_f64(v116), v188), base.F32_convert_i32_s(v113))
						v230 = base.F32_add(base.F32_mul(base.F32_mul(v188, base.F32_demote_f64(v117)), v204), base.F32_convert_i32_s(v111))
						v233 = F_errstart(m, int32(12), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return
						} else {
							if int32(0) <= v111 {
								if v233 == int32(0) {
									if v169|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v307 = base.F32_lt(v226, v221)
									if l0 != int32(2619) {
										v311 = v307
									} else {
										v311 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+40)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = base.F64_promote_f32(v230)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+24)) = base.F64_promote_f32(v222)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+16)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2 + int32(4)
									F_errmsg_internal(m, int32(650856), v32)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										v278 = int32(3138)
										F_errfinish(m, int32(488518), v278, int32(334470))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v169|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v307 = base.F32_lt(v226, v221)
											if l0 != int32(2619) {
												v311 = v307
											} else {
												v311 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
											m.G0 = v32 + int32(112)
											return
										}
									}
								}
							} else {
								if v233 == int32(0) {
									if v169|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v307 = base.F32_lt(v226, v221)
									if l0 != int32(2619) {
										v311 = v307
									} else {
										v311 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+96)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+88)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+80)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+72)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l2 + int32(4)
									F_errmsg_internal(m, int32(650943), v32-int32(-64))
									mBase = m.M
									v276 = m.ExcPending
									if v276 != 0 {
										return
									} else {
										v278 = int32(3142)
										F_errfinish(m, int32(488518), v278, int32(334470))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v169|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v111) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v307 = base.F32_lt(v226, v221)
											if l0 != int32(2619) {
												v311 = v307
											} else {
												v311 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v311)
											m.G0 = v32 + int32(112)
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
}
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_LockRelationOid(m, l0, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_RelationIdGetRelation(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)))
					if v15 == int32(116) {
						v18 = int32(4365684)
						v20 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						*(*int32)(unsafe.Add(mBase, _consts[5])) = v20 | int32(1)
					} else {
					}
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
					switch v26 - int32(83) {
					case 0, 22, 26, 29, 31, 33:
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
						if v34 == int32(0) {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+272))
							if v37 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v37)+128)) = int32(0)
							} else {
							}
							v40 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v40
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v40)
						} else {
							v44 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v44)
						}
					default:
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v29
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v29)
					}
					m.G0 = v6 + int32(16)
					return v12
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg_internal(m, int32(55438), v6)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(487850), int32(61), int32(277006))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
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
	} else {
		v12 = F_RelationIdGetRelation(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)))
				if v15 == int32(116) {
					v18 = int32(4365684)
					v20 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					*(*int32)(unsafe.Add(mBase, _consts[5])) = v20 | int32(1)
				} else {
				}
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
				switch v26 - int32(83) {
				case 0, 22, 26, 29, 31, 33:
					v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
					if v34 == int32(0) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+272))
						if v37 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v37)+128)) = int32(0)
						} else {
						}
						v40 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v40
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v40)
					} else {
						v44 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v44)
					}
				default:
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v29
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v29)
				}
				m.G0 = v6 + int32(16)
				return v12
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(55438), v6)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(487850), int32(61), int32(277006))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
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
func F_relation_openrv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(0)
			v15 = F_RangeVarGetRelidExtended(m, l0, l1, v12, v12, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_RelationIdGetRelation(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
						if v20 == int32(116) {
							v23 = int32(4365684)
							v25 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							*(*int32)(unsafe.Add(mBase, _consts[5])) = v25 | int32(1)
						} else {
						}
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
						switch v31 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
							if v39 == int32(0) {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
								if v42 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v42)+128)) = int32(0)
								} else {
								}
								v45 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v45
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v45)
							} else {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v49)
							}
						default:
							v34 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v34
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v34)
						}
						m.G0 = v6 + int32(16)
						return v17
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
							F_errmsg_internal(m, int32(55438), v6)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(487850), int32(61), int32(277006))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
	} else {
		v12 = int32(0)
		v15 = F_RangeVarGetRelidExtended(m, l0, l1, v12, v12, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_RelationIdGetRelation(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
					if v20 == int32(116) {
						v23 = int32(4365684)
						v25 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						*(*int32)(unsafe.Add(mBase, _consts[5])) = v25 | int32(1)
					} else {
					}
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
					switch v31 - int32(83) {
					case 0, 22, 26, 29, 31, 33:
						v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
						if v39 == int32(0) {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
							if v42 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v42)+128)) = int32(0)
							} else {
							}
							v45 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v45
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v45)
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v49)
						}
					default:
						v34 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v34
						*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v34)
					}
					m.G0 = v6 + int32(16)
					return v17
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
						F_errmsg_internal(m, int32(55438), v6)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(487850), int32(61), int32(277006))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
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
func F_relation_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(0)
			v16 = F_RangeVarGetRelidExtended(m, l0, l1, l2, v14, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 != 0 {
					v18 = F_RelationIdGetRelation(m, v16)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
								F_errmsg_internal(m, int32(55438), v8)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(487850), int32(61), int32(277006))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
							v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
							if v23 == int32(116) {
								v26 = int32(4365684)
								v28 = *(*int32)(unsafe.Add(mBase, _consts[5]))
								*(*int32)(unsafe.Add(mBase, _consts[5])) = v28 | int32(1)
							} else {
							}
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
							switch v34 - int32(83) {
							case 0, 22, 26, 29, 31, 33:
								v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
								if v42 == int32(0) {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
									if v45 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v45)+128)) = int32(0)
									} else {
									}
									v48 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v48
									*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v48)
								} else {
									v52 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v52)
								}
							default:
								v37 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v37
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v37)
							}
							v55 = v18
							m.G0 = v8 + int32(16)
							return v55
						}
					}
				} else {
					v55 = int32(0)
					m.G0 = v8 + int32(16)
					return v55
				}
			}
		}
	} else {
		v14 = int32(0)
		v16 = F_RangeVarGetRelidExtended(m, l0, l1, l2, v14, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v18 = F_RelationIdGetRelation(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v18 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
							F_errmsg_internal(m, int32(55438), v8)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(487850), int32(61), int32(277006))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
						if v23 == int32(116) {
							v26 = int32(4365684)
							v28 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							*(*int32)(unsafe.Add(mBase, _consts[5])) = v28 | int32(1)
						} else {
						}
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
						switch v34 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[6])))
							if v42 == int32(0) {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
								if v45 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v45)+128)) = int32(0)
								} else {
								}
								v48 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v48
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v48)
							} else {
								v52 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v52)
							}
						default:
							v37 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v37
							*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v37)
						}
						v55 = v18
						m.G0 = v8 + int32(16)
						return v55
					}
				}
			} else {
				v55 = int32(0)
				m.G0 = v8 + int32(16)
				return v55
			}
		}
	}
}

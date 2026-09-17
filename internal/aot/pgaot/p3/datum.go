package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumRestore(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5 + int32(4)
	if v6 == int32(-2) {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v16)
		if v6 == int32(-1) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20 + int32(4)
			return v21
		} else {
			v26 = F_palloc(m, v6)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v6 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					base.MemoryCopy(m, v26, v30, v6)
				} else {
				}
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32 + v6
				return v26
			}
		}
	}
}
func F_datumSerialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	if l1 == int32(0) {
		if l2 == int32(0) {
			if l3 != int32(-1) {
				v21 = F_datumGetSize(m, l0, int32(0), l3)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					v27 = v25 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
					v62 = v21
					v63 = v27
					if v62 != 0 {
						base.MemoryCopy(m, v63, l0, v62)
					} else {
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v66 + v62
					return
				}
			} else {
				v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v12 != int32(1) {
					v21 = F_datumGetSize(m, l0, int32(0), l3)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						v27 = v25 + int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
						v62 = v21
						v63 = v27
						if v62 != 0 {
							base.MemoryCopy(m, v63, l0, v62)
						} else {
						}
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v66 + v62
						return
					}
				} else {
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					if v15&int32(254) == int32(2) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
						v38 = F_EOH_get_flat_size(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = v38
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v44 = v42 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
							if v37 == int32(0) {
								v62 = v38
								v63 = v44
								if v62 != 0 {
									base.MemoryCopy(m, v63, l0, v62)
								} else {
								}
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v66 + v62
								return
							} else {
								v48 = F_palloc(m, v38)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_EOH_flatten_into(m, v37, v48, v38)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										if v38 != 0 {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											base.MemoryCopy(m, v52, v48, v38)
										} else {
										}
										v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54 + v38
										F_pfree(m, v48)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v21 = F_datumGetSize(m, l0, int32(0), l3)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v27 = v25 + int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
							v62 = v21
							v63 = v27
							if v62 != 0 {
								base.MemoryCopy(m, v63, l0, v62)
							} else {
							}
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v66 + v62
							return
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(-1)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v32 + int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l0
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v70 + int32(4)
			return
		}
	} else {
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(-2)
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v70 + int32(4)
		return
	}
}
func F_datum_compute_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v8 = base.B2i32(l3 != int32(-1))
	if v8|base.B2i32(l4 == int32(112)) == int32(0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v14&int32(3) != 0 {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v27&int32(1) != 0 {
				v51 = l0
				v52 = v27
				if v52&int32(255) == int32(1) {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v88 = v51
						v90 = int32(6)
						return v88 + v90
					} else {
						v65 = int32(18)
						if v58 == v65 {
							v69 = v65
						} else {
							v69 = int32(2)
						}
						return v69 + v51
					}
				} else {
					if v52&int32(1) != 0 {
						return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
					}
				}
			} else {
				switch l2 - int32(99) {
				case 0:
					v45 = l0
				case 1:
					v45 = (l0 + int32(7)) & int32(-8)
				default:
					v45 = (l0 + int32(1)) & int32(-2)
				case 6:
					v45 = (l0 + int32(3)) & int32(-4)
				}
				if int32(0) < l3 {
					v88 = v45
					v90 = l3
					return v88 + v90
				} else {
					if l3 != int32(-1) {
						v85 = F_strlen(m, l1)
						mBase = m.M
						v88 = v45
						v90 = v85 + int32(1)
						return v88 + v90
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v51 = v45
						v52 = v50
						if v52&int32(255) == int32(1) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v88 = v51
								v90 = int32(6)
								return v88 + v90
							} else {
								v65 = int32(18)
								if v58 == v65 {
									v69 = v65
								} else {
									v69 = int32(2)
								}
								return v69 + v51
							}
						} else {
							if v52&int32(1) != 0 {
								return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v21 = int32(base.Ui32(v17)>>(uint(int32(2))%32)) - int32(3)
			if base.Ui32(int32(127)) < base.Ui32(v21) {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v27&int32(1) != 0 {
					v51 = l0
					v52 = v27
					if v52&int32(255) == int32(1) {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
						if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v88 = v51
							v90 = int32(6)
							return v88 + v90
						} else {
							v65 = int32(18)
							if v58 == v65 {
								v69 = v65
							} else {
								v69 = int32(2)
							}
							return v69 + v51
						}
					} else {
						if v52&int32(1) != 0 {
							return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
						}
					}
				} else {
					switch l2 - int32(99) {
					case 0:
						v45 = l0
					case 1:
						v45 = (l0 + int32(7)) & int32(-8)
					default:
						v45 = (l0 + int32(1)) & int32(-2)
					case 6:
						v45 = (l0 + int32(3)) & int32(-4)
					}
					if int32(0) < l3 {
						v88 = v45
						v90 = l3
						return v88 + v90
					} else {
						if l3 != int32(-1) {
							v85 = F_strlen(m, l1)
							mBase = m.M
							v88 = v45
							v90 = v85 + int32(1)
							return v88 + v90
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							v51 = v45
							v52 = v50
							if v52&int32(255) == int32(1) {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
								if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v88 = v51
									v90 = int32(6)
									return v88 + v90
								} else {
									v65 = int32(18)
									if v58 == v65 {
										v69 = v65
									} else {
										v69 = int32(2)
									}
									return v69 + v51
								}
							} else {
								if v52&int32(1) != 0 {
									return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
								}
							}
						}
					}
				}
			} else {
				return l0 + v21
			}
		}
	} else {
		if l3 != int32(-1) {
			switch l2 - int32(99) {
			case 0:
				v45 = l0
			case 1:
				v45 = (l0 + int32(7)) & int32(-8)
			default:
				v45 = (l0 + int32(1)) & int32(-2)
			case 6:
				v45 = (l0 + int32(3)) & int32(-4)
			}
			if int32(0) < l3 {
				v88 = v45
				v90 = l3
				return v88 + v90
			} else {
				if l3 != int32(-1) {
					v85 = F_strlen(m, l1)
					mBase = m.M
					v88 = v45
					v90 = v85 + int32(1)
					return v88 + v90
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v51 = v45
					v52 = v50
					if v52&int32(255) == int32(1) {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
						if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v88 = v51
							v90 = int32(6)
							return v88 + v90
						} else {
							v65 = int32(18)
							if v58 == v65 {
								v69 = v65
							} else {
								v69 = int32(2)
							}
							return v69 + v51
						}
					} else {
						if v52&int32(1) != 0 {
							return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
						}
					}
				}
			}
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v27&int32(1) != 0 {
				v51 = l0
				v52 = v27
				if v52&int32(255) == int32(1) {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v88 = v51
						v90 = int32(6)
						return v88 + v90
					} else {
						v65 = int32(18)
						if v58 == v65 {
							v69 = v65
						} else {
							v69 = int32(2)
						}
						return v69 + v51
					}
				} else {
					if v52&int32(1) != 0 {
						return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
					}
				}
			} else {
				switch l2 - int32(99) {
				case 0:
					v45 = l0
				case 1:
					v45 = (l0 + int32(7)) & int32(-8)
				default:
					v45 = (l0 + int32(1)) & int32(-2)
				case 6:
					v45 = (l0 + int32(3)) & int32(-4)
				}
				if int32(0) < l3 {
					v88 = v45
					v90 = l3
					return v88 + v90
				} else {
					if l3 != int32(-1) {
						v85 = F_strlen(m, l1)
						mBase = m.M
						v88 = v45
						v90 = v85 + int32(1)
						return v88 + v90
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						v51 = v45
						v52 = v50
						if v52&int32(255) == int32(1) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
							if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v88 = v51
								v90 = int32(6)
								return v88 + v90
							} else {
								v65 = int32(18)
								if v58 == v65 {
									v69 = v65
								} else {
									v69 = int32(2)
								}
								return v69 + v51
							}
						} else {
							if v52&int32(1) != 0 {
								return int32(base.Ui32(v52&int32(254))>>(uint(int32(1))%32)) + v51
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								return int32(base.Ui32(v80)>>(uint(int32(2))%32)) + v51
							}
						}
					}
				}
			}
		}
	}
}
